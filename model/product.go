package model

type Product struct {
	Name  string
	Price float32
}

// Go 中没有特定的 getter 和 setter方法，一般通过方法来实现

func (product *Product) GetName() string {
	return product.Name
}

func (product *Product) SetName(name string) {
	product.Name = name
}

func (product *Product) GetPrice() float32 {
	return product.Price
}

func (product *Product) SetPrice(price float32) {
	product.Price = price
}
