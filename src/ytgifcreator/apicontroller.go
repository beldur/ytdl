package ytgifcreator

import (
    "net/http"
    "rpctypes"
    //"fmt"
    "strconv"
)

// Handles API specific request
type ApiController struct {
    Context HandleContext
}

func (this *ApiController) Dispatch() {
    requestName := this.Context.r.URL.Path[6:]
    FindMethodOr404(this, requestName)
}

// User want to get Gif for Video Id
func (this *ApiController) RequestGif() {
    this.Context.gaeContext.Infof("Api Request: %v", "gif")

    start, _ := strconv.Atoi(this.Context.r.FormValue("start"))
    end, _ := strconv.Atoi(this.Context.r.FormValue("end"))

    var gifToken int

    requestArguments := rpctypes.RequestGifArgs {
        this.Context.r.FormValue("videoId"),
        start,
        end,
    }

    err := RpcCall("GifCreator.RequestGif", requestArguments, &gifToken)
    if err != nil {
        http.Error(this.Context.w, err.Error(), http.StatusInternalServerError)
        return
    }

    JsonResponse(this.Context.w, struct { GifToken int } { gifToken })
}

func (this *ApiController) GetContext() HandleContext {
    return this.Context
}
