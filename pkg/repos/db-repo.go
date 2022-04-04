package repos

import "context"

type Entity interface {
	GetId() ID
	SetId(id ID) Entity
}

type Opts map[string]interface{}
type SortOpts map[string]int32
type Filter map[string]interface{}
type Query struct {
	filter Filter
	sort   map[string]interface{}
}

type ID string

type PaginatedRecord[T Entity] struct {
	results    []T
	totalCount int64
}

type DbRepo[T Entity] interface {
	NewId() ID
	Find(ctx context.Context, query Query) ([]T, error)
	FindPaginated(ctx context.Context, query Query, page int64, size int64, opts ...Opts) (PaginatedRecord[T], error)
	FindById(ctx context.Context, id ID) (T, error)
	Create(ctx context.Context, data T) (T, error)
	UpdateById(ctx context.Context, id ID, updatedData T) (T, error)
	DeleteById(ctx context.Context, id ID) error
	//Delete(ctx context.Context, query Query) ([]ID, error)
}
