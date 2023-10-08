package models

import "github.com/ostafen/clover/v2/document"

type ConfigModel struct {
	Program       ProgramConfig       `json:"program" clover:"Program"`
	Downloader    DownloaderConfig    `json:"downloader" clover:"Downloader"`
	RssParser     RssParserConfig     `json:"rss_parser" clover:"RssParser"`
	BangumiManage BangumiManageConfig `json:"bangumi_manage" clover:"BangumiManage"`
	Log           LogConfig           `json:"log" clover:"Log"`
	Proxy         ProxyConfig         `json:"proxy" clover:"Proxy"`
	Notification  NotificationConfig  `json:"notification" clover:"Notification"`
}

type ProgramConfig struct {
	RssTime     int64 `json:"rss_time" clover:"RssTime"`
	RenameTime  int64 `json:"rename_time" clover:"RenameTime"`
	WebuiPort   int64 `json:"webui_port" clover:"WebuiPort"`
	DataVersion int64 `json:"data_version" clover:"DataVersion"`
}

type DownloaderConfig struct {
	Type     string `json:"type" clover:"Type"`
	Host     string `json:"host" clover:"Host"`
	Username string `json:"username" clover:"Username"`
	Password string `json:"password" clover:"Password"`
	Path     string `json:"path" clover:"Path"`
	Ssl      bool   `json:"ssl" clover:"Ssl"`
}

type RssParserConfig struct {
	Enable     bool          `json:"enable" clover:"Enable"`
	Type       string        `json:"type" clover:"Type"`
	CustomUrl  string        `json:"custom_url" clover:"CustomUrl"`
	Token      string        `json:"token" clover:"Token"`
	EnableTmdb bool          `json:"enable_tmdb" clover:"EnableTmdb"`
	Filter     []interface{} `json:"filter" clover:"Filter"`
	Language   string        `json:"language" clover:"Language"`
}

type BangumiManageConfig struct {
	Enable           bool   `json:"enable" clover:"Enable"`
	EpsComplete      bool   `json:"eps_complete" clover:"EpsComplete"`
	RenameMethod     string `json:"rename_method" clover:"RenameMethod"`
	GroupTag         bool   `json:"group_tag" clover:"GroupTag"`
	RemoveBadTorrent bool   `json:"remove_bad_torrent" clover:"RemoveBadTorrent"`
}

type LogConfig struct {
	DebugEnable bool `json:"debug_enable" clover:"DebugEnable"`
}

type ProxyConfig struct {
	Enable   bool   `json:"enable" clover:"Enable"`
	Type     string `json:"type" clover:"Type"`
	Host     string `json:"host" clover:"Host"`
	Port     int64  `json:"port" clover:"Port"`
	Username string `json:"username" clover:"Username"`
	Password string `json:"password" clover:"Password"`
}

type NotificationConfig struct {
	Enable bool   `json:"enable" clover:"Enable"`
	Type   string `json:"type" clover:"Type"`
	Token  string `json:"token" clover:"Token"`
	ChatId string `json:"chat_id" clover:"ChatId"`
}

func (c *ConfigModel) FromDocument(d *document.Document) ConfigModel {
	return ConfigModel{
		Program: ProgramConfig{
			RssTime:     d.Get("Program.RssTime").(int64),
			RenameTime:  d.Get("Program.RenameTime").(int64),
			WebuiPort:   d.Get("Program.WebuiPort").(int64),
			DataVersion: d.Get("Program.DataVersion").(int64),
		},
		Downloader: DownloaderConfig{
			Type:     d.Get("Downloader.Type").(string),
			Host:     d.Get("Downloader.Host").(string),
			Username: d.Get("Downloader.Username").(string),
			Password: d.Get("Downloader.Password").(string),
			Path:     d.Get("Downloader.Path").(string),
			Ssl:      d.Get("Downloader.Ssl").(bool),
		},
		RssParser: RssParserConfig{
			Enable:     d.Get("RssParser.Enable").(bool),
			Type:       d.Get("RssParser.Type").(string),
			CustomUrl:  d.Get("RssParser.CustomUrl").(string),
			Token:      d.Get("RssParser.Token").(string),
			EnableTmdb: d.Get("RssParser.EnableTmdb").(bool),
			Filter:     d.Get("RssParser.Filter").([]interface{}),
			Language:   d.Get("RssParser.Language").(string),
		},
		BangumiManage: BangumiManageConfig{
			Enable:           d.Get("BangumiManage.Enable").(bool),
			EpsComplete:      d.Get("BangumiManage.EpsComplete").(bool),
			RenameMethod:     d.Get("BangumiManage.RenameMethod").(string),
			GroupTag:         d.Get("BangumiManage.GroupTag").(bool),
			RemoveBadTorrent: d.Get("BangumiManage.RemoveBadTorrent").(bool),
		},
		Log: LogConfig{
			DebugEnable: d.Get("Log.DebugEnable").(bool),
		},
		Proxy: ProxyConfig{
			Enable:   d.Get("Proxy.Enable").(bool),
			Type:     d.Get("Proxy.Type").(string),
			Host:     d.Get("Proxy.Host").(string),
			Port:     d.Get("Proxy.Port").(int64),
			Username: d.Get("Proxy.Username").(string),
			Password: d.Get("Proxy.Password").(string),
		},
		Notification: NotificationConfig{
			Enable: d.Get("Notification.Enable").(bool),
			Type:   d.Get("Notification.Type").(string),
			Token:  d.Get("Notification.Token").(string),
			ChatId: d.Get("Notification.ChatId").(string),
		},
	}
}

func InitConfigModel() *ConfigModel {
	return &ConfigModel{
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
			Filter: []interface{}{
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
