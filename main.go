package main

import (
	"bee-playaround1/models"
	_ "bee-playaround1/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:saota1278@localhost:5435/bee?sslmode=disable")
	orm.RegisterModel(new(models.User), new(models.Post), new(models.Token))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
