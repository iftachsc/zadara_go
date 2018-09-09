package zadara

import (
	"time"
)

type MirrorMember struct {
	Provider      string `json:"provider"`
	VpsaName      string `json:"vpsa_display_name"`
	CgDisplayName string `json:"cg_display_name"`
	CgName        string `json:"cg_name"`
	PoolName      string `json:"pool_display_name"`
}

type Mirror struct {
	Name            string       `json:"job_name"`
	DisplayName     string       `json:"job_display_name"`
	CgName          string       `json:"cg_name"`
	Status          string       `json:"status"`
	Source          MirrorMember `json:"src"`
	Destination     MirrorMember `json:"dst"`
	Rpo             time.Time    `json:"rpo"`
	DataType        string       `json:"data_type"`
	Thin            string       `json:"thin"`
	VirtualCapacity float64      `json:"virtual_capacity"`
	Encryption      string       `json:"encryption"`
	WanOptimization string       `json:"wan_optimization"`
	AverageSyncRate string       `json:"avg_sync_rate_kb_sec"`
	Created         time.Time    `json:"created_at"`
	Modified        time.Time    `json:"created_at"`
}

type MirrorResponse struct {
	Status  int      `json:"status"`
	Mirrors []Mirror `json:"vpsa_mirror_jobs"`
	Count   int      `json:"count"`
}

type RootMirrorResponse struct {
	Response MirrorResponse `json:"response"`
}
