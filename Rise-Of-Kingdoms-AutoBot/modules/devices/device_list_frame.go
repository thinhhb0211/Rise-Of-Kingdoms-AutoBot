package devices

import "Rise-Of-Kingdoms-AutoBot/pkg/config"

type DeviceListFrame struct {
	config *config.Configuration
}

func NewDeviceListFrame(config *config.Configuration) *DeviceListFrame {
	return &DeviceListFrame{
		config: config,
	}
}

