package models

type Bangumi struct {
	ID            int    `json:"id"`
	OfficialTitle string `json:"official_title"`
	Year          string `json:"year"`
	TitleRaw      string `json:"title_raw"`
	Season        int    `json:"season"`
	SeasonRaw     string `json:"season_raw"`
	GroupName     string `json:"group_name"`
	Dpi           string `json:"dpi"`
	Source        string `json:"source"`
	Subtitle      string `json:"subtitle"`
	EpsCollect    bool   `json:"eps_collect"`
	Offset        int    `json:"offset"`
	Filter        string `json:"filter"`
	RssLink       string `json:"rss_link"`
	PosterLink    string `json:"poster_link"`
	Added         bool   `json:"added"`
	RuleName      string `json:"rule_name"`
	SavePath      string `json:"save_path"`
	Deleted       bool   `json:"deleted"`
}

type Notification struct {
	OfficialTitle string
	Season        int
	Episode       int
	PosterPath    string
}

type Episode struct {
	TitleEn    string
	TitleZh    string
	TitleJp    string
	Season     int
	SeasonRaw  string
	Episode    int
	Sub        string
	Group      string
	Resolution string
	Source     string
}

type SeasonInfo struct {
	OfficialTitle string
	TitleRaw      string
	Season        int
	SeasonRaw     string
	Group         string
	Filter        []string
	Offset        int
	Dpi           string
	Source        string
	Subtitle      string
	Added         bool
	EpsCollect    bool
}
