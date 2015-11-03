package dbmanager

import (
	// "encoding/json"
	"fmt"
)

var products []*Product
var productsMap map[int]*Product

var orders []*Order
var activeOrder *Order

var orderId int

type Product struct {
	Id          int
	DisplayName string
	TypeId      int
}

type OrderedProduct struct {
	Product *Product
	Amount  int
}

type Order struct {
	Id          int
	ProductList []*OrderedProduct
}

func FetchAllProducts() (error, []*Product) {
	if len(products) == 0 {
		products = []*Product{
			&Product{Id: 1, DisplayName: "Хляб", TypeId: 1},
			&Product{Id: 2, DisplayName: "Сирене", TypeId: 2},
			&Product{Id: 3, DisplayName: "Мляко", TypeId: 2},
			&Product{Id: 4, DisplayName: "Coca-cola", TypeId: 3},
			&Product{Id: 5, DisplayName: "Ябълков сок", TypeId: 3},
		}
		productsMap = make(map[int]*Product)
		for _, val := range products {
			productsMap[val.Id] = val
		}
		fmt.Print("%v", productsMap)
	}
	return nil, products
}

func CreateNewOrder() {
	activeOrder = &Order{Id: orderId}
	orderId = orderId + 1
}

func SaveOrder() {
	if activeOrder != nil {
		orders = append(orders, activeOrder)
	}
}

func LoadOrder(orderId int) *Order {
	if orderId < 0 {
		return activeOrder
	}
	for _, val := range orders {
		if val.Id == orderId {
			return val
		}
	}
	return nil
}

func LoadAllOrders() []*Order {
	return orders
}

func AddToOrder(productId int, amount int) {
	if activeOrder == nil {
		activeOrder = &Order{Id: orderId}
		orderId = orderId + 1
	}
	activeOrder.ProductList = append(activeOrder.ProductList, &OrderedProduct{Product: productsMap[productId], Amount: amount})
}
