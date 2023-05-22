package backend

type SensorData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SensorCore struct {
	Sensor Sensor
}

func (sensorCore SensorCore) GetSensorModel() (SensorModel, error) {
	sensorDataList, err := sensorCore.Sensor.GetData()
	if err != nil {
		return SensorModel{}, err
	}
	sm := SensorDataListToSensorModel(sensorDataList)
	return sm, nil
}

type Sensor interface {
	GetData() ([]SensorData, error)
}

type SensorModel struct {
	CPUUsage        string `json:"CPUUsage"`
	MemUsage        string `json:"MemUsage"`
	GPUUsage        string `json:"GPUUsage"`
	MemUsed         string `json:"MemUsed"`
	MemTotal        string `json:"MemTotal"`
	GPUMemUsage     string `json:"GPUMemUsage"`
	GPUMemTotal     string `json:"GPUMemTotal"`
	GPUMemUsed      string `json:"GPUMemUsed"`
	CpuTemp         string `json:"CpuTemp"`
	GpuTemp         string `json:"GpuTemp"`
	CPUPower        string `json:"CPUPower"`
	GPUPower        string `json:"GPUPower"`
	FPS             string `json:"FPS"`
	GPUMemAvailable string `json:"GPUMemAvailable"`
}

func SensorDataListToSensorModel(sensorDataList []SensorData) SensorModel {
	sensorModel := SensorModel{}
	for _, sensorData := range sensorDataList {
		switch sensorData.Name {
		case "CPUUsage":
			sensorModel.CPUUsage = sensorData.Value
		case "MemUsage":
			sensorModel.MemUsage = sensorData.Value
		case "GPUUsage":
			sensorModel.GPUUsage = sensorData.Value
		case "MemUsed":
			sensorModel.MemUsed = sensorData.Value
		case "MemTotal":
			sensorModel.MemTotal = sensorData.Value
		case "GPUMemUsage":
			sensorModel.GPUMemUsage = sensorData.Value
		case "GPUMemTotal":
			sensorModel.GPUMemTotal = sensorData.Value
		case "GPUMemUsed":
			sensorModel.GPUMemUsed = sensorData.Value
		case "CPUTemp":
			sensorModel.CpuTemp = sensorData.Value
		case "GPUTemp":
			sensorModel.GpuTemp = sensorData.Value
		case "CPUPower":
			sensorModel.CPUPower = sensorData.Value
		case "GPUPower":
			sensorModel.GPUPower = sensorData.Value
		case "FPS":
			sensorModel.FPS = sensorData.Value
		case "GPUMemAvailable":
			sensorModel.GPUMemAvailable = sensorData.Value
		}
	}
	// sensorModel.GPUMemTotal = "24000MB"
	// sensorModel.MemTotal = "64000MB"
	return sensorModel
}

var Filters = map[string]string{
	"CPUUsage":    "CPUUsage",
	"MemUsage":    "MemUsage",
	"GPUUsage":    "GPUUsage",
	"MemUsed":     "MemUsed",
	"MemTotal":    "MemTotal",
	"GPUMemUsage": "GPUMemUsage",
	"GPUMemTotal": "GPUMemTotal",
	"GPUMemUsed":  "GPUMemUsed",
	"CpuTemp":     "CpuTemp",
	"GpuTemp":     "GpuTemp",
	"CPUPower":    "CPUPower",
	"GPUPower":    "GPUPower",
	"FPS":         "FPS",
}
