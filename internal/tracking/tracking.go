package tracking

type Tracking struct {
	// Define tracking related data models here
	// For example, you can have a struct to represent a shipment with relevant fields like tracking ID, status, location, etc.
}

type TrackingService interface {
	// Define tracking service methods here
	UpdateTrackingData(trackingID, status, location string) error
	GetTrackingData(trackingID string) (Tracking, error)
}
