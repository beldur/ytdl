package ytgifcreator

import (

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
}

func (this *ApiController) GetContext() HandleContext {
    return this.Context
}
