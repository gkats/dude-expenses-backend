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

func (r *Repository) CreateExpense(params ExpenseParams) (Expense, error) {
	expense := NewExpense(params)

	query := `
	INSERT INTO expenses (user_id, price_cents, date, tag, notes)
	VALUES($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(
		query,
		expense.UserId, expense.PriceCents, expense.Date, expense.Tag, expense.Notes,
	).Scan(&expense.Id, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		return expense, err
	}
	return expense, nil
}

func (r *Repository) GetExpenses(params FilterParams) (Expenses, error) {
	expenses := Expenses{Expenses: make([]Expense, 0)}
	var err error
	var rows *sql.Rows

	qb := newQueryBuilder(
		"SELECT id, user_id, price_cents, date, tag, notes, created_at, updated_at FROM expenses",
		params,
	)
	query, args := qb.Build()
	rows, err = r.db.Query(query, args...)
	defer rows.Close()
	if err != nil {
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

func (r *Repository) GetTags(userId string) (Tags, error) {
	tags := Tags{Tags: make([]string, 0)}

	rows, err := r.db.Query("SELECT DISTINCT tag FROM expenses WHERE user_id = $1", userId)
	defer rows.Close()
	if err != nil {
		return tags, err
	}

	var tag string
	for rows.Next() {
		rows.Scan(&tag)
		tags.Tags = append(tags.Tags, tag)
	}
	return tags, nil
}

func (r *Repository) GetExpense(id string) (*Expense, error) {
	expense := &Expense{}
	query := `
	SELECT id, price_cents, date, tag, notes, user_id, created_at, updated_at
	FROM expenses
	WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&expense.Id, &expense.PriceCents, &expense.Date, &expense.Tag, &expense.Notes,
		&expense.UserId, &expense.CreatedAt, &expense.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *Repository) UpdateExpense(id string, params ExpenseParams) (*Expense, error) {
	expense := NewExpense(params)

	query := `
	UPDATE expenses
	SET price_cents = $1, date = $2, tag = $3, notes = $4
	WHERE id = $5 AND user_id = $6
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		expense.PriceCents, expense.Date, expense.Tag, expense.Notes,
		id, expense.UserId,
	).Scan(&expense.Id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &expense, nil
}
