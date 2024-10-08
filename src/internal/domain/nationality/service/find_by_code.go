package service

import (
	"github.com/edukmx/nuitee/internal/domain"
	"github.com/edukmx/nuitee/internal/domain/nationality"
)

type FindByCode struct {
	repo nationality.Repository
}

func NewFindByCode(repo nationality.Repository) *FindByCode {
	return &FindByCode{repo: repo}
}

func (s *FindByCode) Find(code string) (*nationality.Nationality, error) {
	entity, err := s.repo.FindByIso(code)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, domain.ErrNationalityNotSupported
	}
	return entity, nil
}
