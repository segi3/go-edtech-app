package course

import (
	entity "edtech-app/internal/course/entity"
	"edtech-app/pkg/utils"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll(offset int, limit int) []entity.Course
	FindById(id int) (*entity.Course, error)
	FindByProductId(productId int) ([]entity.Course, error)
	Create(entity entity.Course) (*entity.Course, error)
	Update(entity entity.Course) (*entity.Course, error)
	Delete(entity entity.Course) error
}

type CourseRepositoryImpl struct {
	db *gorm.DB
}

// FindAllByProductId implements CourseRepository
func (repository *CourseRepositoryImpl) FindByProductId(productId int) ([]entity.Course, error) {
	var courses []entity.Course

	res := repository.db.
		Preload("Lesson").
		Preload("Product").
		Where("product_id = ?", productId).
		Find(&courses)

	if res.Error != nil {
		return nil, res.Error
	}

	return courses, nil
}

// Create implements CourseRepository
func (repository *CourseRepositoryImpl) Create(entity entity.Course) (*entity.Course, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements CourseRepository
func (repository *CourseRepositoryImpl) Delete(entity entity.Course) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements CourseRepository
func (repository *CourseRepositoryImpl) FindAll(offset int, limit int) []entity.Course {
	var courses []entity.Course

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("Product").Preload("Lesson").Find(&courses)

	return courses
}

// FindById implements CourseRepository
func (repository *CourseRepositoryImpl) FindById(id int) (*entity.Course, error) {
	var course entity.Course

	if err := repository.db.First(&course, id).Error; err != nil {
		return nil, err
	}

	return &course, nil
}

// Update implements CourseRepository
func (repository *CourseRepositoryImpl) Update(entity entity.Course) (*entity.Course, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{db}
}
