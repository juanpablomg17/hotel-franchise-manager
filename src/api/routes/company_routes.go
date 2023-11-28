package routes

import (
	"net/http"

	controller "github.com/flexuxs/clubHubApp/src/api/controller/company"
	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/gin-gonic/gin"
)

type CompanyRoutes struct {
	Controller *controller.Controller
}

func NewCompanyRoutes(Controller *controller.Controller) *CompanyRoutes {
	return &CompanyRoutes{
		Controller: Controller,
	}
}

func (cr *CompanyRoutes) SaveCompany(c *gin.Context) {

	companyDto := &dto.CompanyDTO{}

	if err := c.ShouldBindJSON(&companyDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := cr.Controller.SaveCompany(companyDto)

	c.JSON(company.StatusCode, gin.H{
		"company": company,
	})
}

func (cr *CompanyRoutes) GetCompany(c *gin.Context) {

	getCompanyDto := &dto.CompanyFilterDTO{}

	if err := c.ShouldBind(&getCompanyDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := cr.Controller.GetCompanyByfilter(getCompanyDto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(response.StatusCode, gin.H{
		"message":   response.Message,
		"companies": response.Companies,
	})
}
