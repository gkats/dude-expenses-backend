package expenses

import (
	"strconv"
)

type queryBuilder struct {
	base   string
	params FilterParams
}

func newQueryBuilder(base string, params FilterParams) *queryBuilder {
	return &queryBuilder{base: base, params: params}
}

func (qb *queryBuilder) Build() (string, []interface{}) {
	args := make([]interface{}, 0)
	query := qb.base

	if len(qb.params.UserId) > 0 {
		args = append(args, qb.params.UserId)
		query += " WHERE user_id = $" + strconv.Itoa(len(args))
	}
	if len(qb.params.From) > 0 {
		args = append(args, qb.params.From)
		if len(args) > 1 {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " date >= $" + strconv.Itoa(len(args))
	}
	if len(qb.params.To) > 0 {
		args = append(args, qb.params.To)
		if len(args) > 1 {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " date <= $" + strconv.Itoa(len(args))
	}

	return query, args
}
