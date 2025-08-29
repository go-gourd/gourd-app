package main

import (
	"app/cmd/gorm/methods"
	"app/cmd/gorm/tags"
	"app/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	// 初始化数据库
	dbConfig, err := config.GetDBConfig("main")
	if err != nil {
		panic(err)
	}

	mainDB, err := gorm.Open(mysql.Open(dbConfig.GenerateDsn()))
	if err != nil {
		panic("database connect failed: " + err.Error())
	}

	// 公共属性
	comOpts := []gen.ModelOpt{
		// 自动时间戳字段属性
		gen.FieldGORMTag("create_time", tags.CreateField),
		gen.FieldGORMTag("update_time", tags.UpdateField),
		gen.FieldType("create_time", "int64"),
		gen.FieldType("update_time", "int64"),

		// 软删除字段属性
		gen.FieldType("delete_time", "soft_delete.DeletedAt"),

		// Json序列化
		gen.WithMethod(methods.JsonMethod{}),
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/orm/query",
		ModelPkgPath: "model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(mainDB)

	// 生成所有表
	//g.ApplyBasic(g.GenerateAllTable(comOpts...)...)

	var allTables []any

	// User
	appModel := g.GenerateModel("user", comOpts...)
	allTables = append(allTables, appModel)

	// 生成指定表
	g.ApplyBasic(allTables...)

	// 执行生成
	g.Execute()
}
