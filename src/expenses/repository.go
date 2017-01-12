package expenses

import (
	"database/sql"
	"log"
	_"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) CreateExpense(params ExpenseParams) (Expense, error) {
	var expense Expense

	sql := `
	INSERT INTO expenses (price_cents, date, tag, notes)
	VALUES($1, $2, $3, $4) RETURNING id
	`
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
		return expense, err
	}
	expense = NewExpense(params)
	err = statement.QueryRow(expense.PriceCents, expense.Date, expense.Tag, expense.Notes).Scan(&expense.Id)
	if err != nil {
		return expense, err
	}
	return expense, nil
}