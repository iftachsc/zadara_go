package zadara

import (
	"time"
)

var (
	mockedMirror1 Mirror = Mirror{
		Name:        "srcjvpsa-00000001",
		Status:      "idle",
		Source:      MirrorMember{Provider: "openstack", VpsaName: "myvpsa"},
		Destination: MirrorMember{Provider: "openstack", VpsaName: "myvpsa2"},
		Created:     time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC),
		Modified:    time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC),
	}
)

const (
	RootMirrorResponseJson = `{
		"response": {
		  "status": 0,
		  "vpsa_mirror_jobs": [
			{
			  "job_name": "srcjvpsa-00000001",
			  "job_display_name": "mirror",
			  "src": {
				"provider": "openstack",
				"vpsa_display_name": "myvpsa",
				"cg_display_name": "volume",
				"cg_name": "cg-00000001",
				"pool_display_name": "poolrenamed"
			  },
			  "dst": {
				"provider": "openstack",
				"vpsa_display_name": "myvpsa2",
				"cg_display_name": "remotevolume",
				"cg_name": null,
				"pool_display_name": "pool"
			  },
			  "status": "Idle",
			  "rpo": "2013-09-27 02:10:20",
			  "wan_optimization": "YES",
			  "avg_sync_rate_kb_sec": "0",
				"created_at": "2013-09-27 02:09:38 UTC",
			  "modified_at": "2013-09-27 02:10:21 UTC"
			}
		  ],
		  "count": 1
		}
	  }`
)
