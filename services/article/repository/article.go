package repository

import (
	"server/services/article"
	"github.com/jinzhu/gorm"
)

type mysqlArticleRepository struct {
	Db *gorm.DB
}

func NewMysqlArticleRepository(DB *gorm.DB) article.Repository {
	return &mysqlArticleRepository{DB}
}
