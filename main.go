package main

import (
	"dagger.io/dagger"
)

func main() {
	dagger.ServeCommands(Test)
}

type TestTarget struct{}

func Test(ctx dagger.Context) (TestTarget, error) {
	return TestTarget{}, nil
}

func (TestTarget) One(ctx dagger.Context, arg string) (string, error) {
	return ctx.Client().
		Container().
		From("registry.access.redhat.com/ubi9/ubi-micro").
		WithExec([]string{"echo", arg}).
		Stdout(ctx)
}

func (TestTarget) Two(ctx dagger.Context, aRg string) (string, error) {
	return ctx.Client().
		Container().
		From("registry.access.redhat.com/ubi9/ubi-micro").
		WithExec([]string{"echo", aRg}).
		Stdout(ctx)
}

func (TestTarget) Three(ctx dagger.Context, aRG string) (string, error) {
	return ctx.Client().
		Container().
		From("registry.access.redhat.com/ubi9/ubi-micro").
		WithExec([]string{"echo", aRG}).
		Stdout(ctx)
}
