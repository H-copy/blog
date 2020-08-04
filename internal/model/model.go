package model

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"blog/global"
	"blog/pkg/setting"
)

type Model struct {
	ID uint32 `json:"id" gorm:"primary_key"`
	CreatedBy string `json:"created_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn uint32 `json:"delete_on"`
	IsDel uint8 `json:"is_del"`
}


func NewDBEngine( databaseSetting *setting.DatabaseSetting )(*gorm.DB, error){

	conf := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%t&loc=Local",
	databaseSetting.Username,
	databaseSetting.Password,
	databaseSetting.Host,
	databaseSetting.DBName,
	databaseSetting.Charset,
	databaseSetting.ParseTime,
	)
	
	db, err := gorm.Open(databaseSetting.DBType, conf)

	if err != nil{
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug"{
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
	
}