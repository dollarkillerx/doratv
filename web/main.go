/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:28 2019-09-12
 */
package main

import (
	"doratv/config"
	"doratv/utils/elog"
	"doratv/web/register"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/kataras/iris"
)

func main() {
	// 设置报警
	defer func() {
		if err := recover(); err != nil {
			clog.Println(err)
			e := elog.Elog("梦蓝字幕组ERROR", "兔兔的  服务器宕机了")
			if e != nil {
				clog.PrintWa(e.Error())
				elog.Elog("梦蓝字幕组ERROR", "兔兔的  服务器宕机了")
				panic("error")
			}
			panic("error")
		}
	}()

	app := register.IrisRegister()

	err := app.Run(iris.Addr(config.Basis.App.Host), iris.WithCharset("UTF-8"))
	if err != nil {
		clog.PrintEr("兔兔服务器宕掉了")
	}
}

