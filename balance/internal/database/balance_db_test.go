package database

import (
	"database/sql"
	"github.com/marcohnp/fullcycle_eda/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory")
	s.Nil(err)
	s.db = db
	db.Exec("Create table balances (account_id varchar(255), balance float, created_at date,update_at date)")
	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("Drop table balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance := entity.NewBalance("1", 100)
	err := s.balanceDB.Save(balance)
	s.Nil(err)
}

func (s *BalanceDBTestSuite) TestFindByID() {
	accountID := "1"
	balance := entity.NewBalance(accountID, 100)
	err := s.balanceDB.Save(balance)
	s.Nil(err)

	balance, err = s.balanceDB.FindByID(accountID)
	s.Nil(err)
	s.Equal(accountID, balance.AccountID)
	s.Equal(float64(100), balance.Balance)
	s.WithinDuration(time.Now(), balance.CreatedAt, time.Second)
	s.WithinDuration(time.Now(), balance.UpdatedAt, time.Second)
}

func (s *BalanceDBTestSuite) TestUpdateBalance() {
	accountID := "1"
	balance := entity.NewBalance(accountID, 100)
	err := s.balanceDB.Save(balance)
	s.Nil(err)

	balance.UpdateBalance(200)
	err = s.balanceDB.UpdateBalance(balance)
	s.Nil(err)

	updatedBalance, err := s.balanceDB.FindByID(accountID)
	s.Nil(err)
	s.Equal(accountID, updatedBalance.AccountID)
	s.Equal(float64(200), updatedBalance.Balance)
	s.WithinDuration(time.Now(), updatedBalance.UpdatedAt, time.Second)
}
