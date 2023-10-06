package models

type Bangumi struct {
	ID            int
	OfficialTitle string
	Year          string
	TitleRaw      string
	Season        int
	SeasonRaw     string
	GroupName     string
	Dpi           string
	Source        string
	Subtitle      string
	EpsCollect    bool
	Offset        int
	Filter        string
	RssLink       string
	PosterLink    string
	Added         bool
	RuleName      string
	SavePath      string
	Deleted       bool
}

type BangumiUpdate struct {
	OfficialTitle string
	Year          string
	TitleRaw      string
	Season        int
	SeasonRaw     string
	GroupName     string
	Dpi           string
	Source        string
	Subtitle      string
	EpsCollect    bool
	Offset        int
	Filter        string
	RssLink       string
	PosterLink    string
	Added         bool
	RuleName      string
	SavePath      string
	Deleted       bool
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
