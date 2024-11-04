package use_cases

import "github.com/iamviniciuss/testcontainer_elasticsearch_go/src/domain"

type MyUseCase struct {
	repository domain.DocumentRepository
}

func NewMyUseCase(repo domain.DocumentRepository) *MyUseCase {
	return &MyUseCase{repository: repo}
}

func (uc *MyUseCase) Execute() ([]domain.Document, error) {
	data, err := uc.repository.List()
	if err != nil {
		return nil, err
	}
	return data, nil
}
