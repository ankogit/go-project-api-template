package repository

type UserRepository interface {
	All() ([]User, error)
	Find(id uint) (User, error)
	Create(user *User) (int64, error)
	Update(user *User) (int64, error)
	Patch(user map[string]interface{}) (int64, error)
	Delete(id string) (int64, error)
}
