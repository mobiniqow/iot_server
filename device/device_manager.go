package device

import (
	"errors"
	"github.com/go-kit/kit/log"
	"sync"
)

type Manager struct {
	Devices []Device
	Logger  log.Logger
}

var instance *Manager
var one sync.Once

func GetInstanceManager(logger log.Logger) *Manager {
	if instance == nil {
		one.Do(func() {
			instance = &Manager{
				Devices: make([]Device, 0),
				Logger:  logger,
			}
		})
	}
	return instance
}
func (dm *Manager) Add(device Device) error {
	for _, element := range dm.Devices {
		if element.ClientID == device.ClientID || element.Conn == device.Conn {
			return errors.New("Device is already exist")
		}
	}
	dm.Logger.Log("Adding new device with clientID: ", device.ClientID)
	dm.Devices = append(dm.Devices, device)
	return nil
}

func (dm *Manager) Delete(device Device) error {
	var selectedDevice int = -1
	for index, element := range dm.Devices {
		if element.ClientID == device.ClientID || element.Conn == device.Conn || element.DeviceID == device.DeviceID {
			selectedDevice = index
			break
		}
	}
	if selectedDevice == -1 {
		return errors.New("Device not exist")
	}
	dm.Logger.Log("Delete device with client_id", device.ClientID)
	dm.Devices = append(dm.Devices[:selectedDevice], dm.Devices[selectedDevice+1:]...)

	return nil
}

func (dm *Manager) Update(device Device) error {
	var isExist bool = false
	for _, element := range dm.Devices {
		if element.ClientID == device.ClientID {
			element = device
			isExist = true
			break
		}
	}
	if !isExist {
		return errors.New("device is not exist")
	}
	dm.Logger.Log("Updated device with client_id %s", device.ClientID)
	return nil
}

func (dm *Manager) Get(device Device) (Device, error) {
	for _, element := range dm.Devices {
		if element.ClientID == device.ClientID || element.Conn == device.Conn || element.DeviceID == device.DeviceID {

			return element, nil
		}
	}
	dm.Logger.Log("Get device with client_id %s", device.ClientID)
	return device, errors.New("device not exist")
}
