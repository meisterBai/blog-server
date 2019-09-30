package main

import (
	"context"
	"flag"
	"github.com/valyala/tcplisten"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/iris-contrib/middleware/cors"
	"server/routers"
	"time"
	"server/config"
	"server/storage/orm"
)

var (
	routerFlag = flag.String("router", "", "specify router to use")
	routerDict map[string]routers.IRouter
)

func registerRouters(a *iris.Application) {
	routerDict = map[string]routers.IRouter{
		"article": routers.NewArticleRouter(a),
	}
}

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())


	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// 关闭所有主机
		app.Shutdown(ctx)
	})

	orm.InitDBServcie(config.Conf.Get("mysql.dial").(string))


	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	app.Use(crs)


	registerRouters(app)


	if *routerFlag == "" {
		for k, v := range routerDict {
			v.RegisterHandlers(k)
		}
	} else {
		if r, ok := routerDict[*routerFlag]; ok {
			r.RegisterHandlers(*routerFlag)
		} else {
			panic("router flag error: cannot find router " + *routerFlag)
		}
	}


	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}

	l, err := listenerCfg.NewListener("tcp4", config.Conf.Get("app.addr").(string))
	if err != nil {
		app.Logger().Fatal(err)
	}

	app.Run(iris.Listener(l))
}