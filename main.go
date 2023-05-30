package main

import (
	"fmt"

	"dagger.io/dagger"
)

func main() {
	dagger.ServeCommands(run)
}

func run(ctx dagger.Context) (string, error) {
	client := ctx.Client()

	src := client.Host().Directory(".")

	container := client.Container().From("registry.access.redhat.com/ubi9/ubi-micro").
		WithMountedDirectory("/mnt", src).
		WithEntrypoint([]string{"sh", "-c"}).
		WithExec([]string{"ls -laZ /mnt"})

	out, err := container.Stdout(ctx)
	if err != nil {
		return "", err
	}

	fmt.Println(out)

	return out, nil
}
