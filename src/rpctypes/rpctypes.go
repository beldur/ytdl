package rpctypes

import (
    "fmt"
)

type GifCreator struct {
}

type RequestGifArgs struct {
    VideoId string
}

func (this *GifCreator) RequestGif(args *RequestGifArgs, gifUrl *string) error {
    fmt.Printf("RequestGif %v\n", args)
    *gifUrl = args.VideoId + ".gif"
    return nil
}
