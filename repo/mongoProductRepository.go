package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	collection *mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewMongoProductRepository(collection *mongo.Collection, ctx2 context.Context, cancelFunc context.CancelFunc) *MongoProductRepository {
	return &MongoProductRepository{
		collection: collection,
		ctx:        ctx2,
		cancel:     cancelFunc,
	}
}

func (m *MongoProductRepository) InsertProduct(product *entity.Product) (*mongo.InsertOneResult, error) {
	result, err := m.collection.InsertOne(ctx, product)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *MongoProductRepository) FindProductById(id string) (*entity.Product, error) {
	var product entity.Product
	var idInt int
	_, err := fmt.Sscan(id, &idInt) // Convert string ID to int
	if err != nil {
		return nil, err
	}
	err = m.collection.FindOne(m.ctx, bson.M{"id": idInt}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m *MongoProductRepository) UpdateProduct(id string, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error) {
	var idInt int
	_, err := fmt.Sscan(id, &idInt) // Convert string ID to int
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": idInt}
	update := bson.D{{"$set", productUpdate}}
	updateResult, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (m *MongoProductRepository) FindAllProducts() (*[]entity.Product, error) {
	var products []entity.Product
	cursor, err := m.collection.Find(m.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(m.ctx, &products); err != nil {
		return nil, err
	}
	return &products, nil
}

func (m *MongoProductRepository) DeleteProductById(id string) (*mongo.DeleteResult, error) {
	var idInt int
	_, err := fmt.Sscan(id, &idInt) // Convert string ID to int
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": idInt}
	result, err := m.collection.DeleteOne(m.ctx, filter)
	if err != nil {
		return result, err
	}
	return result, nil
}
