package downloaders

import (
	"github.com/rs/zerolog/log"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/downloaders/qbit"
	"Auto_Bangumi/v2/models"
)

// Global Qbit Client
var Qbit *qbit.Client

func Init() {
	cfg, exists := db.Cache.Get("config")
	if !exists {
		log.Fatal().Msg("Config not in cache while initializing Qbittorrent Client.")
	}
	cli, err := qbit.NewCli(cfg.(models.ConfigModel).Downloader.Host,
		cfg.(models.ConfigModel).Downloader.Username,
		cfg.(models.ConfigModel).Downloader.Password)
	if err != nil {
		log.Fatal().Msg("Error while initializing Qbittorrent Client: " + err.Error())
	}

	Qbit = cli
}

func GetAllTorrent() []qbit.Torrent {
	list, err := Qbit.TorrentList(qbit.Optional{
		"category": "BangumiCollection",
	})
	if err != nil {
		log.Error().Msgf("Error while getting torrent list from Qbit: %s", err)
		return nil
	}

	return list
}
