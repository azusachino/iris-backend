package router

import (
	"iris/controller"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(app *gin.Engine) {
	ac := new(controller.ArticleController)
	doApply(app, ac)
}

func doApply(app *gin.Engine, handlers ...controller.Handler) {
	for _, handler := range handlers {
		app.Handle(handler.Method(), handler.Path(), handler.HandlerFunc)
	}
}