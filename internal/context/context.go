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

func (context *Context) Args() map[string]any {
	return context.arguments
}

func (context *Context) GetArg(name string) any {
	return getValue(context.arguments, name)
}

func (context *Context) MustGetArg(name string) any {
	return mustGetValue(context.arguments, name)
}

func (context *Context) MustGetStringArg(name string) string {
	return mustGetString(context.arguments, name)
}

func (context *Context) MustGetIntArg(name string) int {
	return mustGetInteger(context.arguments, name)
}

func (context *Context) MustGetFloatArg(name string) float64 {
	return mustGetFloat(context.arguments, name)
}

func (context *Context) MustGetBoolArg(name string) bool {
	return mustGetBool(context.arguments, name)
}

func (context *Context) MustGetListArg(name string) []string {
	return mustGetList(context.arguments, name)
}

func (context *Context) MustGetIntListArg(name string) []int {
	return mustGetIntList(context.arguments, name)
}

func (context *Context) MustGetFloatListArg(name string) []float64 {
	return mustGetFloatList(context.arguments, name)
}

func (context *Context) MustGetBoolListArg(name string) []bool {
	return mustGetBoolList(context.arguments, name)
}

func (context *Context) Flags() map[string]any {
	return context.flags
}

func (context *Context) GetFlag(name string) any {
	return getValue(context.flags, name)
}

func (context *Context) MustGetFlag(name string) any {
	return mustGetValue(context.flags, name)
}

func (context *Context) MustGetStringFlag(name string) string {
	return mustGetString(context.flags, name)
}

func (context *Context) MustGetIntFlag(name string) int {
	return mustGetInteger(context.flags, name)
}

func (context *Context) MustGetFloatFlag(name string) float64 {
	return mustGetFloat(context.flags, name)
}

func (context *Context) MustGetBoolFlag(name string) bool {
	return mustGetBool(context.flags, name)
}

func (context *Context) MustGetListFlag(name string) []string {
	return mustGetList(context.flags, name)
}

func (context *Context) MustGetIntListFlag(name string) []int {
	return mustGetIntList(context.flags, name)
}

func (context *Context) MustGetFloatListFlag(name string) []float64 {
	return mustGetFloatList(context.flags, name)
}

func (context *Context) MustGetBoolListFlag(name string) []bool {
	return mustGetBoolList(context.flags, name)
}
