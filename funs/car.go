package funs

type Car struct {
	Price string
	Color string
}

func (c Car) GetArea() string {

	return c.Price + "&" + c.Color

}
