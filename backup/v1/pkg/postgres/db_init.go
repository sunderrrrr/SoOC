package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Initial(db *sqlx.DB) bool {
	fmt.Println("все ок")
	return true
}
