package main

import (
	"context"
	"fmt"
	"log"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	src := client.Host().Directory(".")

	container := client.Container().From("registry.access.redhat.com/ubi9/ubi-micro").
		WithMountedDirectory("/mnt", src).
		WithEntrypoint([]string{"sh", "-c"}).
		WithExec([]string{"ls -laZ /mnt"})

	out, err := container.Stdout(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
