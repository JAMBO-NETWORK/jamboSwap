// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"jamSwap/mvc/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/v1/saveFile", &controllers.UserController{}, "post:SaveFile")
	beego.Router("/v1/saveAdvertisement", &controllers.UserController{}, "post:SaveAdvertisement")
	beego.Router("/v1/hasAdvertisement", &controllers.UserController{}, "get:HasAdvertisement")
	beego.Router("/v1/list", &controllers.UserController{}, "get:List")
	beego.Router("/v1/addressList", &controllers.UserController{}, "get:AddressList")
	beego.Router("/v1/removeAdvertisement", &controllers.UserController{}, "post:RemoveAdvertisement")
	beego.Router("/v1/getTotalTvlAndJamPrice", &controllers.LiquidityController{}, "get:GetTotalTvlAndJamPrice")
	beego.Router("/v1/liquidityInfo", &controllers.LiquidityController{}, "get:LiquidityInfo")
	beego.Router("/v1/updateLiquidityInfo", &controllers.LiquidityController{}, "post:UpdateLiquidityInfo")
	beego.Router("/v1/hideAdvertisement", &controllers.UserController{}, "post:HideAdvertisement")
	beego.Router("/v1/updateSortNum", &controllers.UserController{}, "post:UpdateSortNum")
	beego.Router("/v1/openAdvertisement", &controllers.UserController{}, "post:OpenAdvertisement")

}
