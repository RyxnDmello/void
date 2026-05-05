package flag

var dateLayouts = []string{
	"2006-01-02",
	"2006/01/02",
}

func NewDateFlag(name string) *Flag {
	return NewCustomDateTimeFlag(name, dateLayouts)
}

func NewDateListFlag(name string) *Flag {
	return NewCustomDateTimeListFlag(name, dateLayouts)
}

func NewDateEnumFlag(name string, enum []string) *Flag {
	return NewCustomDateTimeEnumFlag(name, enum, dateLayouts)
}
