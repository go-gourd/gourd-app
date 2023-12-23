package tags

import "gorm.io/gen/field"

func CreateField(tag field.GormTag) field.GormTag {
	tag.Set("autoCreateTime", "")
	return tag
}

func UpdateField(tag field.GormTag) field.GormTag {
	tag.Set("autoUpdateTime", "")
	return tag
}
