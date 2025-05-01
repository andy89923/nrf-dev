package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/nrf/pkg/configurations"
	"github.com/free5gc/openapi/models"
)

func (s *Server) getDynamicConfigRoutes() []Route {
	return []Route{
		{
			Name:    "Dynamic Config",
			Method:  http.MethodPost,
			Pattern: "/post-config",
			APIFunc: s.ConfigPost,
		},
	}
}

func (s *Server) ConfigPost(c *gin.Context) {
	var config configurations.DynamicConfig

	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "Invalid JSON format",
			Detail: err.Error(),
		})
		return
	}

	if err := s.Processor().ConfigUpdate(&config); err != nil {
		c.JSON(http.StatusInternalServerError, models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "Failed to update dynamic config",
			Detail: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, config)
}
