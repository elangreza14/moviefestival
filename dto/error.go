package dto

import "fmt"

type ErrorNotFound struct {
	Entity string
}

func (e ErrorNotFound) Error() string {
	return fmt.Sprintf("cannot found valid %s", e.Entity)
}
