package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/fx"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"kloudlite.io/pkg/errors"
	fn "kloudlite.io/pkg/functions"
	t "kloudlite.io/pkg/types"
)

type dbRepo[T Entity] struct {
	db             *mongo.Database
	collectionName string
	shortName      string
}

var re = regexp.MustCompile(`(\W|_)+/g`)

func toMap(v any) (map[string]any, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func bsonToStruct[T any](r *mongo.SingleResult) (T, error) {
	var m map[string]any
	var result T
	if err := r.Decode(&m); err != nil {
		return result, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (repo *dbRepo[T]) NewId() ID {
	id, e := fn.CleanerNanoid(28)
	if e != nil {
		panic(fmt.Errorf("could not get cleanerNanoid()"))
	}
	return ID(fmt.Sprintf("%s-%s", repo.shortName, strings.ToLower(id)))
}

func (repo *dbRepo[T]) Find(ctx context.Context, query Query) ([]T, error) {
	curr, err := repo.db.Collection(repo.collectionName).Find(
		ctx, query.Filter, &options.FindOptions{
			Sort: query.Sort,
		},
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return make([]T, 0), nil
		}
		return nil, err
	}

	var m []map[string]any

	err = curr.All(ctx, &m)
	if err != nil {
		return nil, err
	}

	results := make([]T, len(m))
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &results); err != nil {
		return nil, err
	}

	return results, err
}

func (repo *dbRepo[T]) findOne(ctx context.Context, filter Filter) (T, error) {
	one := repo.db.Collection(repo.collectionName).FindOne(ctx, filter)
	res := fn.NewTypeFromPointer[T]()
	var m map[string]any
	err := one.Decode(&m)
	if err != nil {
		fmt.Println("(debug/info) ERR: ", err)
		return res, err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return res, err
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (repo *dbRepo[T]) FindOne(ctx context.Context, filter Filter) (T, error) {
	one := repo.db.Collection(repo.collectionName).FindOne(ctx, filter)
	t := make([]T, 1)
	var m map[string]any
	err := one.Decode(&m)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return t[0], nil
		}
		return t[0], err
	}

	b, err := json.Marshal(m)
	if err != nil {
		return t[0], err
	}
	if err := json.Unmarshal(b, &t[0]); err != nil {
		return t[0], err
	}
	return t[0], nil
}

func (repo *dbRepo[T]) FindPaginated(ctx context.Context, filter Filter, pagination *t.CursorPagination) (*PaginatedRecord[T], error) {
	if pagination == nil {
		pagination = &t.CursorPagination{First: 10, After: ""}
	}

	queryFilter := Filter{}

	for k, v := range filter {
		queryFilter[k] = v
	}

	queryFilter["creationTime"] = bson.M{"$gte": pagination.After}

	results := make([]T, 0, pagination.First+1)
	curr, err := repo.db.Collection(repo.collectionName).Find(
		ctx, filter, &options.FindOptions{
			Limit: fn.New(pagination.First + 1),
			Sort:  map[string]any{"creationTime": 1},
		},
	)
	if err != nil {
		return nil, err
	}

	if err = curr.All(ctx, &results); err != nil {
		return nil, err
	}

	total, err := repo.db.Collection(repo.collectionName).CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	pageInfo := PageInfo{}

	if len(results) > 0 {
		pageInfo.StartCursor = results[0].GetCreationTime().Format(time.RFC3339)
		pageInfo.EndCursor = results[len(results)-1].GetCreationTime().Format(time.RFC3339)
		pageInfo.HasNextPage = len(results) > int(pagination.First)
	}

	results = append(results[:len(results)-1])

	edges := make([]RecordEdge[T], len(results))
	for i := range results {
		edges[i] = RecordEdge[T]{Node: results[i], Cursor: results[i].GetCreationTime().Format(time.RFC3339)}
	}

	return &PaginatedRecord[T]{
		Edges:      edges,
		PageInfo:   pageInfo,
		TotalCount: total,
	}, nil
}

// func (repo *dbRepo[T]) FindPaginated(ctx context.Context, query Query, page int64, size int64, opts ...Opts) (PaginatedRecord[T], error) {
// 	results := make([]T, 0)
// 	var offset int64 = (page - 1) * size
// 	curr, e := repo.db.Collection(repo.collectionName).Find(
// 		ctx, query.Filter, &options.FindOptions{
// 			Limit: &size,
// 			Skip:  &offset,
// 			Sort:  query.Sort,
// 		},
// 	)
// 	e = curr.All(ctx, results)
//
// 	total, e := repo.db.Collection(repo.collectionName).CountDocuments(ctx, query.Filter)
//
// 	return PaginatedRecord[T]{
// 		Results:    results,
// 		TotalCount: total,
// 	}, e
// }

func (repo *dbRepo[T]) FindById(ctx context.Context, id ID) (T, error) {
	var result T
	r := repo.db.Collection(repo.collectionName).FindOne(ctx, &Filter{"id": id})

	err := r.Decode(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (repo *dbRepo[T]) withId(data T) {
	if data.GetId() != "" {
		return
	}
	data.SetId(repo.NewId())
}

func (repo *dbRepo[T]) withCreationTime(data T) {
	if !data.GetCreationTime().IsZero() {
		return
	}
	data.SetCreationTime(time.Now())
	data.SetUpdateTime(time.Now())
}

func (repo *dbRepo[T]) withUpdateTime(data T) {
	data.SetUpdateTime(time.Now())
}

func (repo *dbRepo[T]) Create(ctx context.Context, data T) (T, error) {
	repo.withId(data)
	repo.withCreationTime(data)

	b, err := json.Marshal(data)
	if err != nil {
		var x T
		return x, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		var x T
		return x, err
	}

	r, e := repo.db.Collection(repo.collectionName).InsertOne(ctx, m)
	if e != nil {
		var x T
		return x, e
	}
	var result T
	r2 := repo.db.Collection(repo.collectionName).FindOne(ctx, Filter{"_id": r.InsertedID})

	var resultM map[string]any
	e = r2.Decode(&resultM)

	m2, err := json.Marshal(resultM)
	if err := json.Unmarshal(m2, &result); err != nil {
		var x T
		return x, err
	}

	return result, e
}

func (repo *dbRepo[T]) Exists(ctx context.Context, filter Filter) (bool, error) {
	one := repo.db.Collection(repo.collectionName).FindOne(ctx, filter)
	res := fn.NewTypeFromPointer[T]()
	err := one.Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *dbRepo[T]) UpdateMany(ctx context.Context, filter Filter, updatedData map[string]any) error {
	updatedData["updateTime"] = time.Now()

	_, err := repo.db.Collection(repo.collectionName).UpdateMany(
		ctx,
		filter,
		bson.M{"$set": updatedData},
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dbRepo[T]) UpdateOne(ctx context.Context, filter Filter, updatedData T, opts ...UpdateOpts) (T, error) {
	after := options.After
	updateOpts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	if opt := fn.ParseOnlyOption[UpdateOpts](opts); opt != nil {
		updateOpts.Upsert = &opt.Upsert
	}

	m, err := toMap(updatedData)
	if err != nil {
		var x T
		return x, err
	}

	repo.withUpdateTime(updatedData)
	r := repo.db.Collection(repo.collectionName).FindOneAndUpdate(
		ctx,
		filter,
		bson.M{"$set": m},
		updateOpts,
	)

	return bsonToStruct[T](r)
}

func (repo *dbRepo[T]) UpdateById(ctx context.Context, id ID, updatedData T, opts ...UpdateOpts) (T, error) {
	var result T
	after := options.After
	updateOpts := &options.FindOneAndUpdateOptions{ReturnDocument: &after}

	if opt := fn.ParseOnlyOption[UpdateOpts](opts); opt != nil {
		updateOpts.Upsert = &opt.Upsert
	}
	repo.withUpdateTime(updatedData)

	m, err := toMap(updatedData)
	if err != nil {
		return result, err
	}

	r := repo.db.Collection(repo.collectionName).FindOneAndUpdate(
		ctx,
		&Filter{"id": id},
		bson.M{"$set": m},
		updateOpts,
	)
	return bsonToStruct[T](r)
}

func (repo *dbRepo[T]) Upsert(ctx context.Context, filter Filter, data T) (T, error) {
	id := func() ID {
		if data.GetId() != "" {
			return data.GetId()
		}
		if t, err := repo.findOne(ctx, filter); err == nil {
			return t.GetId()
		}
		return repo.NewId()
	}()

	data.SetId(id)
	repo.withCreationTime(data)
	repo.withUpdateTime(data)

	return repo.UpdateById(
		ctx, id, data, UpdateOpts{
			Upsert: true,
		},
	)
}

func (repo *dbRepo[T]) DeleteById(ctx context.Context, id ID) error {
	var result T
	r := repo.db.Collection(repo.collectionName).FindOneAndDelete(ctx, &Filter{"id": id})
	e := r.Decode(&result)
	return e
}

func (repo *dbRepo[T]) DeleteOne(ctx context.Context, filter Filter) error {
	var result T
	r := repo.db.Collection(repo.collectionName).FindOneAndDelete(ctx, filter)
	err := r.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
	}
	return err
}

func (repo *dbRepo[T]) DeleteMany(ctx context.Context, filter Filter) error {
	_, err := repo.db.Collection(repo.collectionName).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func buildIndexName(curr string, indexKey string, indexValue int) string {
	if curr == "" {
		return curr + fmt.Sprintf("%s_%d", indexKey, indexValue)
	}
	return curr + "_" + fmt.Sprintf("%s_%d", indexKey, indexValue)
}

func (repo *dbRepo[T]) IndexFields(ctx context.Context, indices []IndexField) error {
	if len(indices) == 0 {
		return nil
	}

	indices = append(indices, IndexField{
		Field:  []IndexKey{{Key: "creationTime", Value: IndexAsc}},
		Unique: false,
	})
	// var models []mongo.IndexModel
	for _, f := range indices {
		b := bson.D{}

		// This method to create indexes ensure temporary safety, while index modification,
		// which we need to do when index names clash
		// READ MORE @ https://www.mongodb.com/docs/manual/tutorial/manage-indexes/#modify-an-index
		indexName := ""
		for _, field := range f.Field {
			switch field.Value {
			case IndexAsc:
				b = append(b, bson.E{Key: field.Key, Value: 1})
				indexName = buildIndexName(indexName, field.Key, 1)
			case IndexDesc:
				b = append(b, bson.E{Key: field.Key, Value: -1})
				indexName = buildIndexName(indexName, field.Key, -1)
			}
		}

		indexModel := mongo.IndexModel{Keys: b, Options: &options.IndexOptions{Unique: &f.Unique, Name: &indexName}}

		_, err := repo.db.Collection(repo.collectionName).Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			dummyKey := fn.CleanerNanoidOrDie(10)
			b2 := append(b, bson.E{Key: dummyKey, Value: 1})
			dummyIdxName := buildIndexName(indexName, dummyKey, 1)
			_, err := repo.db.Collection(repo.collectionName).Indexes().CreateOne(ctx, mongo.IndexModel{Keys: b2, Options: &options.IndexOptions{Unique: &f.Unique, Name: &dummyIdxName}})
			if err != nil {
				return err
			}
			_, err = repo.db.Collection(repo.collectionName).Indexes().DropOne(ctx, indexName)
			if err != nil {
				return err
			}
			if _, err := repo.db.Collection(repo.collectionName).Indexes().CreateOne(ctx, indexModel); err != nil {
				return err
			}
			if _, err := repo.db.Collection(repo.collectionName).Indexes().DropOne(ctx, dummyIdxName); err != nil {
				return err
			}
		}

		// models = append(
		// 	models, mongo.IndexModel{
		// 		Keys: b,
		// 		Options: &options.IndexOptions{
		// 			Unique: &f.Unique,
		// 		},
		// 	},
		// )
	}

	// for i := range models {
	// 	_, err := repo.db.Collection(repo.collectionName).Indexes().CreateOne(ctx, models[i])
	// 	if err != nil{
	// 		repo.db.Collection(repo.collectionName).Indexes().CreateOne(ctx, models[i]., opts ...*options.CreateIndexesOptions)
	// 	}
	// 	//body
	// }

	// _, err := repo.db.Collection(repo.collectionName).Indexes().CreateMany(ctx, models)
	// return err
	return nil
}

func (repo *dbRepo[T]) SilentUpsert(ctx context.Context, filter Filter, data T) (T, error) {
	id := func() ID {
		if data.GetId() != "" {
			return data.GetId()
		}
		if t, err := repo.findOne(ctx, filter); err == nil {
			return t.GetId()
		}
		return repo.NewId()
	}()
	data.SetId(id)
	if data.GetCreationTime().IsZero() {
		repo.withCreationTime(data)
	}
	return repo.UpdateById(
		ctx, id, data, UpdateOpts{
			Upsert: true,
		},
	)
}
func (repo *dbRepo[T]) SilentUpdateMany(ctx context.Context, filter Filter, updatedData map[string]any) error {
	_, err := repo.db.Collection(repo.collectionName).UpdateMany(
		ctx,
		filter,
		bson.M{"$set": updatedData},
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *dbRepo[T]) SilentUpdateById(ctx context.Context, id ID, updatedData T, opts ...UpdateOpts) (T, error) {
	after := options.After
	updateOpts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	if opt := fn.ParseOnlyOption[UpdateOpts](opts); opt != nil {
		updateOpts.Upsert = &opt.Upsert
	}

	m, err := toMap(updatedData)
	if err != nil {
		var x T
		return x, err
	}

	r := repo.db.Collection(repo.collectionName).FindOneAndUpdate(
		ctx,
		&Filter{"id": id},
		bson.M{"$set": m},
		updateOpts,
	)
	return bsonToStruct[T](r)
}

func (repo *dbRepo[T]) ErrAlreadyExists(err error) bool {
	return mongo.IsDuplicateKeyError(err)
}

type MongoRepoOptions struct {
	IndexFields []string
}

func NewMongoRepo[T Entity](
	db *mongo.Database,
	collectionName string,
	shortName string,
) DbRepo[T] {
	return &dbRepo[T]{
		db,
		collectionName,
		shortName,
	}
}

func NewFxMongoRepo[T Entity](collectionName, shortName string, indexFields []IndexField) fx.Option {
	return fx.Module(
		"repo",
		fx.Provide(
			func(db *mongo.Database) DbRepo[T] {
				return NewMongoRepo[T](
					db,
					collectionName,
					shortName,
				)
			},
		),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, repo DbRepo[T]) {
				lifecycle.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							err := repo.IndexFields(ctx, indexFields)
							if err != nil {
								return errors.NewEf(err, "could not create indexes on DB for repo %T", repo)
							}
							return nil
						},
					},
				)
			},
		),
	)
}
