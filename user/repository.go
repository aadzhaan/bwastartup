package user

import "gorm.io/gorm"

//interface karena untuk package lain / object lain /struct lain dia mengacu ke respository
// menggunakan huruf kapital awal bersifat public
type Repository interface {
	Save(user User) (User, error)
}

type respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *respository {
	return &respository{db}
}

// membuat function bernama Save untuk repository diatas
func (r *respository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
