package account

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Account struct {
	ID      int64           `gorm:"primaryKey"`
	Balance decimal.Decimal `gorm:"type:decimal(10,2);check:balance >= 0"`
}

type Transactions struct {
	ID            int64 `gorm:"primaryKey"`
	FromAccountID int64
	ToAccountID   int64
	Amount        decimal.Decimal `gorm:"type:decimal(10,2)"`
}

func CreateAccountTable(db *gorm.DB) {
	db.AutoMigrate(&Account{})
}

func CreateTransactionsTable(db *gorm.DB) {
	db.AutoMigrate(&Transactions{})
}

func CreateAccountData(db *gorm.DB) {
	a1 := Account{
		Balance: decimal.NewFromFloat(50.00),
	}
	a2 := Account{
		Balance: decimal.NewFromFloat(100.00),
	}
	db.Create([]Account{a1, a2})
}

func UpdateAccountData(db *gorm.DB) {
	var a Account
	db.Where("id = ?", 1).Find(&a)
	a.Balance = decimal.NewFromFloat(100.00)
	db.Save(&a)
}

func QueryAccountData(db *gorm.DB) []Account {
	var a []Account
	db.Find(&a)
	return a
}

func QueryTransactions(db *gorm.DB) []Transactions {
	var t []Transactions
	db.Find(&t)
	return t
}

func Transfer(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		var a1 Account
		var a2 Account
		tx.Where("id = ?", 1).Find(&a1)
		tx.Where("id = ?", 2).Find(&a2)
		a1.Balance = a1.Balance.Sub(decimal.NewFromFloat(100.00))
		if err := tx.Save(&a1).Error; err != nil {
			return err
		}
		a2.Balance = a2.Balance.Add(decimal.NewFromFloat(100.00))

		if err := tx.Save(&a2).Error; err != nil {
			return err
		}
		t := Transactions{
			FromAccountID: a1.ID,
			ToAccountID:   a2.ID,
			Amount:        decimal.NewFromFloat(100.00),
		}
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		return nil
	})
}
