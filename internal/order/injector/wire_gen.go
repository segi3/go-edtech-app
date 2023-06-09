// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package order

import (
	"gorm.io/gorm"
	"edtech-app/internal/cart/repository"
	cart2 "edtech-app/internal/cart/usecase"
	"edtech-app/internal/discount/repository"
	discount2 "edtech-app/internal/discount/usecase"
	"edtech-app/internal/order/delivery/http"
	order2 "edtech-app/internal/order/repository"
	order3 "edtech-app/internal/order/usecase"
	"edtech-app/internal/order_detail/repository"
	order_detail2 "edtech-app/internal/order_detail/usecase"
	"edtech-app/internal/payment/usecase"
	"edtech-app/internal/product/repository"
	product2 "edtech-app/internal/product/usecase"
	"edtech-app/pkg/fileupload/cloudinary"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *order.OrderHandler {
	orderRepository := order2.NewOrderRepository(db)
	cartRepository := cart.NewCartRepository(db)
	cartUseCase := cart2.NewCartUseCase(cartRepository)
	discountRepository := discount.NewDiscountRepository(db)
	discountUseCase := discount2.NewDiscountUseCase(discountRepository)
	productRepository := product.NewProductRepository(db)
	fileUpload := fileupload.NewFileUpload()
	productUseCase := product2.NewProductUseCase(productRepository, fileUpload)
	orderDetailRepository := order_detail.NewOrderDetailRepository(db)
	orderDetailUseCase := order_detail2.NewOrderDetailUseCase(orderDetailRepository)
	paymentUseCase := payment.NewPaymentUseCase()
	orderUseCase := order3.NewOrderUseCase(orderRepository, cartUseCase, discountUseCase, productUseCase, orderDetailUseCase, paymentUseCase)
	orderHandler := order.NewOrderHandler(orderUseCase)
	return orderHandler
}
