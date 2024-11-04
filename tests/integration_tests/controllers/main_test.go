package integration_tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	_ "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain_ExecuteEndpoint(t *testing.T) {
	go func() {
		run_app_for_test("1236")
	}()

	time.Sleep(3 * time.Second)
	resp, err := http.Get("http://localhost:1236/execute")

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result []domain.Document
	err = json.Unmarshal(body, &result)

	assert.Len(t, result, 1)
	assert.Equal(t, "1", result[0].Id)
	assert.Equal(t, "Document 1", result[0].Name)
}
