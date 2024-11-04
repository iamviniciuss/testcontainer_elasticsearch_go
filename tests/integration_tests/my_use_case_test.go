package use_cases

import (
	"testing"

	use_cases "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/application/use_cases"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	setup_test "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"

	"github.com/stretchr/testify/assert"
)

func TestMyUseCase_Execute(t *testing.T) {
	err1 := setup_test.DocumentRepositoryGlobal.Create(domain.Document{
		Id: "1",
	})
	assert.NoError(t, err1)

	useCase := use_cases.NewMyUseCase(setup_test.DocumentRepositoryGlobal)

	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].Id)
}
