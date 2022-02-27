package domain

type (
	CartUsecase interface {
		AddProduct(name string, qty int) error
		GetCart() ([]Cart, error)
		DeleteProduct(name string) error
	}

	CartRepository interface {
		AddProduct(name string, qty int) error
		GetCart() ([]Cart, error)
		DeleteProduct(name string) error
	}
)

type Cart struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
