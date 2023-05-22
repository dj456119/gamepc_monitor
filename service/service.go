package service

import (
	"gamepc_monitor/backend"
	aida64_sse_sensor "gamepc_monitor/backend/aida64_sse"
)

type Service struct {
	SensorCore backend.SensorCore
}

func NewService() Service {
	return Service{
		SensorCore: backend.SensorCore{
			Sensor: &aida64_sse_sensor.Aida64SSESensor{
				Address: "http://192.168.50.218:8085",
				Filters: backend.Filters,
			},
		},
	}
}

func (s Service) GetSensorModel() (backend.SensorModel, error) {
	return s.SensorCore.GetSensorModel()
}
