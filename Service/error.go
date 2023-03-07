package myErr

import "fmt"

func Handle(msg string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Error("%s: %w", msg, err)
}
