package usecase

import "github.com/srcl/shopping-cart/domain"

type cartUsecase struct {
	cartRepository domain.CartRepository
}

func NewCartUsecase(cartRepository domain.CartRepository) domain.CartUsecase {
	return &cartUsecase{
		cartRepository: cartRepository,
	}
}

func (c *cartUsecase) AddProduct(name string, qty int) error {
	err := c.cartRepository.AddProduct(name, qty)
	if err != nil {
		return err
	}

	return nil
}

func (c *cartUsecase) GetCart() ([]domain.Cart, error) {
	carts, err := c.cartRepository.GetCart()
	if err != nil {
		return nil, nil
	}

	return carts, nil
}

func (c *cartUsecase) DeleteProduct(name string) error {
	err := c.cartRepository.DeleteProduct(name)
	if err != nil {
		return nil
	}

	return nil
}
