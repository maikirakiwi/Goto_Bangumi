package models

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
