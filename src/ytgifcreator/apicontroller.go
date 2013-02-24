package ytgifcreator

import (
    "net/http"
    "rpctypes"
    "fmt"
)

// Handles API specific request
type ApiController struct {
    Context HandleContext
}

func (this *ApiController) Dispatch() {
    requestName := this.Context.r.URL.Path[6:]
    FindMethodOr404(this, requestName)
}

func (this *ApiController) RequestGif() {
    this.Context.gaeContext.Infof("Api Request: %v", "gif")
    args := rpctypes.RequestGifArgs{this.Context.r.FormValue("videoId")}
    var reply string

    err := RpcCall("GifCreator.RequestGif", args, &reply)
    if err != nil {
        http.Error(this.Context.w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(this.Context.w, "reply: %v", reply)
}

func (this *ApiController) GetContext() HandleContext {
    return this.Context
}
