package main

import (
    "fmt"
    "net/rpc"
    "net/http"
    "net"
    "rpctypes"
)

var queueConter int = 0


func main() {
    fmt.Println("Starting server...")

    gifCreator := new(rpctypes.GifCreator)
    rpc.Register(gifCreator)
    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", ":8081")
    if err != nil {
        panic("Could not start listen on Port 80")
    }

    http.Serve(listener, nil)
}
    /*ytVideo := new(ytlib.YTVideo).Init("8k9XnnqACdc")
    ytVideo.GetFormatList()
    title := ytVideo.VideoInformation["title"][0]

    fmt.Println(title, ytVideo.FormatList)

    format, _ := ytVideo.GetBestQuality()
    ytVideo.DownloadVideo(title,
        ytlib.DownloadOptions { Format: format, Begin: 60000 },
    )*/
