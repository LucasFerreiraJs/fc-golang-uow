package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/lucasferreirajs/18-uow/internal/repository"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourse(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE IF EXISTS `courses`;")
	dbt.Exec("DROP TABLE IF EXISTS `categories`;")

	dbt.Exec("create table if not exists `categories` (id int primary key AUTO_INCREMENT,name varchar(255) not null);")
	dbt.Exec("create table if not exists `courses` ( id int primary key AUTO_INCREMENT, name varchar(255) not null, category_id integer not null, foreign key (category_id) references categories(id));")

	input := InputUsecase{
		CategoryName:     "Category_1",
		CourseName:       "Course_1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()
	usecase := NewAddCourseUsecase(repository.NewCourseRepository(dbt), repository.NewCategoryRepository(dbt))
	err = usecase.Execute(ctx, input)
	assert.NoError(t, err)

}
