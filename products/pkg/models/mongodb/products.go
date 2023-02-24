package mongodb

import (
	"context"
	"errors"

	"github.com/duard/simples-microservices/products/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductModel represent a mgo database session with a product data model.
type ProductModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the products table.
func (m *ProductModel) All() ([]models.Product, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.Product{}

	// Find all products
	productCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = productCursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	return mm, err
}

// FindByID will be used to find a new product registry by id
func (m *ProductModel) FindByID(id string) (*models.Product, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Find product by id
	var product = models.Product{}
	err = m.C.FindOne(context.TODO(), bson.M{"_id": p}).Decode(&product)
	if err != nil {
		// Checks if the product was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &product, nil
}

// Insert will be used to insert a new product registry
func (m *ProductModel) Insert(product models.Product) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), product)
}

// Delete will be used to delete a product registry
func (m *ProductModel) Delete(id string) (*mongo.DeleteResult, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return m.C.DeleteOne(context.TODO(), bson.M{"_id": p})
}
