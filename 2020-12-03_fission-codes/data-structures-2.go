package main

import "time"

// Node represents the state of a device. UUID is recommended
// for ID to prevent collisions is distributed instances.
type Node struct {
	ID     string  `json:"id" boltholdKey:"ID"`
	Type   string  `json:"type"`
	Points []Point `json:"points"`
}

// Point is a flexible data structure that can be used to represent
// a sensor value or a configuration parameter.
// ID, Type, and Index uniquely identify a point in a device
type Point struct {
	// ID of the sensor that provided the point
	ID string `json:"id,omitempty"`

	// Type of point (voltage, current, key, etc)
	Type string `json:"type,omitempty" boltholdIndex:"Type"`

	// Index is used to specify a position in an array such as
	// which pump, temp sensor, etc.
	Index int `json:"index,omitempty"`

	// Time the point was taken
	Time time.Time `json:"time,omitempty" boltholdKey:"Time" gob:"-"`

	// Duration over which the point was taken. This is useful
	// for averaged values to know what time period the value applies
	// to.
	Duration time.Duration `json:"duration,omitempty"`

	// Average OR
	// Instantaneous analog or digital value of the point.
	// 0 and 1 are used to represent digital values
	Value float64 `json:"value,omitempty"`

	// Optional text value of the point for data that is best represented
	// as a string rather than a number.
	Text string `json:"text,omitempty"`

	// statistical values that may be calculated over the duration of the point
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}
