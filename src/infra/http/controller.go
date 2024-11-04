package http

import (
	"github.com/gofiber/fiber/v2"
	use_cases "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/application/use_cases"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
)

func MyController(repo domain.DocumentRepository) func(c *fiber.Ctx) error {
	myUseCase := use_cases.NewMyUseCase(repo)

	return func(c *fiber.Ctx) error {

		result, err := myUseCase.Execute()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(result)
	}
}
