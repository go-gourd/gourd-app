package gen_tools

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TableRelate struct {
	TableName  string                 //关联的表名
	FieldName  string                 //字段名
	ForeignKey string                 //外键
	LocalKey   string                 //本表键
	Type       field.RelationshipType //关联方式
	Relate     *[]TableRelate
}

type Table struct {
	Name      string         //表名
	ModelName string         //模型名
	Opts      []gen.ModelOpt //表配置选项
	Relate    *[]TableRelate //关联表
}

type Database struct {
	Name         string          //数据库名 用于生成文件夹
	DB           *gorm.DB        //数据库连接
	Tables       []Table         //表
	TablePrefix  string          //表前缀
	ComOpts      *[]gen.ModelOpt //公共配置选项
	OutPath      string          //query code path
	ModelPkgPath string          //generated model code's package name
	config       *gen.Config     //gen配置 如设置将覆盖
}

func (t Table) GetModelName(prefix string) string {
	if t.ModelName != "" {
		return t.ModelName
	}
	ns := schema.NamingStrategy{
		SingularTable: true,   //单数表名
		TablePrefix:   prefix, //表前缀
	}
	return ns.SchemaName(t.Name)
}
