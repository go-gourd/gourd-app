package global

import (
	"gorm.io/gorm"
)

var Dbs map[string]*gorm.DB

func init() {
	Dbs = make(map[string]*gorm.DB)
}

// SetDb sets the database connection for the given name.
func SetDb(name string, db *gorm.DB) {
	Dbs[name] = db
}

// GetDb returns the database connection for the given name.
func GetDb(name string) *gorm.DB {
	return Dbs[name]
}
