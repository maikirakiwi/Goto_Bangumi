package models

import (
	"encoding/json"
)

// Cursed but too lazy
func InitConfigModel() map[string]interface{} {
	default_config := `{
		"program": {
			"sleep_time": 7200,
			"times": 20,
			"webui_port": 7892,
			"data_version": 4
		},
		"downloader": {
			"type": "qbittorrent",
			"host": "127.0.0.1:8080",
			"username": "admin",
			"password": "adminadmin",
			"path": "/downloads/Bangumi",
			"ssl": false
		},
		"rss_parser": {
			"enable": true,
			"type": "mikan",
			"custom_url": "mikanani.me",
			"token": "",
			"enable_tmdb": false,
			"filter": [
				"720",
				"\\d+-\\d+"
			],
			"language": "zh"
		},
		"bangumi_manage": {
			"enable": true,
			"eps_complete": false,
			"rename_method": "pn",
			"group_tag": false,
			"remove_bad_torrent": false
		},
		"log": {
			"debug_enable": false
		},
		"proxy": {
			"enable": false,
			"type": "http",
			"host": "",
			"port": 1080,
			"username": "",
			"password": ""
		},
		"notification": {
			"enable": false,
			"type": "telegram",
			"token": "",
			"chat_id": ""
		}
	}`
	var cfg map[string]interface{}
	json.Unmarshal([]byte(default_config), &cfg)

	return cfg
}
