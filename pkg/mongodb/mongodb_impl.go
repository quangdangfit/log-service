package mongodb

import (
	"context"
	"log"
	"reflect"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/quangdangfit/log-service/pkg/paging"
)

type dbImpl struct {
	db *mongo.Database
}

/**========================================================================
 *                           INTERFACE IMPLEMENTATION
 *========================================================================**/

func (d *dbImpl) GetCollection(collection string) *mongo.Collection {
	return d.db.Collection(collection)
}

func (d *dbImpl) Delete(ctx context.Context, collection string, filter interface{}, opts *options.DeleteOptions) error {
	result, err := d.db.Collection(collection).DeleteMany(ctx, filter, opts)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (d *dbImpl) Close() {
	err := d.db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatalf("failed ro close database connection, error %v", err)
	}
}

// Insert - insert document to database
//
// is type of document is []interface{}, insert many
func (d *dbImpl) Insert(ctx context.Context, collection string, document interface{}) error {
	// check insert many
	s := reflect.ValueOf(document)
	if s.Kind() == reflect.Slice {
		docs := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			docs[i] = s.Index(i).Interface()
		}
		_, err := d.db.Collection(collection).InsertMany(ctx, docs)
		return err
	}

	_, err := d.db.Collection(collection).InsertOne(ctx, document)
	return err
}

// UpdateMany - default upsert is false
//
// return errors.NotFound.New() if filter is not match document
func (d *dbImpl) UpdateMany(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error {
	uFn := d.db.Collection(collection).UpdateMany
	return d.update(ctx, filter, update, uFn, options)
}

// UpdateOne - default upsert is false
//
// return errors.NotFound.New() if filter is not match document
func (d *dbImpl) UpdateOne(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error {
	uFn := d.db.Collection(collection).UpdateOne
	return d.update(ctx, filter, update, uFn, options)
}

// UpdateOneRaw - default upsert is false
func (d *dbImpl) UpdateOneRaw(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error {
	uFn := d.db.Collection(collection).UpdateOne
	return d.updateRaw(ctx, filter, update, uFn, options)
}

// UpdateManyRaw - default upsert is false
func (d *dbImpl) UpdateManyRaw(ctx context.Context, collection string, filter, update interface{}, options *options.UpdateOptions) error {
	uFn := d.db.Collection(collection).UpdateMany
	return d.updateRaw(ctx, filter, update, uFn, options)
}

func (d *dbImpl) FindById(ctx context.Context, collection string, id string, result interface{}) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}
	err = d.db.
		Collection(collection).
		FindOne(ctx, filter).
		Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (d *dbImpl) FindOne(ctx context.Context, collection string, result interface{}, opts ...Option) error {
	opt := getOption(opts...)
	var bOpt *options.FindOneOptions
	if opt.sorter != nil {
		bOpt = options.FindOne().SetSort(opt.sorter)
	}

	err := d.db.
		Collection(collection).
		FindOne(ctx, opt.filter, bOpt).
		Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (d *dbImpl) Find(ctx context.Context, collection string, result interface{}, opts ...Option) error {
	opt := getOption(opts...)

	err := d.find(ctx, collection, result, opt)
	if err != nil {
		return err
	}

	return nil
}

func (d *dbImpl) FindOneAndUpdate(ctx context.Context, collection string, filter, update interface{}, opts *options.FindOneAndUpdateOptions, result interface{}) error {
	err := d.db.Collection(collection).FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, opts).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (d *dbImpl) EnsureIndexes(ctx context.Context, collection string, indexes []mongo.IndexModel) error {
	err := d.ensureIndexes(ctx, d.db, collection, indexes)
	if err != nil {
		return err
	}
	return nil
}

func (d *dbImpl) WithTransaction(callback func(sc mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	var (
		ses mongo.Session
		ctx context.Context
		err error
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if ses, err = d.db.Client().StartSession(); err != nil {
		return nil, err
	}
	defer ses.EndSession(ctx)

	return ses.WithTransaction(ctx, callback)
}

func (d *dbImpl) Count(ctx context.Context, collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	rs, err := d.db.Collection(collection).CountDocuments(ctx, filter, opts...)
	if err != nil {
		return 0, err
	}

	return rs, nil
}

func (d *dbImpl) BulkWriteRaw(ctx context.Context, collection string, operations []mongo.WriteModel, opts ...*options.BulkWriteOptions) error {
	_, err := d.db.Collection(collection).BulkWrite(ctx, operations, opts...)
	return err
}

/**========================================================================
 *                           PRIVATE METHODS
 *========================================================================**/

func (d *dbImpl) update(ctx context.Context, filter, update interface{}, updateFn updateFn, options *options.UpdateOptions) error {
	_, err := updateFn(ctx, filter, bson.M{"$set": update}, options)
	if err != nil {
		return err
	}

	// ignore for case upsert
	//if u.MatchedCount == 0 {
	//	return mongo.ErrNoDocuments
	//}

	return nil
}

func (d *dbImpl) updateRaw(ctx context.Context, filter, update interface{}, updateFn updateFn, options *options.UpdateOptions) error {
	_, err := updateFn(ctx, filter, update, options)
	if err != nil {
		return err
	}

	return nil
}

func (d *dbImpl) find(ctx context.Context, collection string,
	result interface{}, opt option) error {
	bOpt := new(options.FindOptions)
	if opt.sorter != nil {
		bOpt.SetSort(opt.sorter)
	}
	cls := d.db.Collection(collection)

	if opt.paging != nil {
		total, err := cls.CountDocuments(ctx, opt.filter)
		if err != nil {
			return err
		}

		opt.paging = paging.New(opt.page, opt.limit, total)
		bOpt.SetSkip(opt.paging.Skip)
		bOpt.SetLimit(int64(opt.paging.Limit))
	}

	cur, err := cls.Find(ctx, opt.filter, bOpt)
	if err != nil {
		return err
	}

	return cur.All(ctx, result)
}

// dropIndexes :
func (d *dbImpl) dropIndexes(dropIndexes []string, collection *mongo.Collection) error {
	for _, indexStr := range dropIndexes {
		opts := options.DropIndexes().SetMaxTime(10 * time.Second)
		_, err := collection.Indexes().DropOne(context.Background(), indexStr, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

// createIndexes :
func (d *dbImpl) createIndexes(createIndexes []mongo.IndexModel, collection *mongo.Collection) error {
	for _, index := range createIndexes {
		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		_, err := collection.Indexes().CreateOne(context.Background(), index, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *dbImpl) ensureIndexes(_ context.Context, database *mongo.Database, collection string, indexes []mongo.IndexModel) error {
	type MongoIndex struct {
		Name string
		Keys interface{}
	}

	var (
		dropIndexes = make([]string, 0)
	)

	c := database.Collection(collection)
	duration := 10 * time.Second
	batchSize := int32(100)

	cur, err := c.Indexes().List(context.Background(), &options.ListIndexesOptions{
		BatchSize: &batchSize,
		MaxTime:   &duration,
	})
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	sort.Slice(indexes, func(i, j int) bool {
		return *indexes[i].Options.Name <= *indexes[j].Options.Name
	})

	for cur.Next(context.Background()) {
		index := MongoIndex{}
		cur.Decode(&index)

		if index.Name == "_id_" {
			continue
		}

		isDrop := true
		for _, v := range indexes {
			if *v.Options.Name == index.Name {
				isDrop = false
			}
		}

		// Drop all index is not found on which defined
		if isDrop {
			dropIndexes = append(dropIndexes, index.Name)
		}
	}
	err = d.dropIndexes(dropIndexes, c)
	if err != nil {
		return err
	}

	err = d.createIndexes(indexes, c)
	if err != nil {
		return err
	}

	return nil
}

func getOption(opts ...Option) option {
	opt := option{
		sorter: bson.M{_idField: SortDescending},
		filter: bson.D{},
	}

	for _, o := range opts {
		o.apply(&opt)
	}

	return opt
}
