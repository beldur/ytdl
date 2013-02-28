package main

import (
    "rpctypes"
    "fmt"
    "ytlib"
)

type GifCreator struct {
}

func (this *GifCreator) RequestGif(args *rpctypes.RequestGifArgs, status *rpctypes.VideoStatus) error {
    fmt.Printf("RequestGif %v\n", args)

    newStatus, err := downloadManager.StartDownload(args.VideoId, ytlib.DownloadOptions {
        Format: 43,
        Start: args.Start * 1000,
    })

    if err != nil {
        fmt.Printf("Wohoo error here %v\n", err)
        return err
    }

    status.Hash = newStatus.Hash
    status.Status = newStatus.Status

    fmt.Printf("Status after StartDownload %#v \n", status)
    return nil
}

func (this *GifCreator) RequestStatus(args *rpctypes.RequestStatusArgs, status *rpctypes.VideoStatus) error {
    fmt.Printf("RequestStatus %v\n", args)

    newStatus, err := downloadManager.GetVideoStatus(args.Token)
    if err != nil {
        return err
    }

    status.Hash = newStatus.Hash
    status.Status = newStatus.Status

    return nil
}
