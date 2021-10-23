package product

type Service interface {
	AddProduct(CreateProductData) error
}

type Repository interface {
	AddProduct(product Product) error
}
