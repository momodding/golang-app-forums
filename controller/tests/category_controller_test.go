package tests

import (
	"encoding/json"
	controller2 "forum-app/controller"
	"forum-app/entity"
	mocks "forum-app/mocks/service"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type Suite struct {
	suite.Suite

	controller controller2.CategoryController
	service    mocks.CategoryService
}

func (s *Suite) SetupSuite() {

	s.service = mocks.CategoryService{}

	s.controller = controller2.NewCategoryController(&s.service)
}

func (s *Suite) AfterTest(_, _ string) {
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Find_All_then_Return_data() {
	s.service.On("FindAll").Return([]entity.Category{{ID: 1, Name: "baru"}})

	req := httptest.NewRequest("GET", "/categories", nil)
	req.Header.Set("X-Custom-Header", "hi")

	app := fiber.New()
	// http.Response
	resp, _ := app.Test(req)
	var data []entity.Category
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	assert.NotEqual(s.T(), nil, resp)
	assert.NotEqual(s.T(), 0, len(data))
	assert.Equal(s.T(), 1, data[0].ID)
}

func (s *Suite) Test_Create_then_Return_data() {
	s.service.On("Create").Return(entity.Category{ID: 1, Name: "baru"})

	req := httptest.NewRequest("POST", "/categories", nil)
	req.Header.Set("X-Custom-Header", "hi")

	app := fiber.New()
	// http.Response
	resp, _ := app.Test(req)
	var data entity.Category
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	assert.NotEqual(s.T(), nil, resp)
	assert.NotEqual(s.T(), nil, data)
	assert.Equal(s.T(), 1, data.ID)
}
