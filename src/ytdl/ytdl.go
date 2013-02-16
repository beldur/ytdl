package main

import (
    "fmt"
    "ytlib"
)

func main() {
    ytVideo := new(ytlib.YTVideo).Init("8k9XnnqACdc")
    ytVideo.GetFormatList()
    title := ytVideo.VideoInformation["title"][0]

    fmt.Println(title)

    format, _ := ytVideo.GetBestQuality()
    ytVideo.DownloadVideo(title,
        ytlib.DownloadOptions { Format: format, Begin: 10000 },
   )
}
