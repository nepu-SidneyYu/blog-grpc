package mysql

import (
	"sync"

	"gorm.io/gorm"
)

var (
	_db     *gorm.DB
	_dbOnce sync.Once
)
