package domain

type DocumentRepository interface {
	List() ([]Document, error)
	Create(doc Document) error
}
