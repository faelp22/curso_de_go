package service

import (
	"fmt"
	"log"

	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
)

type ProdutoServiceInterface interface {
	GetAll() *entity.ProdutoList
	Get(ID int) *entity.Produto
	Put(produto *entity.Produto) (ok bool)
	Delete(ID int) (ok bool)
}

type produto_service struct {
	dbp database.DatabaseInterface
}

func NewProdutoService(dabase_pool database.DatabaseInterface) *produto_service {
	return &produto_service{
		dabase_pool,
	}
}

func (ps *produto_service) GetAll() *entity.ProdutoList {

	DB := ps.dbp.GetDB()

	rows, err := DB.Query("select id, name, code, price from tb_produto limit 100;")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	lista_produtos := &entity.ProdutoList{}

	for rows.Next() {
		p := entity.Produto{}

		if err := rows.Scan(&p.ID, &p.Name, &p.Code, &p.Price); err != nil {
			fmt.Println(err.Error())
		} else {
			lista_produtos.List = append(lista_produtos.List, &p)
		}

	}

	return lista_produtos

}

func Get(dabase_pool database.DatabaseInterface) {
	DB := dabase_pool.GetDB()

	stmt, err := DB.Prepare("select id, name from tb_produto where id = ? AND name = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	prod := entity.Produto{}

	err = stmt.QueryRow(3, "notebook").Scan(&prod.ID, &prod.Name)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Produto ID %d Name %v \n", prod.ID, prod.Name)
}
