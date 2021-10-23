package category

type Service interface {
	AddCategory(CreateCategoryData) error
}

type Repository interface {
	AddCategory(category Category) error
}
