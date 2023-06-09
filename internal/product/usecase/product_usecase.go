package product

import (
	dto "edtech-app/internal/product/dto"
	entity "edtech-app/internal/product/entity"
	repository "edtech-app/internal/product/repository"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
)

type ProductUseCase interface {
	FindAll(offset int, limit int) []entity.Product
	FindById(id int) (*entity.Product, error)
	FindExist(id int) (bool, error)
	Count() int
	Create(dto dto.ProductRequestBody) (*entity.Product, error)
	Update(id int, dto dto.ProductRequestBody) (*entity.Product, error)
	Delete(id int) error
}

type ProductUseCaseImpl struct {
	repository repository.ProductRepository
	fileUpload fileUpload.FileUpload
}

// FindExist implements ProductUseCase
func (usecase *ProductUseCaseImpl) FindExist(id int) (bool, error) {
	return usecase.repository.FindExist(id)
}

// Count implements ProductUseCase
func (usecase *ProductUseCaseImpl) Count() int {
	return usecase.repository.Count()
}

// Create implements ProductUseCase
func (usecase *ProductUseCaseImpl) Create(dto dto.ProductRequestBody) (*entity.Product, error) {
	dataProduct := entity.Product{
		ProductCategoryID: dto.ProductCategoryID,
		Title:             dto.Title,
		Description:       dto.Description,
		Price:             dto.Price,
		CreatedByID:       dto.CreatedBy,
	}

	// Upload thumbnail image
	if dto.Thumbnail != nil {
		image, err := usecase.fileUpload.Upload(*dto.Thumbnail)

		if err != nil {
			return nil, err
		}

		if image != nil {
			dataProduct.Thumbnail = image
		}
	}

	// Upload video
	if dto.Video != nil {
		video, err := usecase.fileUpload.Upload(*dto.Video)

		if err != nil {
			return nil, err
		}

		if video != nil {
			dataProduct.Video = video
		}
	}

	// Kita akan memanggil repository save
	product, err := usecase.repository.Create(dataProduct)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// Delete implements ProductUseCase
func (usecase *ProductUseCaseImpl) Delete(id int) error {
	// Cari data product berdasarkan id
	product, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*product)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductUseCase
func (usecase *ProductUseCaseImpl) FindAll(offset int, limit int) []entity.Product {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements ProductUseCase
func (usecase *ProductUseCaseImpl) FindById(id int) (*entity.Product, error) {
	return usecase.repository.FindById(id)
}

// Update implements ProductUseCase
func (usecase *ProductUseCaseImpl) Update(id int, dto dto.ProductRequestBody) (*entity.Product, error) {
	// cari data product berdasarkan id
	product, err := usecase.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	product.Title = dto.Title
	product.Description = dto.Description
	product.Price = dto.Price
	product.UpdatedByID = &dto.UpdatedBy

	// Jika terdapat update file image
	if dto.Thumbnail != nil {
		image, err := usecase.fileUpload.Upload(*dto.Thumbnail)

		if err != nil {
			return nil, err
		}

		if product.Thumbnail != nil {
			// Delete image
			_, err := usecase.fileUpload.Delete(*product.Thumbnail)

			if err != nil {
				return nil, err
			}
		}

		product.Thumbnail = image
	}

	// Jika terdapat update file video
	if dto.Video != nil {
		video, err := usecase.fileUpload.Upload(*dto.Video)

		if err != nil {
			return nil, err
		}

		if product.Video != nil {
			// Delete image
			_, err := usecase.fileUpload.Delete(*product.Video)

			if err != nil {
				return nil, err
			}
		}

		product.Video = video
	}

	updateProduct, err := usecase.repository.Update(*product)

	if err != nil {
		return nil, err
	}

	return updateProduct, nil

}

func NewProductUseCase(repository repository.ProductRepository, fileUpload fileUpload.FileUpload) ProductUseCase {
	return &ProductUseCaseImpl{repository, fileUpload}
}
