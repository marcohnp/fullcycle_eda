package database

import (
	"database/sql"
	"github.com/marcohnp/fullcycle_eda/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{DB: db}
}

func (b *BalanceDB) FindByID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := b.DB.Prepare("SELECT account_id, balance, created_at, update_at FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountID)
	err = row.Scan(&balance.AccountID, &balance.Balance, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (b *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("INSERT INTO balances (account_id, balance, created_at, update_at) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.AccountID, balance.Balance, balance.CreatedAt, balance.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceDB) UpdateBalance(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET balance = ?, update_at = ? WHERE account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Balance, balance.UpdatedAt, balance.AccountID)
	if err != nil {
		return err
	}
	return nil
}
