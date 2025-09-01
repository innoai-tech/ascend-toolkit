package v1

type ChipInfo struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Version string `json:"version"`
	NpuName string `json:"npu_name"`
}

type ChipIndex struct {
	PhysicID int32 `json:"physic_id"`
	LogicID  int32 `json:"logic_id"`
	CardID   int32 `json:"card_id"`
	DeviceID int32 `json:"device_id"`
}

type Chip struct {
	ChipIndex

	Info          ChipInfo `json:"info"`
	Memory        Memory   `json:"memory"`
	AICore        AICore   `json:"ai_core"`
	Vector        Vector   `json:"vector"`
	Overall       Overall  `json:"overall"`
	Voltage       float32  `json:"voltage"`
	Temperature   int32    `json:"temperature"`
	Power         float32  `json:"power"`
	Health        uint32   `json:"health"`
	NetworkHealth uint32   `json:"network_health"`
}

type AICore struct {
	Utilization      int    `json:"utilization"`
	CurrentFrequency uint32 `json:"current_frequency"`
	RatedFrequency   uint32 `json:"rated_frequency"`
}

type Vector struct {
	Utilization int `json:"utilization"`
}

type Overall struct {
	Utilization int `json:"utilization"`
}

type Memory struct {
	Size                uint64              `json:"size"`
	Used                uint64              `json:"used"`
	Frequency           uint32              `json:"frequency"`
	Utilization         uint32              `json:"utilization"`
	HighBandwidthMemory HighBandwidthMemory `json:"high_bandwidth,omitzero"`
}

type HighBandwidthMemory struct {
	Size                 uint64 `json:"size"`
	Used                 uint64 `json:"used"`
	Frequency            uint32 `json:"frequency"`
	Temperature          int32  `json:"temperature"`
	BandwidthUtilization uint32 `json:"bandwidth_utilization"`
}
