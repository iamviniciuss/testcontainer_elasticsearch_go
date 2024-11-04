package infra

import (
	"testing"

	"github.com/iamviniciuss/testcontainer_elasticsearch_go/src/infra/elasticsearch"
	setup "github.com/iamviniciuss/testcontainer_elasticsearch_go/tests/setup"
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

func (suite *DocumentRepositoryESSuite) Test_List_Zero_Documents() {
	setup.CreateIndex(suite.ESConnection.Client())

	documentRepositoryES := NewDocumentRepositoryES(suite.ESConnection)
	documents, err := documentRepositoryES.List()

	suite.Nil(err)
	suite.Len(documents, 0)
}

func (suite *DocumentRepositoryESSuite) Test_List_All_Documents() {
	setup.CreateIndex(suite.ESConnection.Client())
	setup.IndexDocuments(suite.ESConnection.Client())

	documentRepositoryES := NewDocumentRepositoryES(suite.ESConnection)
	documents, err := documentRepositoryES.List()

	suite.Nil(err)
	suite.Len(documents, 3)
}
