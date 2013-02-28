package main

import (
    "ytlib"
    "fmt"
    "crypto/sha1"
    "rpctypes"
    "io"
)

// Video Status options
const (
    DOWNLOADING rpctypes.VideoProcessStatus = iota
    DONE
)

type DownloadManager struct {
    downloadDir string
    queueCounter int
    videoList map[string]rpctypes.VideoProcessStatus
}

func (this *DownloadManager) Init(downloadDirectory string) *DownloadManager {
    return &DownloadManager {
        downloadDirectory,
        0,
        map[string]rpctypes.VideoProcessStatus {},
    }
}

// Request a new video for download
func (this *DownloadManager) StartDownload(videoId string, options ytlib.DownloadOptions) (*rpctypes.VideoStatus, error) {
    // Do we have a enough for our guest?
    if this.queueCounter > 3 {
        return nil, fmt.Errorf("Sorry queue is full :(")
    }

    videoHash := this.GetVideoHash(videoId, options)

    // Check if this video known to us
    videoStatus, ok := this.GetVideoStatus(videoHash)
    if ok == nil {
        return videoStatus, nil
    }

    ytVideo := new(ytlib.YTVideo).Init(videoId)
    ytVideo.GetFormatList()

    if !ytVideo.HasFormat(options.Format) {
        return nil, fmt.Errorf("Format %d is not available", options.Format)
    }

    // Start Download asynchronously
    go func () {
        this.queueCounter++
        defer this.reduceCounter()

        ytVideo.DownloadVideo(this.downloadDir + videoId, options)
        this.UpdateStatus(videoHash, DONE)
    }()

    return this.CreateVideoStatus(videoId, options, DOWNLOADING), nil
}

// Create a video status struct
func (this *DownloadManager) CreateVideoStatus(videoId string, options ytlib.DownloadOptions, status rpctypes.VideoProcessStatus) *rpctypes.VideoStatus {
    videoHash := this.GetVideoHash(videoId, options)
    fmt.Println(videoHash)
    return &rpctypes.VideoStatus { videoHash, status }
}

// Get current video status
func (this *DownloadManager) GetVideoStatus(videoHash string) (*rpctypes.VideoStatus, error) {
    videoStatus, exists := this.videoList[videoHash]
    if exists {
        return &rpctypes.VideoStatus { videoHash, videoStatus }, nil
    }

    return nil, fmt.Errorf("Video status not found")
}

func (this *DownloadManager) UpdateStatus(videoHash string, status rpctypes.VideoProcessStatus) {
    this.videoList[videoHash] = status
}

func (this *DownloadManager) reduceCounter () {
    this.queueCounter = this.queueCounter - 1
}

// Create Hash for Video and options
func (this *DownloadManager) GetVideoHash(videoId string, options ytlib.DownloadOptions) string {
    hasher := sha1.New()
    io.WriteString(hasher, fmt.Sprintf("%s-%d-%d-%d", videoId, options.Format, options.Start, options.End))
    return fmt.Sprintf("%x", hasher.Sum(nil))
}
