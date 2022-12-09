package routers

import (
	"ginshop/controllers/front"
	"ginshop/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", front.DefaultController{}.Index)
		defaultRouters.GET("/category:id", front.ProductController{}.Category)
		defaultRouters.GET("/detail", front.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", front.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", front.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", front.CartController{}.AddCart)

		defaultRouters.GET("/cart/successTip", front.CartController{}.AddCartSuccess)
		defaultRouters.GET("/cart/decCart", front.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", front.CartController{}.IncCart)

		defaultRouters.GET("/cart/changeOneCart", front.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", front.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", front.CartController{}.DelCart)

		defaultRouters.GET("/pass/login", front.PassController{}.Login)
		defaultRouters.GET("/pass/captcha", front.PassController{}.Captcha)

		defaultRouters.GET("/pass/registerStep1", front.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", front.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", front.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", front.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", front.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", front.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", front.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", front.PassController{}.LoginOut)
		//判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, front.BuyController{}.Checkout) //判断用户权限
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, front.BuyController{}.DoCheckout)
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, front.BuyController{}.Pay)
		defaultRouters.GET("/buy/orderPayStatus", middlewares.InitUserAuthMiddleware, front.BuyController{}.OrderPayStatus)

		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, front.AddressController{}.AddAddress)
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, front.AddressController{}.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, front.AddressController{}.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, front.AddressController{}.GetOneAddressList)

		defaultRouters.GET("/alipay", middlewares.InitUserAuthMiddleware, front.AlipayController{}.Alipay)
		defaultRouters.POST("/alipayNotify", front.AlipayController{}.AlipayNotify)
		defaultRouters.GET("/alipayReturn", middlewares.InitUserAuthMiddleware, front.AlipayController{}.AlipayReturn)

		defaultRouters.GET("/wxpay", middlewares.InitUserAuthMiddleware, front.WxpayController{}.Wxpay)
		defaultRouters.POST("/wxpay/notify", front.WxpayController{}.WxpayNotify)

		defaultRouters.GET("/user", middlewares.InitUserAuthMiddleware, front.UserController{}.Index)
		defaultRouters.GET("/user/order", middlewares.InitUserAuthMiddleware, front.UserController{}.OrderList)
		defaultRouters.GET("/user/orderinfo", middlewares.InitUserAuthMiddleware, front.UserController{}.OrderInfo)
	}

}
