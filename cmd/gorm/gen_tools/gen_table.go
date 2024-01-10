package gen_tools

import (
	"fmt"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

// GenTable 生成表
func GenTable(db Database) {

	// 生成配置
	packName := ""
	if db.Name != "" {
		packName = db.Name + "_"
	}
	conf := gen.Config{
		// 生成的路径
		OutPath:      "./internal/orm/" + packName + "query",
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

	// 生成商城模型
	g := gen.NewGenerator(conf)
	g.UseDB(db.DB)

	// 如果没有指定表则生成全部表
	if len(db.Tables) == 0 {
		g.ApplyBasic(g.GenerateAllTable(*db.ComOpts...)...)
		g.Execute()
		return
	}

	var allTables []any

	// 定义递归处理关联表方法
	var handRelate func(relates []TableRelate) []gen.ModelOpt

	// 处理关联表
	handRelate = func(relates []TableRelate) []gen.ModelOpt {

		var opts []gen.ModelOpt

		for _, relate := range relates {

			// 关联表是否已生成
			rTab := findTable(db.Tables, relate.TableName)
			fmt.Printf("生成关联表: %s\n", relate.TableName)

			if rTab == nil {
				// 如果关联表不存在则生成
				rTab = &Table{Name: relate.TableName}
			}

			// 生成关联表的外键
			tag := field.GormTag{}
			tag.Set("foreignKey", relate.ForeignKey)
			tag.Set("references", relate.LocalKey)

			// 如果有公共参数
			if db.ComOpts != nil {
				rTab.Opts = append(rTab.Opts, *db.ComOpts...)
			}

			// 如果有关联表的参数
			if relate.Relate != nil {
				rTab.Opts = append(rTab.Opts, handRelate(*relate.Relate)...)
			}

			// 生成关联表
			relateModel := g.GenerateModelAs(db.TablePrefix+rTab.Name, rTab.GetModelName(db.TablePrefix), rTab.Opts...)

			opts = append(opts, gen.FieldRelate(relate.Type, relate.FieldName, relateModel, &field.RelateConfig{
				GORMTag: tag,
			}))
		}
		return opts
	}

	//生成指定表
	for _, tab := range db.Tables {

		// 当前表的所有配置属性
		var tableOpts []gen.ModelOpt

		tableOpts = append(tableOpts, tab.Opts...)

		// 如果有公共参数
		if db.ComOpts != nil {
			tableOpts = append(tableOpts, *db.ComOpts...)
		}

		// 处理关联表
		if tab.Relate != nil {
			tableOpts = append(tableOpts, handRelate(*tab.Relate)...)
		}

		newTable := g.GenerateModelAs(db.TablePrefix+tab.Name, tab.GetModelName(db.TablePrefix), tableOpts...)
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
