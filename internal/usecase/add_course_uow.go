package usecase

import (
	"context"

	"github.com/lucasferreirajs/18-uow/internal/entity"
	"github.com/lucasferreirajs/18-uow/internal/repository"
	"github.com/lucasferreirajs/18-uow/pkg/uow"
)

type InputUsecaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUsecaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUsecaseUow(uow uow.UowInterface) *AddCourseUsecaseUow {
	return &AddCourseUsecaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUsecaseUow) Execute(ctx context.Context, input InputUsecase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {

		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.getCagetoryRepository(ctx)

		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		repoCourse := a.getCourseRepository(ctx)

		err = repoCourse.Insert(ctx, course)
		if err != nil {

			return err
		}

		return nil
	})
}

func (a *AddCourseUsecaseUow) getCagetoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	// cast
	return repo.(repository.CategoryRepositoryInterface)
}

func (a *AddCourseUsecaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	// cast
	return repo.(repository.CourseRepositoryInterface)
}
