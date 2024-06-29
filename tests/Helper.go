package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	esv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
)

const (
	DOCUMENT_INDEX_NAME = "documents"
)

func IndexDocuments(es *esv8.Client) {
	documents := []domain.Document{
		{Id: "123", Name: "Employee List", Status: "FINISH", Type: "PDF", CreateByEmployeeID: "817"},
		{Id: "124", Name: "Account Owner List", Status: "FINISH", Type: "PDF", CreateByEmployeeID: "820"},
		{Id: "125", Name: "Customer List", Status: "FINISH", Type: "PDF", CreateByEmployeeID: "912"},
	}

	for _, doc := range documents {
		docJSON, err := json.Marshal(doc)
		if err != nil {
			log.Printf("Error marshaling document: %s", err)
		}

		req := esapi.IndexRequest{
			Index:      DOCUMENT_INDEX_NAME,
			DocumentID: doc.Id,
			Body:       strings.NewReader(string(docJSON)),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Printf("Error indexing document %s: %s", doc.Id, err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error response from Elasticsearch for document %s: %s", doc.Id, res.String())
		}
	}
	fmt.Println("Documents indexed")
}

func CreateIndex(es *esv8.Client) {
	mapping := `{
        "mappings": {
            "properties": {
                "id": { "type": "keyword" },
                "name": { "type": "text" },
                "status": { "type": "text" },
                "type": { "type": "keyword" },
                "create_by_employee_id": { "type": "keyword" }
            }
        }
    }`

	req := esapi.IndicesCreateRequest{
		Index: DOCUMENT_INDEX_NAME,
		Body:  strings.NewReader(mapping),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error creating index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error response from Elasticsearch: %s", res.String())
		return
	}

	fmt.Println("Index created")
}

func DeleteIndex(es *esv8.Client) {
	req := esapi.IndicesDeleteRequest{
		Index: []string{DOCUMENT_INDEX_NAME},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error deleting index %s: %s", DOCUMENT_INDEX_NAME, err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Error response from Elasticsearch for index %s: %s", DOCUMENT_INDEX_NAME, res.String())
		return
	}

	log.Printf("Index %s deleted successfully", DOCUMENT_INDEX_NAME)
}
