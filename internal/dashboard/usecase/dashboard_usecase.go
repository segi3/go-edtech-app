package dashboard

import (
	adminUseCase "edtech-app/internal/admin/usecase"
	dto "edtech-app/internal/dashboard/dto"
	orderUseCase "edtech-app/internal/order/usecase"
	productUseCase "edtech-app/internal/product/usecase"
	userUseCase "edtech-app/internal/user/usecase"
)

type DashboardUseCase interface {
	GetDataDashboard() dto.DashboardResponseBody
}

type DashboardUseCaseImpl struct {
	userUseCase    userUseCase.UserUseCase
	adminUseCase   adminUseCase.AdminUseCase
	orderUseCase   orderUseCase.OrderUseCase
	productUseCase productUseCase.ProductUseCase
}

// GetDataDashboard implements DashboardUseCase
func (usecase *DashboardUseCaseImpl) GetDataDashboard() dto.DashboardResponseBody {
	totalUser := usecase.userUseCase.Count()
	totalAdmin := usecase.adminUseCase.Count()
	totalOrder := usecase.orderUseCase.Count()
	totalProduct := usecase.productUseCase.Count()

	return dto.DashboardResponseBody{
		TotalUser:    int64(totalUser),
		TotalProduct: int64(totalAdmin),
		TotalOrder:   int64(totalOrder),
		TotalAdmin:   int64(totalProduct),
	}
}

func NewDashboardUseCase(
	userUseCase userUseCase.UserUseCase,
	adminUseCase adminUseCase.AdminUseCase,
	orderUseCase orderUseCase.OrderUseCase,
	productUseCase productUseCase.ProductUseCase,
) DashboardUseCase {
	return &DashboardUseCaseImpl{userUseCase, adminUseCase, orderUseCase, productUseCase}
}
