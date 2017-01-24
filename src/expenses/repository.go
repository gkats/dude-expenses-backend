package expenses

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) CreateExpense(params ExpenseParams) (Expense, error) {
	expense := NewExpense(params)

	query := `
	INSERT INTO expenses (user_id, price_cents, date, tag, notes)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at
	`
	err := repository.db.QueryRow(
		query,
		expense.UserId, expense.PriceCents, expense.Date, expense.Tag, expense.Notes,
	).Scan(&expense.Id, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (repository *Repository) GetExpenses(params FilterParams) (Expenses, error) {
	expenses := Expenses{Expenses: make([]Expense, 0)}
	var err error
	var rows *sql.Rows

	qb := newQueryBuilder("SELECT * FROM expenses", params)
	query, args := qb.Build()
	rows, err = repository.db.Query(query, args...)
	defer rows.Close()
	if err != nil {
		// TODO Log error!
		return expenses, err
	}

	expense := Expense{}
	for rows.Next() {
		rows.Scan(
			&expense.Id, &expense.UserId, &expense.PriceCents, &expense.Date, &expense.Tag,
			&expense.Notes, &expense.CreatedAt, &expense.UpdatedAt,
		)
		expenses.Expenses = append(expenses.Expenses, expense)
	}
	return expenses, nil
}
