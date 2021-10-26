package cart

type Service interface {
	CreateCart(int, CreateCartData) error
	FindCartByUserID(userID int) ([]CartProduct, error)
}

type Repository interface {
	FindCartByUserID(int) (*Cart, error)
	CreateCart(Cart) (*Cart, error)
	UpdateCart(Cart) (*Cart, error)
	FindCartProduct(int, int) (*CartProduct, error)
	FindCartProductByCart(int) ([]CartProduct, error)
	CreateCartProduct(CartProduct) (*CartProduct, error)
	UpdateCartProduct(CartProduct) (*CartProduct, error)
	SumPrice(int) int
}
