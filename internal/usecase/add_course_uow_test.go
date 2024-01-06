package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/lucasferreirajs/18-uow/internal/db"
	"github.com/lucasferreirajs/18-uow/internal/repository"
	"github.com/lucasferreirajs/18-uow/pkg/uow"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS `courses`;")
	dbt.Exec("DROP TABLE IF EXISTS `categories`;")

	dbt.Exec("create table if not exists `categories` (id int primary key AUTO_INCREMENT,name varchar(255) not null);")
	dbt.Exec("create table if not exists `courses` ( id int primary key AUTO_INCREMENT, name varchar(255) not null, category_id integer not null, foreign key (category_id) references categories(id));")

	ctx := context.Background()

	uow := uow.NewUow(ctx, dbt)
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})
	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUsecase{
		CategoryName:     "Category_1",
		CourseName:       "Course_1",
		CourseCategoryID: 1,
	}

	usecase := NewAddCourseUsecaseUow(uow)
	err = usecase.Execute(ctx, input)
	assert.NoError(t, err)

}
