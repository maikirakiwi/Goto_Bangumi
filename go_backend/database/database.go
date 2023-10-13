package database

import (
	"os"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"
	"golang.org/x/crypto/bcrypt"

	"Auto_Bangumi/v2/models/store"
)

var Conn *objectbox.ObjectBox

func initObjectBox() {
	var err error
	Conn, err = objectbox.NewBuilder().Model(store.ObjectBoxModel()).Build()
	if err != nil {
		log.Fatal().Msgf("Error while initializing store: %s", err.Error())
	}
}

func firstRun() {
	initObjectBox()

	// Default user is admin/adminadmin
	password, _ := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	_, err := store.BoxForUser(Conn).Put(&store.User{
		Id:       1,
		Username: "admin",
		Password: password,
	})
	if err != nil {
		log.Fatal().Msgf("Error while creating default user: %s", err.Error())
	}

	// Default settings
	_, err = store.BoxForConfigModel(Conn).Put(store.InitConfigModel())
	if err != nil {
		log.Fatal().Msgf("Error while setting default config: %s", err.Error())
	}
}

func testContent() {
	test := `[
		{
			"id": 6,
			"official_title": "总之就是非常可爱 女子高中篇",
			"year": "",
			"title_raw": "Joshikou",
			"season": 1,
			"season_raw": "",
			"group_name": "云光字幕组",
			"dpi": "1080p",
			"source": "",
			"subtitle": "简体双语",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁,S2,剃须",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202304/3eefbe81.jpg",
			"added": true,
			"rule_name": "总之就是非常可爱 女子高中篇 S1",
			"save_path": "/downloads/emby/总之就是非常可爱 女子高中篇/Season 1",
			"deleted": false
		},
		{
			"id": 5,
			"official_title": "布莱泽奥特曼",
			"year": "",
			"title_raw": "ULTRAMAN BLAZAR",
			"season": 1,
			"season_raw": "",
			"group_name": "星空字幕组",
			"dpi": "1080P",
			"source": "",
			"subtitle": "简日双语",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/5fac8055.jpg",
			"added": true,
			"rule_name": "布莱泽奥特曼 S1",
			"save_path": "/downloads/emby/布莱泽奥特曼/Season 1",
			"deleted": false
		},
		{
			"id": 8,
			"official_title": "不死少女·杀人笑剧",
			"year": "",
			"title_raw": "Undead Girl Murder Farce",
			"season": 1,
			"season_raw": "",
			"group_name": "喵萌奶茶屋",
			"dpi": "1080p",
			"source": "",
			"subtitle": "简日双语",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/d774026d.jpg",
			"added": true,
			"rule_name": "不死少女·杀人笑剧 S1",
			"save_path": "/downloads/emby/不死少女·杀人笑剧/Season 1",
			"deleted": false
		},
		{
			"id": 3,
			"official_title": "僵尸百分百～变成僵尸之前想做的100件事",
			"year": "",
			"title_raw": "Zom 100 - Zombie ni Naru made ni Shitai 100 no Koto",
			"season": 1,
			"season_raw": "",
			"group_name": "漫猫字幕社",
			"dpi": "1080P",
			"source": "",
			"subtitle": "简日双语",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/50f38c71.jpg",
			"added": true,
			"rule_name": "僵尸百分百～变成僵尸之前想做的100件事 S1",
			"save_path": "/downloads/emby/僵尸百分百～变成僵尸之前想做的100件事/Season 1",
			"deleted": false
		},
		{
			"id": 1,
			"official_title": "我的幸福婚姻",
			"year": "",
			"title_raw": "Watashi no Shiawase na Kekkon",
			"season": 1,
			"season_raw": "",
			"group_name": "喵萌奶茶屋",
			"dpi": "1080p",
			"source": "",
			"subtitle": "简日双语",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/2e6aede2.jpg",
			"added": true,
			"rule_name": "我的幸福婚姻 S1",
			"save_path": "/downloads/emby/我的幸福婚姻/Season 1",
			"deleted": false
		},
		{
			"id": 2,
			"official_title": "白圣女与黑牧师",
			"year": "",
			"title_raw": "Shiro Seijo to Kuro Bokushi",
			"season": 1,
			"season_raw": "",
			"group_name": "爱恋字幕社\u0026猫恋汉化组",
			"dpi": "1080p",
			"source": "",
			"subtitle": "简中",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/5f3fad96.jpg",
			"added": true,
			"rule_name": "白圣女与黑牧师 S1",
			"save_path": "/downloads/emby/白圣女与黑牧师/Season 1",
			"deleted": false
		},
		{
			"id": 7,
			"official_title": "能干猫今天也忧郁",
			"year": "",
			"title_raw": "Dekiru Neko wa Kyou mo Yuuutsu",
			"season": 1,
			"season_raw": "",
			"group_name": "ANi",
			"dpi": "1080P",
			"source": "Baha",
			"subtitle": "CHT",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/9aa8e8b0.jpg",
			"added": true,
			"rule_name": "能干猫今天也忧郁 S1",
			"save_path": "/downloads/emby/能干猫今天也忧郁/Season 1",
			"deleted": false
		},
		{
			"id": 9,
			"official_title": "莱莎的炼金工房 ～常暗女王与秘密藏身处～",
			"year": "",
			"title_raw": "The Animation",
			"season": 1,
			"season_raw": "",
			"group_name": "ANi",
			"dpi": "1080P",
			"source": "Baha",
			"subtitle": "CHT",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/d61fd98b.jpg",
			"added": true,
			"rule_name": "莱莎的炼金工房 ～常暗女王与秘密藏身处～ S1",
			"save_path": "/downloads/emby/莱莎的炼金工房 ～常暗女王与秘密藏身处～/Season 1",
			"deleted": false
		},
		{
			"id": 4,
			"official_title": "租借女友",
			"year": "",
			"title_raw": "Kanojo, Okarishimasu (2023)",
			"season": 3,
			"season_raw": "第三季",
			"group_name": "桜都字幕组",
			"dpi": "1080p",
			"source": "",
			"subtitle": "简体内嵌",
			"eps_collect": true,
			"offset": 0,
			"filter": "720,\\d+-\\d,繁",
			"rss_link": "https://mikanani.me/RSS/MyBangumi?token=GAZ4CI23%2fGMzekgJa%2fn8o1fLaC2oO5iP%2fAQ09APoYSs%3d",
			"poster_link": "https://mikanani.me/images/Bangumi/202307/cd277ca7.jpg",
			"added": true,
			"rule_name": "租借女友 S3",
			"save_path": "/downloads/emby/租借女友/Season 3",
			"deleted": false
		}
	]`
	bangumitest := make([]store.Bangumi, 9)
	json.Unmarshal([]byte(test), &bangumitest)
	store.BoxForBangumi(Conn).ObjectBox.RunInWriteTx(func() error {
		for i := 0; i < 8; i++ {
			store.BoxForBangumi(Conn).Put(&bangumitest[i])
		}
		return nil
	})
}

func Init() {
	if _, err := os.Stat("./objectBox"); err != nil {
		firstRun()
	} else if len(os.Args) > 1 && os.Args[1] == "dev" {
		log.Info().Msg("Development mode detected. Skipping JWT auth; Removing data directory; Adding sample data.")
		err := os.RemoveAll("./objectBox")

		if err != nil {
			log.Fatal().Msgf("Error removing data directory: %s", err.Error())
		}

		firstRun()
		testContent()
	} else {
		initObjectBox()
	}
}

func Teardown() {
	Conn.Close()
}
