package service

import (
	"fmt"
	"log"

	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
)

type ProdutoServiceInterface interface {
	GetAll() *entity.ProdutoList
	GetByID(ID *int64) *entity.Produto
	Create(produto *entity.Produto) int64
	Update(ID *int64, produto *entity.Produto) int64
	Delete(ID *int64) int64
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

	rows, err := DB.Query("SELECT id, name, code, price FROM tb_produto LIMIT 100")
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

func (ps *produto_service) GetByID(ID *int64) *entity.Produto {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("SELECT id, name, code, price FROM tb_produto WHERE id = $1")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	p := entity.Produto{}

	err = stmt.QueryRow(ID).Scan(&p.ID, &p.Name, &p.Code, &p.Price)
	if err != nil {
		log.Println(err.Error())
	}

	return &p
}

func (ps *produto_service) Create(produto *entity.Produto) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("INSERT INTO tb_produto (name, code, price) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price)
	if err != nil {
		log.Println(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
	}

	return lastId
}

func (ps *produto_service) Update(ID *int64, produto *entity.Produto) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("UPDATE tb_produto SET name = $1, code = $2, price = $3 WHERE id = $4")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(produto.Name, produto.Code, produto.Price, ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}

func (ps *produto_service) Delete(ID *int64) int64 {
	DB := ps.dbp.GetDB()

	stmt, err := DB.Prepare("DELETE FROM tb_produto WHERE id = $1")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	result, err := stmt.Exec(ID)
	if err != nil {
		log.Println(err.Error())
	}

	rowsaff, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
	}

	return rowsaff
}
