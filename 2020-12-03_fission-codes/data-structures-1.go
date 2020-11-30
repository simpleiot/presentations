package main

import "time"

// Device represents the state of a device
// The config is typically updated by the portal/UI, and the
// State is updated by the device. Keeping these datastructures
// separate reduces the possibility that one update will step
// on another.
type Device struct {
	ID            string        `json:"id" boltholdKey:"ID"`
	Config        DeviceConfig  `json:"config"`
	State         DeviceState   `json:"state"`
	CmdPending    bool          `json:"cmdPending"`
	SwUpdateState SwUpdateState `json:"swUpdateState"`
	Groups        []uuid.UUID   `json:"groups"`
	Rules         []uuid.UUID   `json:"rules"`
}

// DeviceState represents information about a device that is
// collected, vs set by user.
type DeviceState struct {
	Version  DeviceVersion `json:"version"`
	Ios      []Sample      `json:"ios"`
	LastComm time.Time     `json:"lastComm"`
	SysState SysState      `json:"sysState"`
}

// SwUpdateState represents the state of an update
type SwUpdateState struct {
	Running     bool   `json:"running"`
	Error       string `json:"error"`
	PercentDone int    `json:"percentDone"`
}

// Sample represents a value in time and should include data that may be
// graphed.
type Sample struct {
	// Type of sample (voltage, current, key, etc)
	Type string `json:"type,omitempty"`

	// ID of the sensor that provided the sample
	ID string `json:"id,omitempty"`

	// Average OR
	// Instantaneous analog or digital value of the sample.
	// 0 and 1 are used to represent digital values
	Value float64 `json:"value,omitempty"`

	// statistical values that may be calculated
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`

	// Time the sample was taken
	Time time.Time `json:"time,omitempty" boltholdKey:"Time" gob:"-"`

	// Duration over which the sample was taken
	Duration time.Duration `json:"duration,omitempty"`

	// Tags are additional attributes used to describe the sample
	// You might add things like friendly name, etc.
	Tags map[string]string `json:"tags,omitempty"`

	// Attributes are additional numerical values
	Attributes map[string]float64 `json:"attributes,omitempty"`
}
