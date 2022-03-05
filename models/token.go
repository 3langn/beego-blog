package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	Id    int64  `json:"id"`
	Token string `orm:"size(256)" json:"token"`
	Type  string `orm:"size(50)" json:"type"`
	User  *User  `orm:"rel(fk)" json:"user"`
}

type jwtCustomClaim struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

func (t *Token) GenerateToken(UserID int64, Type string) (string, error) {
	claims := &jwtCustomClaim{
		UserID,
		Type,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    "go-jwt",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := token.SignedString(getSecretKey())
	if err != nil {
		panic(err)
	}
	return tk, nil
}

func (t *Token) GenerateAuthToken(UserID int64) (*AuthToken, error) {
	accessToken, err := t.GenerateToken(UserID, AccessTokenType)
	if err != nil {
		return nil, err
	}
	refreshToken, err := t.GenerateToken(UserID, RefreshTokenType)
	if err != nil {
		return nil, err
	}

	*t = Token{
		Token: refreshToken,
		Type:  RefreshTokenType,
		User: &User{
			Id: UserID,
		},
	}

	o := orm.NewOrm()
	_, err = o.Insert(t)
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

//func (t *Token) ValidateToken(token string) (*jwt.Token, error) {
//	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
//		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("Unexpected signing method %v ", t_.Header["alg"])
//		}
//		return getSecretKey(), nil
//	})
//}

func getSecretKey() []byte {
	return []byte("secret")
}
