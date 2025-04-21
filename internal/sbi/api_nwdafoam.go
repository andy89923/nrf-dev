package sbi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/nrf/internal/logger"
	"github.com/free5gc/nwdaf/pkg/components"
	"github.com/free5gc/openapi/models"
)

func (s *Server) getNwdafOamRoutes() []Route {
	return []Route{
		{
			Name:    "Health Check",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.String(http.StatusOK, "AMF NWDAF-OAM woking!")
			},
		},
		{
			Name:    "NfResourceGet",
			Method:  http.MethodGet,
			Pattern: "/nf-resource",
			APIFunc: s.NrfOamNfResourceGet,
		},
		{
			Name:    "NfLoadLevelAnalyticsNotification",
			Method:  http.MethodPost,
			Pattern: "/callback",
			APIFunc: s.NfLoadLevelAnalyticsNotification,
		},
	}
}

func (s *Server) NrfOamNfResourceGet(c *gin.Context) {
	nfResource, err := components.GetNfResouces(context.Background())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, *nfResource)
}

func (s *Server) NfLoadLevelAnalyticsNotification(c *gin.Context) {
	logger.SBILog.Infoln("Receive NfLoadLevelAnalyticsNotification")

	var notification []models.NnwdafEventsSubscriptionNotification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "Invalid JSON format",
			Detail: err.Error(),
		})
		return
	}
	if len(notification) == 0 {
		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Status: http.StatusBadRequest,
			Cause:  "Notification is empty",
			Detail: "Empty notification",
		})
		return
	}
	logger.SBILog.Infoln("Notification:", notification)

	// TODO: process notification

	c.Status(http.StatusNoContent)
}
