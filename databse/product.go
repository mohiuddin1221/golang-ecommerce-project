package databse

var productList []Product

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p

}
func List() []Product {
	return productList
}
func Get(productId int) *Product {

	for _, product := range productList {
		if productId == product.ID {
			return &product

		}
	}
	return nil

}

func Put(product Product) {
	for index, p := range productList {
		if p.ID == product.ID {
			productList[index] = product
		}

	}
}
func Delete(productId int) {
	var tempstore []Product
	for _, p := range productList {
		if p.ID != productId {
			tempstore = append(tempstore, p)
		}
	}
	productList = tempstore

}

func init() {

	prd1 := Product{
		ID:          1,
		Title:       "mango",
		Description: "mango is read",
		Price:       200,
	}
	prd2 := Product{
		ID:          2,
		Title:       "Orange",
		Description: "mango is read",
		Price:       200,
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banan is yellow................",
		Price:       200,
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}
