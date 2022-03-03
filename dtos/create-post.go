package dtos

type CreatePostDto struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	UserId  int    `json:"user_id"`
}
