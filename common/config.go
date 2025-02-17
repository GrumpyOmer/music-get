package common

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

var (
	MP3DownloadDir                   string
	MP3DownloadBr                    int
	MP3ConcurrentDownloadTasksNumber int
	// get cpu cores number
	MaxConcurrentDownloadTasksNumber = runtime.NumCPU()
)

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "."
	}
	downloadDir := filepath.Join(homedir, "Music-Get")
	flag.StringVar(&MP3DownloadDir, "o", downloadDir, "MP3 download directory")
	flag.IntVar(&MP3DownloadBr, "br", 128, "MP3 prior download bit rate, 128|192|320")
	flag.IntVar(&MP3ConcurrentDownloadTasksNumber, "n", 1, "MP3 concurrent download tasks number, max 16")
	flag.Parse()
}
