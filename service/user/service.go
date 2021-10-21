package user

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

type CreateUserData struct {
	Name         string `validate:"required"`
	Email        string `validate:"required"`
	Phone_number string `validate:"required,number"`
	Password     string `validate:"required"`
	Address      string
}

func (s *service) ServiceFuncForUser() error {

	return nil
}
