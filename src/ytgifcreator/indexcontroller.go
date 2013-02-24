package ytgifcreator

import (
)

type Controller interface {
    Dispatch()
    GetContext() HandleContext
}

type IndexController struct {
    Context HandleContext
}

func (this *IndexController) Dispatch() {
    this.Context.RenderTemplate("index.html", struct {}{})
}

func (this *IndexController) GetContext() HandleContext {
    return this.Context
}

