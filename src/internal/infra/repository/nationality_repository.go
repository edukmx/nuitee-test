package repository

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/edukmx/nuitee/config"
	"github.com/edukmx/nuitee/internal/domain/nationality"
)

type NationalityRepository struct {
	Nationalities []nationality.Nationality
}

func NewNationalityRepository(config *config.Config) (nationality.Repository, error) {
	file, err := os.Open(config.NationalitiesPath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}
	var nationalities []nationality.Nationality
	for i, record := range records {
		if i == 0 {
			continue
		}
		entity := nationality.Nationality{
			NationalityISO: record[0],
			LanguageISO:    record[1],
		}
		nationalities = append(nationalities, entity)
	}
	return &NationalityRepository{Nationalities: nationalities}, nil
}

func (n NationalityRepository) FindByIso(iso string) (*nationality.Nationality, error) {
	for _, entity := range n.Nationalities {
		if entity.NationalityISO == iso {
			return &entity, nil
		}
	}
	return nil, fmt.Errorf("nationality %s not found", iso)
}
