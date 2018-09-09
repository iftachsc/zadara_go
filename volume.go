package zadara

type Volume struct {
	Name            string  `json:"name"`
	DisplauName     string  `json:"display_name"`
	CgName          string  `json:"cg_name"`
	PoolName        string  `json:"pool_name"`
	PoolDisplayName string  `json:"pool_display_name"`
	DataType        string  `json:"data_type"`
	Thin            string  `json:"thin"`
	VirtualCapacity float64 `json:"virtual_capacity"`
	Encryption      string  `json:"encryption"`
}

func (v Volume) GetName() string {
	return v.Name
}

type VolumeResponse struct {
	Status  int      `json:"status"`
	Volumes []Volume `json:"volumes"`
	Count   int      `json:"count"`
}

type RootVolumeResponse struct {
	Response VolumeResponse `json:"response"`
}
