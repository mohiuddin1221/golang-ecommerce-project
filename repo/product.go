package repo

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductRepo interface {
	Craete(p Product) (*Product, error)
	List() ([]*Product, error)
	Get(productId int) (*Product, error)
	Update(product Product) error
	Delete(productId int) error
}

type productRepo struct {
	productList []*Product
}

// /constructor is a function
// Creating a constuctor for struct
func NewproductRepo() ProductRepo {
	//first of all constructor banaici
	repo := &productRepo{}
	//then Someproduct pass by this constuctor
	generateProduct(repo)
	return repo

}

func (r *productRepo) Craete(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil

}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil

}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, product := range r.productList {
		if productId == product.ID {
			return product, nil

		}
	}
	return nil, nil

}

func (r *productRepo) Update(product Product) error {
	for index, p := range r.productList {
		if p.ID == product.ID {
			r.productList[index] = &product
		}

	}
	return nil

}

func (r *productRepo) Delete(productId int) error {
	var tempstore []*Product
	for _, p := range r.productList {
		if p.ID != productId {
			tempstore = append(tempstore, p)
		}
	}
	r.productList = tempstore
	return nil

}

func generateProduct(r *productRepo) {

	prd1 := &Product{
		ID:          1,
		Title:       "mango",
		Description: "mango is read",
		Price:       200,
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Orange",
		Description: "mango is read",
		Price:       200,
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banan is yellow................",
		Price:       200,
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
}
