
package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	_dbService *DBService
)

type DBService struct {
	DB *gorm.DB
}

func InitDBServcie(dialStr string) {
	if _dbService == nil {
		db, err := gorm.Open("mysql", dialStr)
		if err != nil {
			panic(err)
		}

		_dbService = &DBService{
			DB: db,
		}
	}
}

func GetDBService() *DBService {
	if _dbService == nil {
		panic("db service must initialize first")
	}
	return _dbService
}

func (db *DBService) Close() {
	_dbService.DB.Close()
}