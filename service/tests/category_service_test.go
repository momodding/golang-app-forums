package tests

import (
	"database/sql"
	"fmt"
	"forum-app/entity"
	service2 "forum-app/service"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
	DB   *sql.DB
	mock sqlmock.Sqlmock

	service service2.CategoryService
}

func (s *Suite) SetupSuite() {
	var (
		err error
	)

	s.DB, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	gdb, err := gorm.Open(postgres.New(postgres.Config{Conn: s.DB}), &gorm.Config{Logger: newLogger})
	require.NoError(s.T(), err)

	s.service = service2.NewCategoryService(gdb)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_Find_All_then_Return_data() {
	var (
		id = uint64(1)
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "category"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(id, "test", "test"))

	res := s.service.FindAll()
	assert.Equal(s.T(), id, res[0].ID)
}

func (s *Suite) Test_Save_then_Return_data() {
	var (
		id       = uint64(1)
		currTime = time.Now()
		category = &entity.Category{Name: "test", Description: "test", CreatedAt: &currTime, UpdatedAt: &currTime}
	)
	fmt.Println(currTime)
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "category" ("name","description","created_at","updated_at","deleted_at") 
						VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(category.Name, category.Description, category.CreatedAt, category.UpdatedAt, category.DeletedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(id))
	s.mock.ExpectCommit()

	res := s.service.Save(*category)
	assert.Equal(s.T(), id, res.ID)
}
