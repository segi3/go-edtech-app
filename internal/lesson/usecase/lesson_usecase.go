package lesson

import (
	dto "edtech-app/internal/lesson/dto"
	entity "edtech-app/internal/lesson/entity"
	repository "edtech-app/internal/lesson/repository"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"errors"
)

type LessonUseCase interface {
	FindAll(offset int, limit int) []entity.Lesson
	FindById(id int) (*entity.Lesson, error)
	FindExist(id int) (bool, error)
	Create(dto dto.ProductRequestBody) (*entity.Lesson, error)
	Update(id int, dto dto.ProductRequestBody) (*entity.Lesson, error)
	Delete(id int) error
}

type LessonUseCaseImpl struct {
	repository repository.LessonRepository
	fileUpload fileUpload.FileUpload
}

// FindExist implements LessonUseCase
func (usecase *LessonUseCaseImpl) FindExist(id int) (bool, error) {
	return usecase.repository.FindExist(id)
}

// Create implements LessonUseCase
func (usecase *LessonUseCaseImpl) Create(dto dto.ProductRequestBody) (*entity.Lesson, error) {
	lessonData := entity.Lesson{
		Title:       dto.Title,
		Description: dto.Description,
		TextContent: dto.TextContent,
		CreatedByID: dto.CreatedBy,
	}

	if dto.VideoContent == nil {
		return nil, errors.New("create must include video content")
	}

	// Upload video content
	videoUrl, err := usecase.fileUpload.Upload(*dto.VideoContent)

	if err != nil {
		return nil, err
	}

	if videoUrl != nil {
		lessonData.VideoContent = videoUrl
	} else {
		return nil, errors.New("Failed to get resource URL") // videoURL must not be empty
	}

	// save lesson data
	lesson, err := usecase.repository.Create(lessonData)

	if err != nil {
		return nil, err
	}

	return lesson, nil
}

// Delete implements LessonUseCase
func (usecase *LessonUseCaseImpl) Delete(id int) error {
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

// FindAll implements LessonUseCase
func (usecase *LessonUseCaseImpl) FindAll(offset int, limit int) []entity.Lesson {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements LessonUseCase
func (usecase *LessonUseCaseImpl) FindById(id int) (*entity.Lesson, error) {
	return usecase.repository.FindById(id)
}

// Update implements LessonUseCase
func (usecase *LessonUseCaseImpl) Update(id int, dto dto.ProductRequestBody) (*entity.Lesson, error) {
	// get lesson  by id
	lesson, err := usecase.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	lesson.Title = dto.Title
	lesson.Description = dto.Description
	lesson.TextContent = dto.TextContent
	lesson.UpdatedByID = &dto.UpdatedBy

	// if video content exists
	if dto.VideoContent != nil {
		video, err := usecase.fileUpload.Upload(*dto.VideoContent)

		if err != nil {
			return nil, err
		}

		if lesson.VideoContent != nil {
			// Delete image
			_, err := usecase.fileUpload.Delete(*lesson.VideoContent)

			if err != nil {
				return nil, err
			}
		}

		lesson.VideoContent = video
	}

	updatedLesson, err := usecase.repository.Update(*lesson)

	if err != nil {
		return nil, err
	}

	return updatedLesson, nil

}

func NewLessonUseCase(repository repository.LessonRepository, fileUpload fileUpload.FileUpload) LessonUseCase {
	return &LessonUseCaseImpl{repository, fileUpload}
}
