package tracking

import (
	"github.com/micro/go-micro/v3/api"
	"net/http"
)

type TrackingHandler struct {
	// Tracking service instance
	trackingService TrackingService
}

func NewTrackingHandler(service TrackingService) *TrackingHandler {
	return &TrackingHandler{
		trackingService: service,
	}
}

func (h *TrackingHandler) RealTimeTracking(ctx api.Context, req api.Request, rsp api.Response) error {
	// Implement the handler to update the real-time tracking data for a shipment.
	// Extract data from the request and call the tracking service to update the tracking data.
	// Handle error and return appropriate response.
	trackingID := req.Get("tracking_id").String()
	status := req.Get("status").String()
	location := req.Get("location").String()

	if err := h.trackingService.UpdateTrackingData(trackingID, status, location); err != nil {
		return rsp.Write(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to update tracking data",
		})
	}

	return rsp.Write(http.StatusOK, map[string]interface{}{
		"message": "Real-time tracking data updated successfully",
	})
}

func (h *TrackingHandler) GetTrackingData(ctx api.Context, req api.Request, rsp api.Response) error {
	// Implement the handler to retrieve the real-time tracking data for a shipment.
	// Extract data from the request and call the tracking service to get the tracking data.
	// Handle error and return appropriate response.
	trackingID := req.Get("tracking_id").String()

	trackingData, err := h.trackingService.GetTrackingData(trackingID)
	if err != nil {
		return rsp.Write(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to get tracking data",
		})
	}

	return rsp.Write(http.StatusOK, trackingData)
}
