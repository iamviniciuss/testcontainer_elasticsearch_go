package tests_setup

import (
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
	infra "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/repository"
	setup "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests/setup"
)

var (
	DocumentRepositoryGlobal domain.DocumentRepository
)

func init() {
	// setup.NewESConnectionTests("infra2")
	// DocumentRepositoryGlobal = infra.NewDocumentRepositoryES(elasticsearch.NewESConnection())

	// // defer teardown()
	teardown := setup.NewESConnectionTests("infra2")
	DocumentRepositoryGlobal = infra.NewDocumentRepositoryES(elasticsearch.NewESConnection())

	defer teardown()
}
