package rpctypes

import (
)

type RequestGifArgs struct {
    VideoId string
    Start int
    End int
}

type RequestStatusArgs struct {
    Token string
}

type VideoProcessStatus int

type VideoStatus struct {
    Hash string
    Status VideoProcessStatus
}
