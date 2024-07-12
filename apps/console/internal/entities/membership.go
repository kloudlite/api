package entities

import (
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/pkg/repos"
)

type EnvMembership struct {
	EnvName string    `json:"envName"`
	UserId  repos.ID  `json:"userId"`
	Role    iamT.Role `json:"role"`
}
