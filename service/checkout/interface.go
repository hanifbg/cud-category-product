package checkout

type Service interface {
	CreateCheckout(int, CreateCheckoutData) error
	UpdateCheckout(int, UpdateCheckout) error
}

type Repository interface {
	CreateCheckout(Checkout) error
	FindCheckoutById(int) (*Checkout, error)
	UpdateCheckout(checkout Checkout) error
}
