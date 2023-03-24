package gen_tools

import (
	"fmt"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

// GenTable 生成表
func GenTable(db Database) {

	//生成配置
	packName := ""
	if db.Name != "" {
		packName = db.Name + "_"
	}
	conf := gen.Config{
		OutPath:      "./app/orm/" + packName + "query",
		ModelPkgPath: packName + "model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	}
	if db.OutPath != "" {
		conf.OutPath = db.OutPath
	}
	if db.ModelPkgPath != "" {
		conf.ModelPkgPath = db.ModelPkgPath
	}
	if db.config != nil {
		conf = *db.config
	}

	//生成商城模型
	g := gen.NewGenerator(conf)
	g.UseDB(db.DB)

	//如果没有表则生成全部表
	if len(db.Tables) == 0 {
		g.ApplyBasic(g.GenerateAllTable()...)
		g.Execute()
		return
	}

	var allTables []any

	//生成指定表
	for _, tab := range db.Tables {

		//切片存入当前表的所有选项
		var tableOpts []gen.ModelOpt

		tableOpts = append(tableOpts, tab.Opts...)

		//如果有公共参数
		if len(db.ComOpts) > 0 {
			tableOpts = append(tableOpts, db.ComOpts...)
		}

		if len(tab.Relate) > 0 {
			for _, relate := range tab.Relate {

				rTab := findTable(db.Tables, relate.TableName)
				fmt.Printf("生成关联表: %s\n", relate.TableName)

				if rTab == nil {
					//如果关联表不存在则生成
					rTab = &Table{Name: relate.TableName}
				}

				//生成关联表
				relateModel := g.GenerateModelAs(rTab.Name, rTab.GetModelName(db.TablePrefix), rTab.Opts...)
				tableOpts = append(tableOpts, gen.FieldRelate(relate.Type, relate.FieldName, relateModel, &field.RelateConfig{
					GORMTag: "foreignKey:" + relate.ForeignKey + ";references:" + relate.LocalKey,
				}))

			}
		}

		newTable := g.GenerateModelAs(tab.Name, tab.GetModelName(db.TablePrefix), tableOpts...)
		allTables = append(allTables, newTable)
	}

	g.ApplyBasic(allTables...)

	g.Execute()
}

// 查找表
func findTable(tabs []Table, name string) *Table {
	for _, tab := range tabs {
		if tab.Name == name || tab.ModelName == name {
			return &tab
		}
	}
	return nil
}
