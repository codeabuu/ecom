package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("product not found")
	ErrCantDecoderoducts  = errors.New("product not found")
	ErrUserIdIsNotValid   = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("cant add this product to cart")
	ErrCantRemoveItemCart = errors.New("cant remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("cant update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
