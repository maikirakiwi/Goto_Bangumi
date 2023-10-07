package models

import "github.com/ostafen/clover/v2/document"

type Bangumi struct {
	ID            int64  `json:"id" clover:"ID"`
	OfficialTitle string `json:"official_title" clover:"OfficialTitle"`
	Year          string `json:"year" clover:"Year"`
	TitleRaw      string `json:"title_raw" clover:"TitleRaw"`
	Season        int64  `json:"season" clover:"Season"`
	SeasonRaw     string `json:"season_raw" clover:"SeasonRaw"`
	GroupName     string `json:"group_name" clover:"GroupName"`
	Dpi           string `json:"dpi" clover:"Dpi"`
	Source        string `json:"source" clover:"Source"`
	Subtitle      string `json:"subtitle" clover:"Subtitle"`
	EpsCollect    bool   `json:"eps_collect" clover:"EpsCollect"`
	Offset        int64  `json:"offset" clover:"Offset"`
	Filter        string `json:"filter" clover:"Filter"`
	RssLink       string `json:"rss_link" clover:"RssLink"`
	PosterLink    string `json:"poster_link" clover:"PosterLink"`
	Added         bool   `json:"added" clover:"Added"`
	RuleName      string `json:"rule_name" clover:"RuleName"`
	SavePath      string `json:"save_path" clover:"SavePath"`
	Deleted       bool   `json:"deleted" clover:"Deleted"`
}

func (b *Bangumi) FromDocument(d *document.Document) *Bangumi {
	return &Bangumi{
		ID:            d.Get("ID").(int64),
		OfficialTitle: d.Get("OfficialTitle").(string),
		Year:          d.Get("Year").(string),
		TitleRaw:      d.Get("TitleRaw").(string),
		Season:        d.Get("Season").(int64),
		SeasonRaw:     d.Get("SeasonRaw").(string),
		GroupName:     d.Get("GroupName").(string),
		Dpi:           d.Get("Dpi").(string),
		Source:        d.Get("Source").(string),
		Subtitle:      d.Get("Subtitle").(string),
		EpsCollect:    d.Get("EpsCollect").(bool),
		Offset:        d.Get("Offset").(int64),
		Filter:        d.Get("Filter").(string),
		RssLink:       d.Get("RssLink").(string),
		PosterLink:    d.Get("PosterLink").(string),
		Added:         d.Get("Added").(bool),
		RuleName:      d.Get("RuleName").(string),
		SavePath:      d.Get("SavePath").(string),
		Deleted:       d.Get("Deleted").(bool),
	}
}

type BangumiUpdate struct {
	OfficialTitle string `json:"official_title" clover:"OfficialTitle"`
	Year          string `json:"year" clover:"Year"`
	TitleRaw      string `json:"title_raw" clover:"TitleRaw"`
	Season        int64  `json:"season" clover:"Season"`
	SeasonRaw     string `json:"season_raw" clover:"SeasonRaw"`
	GroupName     string `json:"group_name" clover:"GroupName"`
	Dpi           string `json:"dpi" clover:"Dpi"`
	Source        string `json:"source" clover:"Source"`
	Subtitle      string `json:"subtitle" clover:"Subtitle"`
	EpsCollect    bool   `json:"eps_collect" clover:"EpsCollect"`
	Offset        int64  `json:"offset" clover:"Offset"`
	Filter        string `json:"filter" clover:"Filter"`
	RssLink       string `json:"rss_link" clover:"RssLink"`
	PosterLink    string `json:"poster_link" clover:"PosterLink"`
	Added         bool   `json:"added" clover:"Added"`
	RuleName      string `json:"rule_name" clover:"RuleName"`
	SavePath      string `json:"save_path" clover:"SavePath"`
	Deleted       bool   `json:"deleted" clover:"Deleted"`
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
