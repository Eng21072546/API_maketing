package repo

import (
	"context"
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

func (m *MongoProductRepository) FindProductById(id int) (*entity.Product, error) {
	var product entity.Product
	err = m.collection.FindOne(m.ctx, bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m *MongoProductRepository) UpdateProduct(id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": id}
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

func (m *MongoProductRepository) DeleteProductById(id int) (*mongo.DeleteResult, error) {
	filter := bson.M{"id": id}
	result, err := m.collection.DeleteOne(m.ctx, filter)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *MongoProductRepository) DecreaseStock(order entity.Order) error {
	for _, productOrder := range order.ProductList {
		product, _ := m.FindProductById(productOrder.ProductID)
		currenStock := product.Stock
		newStock := currenStock - productOrder.Quantity
		err := m.UpdateStock(productOrder.ProductID, newStock)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MongoProductRepository) UpdateStock(productID int, quantity int) error {
	collection := client.Database("market").Collection("product")
	filter := bson.M{"id": productID}
	update := bson.M{"$set": bson.M{"stock": quantity}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
