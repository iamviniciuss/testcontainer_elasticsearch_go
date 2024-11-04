package infra

import (
	"fmt"
	"testing"
	// "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/testsetup"
)

func TestMain(m *testing.M) {
	fmt.Println(" #### Starting Elasticsearch Container before all tests ####")
	// testsetup.TestMain(m)
	fmt.Println(" #### Elasticsearch container is ready ####")
}
