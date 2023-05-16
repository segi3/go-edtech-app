package course

import (
	dto "edtech-app/internal/course/dto"
	entity "edtech-app/internal/course/entity"
	repository "edtech-app/internal/course/repository"
	lessonEntity "edtech-app/internal/lesson/entity"
	lessonUseCase "edtech-app/internal/lesson/usecase"
	productEntity "edtech-app/internal/product/entity"
	productUseCase "edtech-app/internal/product/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"errors"
	"fmt"
	"strconv"
)

type CourseUseCase interface {
	FindAll(offset int, limit int) []productEntity.Product
	FindByProductId(id int) (*dto.CourseResponseBody, error)
	Create(dto dto.CourseBindingRequestBody) (*dto.CourseResponseBody, error)
	Update(id int, dto dto.CourseBindingRequestBody) (*entity.Course, error)
	Delete(id int) error
}

type CourseUseCaseImpl struct {
	repository     repository.CourseRepository
	fileUpload     fileUpload.FileUpload
	productUseCase productUseCase.ProductUseCase
	lessonUseCase  lessonUseCase.LessonUseCase
}

// Create implements CourseUseCase
func (usecase *CourseUseCaseImpl) Create(dtoInput dto.CourseBindingRequestBody) (*dto.CourseResponseBody, error) {

	// validate if product id
	exist, err := usecase.productUseCase.FindExist(int(dtoInput.ProductID))

	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errors.New("product id:" + strconv.Itoa(int(dtoInput.ProductID)) + " not found")
	}

	// validate if lesson ids exists
	for _, lessonId := range dtoInput.LessonIDs {

		exist, err = usecase.lessonUseCase.FindExist(int(lessonId))

		if err != nil {
			return nil, err
		}

		if !exist {
			return nil, errors.New("lesson id:" + strconv.Itoa(int(lessonId)) + " not found")
		}
	}

	// check product is already have course
	exist, err = usecase.repository.FindByProductIdExists(int(dtoInput.ProductID))

	if err != nil {
		return nil, err
	}

	if exist {
		return nil, errors.New("product id:" + strconv.Itoa(int(dtoInput.ProductID)) + " already have a course")
	}

	// loop lesson id array and create entity within iterations
	for _, lessonId := range dtoInput.LessonIDs {

		courseData := entity.Course{
			LessonID:    lessonId,
			ProductID:   dtoInput.ProductID,
			CreatedByID: &dtoInput.CreatedBy,
		}

		usecase.repository.Create(courseData)
	}

	// return response course
	courses, _ := usecase.repository.FindByProductId(int(dtoInput.ProductID))

	var lessons []lessonEntity.Lesson
	for _, course := range courses {
		lessons = append(lessons, *course.Lesson)
	}
	fmt.Println(lessons)

	courseResponse := dto.CourseResponseBody{
		Product: *courses[0].Product,
		Lessons: lessons,
	}

	return &courseResponse, nil

}

// Delete implements CourseUseCase
func (usecase *CourseUseCaseImpl) Delete(id int) error {
	// get lesson by id
	lesson, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*lesson)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements CourseUseCase
func (usecase *CourseUseCaseImpl) FindAll(offset int, limit int) []productEntity.Product {
	return usecase.productUseCase.FindAll(offset, limit)
}

// FindById implements CourseUseCase
func (usecase *CourseUseCaseImpl) FindByProductId(id int) (*dto.CourseResponseBody, error) {

	// check if product exists
	exist, err := usecase.repository.FindByProductIdExists(id)

	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errors.New("product id:" + strconv.Itoa(id) + " does not exists")
	}

	// return response course
	courses, _ := usecase.repository.FindByProductId(id)

	var lessons []lessonEntity.Lesson
	for _, course := range courses {
		lessons = append(lessons, *course.Lesson)
	}
	fmt.Println(lessons)

	courseResponse := dto.CourseResponseBody{
		Product: *courses[0].Product,
		Lessons: lessons,
	}

	return &courseResponse, nil
}

// Update implements CourseUseCase
func (usecase *CourseUseCaseImpl) Update(id int, dto dto.CourseBindingRequestBody) (*entity.Course, error) {
	panic("unimplemented")

}

func NewCourseUseCase(repository repository.CourseRepository, fileUpload fileUpload.FileUpload, productUseCase productUseCase.ProductUseCase, lessonUseCase lessonUseCase.LessonUseCase) CourseUseCase {
	return &CourseUseCaseImpl{repository, fileUpload, productUseCase, lessonUseCase}
}
