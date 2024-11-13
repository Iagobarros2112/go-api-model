package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

//tipo basico repositorio
type ProductRepository struct {
	connection *sql.DB
}

//funcao repositorio
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

//funcao buscar todos os produtos
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

//funcao criar um produto
func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

//funçao achar um produto
func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

//funcao deletar um produto
func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
     
	//query para eu achar tal produto
	query, err := pr.connection.Prepare("DELETE * FROM product WHERE id = $1")
	//se nao achar mostre um erro
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
   
	
	
}
