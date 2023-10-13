package store

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type ConfigModel struct {
	Id            uint64
	Program       ProgramConfig       `json:"program"`
	Downloader    DownloaderConfig    `json:"downloader"`
	RssParser     RssParserConfig     `json:"rss_parser"`
	BangumiManage BangumiManageConfig `json:"bangumi_manage"`
	Log           LogConfig           `json:"log"`
	Proxy         ProxyConfig         `json:"proxy"`
	Notification  NotificationConfig  `json:"notification"`
}

type Cache struct {
	Id         uint64
	activeUser string `objectbox:"-"`
}

type User struct {
	Id       uint64
	Username string `objectbox:"unique"`
	Password []byte
}
