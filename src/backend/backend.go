package main

import (
    "fmt"
    "net/rpc"
    "net/http"
    "net"
    "rpctypes"
    "flag"
)

var queueConter int = 0
var port = flag.Int("port", 80, "The Port to run the RPC Service on")

func main() {
    flag.Parse()
    fmt.Printf("Starting server on port %v...\n", *port)

    gifCreator := new(rpctypes.GifCreator)
    rpc.Register(gifCreator)
    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        panic(fmt.Sprintf("Could not start listen on Port %d", *port))
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
