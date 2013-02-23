package ytgifcreator

import (
)

type Controller interface {
    Dispatch()
}

type IndexController struct {
    Context HandleContext
}

func (this *IndexController) Dispatch() {
    this.Context.RenderTemplate("index.html", struct {}{})
}

// Handles API specific request
type ApiController struct {
    Context HandleContext
}

func (this *ApiController) Dispatch() {

}
