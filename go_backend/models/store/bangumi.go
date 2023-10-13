package store

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Bangumi struct {
	ID            uint64 `json:"id" objectbox:"id"`
	OfficialTitle string `json:"official_title"`
	Year          string `json:"year"`
	TitleRaw      string `json:"title_raw"`
	Season        int64  `json:"season"`
	SeasonRaw     string `json:"season_raw"`
	GroupName     string `json:"group_name"`
	Dpi           string `json:"dpi"`
	Source        string `json:"source"`
	Subtitle      string `json:"subtitle"`
	EpsCollect    bool   `json:"eps_collect"`
	Offset        int64  `json:"offset"`
	Filter        string `json:"filter"`
	RssLink       string `json:"rss_link"`
	PosterLink    string `json:"poster_link"`
	Added         bool   `json:"added"`
	RuleName      string `json:"rule_name"`
	SavePath      string `json:"save_path"`
	Deleted       bool   `json:"deleted"`
}
