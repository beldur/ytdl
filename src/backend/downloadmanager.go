package main

import (
    "ytlib"
    "fmt"
)

type DownloadManager struct {
    downloadDir string
    queueCounter int
    videoList map[string]int
}

func (this *DownloadManager) Init(downloadDirectory string) *DownloadManager {
    return &DownloadManager {
        downloadDirectory,
        0,
        map[string]int {},
    }
}

func (this *DownloadManager) StartDownload(videoId string, options ytlib.DownloadOptions) (int, error) {
    ytVideo := new(ytlib.YTVideo).Init(videoId)
    ytVideo.GetFormatList()

    if !ytVideo.HasFormat(options.Format) {
        return -1, fmt.Errorf("Format %d is not available", options.Format)
    }

    fmt.Printf("Downloading with %v\n", options)
    go ytVideo.DownloadVideo(this.downloadDir + videoId, options)

    return 123, nil
}
