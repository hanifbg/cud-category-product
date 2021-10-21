package user

type Service interface {
	ServiceFuncForUser() error
}

type Repository interface {
	RepoFuncForUser() error
}
