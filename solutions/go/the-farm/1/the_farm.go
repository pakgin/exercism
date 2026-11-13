package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	Err int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.Err)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	switch err {
	case ErrScaleMalfunction:
		if fodder > 0 {
			fodder *= 2
		} else if fodder == 0 {
			return 0.0, err
		}
	case nil:
		break
	default:
		return 0.0, err
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{Err: cows}
	} else if fodder < 0 {
		return 0.0, errors.New("negative fodder")
	}
	return fodder / float64(cows), nil
}
