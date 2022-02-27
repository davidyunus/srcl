package repository

import "github.com/srcl/shopping-cart/domain"

func remove(s []domain.Cart, i int) []domain.Cart {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
