package main

import (
	mysql "edtech-app/pkg/db/mysql"

	"github.com/gin-gonic/gin"

	admin "edtech-app/internal/admin/injector"
	cart "edtech-app/internal/cart/injector"
	classRoom "edtech-app/internal/class_room/injector"
	dashboard "edtech-app/internal/dashboard/injector"
	discount "edtech-app/internal/discount/injector"
	oauth "edtech-app/internal/oauth/injector"
	order "edtech-app/internal/order/injector"
	product "edtech-app/internal/product/injector"
	productCategory "edtech-app/internal/product_category/injector"
	profile "edtech-app/internal/profile/injector"
	register "edtech-app/internal/register/injector"
	user "edtech-app/internal/user/injector"
	webhook "edtech-app/internal/webhook/injector"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)
	db := mysql.DB()

	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitiliazedService(db).Route(&r.RouterGroup)
	cart.InitiliazedService(db).Route(&r.RouterGroup)
	discount.InitializedService(db).Route(&r.RouterGroup)
	order.InitializedService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)
	classRoom.InitializedService(db).Route(&r.RouterGroup)
	dashboard.InitializedService(db).Route(&r.RouterGroup)
	user.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
