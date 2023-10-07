package models

type Config struct {
	Program       map[string]interface{} `json:"program"`
	Downloader    map[string]interface{} `json:"downloader"`
	RssParser     map[string]interface{} `json:"rss_parser"`
	BangumiManage map[string]interface{} `json:"bangumi_manage"`
	Log           map[string]interface{} `json:"log"`
	Proxy         map[string]interface{} `json:"proxy"`
	Notification  map[string]interface{} `json:"notification"`
}

func InitConfigModel() map[string]interface{} {
	return map[string]interface{}{
		"program": map[string]interface{}{
			"sleep_time":   7200,
			"times":        20,
			"webui_port":   7892,
			"data_version": 4,
		},
		"downloader": map[string]interface{}{
			"type":     "qbittorrent",
			"host":     "127.0.0.1:8080",
			"username": "admin",
			"password": "adminadmin",
			"path":     "/downloads/Bangumi",
			"ssl":      false,
		},
		"rss_parser": map[string]interface{}{
			"enable":      true,
			"type":        "mikan",
			"custom_url":  "mikanani.me",
			"token":       "",
			"enable_tmdb": false,
			"filter": []interface{}{
				"720",
				"\\d+-\\d+",
			},
			"language": "zh",
		},
		"bangumi_manage": map[string]interface{}{
			"enable":             true,
			"eps_complete":       false,
			"rename_method":      "pn",
			"group_tag":          false,
			"remove_bad_torrent": false,
		},
		"log": map[string]interface{}{
			"debug_enable": false,
		},
		"proxy": map[string]interface{}{
			"enable":   false,
			"type":     "http",
			"host":     "",
			"port":     1080,
			"username": "",
			"password": "",
		},
		"notification": map[string]interface{}{
			"enable":  false,
			"type":    "telegram",
			"token":   "",
			"chat_id": "",
		},
	}
}
