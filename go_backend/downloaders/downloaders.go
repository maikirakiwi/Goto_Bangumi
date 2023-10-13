package downloaders

import (
	"github.com/rs/zerolog/log"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/downloaders/qbit"
	"Auto_Bangumi/v2/models/store"
)

// Global Qbit Client
var Qbit *qbit.Client

func Init() {
	cfg, err := store.BoxForConfigModel(db.Conn).Get(1)
	if err != nil {
		log.Fatal().Msg("Config not in cache while initializing Qbittorrent Client.")
	}

	var host string
	if !cfg.Downloader.Ssl {
		host = "http://" + cfg.Downloader.Host
	} else {
		host = "https://" + cfg.Downloader.Host
	}

	cli, err := qbit.NewCli(host,
		cfg.Downloader.Username,
		cfg.Downloader.Password)
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
