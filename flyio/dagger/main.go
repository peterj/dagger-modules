// Deploy apps to Fly.io
//
// Basic Dagger module for deploying apps to Fly.io

package main

import (
	"context"
)

type Flyio struct{}

// Deploy deploys an app from the src folder to Fly.io
func (m *Flyio) Deploy(ctx context.Context,
	// +required
	src *Directory,
	// +required
	token *Secret) (string, error) {
	return m.FlyContainer(ctx, token).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"/root/.fly/bin/flyctl", "deploy"}).
		Stdout(ctx)
}

// FlyContainer creates a container with the flyctl CLI installed
func (m *Flyio) FlyContainer(ctx context.Context, token *Secret) *Container {
	return dag.Container().
		From("alpine:3.20.0").
		WithExec([]string{"apk", "add", "curl"}).
		WithExec([]string{"curl", "-LO", "https://fly.io/install.sh"}).
		WithExec([]string{"sh", "install.sh"}).
		WithSecretVariable("FLY_API_TOKEN", token)
}
