package models

import "log"

type Product struct {
	Name string
	Photo string
	Characteristics string
	Description string
}


type ProductMethods interface {
	GetProdName(id int) string
	GetProduct(id int) (p Product)
}

func (c *Connection) GetProdName(id int) string {
	var prod Product
	err := c.DB.QueryRow("select name from products where id=$1 ", id).Scan(&prod.Name)
	if err != nil {
		log.Println(err)
	}
	return prod.Name
}
func (c *Connection) GetProduct(id int) (p Product) {
	var prod Product
	err := c.DB.QueryRowx("select name, photo, characteristics, description from products where id=$1 ", id).
		Scan(&prod.Name,&prod.Photo,&prod.Characteristics,&prod.Description)
	if err != nil {
		log.Println(err)
	}
	return prod
}
