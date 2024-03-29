package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	driverpostgres "github.com/moura95/go-ddd/internal/domain/driver/postgres"
	vehiclepostgres "github.com/moura95/go-ddd/internal/domain/vehicle/postgres"
	driverrouter "github.com/moura95/go-ddd/internal/infra/http/gin/router/driver"
	vehiclerouter "github.com/moura95/go-ddd/internal/infra/http/gin/router/vehicle"
	"github.com/moura95/go-ddd/internal/service/driver"
	"github.com/moura95/go-ddd/internal/service/vehicle"
	"go.uber.org/zap"
)

func (s *Server) createRoutesV1(router *gin.Engine, log *zap.SugaredLogger) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	routes := router.Group("/")
	// Instance Driver Repository Postgres
	driverRepository := driverpostgres.NewDriverRepository(s.store, log)
	// Instance Driver Service
	driverService := driver.NewDriverService(s.store, driverRepository, *s.config, log)

	// Instance VehicleRouter Repository
	vehicleRepository := vehiclepostgres.NewVehicleRepository(s.store, log)
	// Instance VehicleRouter Service
	vehicleService := vehicle.NewVehicleService(s.store, vehicleRepository, *s.config, log)

	vehiclerouter.NewVehicleRouter(vehicleService, log).SetupVehicleRoute(routes)
	driverrouter.NewDriverRouter(driverService, log).SetupDriverRoute(routes)
}
