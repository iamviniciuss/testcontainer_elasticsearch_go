package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	elasticsearch "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
	my_controller "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/http"
	repository "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/repository"
)

func main() {
	app := fiber.New()

	documentRepository := repository.NewDocumentRepositoryES(elasticsearch.NewESConnection())
	documentRepository.Create(domain.Document{Id: "1", Name: "Document 1"})

	app.Get("/execute", my_controller.MyController(documentRepository))

	app.Listen(":3000")
}
