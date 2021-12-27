package command

import (
	"errors"
	"fmt"
)

var ErrUniqueConstraintViolated = errors.New("unique constraint violated")

type uniqueConstraints struct {
	m map[string]bool
}

func (uc *uniqueConstraints) checkUniqueConstraint(value string) error {
	if uc.m[value] {
		return fmt.Errorf("value %s violates unique constraint: %w", value, ErrUniqueConstraintViolated)
	}
	return nil
}

func (uc *uniqueConstraints) removeUniqueConstraint(value string) {
	delete(uc.m, value)
}
