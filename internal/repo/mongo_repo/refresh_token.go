package mongo_repo

import (
	"context"
	"fmt"
	"sso/internal/def"
	"sso/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefreshToken struct {
	collection *mongo.Collection
}

func NewRefreshToken(db *mongo.Database) *RefreshToken {
	return &RefreshToken{
		collection: db.Collection(def.TableRefreshTokens.String()),
	}
}

func (r *RefreshToken) DeleteByUser(ctx context.Context, user *model.User) error {
	const op = "mongo_repo.RefreshToken.DeleteByUser"

	filter := bson.M{"user_id": user.ID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("%s: %w", op, def.ErrNotFound)
	}

	return nil
}

func (r *RefreshToken) Create(ctx context.Context, refreshToken *model.RefreshToken) error {
	const op = "mongo_repo.RefreshToken.Create"

	refreshToken.ID = primitive.NewObjectID()
	refreshToken.CreatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, refreshToken)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
