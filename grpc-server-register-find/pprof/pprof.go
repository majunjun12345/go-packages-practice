package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
)

const (
	defaultPrefix string = "/debug"
)

// Router for http handler
func Router(router *gin.Engine) {

	statsvizAPI := router.Group(defaultPrefix)
	{
		statsvizAPI.GET("/statsviz/", pprofHandler(statsviz.Index.ServeHTTP))
		statsvizAPI.GET("/statsviz/ws", pprofHandler(statsviz.Ws))
	}

	pprofAPI := router.Group(defaultPrefix)
	{
		pprofAPI.GET("/", pprofHandler(pprof.Index))
		pprofAPI.GET("/cmdline", pprofHandler(pprof.Cmdline))
		pprofAPI.GET("/profile", pprofHandler(pprof.Profile))
		pprofAPI.POST("/symbol", pprofHandler(pprof.Symbol))
		pprofAPI.GET("/symbol", pprofHandler(pprof.Symbol))
		pprofAPI.GET("/trace", pprofHandler(pprof.Trace))
		pprofAPI.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		pprofAPI.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		pprofAPI.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		pprofAPI.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		pprofAPI.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		pprofAPI.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
		return
	}
}
