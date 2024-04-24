package domain

import (
	"fmt"

	"github.com/kloudlite/api/pkg/errors"
)

type ErrGRPCCall struct {
	Err error
}

func (e ErrGRPCCall) Error() string {
	return fmt.Sprintf("grpc call failed with error: %v", errors.NewE(e.Err))
}
