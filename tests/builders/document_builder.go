package builders

import "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"

type DocumentBuilder struct {
	document domain.Document
}

func NewDocumentBuilder() *DocumentBuilder {
	return &DocumentBuilder{
		document: domain.Document{},
	}
}

func (b *DocumentBuilder) WithId(id string) *DocumentBuilder {
	b.document.Id = id
	return b
}

func (b *DocumentBuilder) WithName(name string) *DocumentBuilder {
	b.document.Name = name
	return b
}

func (b *DocumentBuilder) Build() domain.Document {
	return b.document
}
