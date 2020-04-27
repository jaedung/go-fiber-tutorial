package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // comment
)

// comment
var (
	DBConn *gorm.DB
)
