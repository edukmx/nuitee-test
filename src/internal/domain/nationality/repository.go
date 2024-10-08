package nationality

//go:generate mockery --case=snake --outpkg=mocks --output=../../mocks --name=Repository
type Repository interface {
	FindByIso(iso string) (*Nationality, error)
}
