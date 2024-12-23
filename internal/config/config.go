package config

import "database/sql"

func InitDB(db *sql.DB) {
	queryMenuItem1 := `
	Insert into menu_items (Name, Description, Price) values
    ('Caffe Latte', 'Espresso with steamed milk', 3.50)
	`
	queryMenuItem2 := `
	Insert into menu_items (Name, Description, Price) values
    ('Blueberry Muffin', 'Freshly baked muffin with blueberries', 2.00)
	`
	queryMenuItem3 := `
	Insert into menu_items (Name, Description, Price) values
    ('Espresso', 'Strong and bold coffee', 2.50)
	`
	db.Exec(queryMenuItem1)
	db.Exec(queryMenuItem2)
	db.Exec(queryMenuItem3)
	queryInventory1 := `
	Insert into inventory (Name, Quantity, Unit) values
	('Espresso Shot', 500, 'shots')
	`
	queryInventory2 := `
	Insert into inventory (Name, Quantity, Unit) values
	('Milk', 5000, 'ml')
	`
	queryInventory3 := `
	Insert into inventory (Name, Quantity, Unit) values
	('Flour', 10000, 'g')
	`
	queryInventory4 := `
	Insert into inventory (Name, Quantity, Unit) values
	('Blueberries', 2000, 'g')
	`
	queryInventory5 := `
	Insert into inventory (Name, Quantity, Unit) values
	('Sugar', 5000, 'g')
	`
	db.Exec(queryInventory1)
	db.Exec(queryInventory2)
	db.Exec(queryInventory3)
	db.Exec(queryInventory4)
	db.Exec(queryInventory5)

	queryMenuItemsIngridients1 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(1, 1, 1)
	`
	queryMenuItemsIngridients2 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(1, 2, 200)
	`
	queryMenuItemsIngridients3 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(2, 3, 100)
	`
	queryMenuItemsIngridients4 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(2, 4, 20)
	`
	queryMenuItemsIngridients5 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(2, 5, 30)
	`
	queryMenuItemsIngridients6 := `
	insert into menu_item_ingredients (MenuID, IngredientID, Quantity) values
	(3, 1, 1)
	`
	db.Exec(queryMenuItemsIngridients1)
	db.Exec(queryMenuItemsIngridients2)
	db.Exec(queryMenuItemsIngridients3)
	db.Exec(queryMenuItemsIngridients4)
	db.Exec(queryMenuItemsIngridients5)
	db.Exec(queryMenuItemsIngridients6)
}
