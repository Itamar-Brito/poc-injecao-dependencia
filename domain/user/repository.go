package user

type Repository interface {
	FindByID(id int) (*User, error)
}
