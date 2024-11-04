package use_cases

import (
	"testing"

	use_cases "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/application/use_cases"
	setup_test "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/tests/builders"

	"github.com/stretchr/testify/assert"
)

func setupTestData(t *testing.T) {
	document := builders.NewDocumentBuilder().
		WithId("1").
		WithName("Document 1").
		Build()

	err := setup_test.DocumentRepositoryGlobal.Create(document)
	assert.NoError(t, err)
}

func TestMyUseCase_Execute(t *testing.T) {
	setupTestData(t)

	useCase := use_cases.NewMyUseCase(setup_test.DocumentRepositoryGlobal)

	result, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].Id)
	assert.Equal(t, "Document 1", result[0].Name)
}
