package lesson

import (
	entity "edtech-app/internal/lesson/entity"
	"edtech-app/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type LessonRepository interface {
	FindAll(offset int, limit int) []entity.Lesson
	FindById(id int) (*entity.Lesson, error)
	FindExist(id int) (bool, error)
	Create(entity entity.Lesson) (*entity.Lesson, error)
	Update(entity entity.Lesson) (*entity.Lesson, error)
	Delete(entity entity.Lesson) error
}

type LessonRepositoryImpl struct {
	db *gorm.DB
}

// FindExist implements LessonRepository
func (repository *LessonRepositoryImpl) FindExist(id int) (bool, error) {
	var lesson entity.Lesson

	res := repository.db.Model(&entity.Lesson{}).
		Where("id = ? ", id).
		First(&lesson)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// user does not exists
		return false, nil
	} else if res.Error != nil {
		// some other problem
		return true, res.Error
	}

	return true, nil // base return set to exist behaviour
}

// Create implements LessonRepository
func (repository *LessonRepositoryImpl) Create(entity entity.Lesson) (*entity.Lesson, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements LessonRepository
func (repository *LessonRepositoryImpl) Delete(entity entity.Lesson) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements LessonRepository
func (repository *LessonRepositoryImpl) FindAll(offset int, limit int) []entity.Lesson {
	var lessons []entity.Lesson

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&lessons)

	return lessons
}

// FindById implements LessonRepository
func (repository *LessonRepositoryImpl) FindById(id int) (*entity.Lesson, error) {
	var lesson entity.Lesson

	if err := repository.db.First(&lesson, id).Error; err != nil {
		return nil, err
	}

	return &lesson, nil
}

// Update implements LessonRepository
func (repository *LessonRepositoryImpl) Update(entity entity.Lesson) (*entity.Lesson, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &LessonRepositoryImpl{db}
}
