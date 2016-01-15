package main

import (
	"github.com/winespace/coa/web"
	"github.com/zenazn/goji"
)

func main() {
	hb := web.GojiHandlerBuilder{NewContext}

	goji.Get("/user", hb.Build(GetUser{}))
	goji.Post("/user", hb.Build(UpdateUser{}))
	goji.Serve()
}
