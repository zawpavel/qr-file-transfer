package filetransfer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/zawpavel/qrcp/application"
	"github.com/zawpavel/qrcp/config"
	"github.com/zawpavel/qrcp/payload"
	"github.com/zawpavel/qrcp/server"
)

func GetDownloadLink(filepath string) string {
	payload, err := payload.FromArgs([]string{filepath}, false)
	if err != nil {
		log.Fatal(err)
	}
	app := application.New()
	cfg := config.New(app)
	srv, err := server.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	srv.Send(payload)
	return srv.SendURL
}

func ReceiveFiles(receiveUrlChannel, outputDirChannel chan string) {
	app := application.New()
	cfg := config.New(app)
	cfg.Output = getDownloadPath()
	srv, err := server.New(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.ReceiveTo(cfg.Output); err != nil {
		log.Fatal(err)
	}
	receiveUrlChannel <- srv.ReceiveURL
	if err := srv.Wait(); err != nil {
		log.Fatal(err)
	}
	outputDirChannel <- cfg.Output
}

func getDownloadPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	downloadDir := filepath.Join(homeDir, "Downloads")
	if _, err := os.Stat(downloadDir); err == nil {
		return downloadDir
	}
	if _, err := os.Stat(homeDir); err == nil {
		return homeDir
	}
	return ""
}
