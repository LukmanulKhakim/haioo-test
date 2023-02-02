package repository

import (
	"hiao-test/soal6/products/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Qty         int
	Price       int
}

type User struct {
	gorm.Model
	Fullname string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Profile  string
	Location string
	Products []Product `gorm:"foreignKey:UserID"`
	Carts    []Cart    `gorm:"foreignKey:UserID"`
}

type Cart struct {
	gorm.Model
	IdProduct uint
	UserID    uint
	Name      string `gorm:"-:migration" gorm:"->"`
	Qty       int
	Price     int    `gorm:"-:migration" gorm:"->"`
	Image     string `gorm:"-:migration" gorm:"->"`
}

func FromDomain(dc domain.Core) Product {
	return Product{
		Model:       gorm.Model{ID: dc.ID},
		Name:        dc.Name,
		Description: dc.Description,
		Qty:         dc.Qty,
		Price:       dc.Price,
	}
}

func ToDomain(p Product) domain.Core {
	return domain.Core{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Qty:         p.Qty,
		Price:       p.Price,
	}
}

func ToDomainArray(listProduct []Product) []domain.Core {
	var res []domain.Core
	for _, val := range listProduct {
		res = append(res, domain.Core{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Qty:         val.Qty,
			Price:       val.Price,
		})
	}

	return res
}

func ToDomainDetail(p Product) domain.Core {
	return domain.Core{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Qty:         p.Qty,
		Price:       p.Price,
	}
}
