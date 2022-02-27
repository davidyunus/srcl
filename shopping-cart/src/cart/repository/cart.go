package repository

import (
	"github.com/srcl/shopping-cart/src/domain"
)

type cartRepository struct {
	carts []domain.Cart
}

func NewCartRepository() domain.CartRepository {
	return &cartRepository{}
}

func (c *cartRepository) AddProduct(name string, qty int) error {
	noProducts := true
	for i, cart := range c.carts {
		if cart.Name == name {
			c.carts[i].Quantity += qty
			noProducts = false
		}
	}

	if noProducts {
		c.carts = append(c.carts, domain.Cart{
			Name:     name,
			Quantity: qty,
		})
	}
	return nil
}

func (c *cartRepository) GetCart() ([]domain.Cart, error) {
	return c.carts, nil
}

func (c *cartRepository) DeleteProduct(name string) error {
	for i, cart := range c.carts {
		if cart.Name == name {
			c.carts = remove(c.carts, i)
		}
	}
	return nil
}
