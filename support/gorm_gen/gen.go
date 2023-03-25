package main

import (
	"github.com/go-gourd/mysql"
	"gorm.io/gen"
	"gourd/support/gorm_gen/gen_tools"
	"gourd/support/gorm_gen/methods"
)

// main 模型代码生成
func main() {

	dbMysql, err := mysql.GetDb("mysql")
	if err != nil {
		panic(err.Error())
	}

	//自动时间戳字段属性
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;default:0;autoCreateTime")
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;default:0;autoUpdateTime")
	softDeleteTimeField := gen.FieldGORMTag("delete_time", "column:delete_time;type:int unsigned;default:0;autoCreateTime")

	//软删除支持
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	//定义数据库
	shopDb := gen_tools.Database{
		DB: dbMysql,
		ComOpts: []gen.ModelOpt{
			gen.WithMethod(methods.JsonMethod{}),
			autoCreateTimeField, autoUpdateTimeField, softDeleteTimeField, softDeleteField,
		},
		Tables: []gen_tools.Table{
			{
				Name:   "user",
				Relate: []gen_tools.TableRelate{
					//关联账户表
					//{TableName: "account", FieldName: "Account", Type: field.HasOne, ForeignKey: "user_id", LocalKey: "id"},
				},
			},
		},
	}

	// 开始生成orm代码
	gen_tools.GenTable(shopDb)

}
