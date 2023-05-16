//go:build wireinject
// +build wireinject

package order

import (
	cartRepository "edtech-app/internal/cart/repository"
	cartUseCase "edtech-app/internal/cart/usecase"
	discountRepository "edtech-app/internal/discount/repository"
	discountUseCase "edtech-app/internal/discount/usecase"
	handler "edtech-app/internal/order/delivery/http"
	repository "edtech-app/internal/order/repository"
	useCase "edtech-app/internal/order/usecase"
	orderDetailRepository "edtech-app/internal/order_detail/repository"
	orderDetailUseCase "edtech-app/internal/order_detail/usecase"
	paymentUseCase "edtech-app/internal/payment/usecase"
	productRepository "edtech-app/internal/product/repository"
	productUseCase "edtech-app/internal/product/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		useCase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
	)

	return &handler.OrderHandler{}
}
