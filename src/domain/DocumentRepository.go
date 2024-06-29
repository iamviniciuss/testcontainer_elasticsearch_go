package domain

type DocumentRepository interface {
	List() (*Document, error)
}
