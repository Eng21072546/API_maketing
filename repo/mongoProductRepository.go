package repo

import (
	"context"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepository struct {
	client   *mongo.Client
	ctxMongo context.Context
}

func NewMongoProductRepository(client *mongo.Client, ctx context.Context) *MongoProductRepository {
	return &MongoProductRepository{client, ctx}
}

func (m *MongoProductRepository) InsertProduct(ctx context.Context, product *entity.Product) (*mongo.InsertOneResult, error) {
	result, err := m.client.Database("market").Collection("product").InsertOne(ctx, product)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *MongoProductRepository) FindProductById(ctx context.Context, id int) (*entity.Product, error) {
	var product entity.Product
	err = m.client.Database("market").Collection("product").FindOne(m.ctxMongo, bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (m *MongoProductRepository) UpdateProduct(ctx context.Context, id int, productUpdate *entity.ProductUpdate) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": id}
	update := bson.D{{"$set", productUpdate}}
	updateResult, err := m.client.Database("market").Collection("product").UpdateOne(m.ctxMongo, filter, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (m *MongoProductRepository) FindAllProducts(ctx context.Context) (*[]entity.Product, error) {
	var products []entity.Product
	cursor, err := m.client.Database("market").Collection("product").Find(m.ctxMongo, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(m.ctxMongo, &products); err != nil {
		return nil, err
	}
	return &products, nil
}

func (m *MongoProductRepository) DeleteProductById(ctx context.Context, id int) (*mongo.DeleteResult, error) {
	filter := bson.M{"id": id}
	result, err := m.client.Database("market").Collection("product").DeleteOne(m.ctxMongo, filter)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *MongoProductRepository) DecreaseStock(ctx context.Context, productOrder []entity.ProductOrder) error {

	for _, productOrder := range productOrder {
		product, _ := m.FindProductById(ctx, productOrder.ProductID)
		currenStock := product.Stock
		newStock := currenStock - productOrder.Quantity
		err := m.UpdateStock(ctx, productOrder.ProductID, newStock)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MongoProductRepository) UpdateStock(ctx context.Context, productID int, quantity int) error {
	filter := bson.M{"id": productID}
	update := bson.M{"$set": bson.M{"stock": quantity}}
	_, err := m.client.Database("market").Collection("product").UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoProductRepository) CheckStock(ctx context.Context, productID int, quantity int) error {
	if quantity == 0 {
		return nil
	}
	product, err := m.FindProductById(ctx, productID)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("product with ID %d not found", productID)
		}
		return fmt.Errorf("error finding product: %w", err)
	}
	stock := entity.Stock{ID: product.ID, Quantities: product.Stock}

	if stock.Quantities < quantity {
		return fmt.Errorf("insufficient stock for product ID %d, only %d available", productID, stock.Quantities)
	}
	return nil
}
