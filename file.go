package tango

import "path/filepath"

func File(path string) func(ctx *Context) {
	return func(ctx *Context) {
		ctx.ServeFile(path)
	}
}

func Dir(dir string) func(ctx *Context) {
	return func(ctx *Context) {
		params := ctx.Params()
		if len(*params) <= 0 {
			ctx.Result = NotFound()
			ctx.HandleError()
			return
		}
		ctx.ServeFile(filepath.Join(dir, (*params)[0].Value))
	}
}