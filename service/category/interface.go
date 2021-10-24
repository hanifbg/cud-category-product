package category

type Service interface {
	AddCategory(CreateCategoryData) error
	UpdateCategory(id int, name string, isActive bool) error
}

type Repository interface {
	FindCategoryById(id int) (*Category, error)
	AddCategory(category Category) error
	UpdateCategory(category Category) error
}
