package terminal

import (
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/repository"
)

type Terminal struct {
	usersRepo       repository.UserRepository
	productsRepo    repository.ProductRepository
	cartItemRepo    repository.CartItemRepository
	transactionRepo repository.TransactionRepository
}

func NewTerminal(usersRepo repository.UserRepository, productsRepo repository.ProductRepository, cartItemRepo repository.CartItemRepository, transactionRepo repository.TransactionRepository) Terminal {
	return Terminal{
		usersRepo, productsRepo, cartItemRepo, transactionRepo,
	}
}

func (terminal *Terminal) Start() {
	terminal.UserListPage()
}
