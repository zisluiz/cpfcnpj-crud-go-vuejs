package dao

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/exception"
	"zisluiz.com/cpfcnpj-crud-api/command/domain/factory"
	"zisluiz.com/cpfcnpj-crud-api/infra/config"
	"zisluiz.com/cpfcnpj-crud-api/infra/repository"
	"zisluiz.com/cpfcnpj-crud-api/query/dao"
	"zisluiz.com/cpfcnpj-crud-api/query/model"
)

type DocumentDAOImpl struct {
	dao.DocumentDAO
}

type DocumentQueryViewModel struct {
	Uuid           string `json:"uuid"`
	Name           string `json:"name"`
	IdentityNumber string `json:"identityNumber"`
	IdentityType   string `json:"identityType"`
	Blocked        bool   `json:"blocked"`
}

func (daoImpl *DocumentDAOImpl) FindDocumentsBy(query *model.Query) *model.Result {
	queryJson, _ := json.Marshal(query)
	log.Printf("Querying documents %s", string(queryJson))

	var database, ctx, close, error = config.GetClient()

	if error != nil {
		return &model.Result{Validations: exception.NewError(error.Error())}
	}

	defer close()

	opts := options.FindOptions{
		Skip:  &query.Page,
		Limit: &query.ResultsPerPage,
	}

	var filters map[string]interface{}

	if query.Filters != nil {
		filters = make(map[string]interface{})

		for _, filter := range query.Filters {
			if reflect.TypeOf(filter.Value).Kind().String() == "string" {
				filters[filter.Name] = primitive.Regex{Pattern: filter.Value.(string), Options: "-i"}
			} else {
				filters[filter.Name] = filter.Value
			}
		}
	}

	if query.Sorts != nil {
		sorts := make(map[string]int)

		for _, sort := range query.Sorts {
			var asc = 1

			if sort.Order == model.SortFieldOrderDESC {
				asc = -1
			}

			sorts[sort.Name] = asc
		}

		opts.SetSort(sorts)
	}

	collection := database().Collection(repository.DocumentCollectionName)

	totalCount, error := collection.CountDocuments(ctx, filters)

	if error != nil {
		return &model.Result{Validations: exception.NewError(error.Error())}
	}

	cursor, error := collection.Find(ctx, filters, &opts)

	if error != nil {
		return &model.Result{Validations: exception.NewError(error.Error())}
	}

	var results []bson.M
	if error = cursor.All(context.TODO(), &results); error != nil {
		return &model.Result{Validations: exception.NewError(error.Error())}
	}

	documents := []*DocumentQueryViewModel{}

	for _, result := range results {
		blocked := false

		if result[repository.DocumentFieldBlocked] != nil {
			blocked = result[repository.DocumentFieldBlocked].(bool)
		}

		document, errorConstruct := factory.NewDocumentWithBlocked(result[repository.DocumentFieldId].(primitive.ObjectID).Hex(), result[repository.DocumentFieldName].(string),
			result[repository.DocumentFieldIdentityNumber].(string), blocked)

		if errorConstruct != nil {
			return &model.Result{Validations: errorConstruct}
		}

		documents = append(documents, &DocumentQueryViewModel{Uuid: document.Uuid, Name: document.Name, IdentityNumber: document.Identity.Value(), IdentityType: document.Identity.Type(), Blocked: document.Blocked})
	}

	return &model.Result{Content: documents, Page: query.Page, ResultsPerPage: query.ResultsPerPage, TotalResults: totalCount}
}
