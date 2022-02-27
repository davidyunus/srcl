package main

import (
	"log"

	"github.com/srcl/shopping-cart/src/domain"

	cartrepo "github.com/srcl/shopping-cart/src/cart/repository"

	cartuc "github.com/srcl/shopping-cart/src/cart/usecase"
)

var (
	cartUsecase    domain.CartUsecase
	cartRepository domain.CartRepository
)

func main() {
	cartRepository = cartrepo.NewCartRepository()
	cartUsecase = cartuc.NewCartUsecase(cartRepository)

	cartUsecase.AddProduct("Pisang Hijau", 2)

	cartUsecase.AddProduct("Semangka Kuning", 3)

	cartUsecase.AddProduct("Apel Merah", 1)
	cartUsecase.AddProduct("Apel Merah", 4)
	cartUsecase.AddProduct("Apel Merah", 2)

	cartUsecase.DeleteProduct("Semangka Kuning")

	cartUsecase.DeleteProduct("Semangka Merah")

	c, _ := cartUsecase.GetCart()

	log.Println(c)

}
