package context

type Context struct {
	path      string
	flags     map[string]any
	arguments map[string]any
}

func NewContext(path string, arguments map[string]any, flags map[string]any) *Context {
	return &Context{
		path:      path,
		flags:     flags,
		arguments: arguments,
	}
}

func (context *Context) Path() string {
	return context.path
}

func (context *Context) Args() map[string]any {
	return context.arguments
}

func (context *Context) Flags() map[string]any {
	return context.flags
}
