package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/kloudlite/api/apps/comms/internal/app/graph/generated"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/pkg/repos"
)

// AccountName is the resolver for the accountName field.
func (r *subscriptionResolver) AccountName(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: AccountName - accountName"))
}

// CreatedBy is the resolver for the createdBy field.
func (r *subscriptionResolver) CreatedBy(ctx context.Context) (<-chan *common.CreatedOrUpdatedBy, error) {
	panic(fmt.Errorf("not implemented: CreatedBy - createdBy"))
}

// CreationTime is the resolver for the creationTime field.
func (r *subscriptionResolver) CreationTime(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: CreationTime - creationTime"))
}

// Enabled is the resolver for the enabled field.
func (r *subscriptionResolver) Enabled(ctx context.Context) (<-chan bool, error) {
	panic(fmt.Errorf("not implemented: Enabled - enabled"))
}

// ID is the resolver for the id field.
func (r *subscriptionResolver) ID(ctx context.Context) (<-chan repos.ID, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// LastUpdatedBy is the resolver for the lastUpdatedBy field.
func (r *subscriptionResolver) LastUpdatedBy(ctx context.Context) (<-chan *common.CreatedOrUpdatedBy, error) {
	panic(fmt.Errorf("not implemented: LastUpdatedBy - lastUpdatedBy"))
}

// MailAddress is the resolver for the mailAddress field.
func (r *subscriptionResolver) MailAddress(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: MailAddress - mailAddress"))
}

// MarkedForDeletion is the resolver for the markedForDeletion field.
func (r *subscriptionResolver) MarkedForDeletion(ctx context.Context) (<-chan *bool, error) {
	panic(fmt.Errorf("not implemented: MarkedForDeletion - markedForDeletion"))
}

// RecordVersion is the resolver for the recordVersion field.
func (r *subscriptionResolver) RecordVersion(ctx context.Context) (<-chan int, error) {
	panic(fmt.Errorf("not implemented: RecordVersion - recordVersion"))
}

// UpdateTime is the resolver for the updateTime field.
func (r *subscriptionResolver) UpdateTime(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: UpdateTime - updateTime"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
