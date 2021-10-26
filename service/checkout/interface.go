package checkout

type Service interface {
	CreateCheckout(int, CreateCheckoutData) error
}

type Repository interface {
	CreateCheckout(Checkout) error
}
