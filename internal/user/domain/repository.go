package domain

type Repository interface {
	FindByID(id int) (*User, error)
}
