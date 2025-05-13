package modules

type AdbDevice struct {
	Serial string `json:"serial"`
}

func (d *AdbDevice) Connect() error {
	// Implement the connection logic here
	return nil
}

func (d *AdbDevice) Disconnect() error {
	// Implement the disconnection logic here
	return nil
}
