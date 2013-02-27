package main

import (
    "fmt"
    "net/rpc"
    "net/http"
    "net"
    "flag"
)

var port = flag.Int("port", 80, "The Port to run the RPC Service on")
var downloadDirectory = flag.String("download-dir", "/tmp/", "Where to download videos")
var downloadManager *DownloadManager

func main() {
    flag.Parse()
    fmt.Printf("Starting server on port %v...\n", *port)

    gifCreator := new(GifCreator)
    downloadManager = new(DownloadManager).Init(*downloadDirectory)

    rpc.Register(gifCreator)
    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        panic(fmt.Sprintf("Could not start listen on Port %d", *port))
    }

    http.Serve(listener, nil)
}

