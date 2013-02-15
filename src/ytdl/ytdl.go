package main

import (
    "fmt"
    "ytlib"
)


func main() {
    videoId := "8k9XnnqACdc"
    ytVideo := new(ytlib.YTVideo).Init(videoId)
    ytVideo.GetFormatList()

    fmt.Println("")
    ytVideo.DownloadWorstQuality(videoId + "Worst")
    ytVideo.DownloadBestQuality(videoId + "Best")
}
