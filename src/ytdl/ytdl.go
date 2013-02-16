package main

import (
    "fmt"
    "ytlib"
)


func main() {
    videoId := "8k9XnnqACdc"
    ytVideo := new(ytlib.YTVideo).Init(videoId)
    list, _ := ytVideo.GetFormatList()
    title := ytVideo.VideoInformation["title"][0]
    fmt.Println(list, title)

    format, _ := ytVideo.GetBestQuality()
    ytVideo.DownloadVideo(title,
        ytlib.DownloadOptions{ Format: format },
    )
}
