package main

import (
    "rpctypes"
    "fmt"
    "ytlib"
)

type GifCreator struct {
}

func (this *GifCreator) RequestGif(args *rpctypes.RequestGifArgs, gifToken *int) error {
    fmt.Printf("RequestGif %v\n", args)

    var err error
    *gifToken, err = downloadManager.StartDownload(args.VideoId, ytlib.DownloadOptions {
        Format: 43,
        Start: args.Start * 1000,
    })

    if err != nil {
        return err
    }

    return nil
}
