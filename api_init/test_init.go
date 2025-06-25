package api_init

import (
	"fmt"
	"github.com/pressly/goose/v3"
)

func TestInit(basePath string) {
	fmt.Println("Init tests ...")
	err := MainInit(basePath)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		InitGlobal.Cfg.DbHost,
		InitGlobal.Cfg.DbUser,
		InitGlobal.Cfg.DbPassword,
		InitGlobal.Cfg.DbName,
		InitGlobal.Cfg.DbPort,
		InitGlobal.Cfg.DbSSLMode,
	)

	db, err := goose.OpenDBWithDriver(InitGlobal.Cfg.DbDriver, dsn)

	if err != nil {
		panic(err)
	}
	err = goose.Up(db, basePath+"/db/migrations")
	if err != nil {
		panic(err)
	}
}
