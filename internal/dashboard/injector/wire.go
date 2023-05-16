//go:build wireinject
// +build wireinject

package webhook

import (
	adminRepository "edtech-app/internal/admin/repository"
	adminUseCase "edtech-app/internal/admin/usecase"
	cartRepository "edtech-app/internal/cart/repository"
	cartUseCase "edtech-app/internal/cart/usecase"
	handler "edtech-app/internal/dashboard/delivery/http"
	useCase "edtech-app/internal/dashboard/usecase"
	discountRepository "edtech-app/internal/discount/repository"
	discountUseCase "edtech-app/internal/discount/usecase"
	orderRepository "edtech-app/internal/order/repository"
	orderUseCase "edtech-app/internal/order/usecase"
	orderDetailRepository "edtech-app/internal/order_detail/repository"
	orderDetailUseCase "edtech-app/internal/order_detail/usecase"
	paymentUseCase "edtech-app/internal/payment/usecase"
	productRepository "edtech-app/internal/product/repository"
	productUseCase "edtech-app/internal/product/usecase"
	userRepository "edtech-app/internal/user/repository"
	userUseCase "edtech-app/internal/user/usecase"
	fileUpload "edtech-app/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.DashboardHandler {
	wire.Build(
		handler.NewDashboardHandler,
		useCase.NewDashboardUseCase,
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
		adminRepository.NewAdminRepository,
		adminUseCase.NewAdminUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &handler.DashboardHandler{}
}
