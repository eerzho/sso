package mongo_repo

import (
	"context"
	"errors"
	"fmt"
	"sso/internal/def"
	"sso/internal/dto"
	"sso/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Permission struct {
	maxListCount int
	collection   *mongo.Collection
}

func NewPermission(db *mongo.Database) *Permission {
	return &Permission{
		maxListCount: 200,
		collection:   db.Collection(def.TablePermissions.String()),
	}
}

func (r *Permission) List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.Permission, *dto.Pagination, error) {
	const op = "mongo_repo.Permission.List"

	if count > r.maxListCount {
		count = r.maxListCount
	}

	filter := bson.M{}
	for key, value := range filters {
		if key == "name" {
			filter[key] = bson.M{"$regex": value, "$options": "i"}
		} else if key == "slug" {
			filter[key] = value
		}
	}

	sort := bson.D{}
	for key, value := range sorts {
		if key == "created_at" ||
			key == "updated_at" ||
			key == "name" ||
			key == "slug" {
			if value == "asc" {
				sort = append(sort, bson.E{Key: key, Value: 1})
			} else if value == "desc" {
				sort = append(sort, bson.E{Key: key, Value: -1})
			}
		}
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * count))
	findOptions.SetLimit(int64(count))
	findOptions.SetSort(sort)

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	var permissions []model.Permission
	err = cursor.All(ctx, &permissions)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	pagination := dto.Pagination{
		Page:  page,
		Count: count,
		Total: int(total),
	}

	return permissions, &pagination, nil
}

func (r *Permission) Create(ctx context.Context, permission *model.Permission) error {
	const op = "mongo_repo.Permission.Create"

	permission.ID = primitive.NewObjectID()
	permission.CreatedAt = time.Now()
	permission.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, &permission)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Permission) IsExistsSlug(ctx context.Context, slug string) (bool, error) {
	const op = "mongo_repo.Permission.IsExistsSlug"

	filter := bson.M{"slug": slug}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return count > 0, nil
}

func (r *Permission) GetByID(ctx context.Context, id string) (*model.Permission, error) {
	const op = "mongo_repo.Permission.GetByID"

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{"_id": idObj}
	var permission model.Permission

	err = r.collection.FindOne(ctx, filter).Decode(&permission)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s: %w", op, def.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &permission, nil
}

func (r *Permission) Delete(ctx context.Context, id string) error {
	const op = "mongo_repo.Permission.Delete"

	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{"_id": idObj}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%s: %w", op, def.ErrNotFound)
	}

	return nil
}
