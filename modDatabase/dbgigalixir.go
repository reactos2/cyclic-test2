package modDatabase

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"cyclictest2/modData"
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

	err = instSelf.dbinst.AutoMigrate(&modData.SystemKVData{})
	if err != nil {
		return err
	}

	result := db.First(&modData.SystemKVData{Name: "key1"})

	if result.Error != nil {
		err = instSelf.dbinst.Save(&modData.SystemKVData{Name: "key1", ValueInt: 0}).Error
		//err = modData.SystemKVData{Name: "key1", ValueInt: 0}.Save()
		if err != nil {
			return err
		}
	}


	return err
}


func (instSelf *CDBGiga) AddVisitTimes() (error, int) {
	data1 := modData.SystemKVData{Name: "key1"}
	result := instSelf.dbinst.First(&data1)
	if result.Error != nil {
		return result.Error, 0
	}
	data1.ValueInt += 1
	err := instSelf.dbinst.Save(&data1).Error
	//err := data1.Save(instSelf.dbinst)
	return err, data1.ValueInt
}