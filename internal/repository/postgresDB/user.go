package postgresDB

import (
	q "github.com/core-go/sql"
	"gorm.io/gorm"
	"reflect"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) All() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) Find(id uint) (User, error) {
	var user User
	err := r.DB.First(&user, "id = ?", id).Error
	return user, err
}

func (r *UserRepository) Create(user *User) (int64, error) {
	res := r.DB.Create(&user)
	return res.RowsAffected, res.Error
}

func (r *UserRepository) Update(user *User) (int64, error) {
	res := r.DB.Save(&user)
	return res.RowsAffected, res.Error
}

func (r *UserRepository) Patch(user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	jsonColumnMap := q.MakeJsonColumnMap(userType)
	colMap := q.JSONToColumns(user, jsonColumnMap)
	var userModel User
	res := r.DB.Model(&userModel).Where("id = ?", user["id"]).Updates(colMap)
	return res.RowsAffected, res.Error
}

func (r *UserRepository) Delete(id string) (int64, error) {
	var user User
	res := r.DB.Where("id = ?", id).Delete(&user)
	return res.RowsAffected, res.Error
}
