package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/model"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/repository"
	"zisluiz.com/cpfcnpj-crud-api/infra/config"
)

const (
	collectionName = "document"
)

type DocumentRepositoryImpl struct {
	repository.DocumentRepository
}

type DocumentStructure struct {
	Id             string `bson:"_id"`
	Name           string `bson:"name"`
	IdentityNumber string `bson:"identityNumber"`
}

func (repository *DocumentRepositoryImpl) Get(uuid string) (*model.Document, *exception.Validations) {
	log.Printf("Finding document %s", uuid)

	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return nil, exception.NewError("Invalid id value!")
	}

	var database, ctx, close, error = config.GetClient()

	if error != nil {
		return nil, exception.NewError(error.Error())
	}

	defer close()

	var result = database().Collection(collectionName).FindOne(ctx, bson.M{"_id": objectId})
	documentStructure := &DocumentStructure{}

	err = result.Decode(documentStructure)

	if err != nil {
		return nil, exception.NewError(err.Error())
	}

	var document, validations = factory.NewDocument(documentStructure.Id, documentStructure.Name, documentStructure.IdentityNumber)

	if validations != nil {
		return nil, validations
	}

	return document, nil
}

func (repository *DocumentRepositoryImpl) Delete(uuid string) *exception.Validations {
	log.Printf("Deleting document %s", uuid)

	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return exception.NewError("Invalid id value!")
	}

	var database, ctx, close, error = config.GetClient()

	if error != nil {
		return exception.NewError(error.Error())
	}

	defer close()

	result, error := database().Collection(collectionName).DeleteOne(ctx, bson.M{"_id": objectId})

	if error != nil {
		return exception.NewError(error.Error())
	}

	if result.DeletedCount == 0 {
		return exception.NewError("No document found for uuid " + uuid + ".")
	}

	return nil
}

func (repository *DocumentRepositoryImpl) Insert(document *model.Document) *exception.Validations {
	log.Print("Preparing to insert a new document.")

	var database, ctx, close, error = config.GetClient()

	if error != nil {
		return exception.NewError(error.Error())
	}

	defer close()

	var _, errDb = database().Collection(collectionName).InsertOne(ctx, bson.D{
		{Key: "name", Value: document.Name},
		{Key: "identityNumber", Value: document.Identity.Value()},
	})

	if errDb != nil {
		return exception.NewError(errDb.Error())
	}

	return nil
}

func (repository *DocumentRepositoryImpl) Update(document *model.Document) *exception.Validations {
	log.Printf("Preparing to update document %s.", document.Uuid)

	objectId, err := primitive.ObjectIDFromHex(document.Uuid)
	if err != nil {
		return exception.NewError("Invalid id value!")
	}

	var database, ctx, close, error = config.GetClient()

	if error != nil {
		return exception.NewError(error.Error())
	}

	defer close()

	var result, errDb = database().Collection(collectionName).UpdateByID(ctx, bson.M{"_id": objectId}, bson.D{
		{Key: "name", Value: document.Name},
		{Key: "identityNumber", Value: document.Identity.Value()},
	})

	if errDb != nil {
		return exception.NewError(errDb.Error())
	}

	if result.ModifiedCount == 0 {
		return exception.NewError("No document found for uuid " + document.Uuid + ".")
	}

	return nil
}
