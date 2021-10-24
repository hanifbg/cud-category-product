package product

type Service interface {
	AddProduct(CreateProductData) error
	UpdateProduct(int, UpdateProduct) error
}

type Repository interface {
	FindProductById(id int) (*Product, error)
	AddProduct(product Product) error
	UpdateProduct(product Product) error
}
