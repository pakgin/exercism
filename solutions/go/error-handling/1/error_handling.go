package erratum

import (
	"errors"
	"fmt"
)

func Use(opener ResourceOpener, input string) error {
	resource, err := opener()
	defer resource.Close()

	if err != nil && !errors.Is(err, TransientError{}) {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	resource.Frob(input)

	return err
}
