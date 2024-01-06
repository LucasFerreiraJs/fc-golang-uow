package repository


import (
	"context"
	"database/sql"

	"github.com/lucasferreirajs/18-uow/internal/db"
	"github.com/lucasferreirajs/18-uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error

}

type CourseRepository struct {
	DB *sql.DB
	Queries *db.Queries
}


func (r *CourseRepository) Insert(ctx context.Context,course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name: course.Name,
		CategoryID: int32(course.CategoryID),
	})
}

