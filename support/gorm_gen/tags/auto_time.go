package tags

import "gorm.io/gen/field"

func CreateField(tag field.GormTag) field.GormTag {
	tag.Set("type", "int unsigned")
	tag.Set("default", "0")
	tag.Set("autoCreateTime", "")
	return tag
}

func UpdateField(tag field.GormTag) field.GormTag {
	tag.Set("type", "int unsigned")
	tag.Set("default", "0")
	tag.Set("autoUpdateTime", "")
	return tag
}

func DeleteField(tag field.GormTag) field.GormTag {
	tag.Set("type", "int unsigned")
	tag.Set("default", "0")
	return tag
}
