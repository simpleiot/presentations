---
title: Simple IoT
author: Cliff Brake
---

# IoT space is exploding right now

1. Cloud
1. Compute resources at edge
1. Connectivity
1. Connected devices

---

## Cloud

- Digital Ocean
- Linode
- Vultr
- AWS
- GCP

---

## Compute resources at edge

- Embedded Linux
- many SOC/SOM/SBC options
  - i.MX6 SOM -- dozens available
- \$10 Linux SOM

---

## Connectivty

- CAT-M plans (< \$5/mo with generous data allowance, 10's MiB)
- LoRaWAN
- BLE Mesh

---

## Connected devices

- ESP8266/ESP32 (couple \$ WiFi connectivity)
- LoRaWAN devices
- BLE Mesh devices

![](esp8266.png)

---

# Many problems to be solved

- HVAC
- machine monitoring
- infrastructure monitoring
- aggriculture
- many other forms of automation

---

# Data Structures -- 1st pass at device config

```go
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

```
