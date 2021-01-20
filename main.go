package main

import (
	"fmt"

	"github.com/beego/beego/v2"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/plugins/cors"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	//orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))

	orm.RegisterDataBase("default", "postgres", "postgres://"+
		LimpiarVariableentorno("PGuser")+":"+
		LimpiarVariableentorno("PGpass")+"@"+
		LimpiarVariableentorno("PGhost")+":"+
		LimpiarVariableentorno("PGport")+"/"+
		LimpiarVariableentorno("PGdb")+"?sslmode=disable&search_path="+
		LimpiarVariableentorno("PGschema")+"")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}

func LimpiarVariableentorno(key string) (respuesta string) {
	respuesta, _ = beego.AppConfig.String(key)
	fmt.Println(respuesta)
	return
}
