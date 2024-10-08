package repository_test

import (
	"testing"

	"github.com/edukmx/nuitee/internal/domain/nationality"
	"github.com/edukmx/nuitee/internal/infra/repository"
)

func TestFindByIso(t *testing.T) {
	tests := []struct {
		name          string
		nationalities []nationality.Nationality
		iso           string
		expected      *nationality.Nationality
		expectErr     bool
	}{
		{
			name: "it should return with no errors",
			nationalities: []nationality.Nationality{
				{NationalityISO: "AR", LanguageISO: "es"},
				{NationalityISO: "US", LanguageISO: "en"},
			},
			iso:       "AR",
			expected:  &nationality.Nationality{NationalityISO: "AR", LanguageISO: "es"},
			expectErr: false,
		},
		{
			name: "it should return not found error",
			nationalities: []nationality.Nationality{
				{NationalityISO: "AR", LanguageISO: "es"},
				{NationalityISO: "US", LanguageISO: "en"},
			},
			iso:       "FR",
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NationalityRepository{
				Nationalities: tt.nationalities,
			}

			result, err := repo.FindByIso(tt.iso)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if !tt.expectErr && (result.NationalityISO != tt.expected.NationalityISO || result.LanguageISO != tt.expected.LanguageISO) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
