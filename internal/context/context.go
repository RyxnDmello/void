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

func (context *Context) MustGetArg(name string) any {
	return mustGetValue(context.arguments, name)
}

func (context *Context) MustGetStringArg(name string) string {
	return mustGetString(context.arguments, name)
}

func (context *Context) MustGetNumberArg(name string) float64 {
	return mustGetNumber(context.arguments, name)
}

func (context *Context) MustGetBoolArg(name string) bool {
	return mustGetBool(context.arguments, name)
}

func (context *Context) MustGetListArg(name string) []string {
	return mustGetList(context.arguments, name)
}

func (context *Context) MustGetFlag(name string) any {
	return mustGetValue(context.flags, name)
}

func (context *Context) MustGetStringFlag(name string) string {
	return mustGetString(context.flags, name)
}

func (context *Context) MustGetNumberFlag(name string) float64 {
	return mustGetNumber(context.flags, name)
}

func (context *Context) MustGetBoolFlag(name string) bool {
	return mustGetBool(context.flags, name)
}

func (context *Context) MustGetListFlag(name string) []string {
	return mustGetList(context.flags, name)
}
