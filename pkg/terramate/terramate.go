package terramate

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

// Terramate represents the terramate binary and working directory
type Terramate struct {
	execPath   string
	workingDir string
	stdout     io.Writer
	stderr     io.Writer
}

// Option represents configuration options for Terramate
type Option func(*Terramate)

// WithWorkingDir sets the working directory for terramate commands
func WithWorkingDir(workingDir string) Option {
	return func(tm *Terramate) {
		tm.workingDir = workingDir
	}
}

// WithStdout sets the stdout writer
func WithStdout(w io.Writer) Option {
	return func(tm *Terramate) {
		tm.stdout = w
	}
}

// WithStderr sets the stderr writer
func WithStderr(w io.Writer) Option {
	return func(tm *Terramate) {
		tm.stderr = w
	}
}

// NewTerramate creates a new Terramate instance
func NewTerramate(execPath string, opts ...Option) (*Terramate, error) {
	if execPath == "" {
		return nil, fmt.Errorf("terramate exec path cannot be empty")
	}

	tm := &Terramate{
		execPath: execPath,
	}

	for _, opt := range opts {
		opt(tm)
	}

	return tm, nil
}

// Run executes a terramate command with the given arguments
func (tm *Terramate) Run(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, tm.execPath, args...)

	if tm.workingDir != "" {
		cmd.Dir = tm.workingDir
	}

	if tm.stdout != nil {
		cmd.Stdout = tm.stdout
	}

	if tm.stderr != nil {
		cmd.Stderr = tm.stderr
	}

	return cmd.Run()
}

// Version returns the version of the terramate binary
func (tm *Terramate) Version(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, tm.execPath, "version")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running terramate version: %w", err)
	}
	return string(output), nil
}
