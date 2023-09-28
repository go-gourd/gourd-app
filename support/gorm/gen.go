package main

import (
	"github.com/go-gourd/database"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gourd/support/gorm/gen_tools"
	"gourd/support/gorm/methods"
	"gourd/support/gorm/tags"
)

// main 模型代码生成
func main() {

	// 初始化数据库
	dbConfig := database.GetConfig("mysql")
	if dbConfig == nil {
		panic("mysql config is nil")
	}
	mysqlDb, err := gorm.Open(mysql.Open(dbConfig.GenerateDsn()))
	if err != nil {
		panic("mysql connect failed: " + err.Error())
	}

	comOpts := []gen.ModelOpt{
		//自动时间戳字段属性
		gen.FieldGORMTag("create_time", tags.CreateField),
		gen.FieldGORMTag("update_time", tags.UpdateField),
		gen.FieldType("create_time", "uint"),
		gen.FieldType("update_time", "uint"),

		//软删除字段属性
		gen.FieldType("delete_time", "soft_delete.DeletedAt"),
		//Json序列化
		gen.WithMethod(methods.JsonMethod{}),
	}

	//定义数据库
	shopDb := gen_tools.Database{
		DB:      mysqlDb,
		ComOpts: &comOpts,
		// 需要生成代码的数据表，需要被关联的模型需要放在前面
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
