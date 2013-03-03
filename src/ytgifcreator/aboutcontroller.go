package ytgifcreator

import (
)

type AboutController struct {
    Context HandleContext
}

func (this *AboutController) Dispatch() {
    this.Context.RenderTemplate("about.html", struct {}{})
}

func (this *AboutController) GetContext() HandleContext {
    return this.Context
}

