package model

import "time"

type Transaction struct {
    ID          int       `json:"id"`
    Date        time.Time `json:"date"`
    Amount      float64   `json:"amount"`
    Type        string    `json:"type"`       // INCOME, EXPENSE, INVESTMENT
    Description string    `json:"description"`
    MonthYear   time.Time `json:"month_year"` // stored as the first day of the month
}
