package v1alpha1

// CaptureStatus holds CDC server information
type CaptureStatus struct {
	ID      string `json:"id"`
	Address string `json:"address"`
}

// CaptureInfo holds CDC captured info
type CaptureInfo struct {
	ID            string `json:"id"`
	IsOwner       bool   `json:"is-owner"`
	AdvertiseAddr string `json:"address"`
}

// REMOVE duplicate definitions for SinkSpec and SecretKeySelector if present
