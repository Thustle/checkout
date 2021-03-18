package checkout

// Interface to adhere to
type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {

}

func (c Checkout) Scan(item string) {
	//TODO
}

func (c Checkout) GetTotalPrice() int {
	//TODO
	return 0
}
