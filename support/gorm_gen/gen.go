package main

import (
	"github.com/go-gourd/mysql"
	"gorm.io/gen"
	"gourd/support/gorm_gen/methods"
)

// main 模型代码生成
func main() {

	dbMysql, err := mysql.GetDb("mysql")
	if err != nil {
		panic(err.Error())
	}

	// 构造生成器实例
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath:      "./app/orm/query", //curd代码的输出路径
		ModelPkgPath: "model",           //model代码的输出路径

		Mode: gen.WithDefaultQuery | gen.WithoutContext,
	})
	// 设置目标 db
	g.UseDB(dbMysql)

	//// 自定义字段的数据类型
	//dataMap := map[string]func(detailType string) (dataType string){
	//	"tinyint": func(detailType string) (dataType string) { return "int64" },
	//}
	//// 要先于`ApplyBasic`执行
	//g.WithDataTypeMap(dataMap)

	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;default:0;autoCreateTime")
	autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;default:0;autoUpdateTime")
	softDeleteTimeField := gen.FieldGORMTag("delete_time", "column:delete_time;type:int unsigned;default:0;autoCreateTime")
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteTimeField, softDeleteField}

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	//User := g.GenerateModel("user")

	//需要支持json序列化的模型
	fieldOpts = append(fieldOpts, gen.WithMethod(methods.JsonMethod{}))

	// 创建全部模型文件, 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable(fieldOpts...)

	// 创建有关联关系的模型文件
	// 可以用于指定外键
	//Score := g.GenerateModel("score",
	// append(
	//    fieldOpts,
	//    // user 一对多 address 关联, 外键`uid`在 address 表中
	//    gen.FieldRelate(field.HasMany, "user", User, &field.RelateConfig{GORMTag: "foreignKey:UID"}),
	// )...,
	//)

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(allModel...)

	g.Execute()
}
