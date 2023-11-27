package mapper

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/stretchr/testify/assert"
)

func TestFromCompanyDtoToCompanyAggregate(t *testing.T) {
	// Arrange
	companyDto := &dto.CompanyDTO{}

	gofakeit.Struct(companyDto)

	// Act
	companyAggregate := FromCompanyDtoToCompanyAggregate(companyDto)

	// Assert
	assert.Equal(t, companyDto.Company.Owner.FirstName, companyAggregate.Company.Owner.FirstName)
	assert.Equal(t, companyDto.Company.Owner.LastName, companyAggregate.Company.Owner.LastName)
	assert.Equal(t, companyDto.Company.Owner.Contact.Email, companyAggregate.Company.Owner.Contact.Email)
	assert.Equal(t, companyDto.Company.Owner.Contact.Phone, companyAggregate.Company.Owner.Contact.Phone)
	assert.Equal(t, companyDto.Company.Owner.Contact.Location.City, companyAggregate.Company.Owner.Contact.Location.City)
	assert.Equal(t, companyDto.Company.Owner.Contact.Location.Country, companyAggregate.Company.Owner.Contact.Location.Country)
	assert.Equal(t, companyDto.Company.Owner.Contact.Location.Address, companyAggregate.Company.Owner.Contact.Location.Address)
	assert.Equal(t, companyDto.Company.Owner.Contact.Location.ZipCode, companyAggregate.Company.Owner.Contact.Location.ZipCode)
	assert.Equal(t, companyDto.Company.Information.Name, companyAggregate.Company.Information.Name)
	assert.Equal(t, companyDto.Company.Information.TaxNumber, companyAggregate.Company.Information.TaxNumber)
	assert.Equal(t, companyDto.Company.Information.Location.City, companyAggregate.Company.Information.Location.City)
	assert.Equal(t, companyDto.Company.Information.Location.Country, companyAggregate.Company.Information.Location.Country)
	assert.Equal(t, companyDto.Company.Information.Location.Address, companyAggregate.Company.Information.Location.Address)
	assert.Equal(t, companyDto.Company.Information.Location.ZipCode, companyAggregate.Company.Information.Location.ZipCode)
	assert.Equal(t, len(companyDto.Company.Franchises), len(companyAggregate.Company.Franchises))

	for i, franchiseDto := range companyDto.Company.Franchises {
		assert.Equal(t, franchiseDto.Name, companyAggregate.Company.Franchises[i].Name)
		assert.Equal(t, franchiseDto.URL, companyAggregate.Company.Franchises[i].URL)
		assert.Equal(t, franchiseDto.Location.City, companyAggregate.Company.Franchises[i].Location.City)
		assert.Equal(t, franchiseDto.Location.Country, companyAggregate.Company.Franchises[i].Location.Country)
		assert.Equal(t, franchiseDto.Location.Address, companyAggregate.Company.Franchises[i].Location.Address)
		assert.Equal(t, franchiseDto.Location.ZipCode, companyAggregate.Company.Franchises[i].Location.ZipCode)
	}
}
