package main

import (
	_ "DNAChainChallenge/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"os"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		fmt.Sprintf("postgres://postgres:postgres@(%s)/postgres", os.Getenv("DB_HOST")),
	)
	beego.Run()
}
