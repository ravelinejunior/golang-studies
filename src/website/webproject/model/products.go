package model

import "module/src/website/db"

func SearchAllProducts() []Product {
	db := db.ConnectDatabase()
	selectAllProdcutsFromDatabase, err := db.Query("SELECT * FROM PRODUCTS_TABLE ORDER BY ID ASC")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}

	products := []Product{}

	for selectAllProdcutsFromDatabase.Next() {
		var id, amount int
		var p_name, description string
		var price float64

		err = selectAllProdcutsFromDatabase.Scan(&id, &p_name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = int32(id)
		p.Name = p_name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name string, description string, price float64, amount int) {
	db := db.ConnectDatabase()

	insertDataOnDatabase, err := db.Prepare("INSERT INTO PRODUCTS_TABLE(P_NAME, DESCRIPTION, PRICE, AMOUNT) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertDataOnDatabase.Exec(name, description, price, amount)

	defer db.Close()
}

func EditProduct(productId string) Product {
	db := db.ConnectDatabase()

	databaseProduct, err := db.Query("SELECT * FROM PRODUCTS_TABLE WHERE Id= $1", productId)

	if err != nil {
		panic(err.Error())
	}

	selectedProduct := Product{}

	for databaseProduct.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = databaseProduct.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		selectedProduct.Id = int32(id)
		selectedProduct.Name = name
		selectedProduct.Description = description
		selectedProduct.Price = price
		selectedProduct.Amount = amount

	}
	defer db.Close()
	return selectedProduct
}

func UpdateProduct(id, amount int, name, description string, price float64) {
	db := db.ConnectDatabase()

	updateProductDatabase, err := db.Prepare("UPDATE PRODUCTS_TABLE SET P_NAME=$1, DESCRIPTION=$2, AMOUNT=$3, PRICE=$4 WHERE ID=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProductDatabase.Exec(name, description, amount, price, id)

	defer db.Close()
}

func DeleteProduct(productId string) {
	db := db.ConnectDatabase()

	deleteDataOnDatabase, err := db.Prepare("DELETE FROM PRODUCTS_TABLE WHERE Id= $1")

	if err != nil {
		panic(err.Error())
	}

	deleteDataOnDatabase.Exec(productId)

	defer db.Close()
}

type Product struct {
	Id          int32
	Name        string
	Description string
	Price       float64
	Amount      int
}
