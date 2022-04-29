package harbor

import (
	"context"
)

type Harbor interface {
	CreateProject(ctx context.Context, name string) error
	CreateUserAccount(ctx context.Context, projectName string, name string) error
	DeleteProject(ctx context.Context, name string) error
}
