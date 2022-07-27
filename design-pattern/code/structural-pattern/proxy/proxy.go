package proxy

import "fmt"

type Client struct {
	db *DbClient
}

func (dao Client) update() (int, error) {
	db := dao.db
	tx := db.tx
	tx.open()
	result, err := db.update()
	if err != nil {
		tx.rollback()
	} else if result > 0 {
		tx.commit()
	}
	tx.close()
	return result, err
}

type Tx interface {
	open()
	rollback()
	commit()
	close()
}
type Operation interface {
	update() (int, error)
}
type DbClient struct {
	tx Tx
}

func (db DbClient) update() (int, error) {
	fmt.Println("do update")
	return 1, nil
}

type TransactionManager struct {
}

func (TransactionManager) open() {
	fmt.Println("open a transaction session")
}
func (TransactionManager) rollback() {
	fmt.Println("do rollback")
}
func (TransactionManager) commit() {
	fmt.Println("do commit")
}
func (TransactionManager) close() {
	fmt.Println("close the transaction session")
}

func NewClient() Client {
	return Client{
		db: &DbClient{
			tx: &TransactionManager{},
		},
	}
}
