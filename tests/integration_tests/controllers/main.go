package integration_tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	my_controller "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/http"
	tests "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
)

func run_app_for_test(port string) {
	app := fiber.New()

	err := tests.DocumentRepositoryGlobal.Create(domain.Document{Id: "1", Name: "Document 1"})
	if err != nil {
		panic(err)
	}

	app.Get("/execute", my_controller.MyController(tests.DocumentRepositoryGlobal))

	app.Listen(":" + port)
}
