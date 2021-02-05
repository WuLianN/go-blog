package model

import (
	// "fmt"
	// "time"

	// otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/pkg/setting"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// dsn := databaseSetting.UserName + ":" + databaseSetting.Password + "@tcp(" + databaseSetting.Host + ")/" + databaseSetting.DBName + "?charset=" + databaseSetting.Charset + "&parseTime=" + databaseSetting.ParseTime + "&loc=Local"
	dsn := databaseSetting.UserName + ":" + databaseSetting.Password + "@tcp(" + databaseSetting.Host + ")/" + databaseSetting.DBName + "?charset=" + databaseSetting.Charset + "&parseTime=True" + "&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		
	})


	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		// db.LogMode(true) // v2 版本 已移除该api
	}

	// db.SingularTable(true) // v2 版本 已移除该api

	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns) // v2 版本 已移除该api
	// db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns) // v2 版本 已移除该api
	// otgorm.AddGormCallbacks(db)
	return db, nil
}

// func updateTimeStampForCreateCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		nowTime := time.Now().Unix()
// 		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
// 			if createTimeField.IsBlank {
// 				_ = createTimeField.Set(nowTime)
// 			}
// 		}

// 		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
// 			if modifyTimeField.IsBlank {
// 				_ = modifyTimeField.Set(nowTime)
// 			}
// 		}
// 	}
// }

// func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
// 	if _, ok := scope.Get("gorm:update_column"); !ok {
// 		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	}
// }

// func deleteCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		var extraOption string
// 		if str, ok := scope.Get("gorm:delete_option"); ok {
// 			extraOption = fmt.Sprint(str)
// 		}

// 		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
// 		isDelField, hasIsDelField := scope.FieldByName("IsDel")
// 		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
// 			now := time.Now().Unix()
// 			scope.Raw(fmt.Sprintf(
// 				"UPDATE %v SET %v=%v,%v=%v%v%v",
// 				scope.QuotedTableName(),
// 				scope.Quote(deletedOnField.DBName),
// 				scope.AddToVars(now),
// 				scope.Quote(isDelField.DBName),
// 				scope.AddToVars(1),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		} else {
// 			scope.Raw(fmt.Sprintf(
// 				"DELETE FROM %v%v%v",
// 				scope.QuotedTableName(),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		}
// 	}
// }

// func addExtraSpaceIfExist(str string) string {
// 	if str != "" {
// 		return " " + str
// 	}
// 	return ""
// }