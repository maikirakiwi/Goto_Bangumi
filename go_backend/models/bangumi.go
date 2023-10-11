package models

import (
	"github.com/ostafen/clover/v2/document"
)


//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Bangumi struct {
	ID            int64  `json:"id" objectbox:"index"`
	OfficialTitle string `json:"official_title"`
	Year          string `json:"year"`
	TitleRaw      string `json:"title_raw"`
	Season        int64  `json:"season"
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
	SavePath      string `json:"save_path" objectbox:"unique"`
	Deleted       bool   `json:"deleted"`
}

func (b *Bangumi) FromDocument(d *document.Document) Bangumi {
	return Bangumi{
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
	OfficialTitle string `json:"official_title" objectbox:"OfficialTitle"`
	Year          string `json:"year" objectbox:"Year"`
	TitleRaw      string `json:"title_raw" objectbox:"TitleRaw"`
	Season        int64  `json:"season" objectbox:"Season"`
	SeasonRaw     string `json:"season_raw" objectbox:"SeasonRaw"`
	GroupName     string `json:"group_name" objectbox:"GroupName"`
	Dpi           string `json:"dpi" objectbox:"Dpi"`
	Source        string `json:"source" objectbox:"Source"`
	Subtitle      string `json:"subtitle" objectbox:"Subtitle"`
	EpsCollect    bool   `json:"eps_collect" objectbox:"EpsCollect"`
	Offset        int64  `json:"offset" objectbox:"Offset"`
	Filter        string `json:"filter" objectbox:"Filter"`
	RssLink       string `json:"rss_link" objectbox:"RssLink"`
	PosterLink    string `json:"poster_link" objectbox:"PosterLink"`
	Added         bool   `json:"added" objectbox:"Added"`
	RuleName      string `json:"rule_name" objectbox:"RuleName"`
	SavePath      string `json:"save_path" objectbox:"SavePath"`
	Deleted       bool   `json:"deleted" objectbox:"Deleted"`
}

func (b *BangumiUpdate) FromDocument(d *document.Document) BangumiUpdate {
	return BangumiUpdate{
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
