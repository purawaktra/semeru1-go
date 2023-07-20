package modules

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin/v2"
)

type Semeru1Router struct {
	engine *gin.Engine
	rh     Semeru1RequestHandler
}

func CreateSemeru1Router(engine *gin.Engine, rh Semeru1RequestHandler) Semeru1Router {
	return Semeru1Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Semeru1Router) Init(path string) {
	pathGroup := r.engine.Group(path, apmgin.Middleware(r.engine))
	pathGroup.POST("/select/city/id", r.rh.SelectCityById)
	pathGroup.POST("/select/city/name", r.rh.SelectCityByName)
	pathGroup.POST("/select/city/province", r.rh.SelectCityByProvince)
	pathGroup.POST("/select/city/all", r.rh.SelectAllCity)

	pathGroup.POST("/select/province/id", r.rh.SelectProvinceById)
	pathGroup.POST("/select/province/all", r.rh.SelectAllProvince)
}
