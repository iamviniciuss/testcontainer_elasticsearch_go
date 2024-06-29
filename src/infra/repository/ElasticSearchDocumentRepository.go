package infra

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	elasticsearch "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
)

type DocumentRepositoryES struct {
	esConnection *elasticsearch.ESConnection
}

func NewDocumentRepositoryES(conn elasticsearch.Connection) *DocumentRepositoryES {
	ar := new(DocumentRepositoryES)
	ar.esConnection = conn.(*elasticsearch.ESConnection)
	return ar
}

func (cr *DocumentRepositoryES) List() ([]domain.Document, error) {
	res, err := cr.esConnection.Client().Search(
		cr.esConnection.Client().Search.WithContext(context.Background()),
		cr.esConnection.Client().Search.WithIndex("documents"),
		cr.esConnection.Client().Search.WithBody(strings.NewReader(`{"query": { "match_all": {} }}`)),
	)

	if err != nil {
		return []domain.Document{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return []domain.Document{}, errors.New(res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return []domain.Document{}, err
	}

	list := []domain.Document{}
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		list = append(list, domain.Document{
			Id:                 source["id"].(string),
			Name:               source["name"].(string),
			Status:             source["status"].(string),
			Type:               source["type"].(string),
			CreateByEmployeeID: source["create_by_employee_id"].(string),
		})
	}

	return list, nil
}
