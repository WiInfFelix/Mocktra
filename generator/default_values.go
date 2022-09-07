package generator

var DefaultLocations []Location
var DefaultItems []Item

func InitDefaults() {

	loc1 := Location{Id: 1, Location: "Hamburg"}
	loc2 := Location{Id: 2, Location: "Berlin"}
	loc3 := Location{Id: 3, Location: "Frankfurt"}

	DefaultLocations = append(DefaultLocations, loc1, loc2, loc3)

	item1 := Item{Id: 1, Name: "Red wine", Value: 12.0, Category: Alcohol}
	item2 := Item{Id: 2, Name: "White wine", Value: 9.0, Category: Alcohol}
	item3 := Item{Id: 3, Name: "Beer", Value: 2.0, Category: Alcohol}
	item4 := Item{Id: 4, Name: "Apple", Value: 1.0, Category: Fruit}
	item5 := Item{Id: 5, Name: "Banana", Category: Fruit}
	item6 := Item{Id: 6, Name: "Kiwi", Category: Fruit}
	item7 := Item{Id: 7, Name: "Shampoo", Category: Hygiene}
	item8 := Item{Id: 8, Name: "Soap", Category: Hygiene}

	DefaultItems = append(DefaultItems, item1, item2, item3, item4, item5, item6, item7, item8)

}
