package modDatabase

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CDBGiga struct {
	dbinst *gorm.DB
}

var g_SingleGiga *CDBGiga = &CDBGiga{}

func GetSingleGiga() *CDBGiga {
  return g_SingleGiga
}

func (instSelf *CDBGiga) Initialize(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}

	instSelf.dbinst = db

	err = instSelf.dbinst.AutoMigrate(&SystemKVData{})
	if err != nil {
		return err
	}


	return err
}