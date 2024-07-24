module gourd

go 1.22

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-gourd/database v1.0.2
	github.com/go-gourd/gourd v1.2.0
	github.com/pelletier/go-toml/v2 v2.2.2
	gorm.io/driver/mysql v1.5.7
	gorm.io/gen v0.3.26
	gorm.io/gorm v1.25.11
	gorm.io/plugin/dbresolver v1.5.2
	gorm.io/plugin/soft_delete v1.2.1
)

replace github.com/go-gourd/database => ../database

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sevlyar/go-daemon v0.1.6 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	gorm.io/datatypes v1.2.1 // indirect
	gorm.io/hints v1.1.2 // indirect
)
