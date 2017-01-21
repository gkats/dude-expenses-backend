package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"auth"
)

type Expense struct {
	Id int64 "json:id"
	ExpenseParams
}

type ExpenseParams struct {
	Notes string `json:"notes"`
	Tag   string `json:"tag"`
}

func NewExpense(params ExpenseParams) Expense {
	return Expense{ExpenseParams: params}
}

func main() {
	var jsonStream = `{"id": 1, "notes": "The notes", "tag": "thetag"}`
	var expenseParams ExpenseParams

	var expense Expense

	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	err := decoder.Decode(&expenseParams)

	expense = NewExpense(expenseParams)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(expense.Tag)

	// fmt.Println("Crypto")
	// password := "123456"
	// encrypted_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(encrypted_password))
	// err = bcrypt.CompareHashAndPassword(encrypted_password, []byte(password))
	// if err != nil {
	// 	fmt.Println("Invalid password")
	// } else {
	// 	fmt.Println("Password match")
	// }

	fmt.Println("JWT tokens")
	service := auth.NewTokenService()
	token, err := service.CreateToken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
	err = service.ValidateToken(token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Token valid")
	}
}

