/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:55 2019-09-12
 */
package register

import (
	"doratv/config"
	"doratv/web/routers"
	"github.com/kataras/iris"
)

func IrisRegister() *iris.Application {
	var app *iris.Application
	if config.Basis.App.Debug {
		app = iris.New()
	}else {
		app = iris.Default()
	}

	// 注册路由
	routers.Register(app)

	return app
}

