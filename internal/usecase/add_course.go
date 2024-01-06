package usecase


import (
	"context"

	"github.com/lucasferreirajs/18-uow/internal/entity"
	"github.com/lucasferreirajs/18-uow/internal/repository"
)

type InputUsecase struct {

	CategoryName string
	CourseName string
	CourseCategoryID int
}

type AddCourseUsecase struct {
	CourseRepository repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}


func NewAddCourseUsecase(courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepositoryInterface) *AddCourseUsecase {
	return &AddCourseUsecase {
		CourseRepository: courseRepository,
		CategoryRepository: categoryRepository,
	}

}


func (a *AddCourseUsecase) Execute(ctx context.Context, input InputUsecase) error {
	category := entity.Category {
		Name : input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course {
		Name: input.CourseName,
		CategoryID: input.CourseCategoryID,
	}
	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {

		return err
	}

	return nil
}


