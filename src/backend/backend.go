package main

import (
    "fmt"
    "net/rpc"
    "net/http"
    "net"
    "flag"
    "rpctypes"
    "ytlib"
)

var queueCounter int = 0
var port = flag.Int("port", 80, "The Port to run the RPC Service on")
var downloadDirectory = flag.String("download-dir", "/tmp/", "Where to download videos")

func main() {
    flag.Parse()
    fmt.Printf("Starting server on port %v...\n", *port)

    gifCreator := new(GifCreator)
    rpc.Register(gifCreator)
    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        panic(fmt.Sprintf("Could not start listen on Port %d", *port))
    }

    http.Serve(listener, nil)
}

type GifCreator struct {
}

func (this *GifCreator) RequestGif(args *rpctypes.RequestGifArgs, gifToken *int) error {
    fmt.Printf("RequestGif %v\n", args)

    ytVideo := new(ytlib.YTVideo).Init(args.VideoId)
    ytVideo.GetFormatList()

    if ytVideo.HasFormat(43) {
        dOptions := ytlib.DownloadOptions { Format: 43, Start: args.Start * 1000 }
        fmt.Printf("Downloading with %v\n", dOptions)
        ytVideo.DownloadVideo(*downloadDirectory + args.VideoId, dOptions)
    }

    *gifToken = 123
    return nil
}
