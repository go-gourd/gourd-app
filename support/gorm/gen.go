package main

import (
	"github.com/go-gourd/mysql"
	"gorm.io/gen"
	"gourd/support/gorm/gen_tools"
	"gourd/support/gorm/methods"
	"gourd/support/gorm/tags"
)

// main 模型代码生成
func main() {

	dbMysql, err := mysql.GetDb("mysql")
	if err != nil {
		panic(err.Error())
	}

	comOpts := []gen.ModelOpt{
		//自动时间戳字段属性
		gen.FieldGORMTag("create_time", tags.CreateField),
		gen.FieldGORMTag("update_time", tags.UpdateField),
		gen.FieldType("create_time", "uint32"),
		gen.FieldType("update_time", "uint32"),

		//软删除字段属性
		gen.FieldGORMTag("delete_time", tags.DeleteField),
		gen.FieldType("delete_time", "soft_delete.DeletedAt"),
		//Json序列化
		gen.WithMethod(methods.JsonMethod{}),
	}

	//定义数据库
	shopDb := gen_tools.Database{
		DB:      dbMysql,
		ComOpts: &comOpts,
		Tables: []gen_tools.Table{
			{
				Name:   "user",
				Relate: &[]gen_tools.TableRelate{
					//关联账户表
					//{TableName: "account", FieldName: "Account", Type: field.HasOne, ForeignKey: "user_id", LocalKey: "id"},
				},
			},
		},
	}

	// 开始生成orm代码
	gen_tools.GenTable(shopDb)

}
