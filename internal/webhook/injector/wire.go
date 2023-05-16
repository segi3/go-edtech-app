//go:build wireinject
// +build wireinject

package webhook

import (
	cartRepository "edtech-app/internal/cart/repository"
	cartUseCase "edtech-app/internal/cart/usecase"
	classRoomRepository "edtech-app/internal/class_room/repository"
	classRoomUseCase "edtech-app/internal/class_room/usecase"
	discountRepository "edtech-app/internal/discount/repository"
	discountUseCase "edtech-app/internal/discount/usecase"
	orderRepository "edtech-app/internal/order/repository"
	orderUseCase "edtech-app/internal/order/usecase"
	orderDetailRepository "edtech-app/internal/order_detail/repository"
	orderDetailUseCase "edtech-app/internal/order_detail/usecase"
	paymentUseCase "edtech-app/internal/payment/usecase"
	productRepository "edtech-app/internal/product/repository"
	productUseCase "edtech-app/internal/product/usecase"
	handler "edtech-app/internal/webhook/delivery/http"
	useCase "edtech-app/internal/webhook/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.WebhookHandler {
	wire.Build(
		handler.NewWebHookHandler,
		useCase.NewWebhookUseCase,
		classRoomRepository.NewClassRoomRepository,
		classRoomUseCase.NewClassRoomUseCase,
		orderRepository.NewOrderRepository,
		orderUseCase.NewOrderUseCase,
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
	)

	return &handler.WebhookHandler{}
}
