package cart

import "time"

type CartWithProduct struct {
	ID        int
	CartID    int
	Product   []Product
	Count     int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Product struct {
	ID    int
	Name  string
	Image string
}

func ConvertProduct(id int, name string, image string) Product {
	var product Product
	product.ID = id
	product.Name = name
	product.Image = image

	return product
}
