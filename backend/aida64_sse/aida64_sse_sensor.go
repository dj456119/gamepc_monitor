package aida64_sse_sensor

import (
	"errors"
	"gamepc_monitor/backend"
	"gamepc_monitor/dlog"
	"strings"

	"github.com/donovanhide/eventsource"
)

type Aida64SSESensor struct {
	Address        string
	Filters        map[string]string
	SensorDataList []backend.SensorData
	InitFlag       bool
}

func (aida64SSESensor *Aida64SSESensor) GetData() ([]backend.SensorData, error) {
	if !aida64SSESensor.InitFlag {
		go aida64SSESensor.UpdateSensorData()
		aida64SSESensor.InitFlag = true
	}
	return aida64SSESensor.SensorDataList, nil
}

func (aida64SSESensor *Aida64SSESensor) UpdateSensorData() {
	sseURL := aida64SSESensor.Address + "/sse"

	stream, err := eventsource.Subscribe(sseURL, "")
	if err != nil {
		dlog.Error("Failed to subscribe to ", sseURL, err)
	}

	for ev := range stream.Events {
		data := ev.Data()
		aida64SSESensor.SensorDataList, err = DataToSensorDataList(data)
		if err != nil {
			dlog.Error(err)
			aida64SSESensor.InitFlag = false
			break
		}
	}
}

func DataToSensorDataList(data string) ([]backend.SensorData, error) {
	//data: Page0|{|}Simple1|CPUUsage 11%{|}Simple2|MemUsage 51%{|}Simple3|GPUUsage 1%{|}Simple4|MemUsed 33021 MB{|}Simple5|GPUMemUsed 3433 MB{|}Simple6|CPUTemp 51C{|}Simple7|GPUTemp 41C{|}Simple8|CPUPower N/A W{|}Simple9|GPUPower 36.24 W{|}
	q := strings.Split(data, "{|}")
	result := []backend.SensorData{}
	for i := range q {
		if i == 0 {
			if strings.Contains(q[i], "Reload") {
				dlog.Debug("Reload")
				return []backend.SensorData{}, errors.New("Reload")
			}
		}
		f := strings.Split(q[i], "|")
		if len(f) != 2 {
			dlog.Warn(f)
			continue
		}
		z := strings.Split(f[1], " ")
		sensorData := backend.SensorData{}
		if len(z) == 2 {
			sensorData.Name = z[0]
			sensorData.Value = z[1]
		}
		if len(z) == 3 {
			sensorData.Name = z[0]
			sensorData.Value = z[1] + z[2]
		}
		result = append(result, sensorData)
	}
	dlog.Debug(q, result)
	return result, nil
}
