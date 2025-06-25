package api_init

import (
	"gorm.io/gorm"
	"user-service/db"
	"user-service/pkg/config"
)

type InitGlobalStruct struct {
	Dbh *gorm.DB
	Cfg *config.Config
}

var InitGlobal *InitGlobalStruct

func MainInit(basePath string) error {
	cfg := &config.Config{}
	err := cfg.InitConfig(basePath)
	if err != nil {
		return err
	}

	dbh, err := db.ConnectionFactory(cfg)

	if err != nil {
		return err
	}

	InitGlobal = &InitGlobalStruct{
		Dbh: dbh,
		Cfg: cfg,
	}

	return nil
}

func GetDbh() *gorm.DB {
	return InitGlobal.Dbh
}
