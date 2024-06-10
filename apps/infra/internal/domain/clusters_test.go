package domain_test

import (
	"testing"

	"github.com/kloudlite/api/apps/infra/internal/domain"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	"github.com/kloudlite/api/apps/infra/internal/env"

	"github.com/kloudlite/api/pkg/logging"
)

func TestOnClusterDeleteMessage(t *testing.T) {
	type test struct {
		name string
		fn   func(d domain.ClusterDomain, t *testing.T)
	}

	logerr := func(t *testing.T, gotErr error, wantErr error) {
		t.Errorf("CreateCluster() errored, got error = %v, want error = %v", gotErr, wantErr)
	}

	tests := []test{
		{
			name: "1. cluster is not found",
			fn: func(d domain.ClusterDomain, t *testing.T) {
				if err := d.OnClusterDeleteMessage(InfraContext{}, entities.Cluster{}); err != nil {
					logerr(t, err, nil)
					t.Error(err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &domain{
				logger: logging.EmptyLogger,
				env:    &env.Env{},
			}
			tt.fn(d, t)
		})
	}
}
