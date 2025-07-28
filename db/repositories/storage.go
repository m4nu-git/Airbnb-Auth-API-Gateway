package db

type Storage struct { // faculitates dependency injection for repository
	UserRepository UserRepository
}

func NewStorage() *Storage {
	return &Storage{
		UserRepository: &UserRepositoryImpl{},
	}
}