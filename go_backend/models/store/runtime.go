package store

func InitConfigModel() *ConfigModel {
	return &ConfigModel{
		Id: 1,
		Program: ProgramConfig{
			RssTime:     900,
			RenameTime:  60,
			WebuiPort:   7892,
			DataVersion: 4,
		},
		Downloader: DownloaderConfig{
			Type:     "qbittorrent",
			Host:     "127.0.0.1:8080",
			Username: "admin",
			Password: "adminadmin",
			Path:     "/downloads/Bangumi",
			Ssl:      false,
		},
		RssParser: RssParserConfig{
			Enable:     true,
			Type:       "mikan",
			CustomUrl:  "mikanani.me",
			Token:      "",
			EnableTmdb: false,
			Filter: []string{
				"720",
				"\\d+-\\d+",
			},
			Language: "zh",
		},
		BangumiManage: BangumiManageConfig{
			Enable:           true,
			EpsComplete:      false,
			RenameMethod:     "pn",
			GroupTag:         false,
			RemoveBadTorrent: false,
		},
		Log: LogConfig{
			DebugEnable: false,
		},
		Proxy: ProxyConfig{
			Enable:   false,
			Type:     "http",
			Host:     "",
			Port:     1080,
			Username: "",
			Password: "",
		},
		Notification: NotificationConfig{
			Enable: false,
			Type:   "telegram",
			Token:  "",
			ChatId: "",
		},
	}
}

type ProgramConfig struct {
	RssTime     int64 `json:"rss_time"`
	RenameTime  int64 `json:"rename_time"`
	WebuiPort   int64 `json:"webui_port"`
	DataVersion int64 `json:"data_version"`
}

type DownloaderConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
	Ssl      bool   `json:"ssl"`
}

type RssParserConfig struct {
	Enable     bool     `json:"enable"`
	Type       string   `json:"type"`
	CustomUrl  string   `json:"custom_url"`
	Token      string   `json:"token"`
	EnableTmdb bool     `json:"enable_tmdb"`
	Filter     []string `json:"filter"`
	Language   string   `json:"language"`
}

type BangumiManageConfig struct {
	Enable           bool   `json:"enable"`
	EpsComplete      bool   `json:"eps_complete"`
	RenameMethod     string `json:"rename_method"`
	GroupTag         bool   `json:"group_tag"`
	RemoveBadTorrent bool   `json:"remove_bad_torrent"`
}

type LogConfig struct {
	DebugEnable bool `json:"debug_enable"`
}

type ProxyConfig struct {
	Enable   bool   `json:"enable"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type NotificationConfig struct {
	Enable bool   `json:"enable"`
	Type   string `json:"type"`
	Token  string `json:"token"`
	ChatId string `json:"chat_id"`
}
