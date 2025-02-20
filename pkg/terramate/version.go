package terramate

import (
	"context"
	"fmt"
	"strings"

	version "github.com/hashicorp/go-version"
)

// CheckVersion checks if the terramate version satisfies the given constraint
func (tm *Terramate) CheckVersion(ctx context.Context, constraint string) error {
	v, err := tm.Version(ctx)
	if err != nil {
		return err
	}

	cleanV := strings.TrimSpace(v)
	actual, err := version.NewVersion(cleanV)
	if err != nil {
		return fmt.Errorf("error parsing terramate version: %w", err)
	}

	constraints, err := version.NewConstraint(constraint)
	if err != nil {
		return fmt.Errorf("error parsing version constraint: %w", err)
	}

	if !constraints.Check(actual) {
		return &ErrVersionMismatch{
			Constraint: constraint,
			Version:    cleanV,
		}
	}

	return nil
}
