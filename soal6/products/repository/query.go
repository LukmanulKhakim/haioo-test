package repository

import (
	"hiao-test/soal6/config"
	"hiao-test/soal6/products/domain"

	"errors"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) ShowAll(category, name string, page int) ([]domain.Core, error) {
	var resQry []Product
	if page != 0 {
		ofst := (page - 1) * 10
		if err := rq.db.Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else if name != "" {
		ofst := (page - 1) * 10
		if err := rq.db.Where("name like ?", "%"+name+"%").
			Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else if category != "" {
		ofst := (page - 1) * 10
		if err := rq.db.Where("category = ?", category).
			Offset(ofst).Limit(10).Order("created_at desc").
			Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	} else {
		if err := rq.db.Find(&resQry).Error; err != nil {
			return nil, errors.New(config.DATABASE_ERROR)
		}
	}

	return ToDomainArray(resQry), nil
}

func (rq *repoQuery) Insert(newProduct domain.Core) (domain.Core, error) {
	input := FromDomain(newProduct)

	if err := rq.db.Create(&input).Error; err != nil {
		return domain.Core{}, nil
	}

	newProduct = ToDomain(input)

	return newProduct, nil
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry Product
	err := rq.db.Where("id = ?", ID).Delete(&resQry)
	if err != nil {
		return errors.New("cant delete data")
	}

	if err.RowsAffected < 1 {
		return errors.New("row isnt affected")
	}

	return nil
}
