package main

import (
	"github.com/go-gourd/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gourd/support/gorm_gen/gen_tools"
	"gourd/support/gorm_gen/methods"
)

// main 模型代码生成
func main() {

	dbMysql, err := mysql.GetDb("mysql")
	if err != nil {
		panic(err.Error())
	}

	tagCreate := func(tag field.GormTag) field.GormTag {
		tag.Set("type", "int unsigned")
		tag.Set("default", "0")
		tag.Set("autoCreateTime", "")
		return tag
	}

	tagUpdate := func(tag field.GormTag) field.GormTag {
		tag.Set("type", "int unsigned")
		tag.Set("default", "0")
		tag.Set("autoUpdateTime", "")
		return tag
	}

	tagDelete := func(tag field.GormTag) field.GormTag {
		tag.Set("type", "int unsigned")
		tag.Set("default", "0")
		return tag
	}

	comOpts := []gen.ModelOpt{
		//自动时间戳字段属性
		gen.FieldGORMTag("create_time", tagCreate),
		gen.FieldGORMTag("update_time", tagUpdate),
		gen.FieldType("create_time", "uint32"),
		gen.FieldType("update_time", "uint32"),

		//软删除字段属性
		gen.FieldGORMTag("delete_time", tagDelete),
		gen.FieldType("delete_time", "soft_delete.DeletedAt"),
		//Json序列化
		gen.WithMethod(methods.JsonMethod{}),
	}

	//定义数据库
	shopDb := gen_tools.Database{
		DB:      dbMysql,
		ComOpts: comOpts,
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
