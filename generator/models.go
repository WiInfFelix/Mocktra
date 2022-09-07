package generator

type ItemCategory int

const (
	Alcohol ItemCategory = iota
	Fruit
	Hygiene
)

type Item struct {
	Id       int
	Name     string
	Value    float32
	Category ItemCategory
}

type ItemPurchase struct {
	Item   Item
	Amount int
}

type Transaction struct {
	Id            int
	ItemPurchases []ItemPurchase
}

type Location struct {
	Id           int
	Location     string
	Transactions []Transaction
}

type BalanceSheet struct {
	Locations []Location
}
