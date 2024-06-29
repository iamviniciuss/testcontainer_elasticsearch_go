package infra

import (
	"testing"

	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
	"github.com/iamviniciuss/testcontainer_elasticsearch_go/tests"
	"github.com/stretchr/testify/suite"
)

const (
	DOCUMENT_INDEX_NAME = "documents"
)

type DocumentRepositoryESSuite struct {
	suite.Suite
	*elasticsearch.ESConnection
}

func TestDocumentRepositoryESSuite(t *testing.T) {
	suite.Run(t, &DocumentRepositoryESSuite{
		ESConnection: elasticsearch.NewESConnection(),
	})
}

func (suite *DocumentRepositoryESSuite) Test_Create_A_NPS_Campaign() {
	tests.CreateIndex(suite.ESConnection.Client())

	documentRepositoryES := NewDocumentRepositoryES(suite.ESConnection)
	documents, err := documentRepositoryES.List()

	suite.Nil(err)
	suite.Len(documents, 0)
}

func (suite *DocumentRepositoryESSuite) Test_Create_A_NPS_Campaign2() {
	tests.CreateIndex(suite.ESConnection.Client())
	tests.IndexDocuments(suite.ESConnection.Client())

	documentRepositoryES := NewDocumentRepositoryES(suite.ESConnection)
	documents, err := documentRepositoryES.List()

	suite.Nil(err)
	suite.Len(documents, 3)
}
