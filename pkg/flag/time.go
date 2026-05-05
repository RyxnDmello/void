package flag

var timeLayouts = []string{
	"15:04",
	"15:04:05",
}

func NewTimeFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, timeLayouts)
}

func NewTimeListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, timeLayouts)
}

func NewTimeEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, timeLayouts)
}
