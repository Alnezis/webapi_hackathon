package app

import (
	"api_hackathon/api"
	"api_hackathon/app/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

var DB *sqlx.DB

var CFG *config.Config

func init() {
	CFG = config.InitCfg()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", CFG.Db.UserName, CFG.Db.Password, CFG.Db.DbName))
	if err != nil {
		api.CheckErrInfo(err, "Connect(\"mysql\"")
		os.Exit(1)
	}
	DB = db
}
