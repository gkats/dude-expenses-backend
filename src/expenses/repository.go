package expenses

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (this *Repository) CreateExpense(params ExpenseParams) (Expense, error) {
	expense := NewExpense(params)

	query := `
	INSERT INTO expenses (price_cents, date, tag, notes)
	VALUES($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	err := this.db.QueryRow(query, expense.PriceCents, expense.Date, expense.Tag, expense.Notes).Scan(&expense.Id, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (this *Repository) GetExpenses(params FilterParams) (Expenses, error) {
	expenses := Expenses{}
	query := "SELECT * FROM expenses"
	var err error
	var rows *sql.Rows

	if len(params.From) > 0 && len(params.To) > 0 {
		rows, err = this.db.Query(query+" WHERE date >= $1 AND date <= $2", params.From, params.To)
	} else {
		rows, err = this.db.Query(query)
	}
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return expenses, err
	}

	expense := Expense{}
	for rows.Next() {
		rows.Scan(&expense.Id, &expense.PriceCents, &expense.Date, &expense.Tag, &expense.Notes, &expense.CreatedAt, &expense.UpdatedAt)
		expenses.Expenses = append(expenses.Expenses, expense)
	}
	return expenses, nil
}
