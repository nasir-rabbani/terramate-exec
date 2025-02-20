package terramate

import (
	"fmt"
)

// ErrNoTerramate represents an error when terramate binary cannot be found
type ErrNoTerramate struct {
	err error
}

func (e *ErrNoTerramate) Error() string {
	return fmt.Sprintf("terramate not found: %s", e.err)
}

// ErrVersionMismatch represents an error when terramate version doesn't match requirements
type ErrVersionMismatch struct {
	Constraint string
	Version    string
}

func (e *ErrVersionMismatch) Error() string {
	return fmt.Sprintf("terramate version %s doesn't match constraint %s", e.Version, e.Constraint)
}
