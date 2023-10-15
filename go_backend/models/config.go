package models

import (
	"database/sql/driver"

	json "github.com/sugawarayuuta/sonnet"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password []byte `json:"password"`
}

func (c *ConfigContent) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), &c)
}

func (c ConfigContent) Value() (driver.Value, error) {
	val, err := json.Marshal(c)
	return string(val), err
}

type Config struct {
	gorm.Model    `json:"-"`
	ConfigContent `gorm:"embedded"`
}

type ConfigContent struct {
	Program       ProgramConfig       `json:"program"`
	Downloader    DownloaderConfig    `json:"downloader"`
	RssParser     RssParserConfig     `json:"rss_parser"`
	BangumiManage BangumiManageConfig `json:"bangumi_manage"`
	Log           LogConfig           `json:"log"`
	Proxy         ProxyConfig         `json:"proxy"`
	Notification  NotificationConfig  `json:"notification"`
}

type ProgramConfig struct {
	RssTime     int64  `json:"rss_time"`
	RenameTime  int64  `json:"rename_time"`
	WebuiPort   int64  `json:"webui_port"`
	DataVersion string `json:"data_version"`
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
	Enable     bool          `json:"enable"`
	Type       string        `json:"type"`
	CustomUrl  string        `json:"custom_url"`
	Token      string        `json:"token"`
	EnableTmdb bool          `json:"enable_tmdb"`
	Filter     []interface{} `json:"filter"`
	Language   string        `json:"language"`
}

type BangumiManageConfig struct {
	Enable           bool   `json:"enable"`
	EpsComplete      bool   `json:"eps_complete"`
	RenameMethod     string `json:"rename_method"`
	GroupTag         bool   `json:"group_tag"`
	RemoveBadTorrent bool   `json:"remove_bad_torrent"`
}

type LogConfig struct {
	ConfigID    int  `json:"config_id"`
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
	ConfigID int    `json:"config_id"`
	Enable   bool   `json:"enable"`
	Type     string `json:"type"`
	Token    string `json:"token"`
	ChatId   string `json:"chat_id"`
}

func InitConfigModel() *Config {
	return &Config{
		ConfigContent: ConfigContent{
			Program: ProgramConfig{
				RssTime:     900,
				RenameTime:  60,
				WebuiPort:   7892,
				DataVersion: "3.1.7",
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
				ConfigID:    0,
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
		},
	}
}
