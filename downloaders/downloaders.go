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

	var host string
	if !cfg.(models.Config).Downloader.Ssl {
		host = "http://" + cfg.(models.Config).Downloader.Host
	} else {
		host = "https://" + cfg.(models.Config).Downloader.Host
	}

	cli, err := qbit.NewCli(host,
		cfg.(models.Config).Downloader.Username,
		cfg.(models.Config).Downloader.Password)
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
