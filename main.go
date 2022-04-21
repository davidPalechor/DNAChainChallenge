package main

import (
	_ "DNAChainChallenge/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//orm.RegisterDriver("postgres", orm.DRPostgres)
	//orm.RegisterDataBase("default",
	//		"postgres",
	//		fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", os.Getenv("DB_HOST")),
	//)
	beego.Run()
}
