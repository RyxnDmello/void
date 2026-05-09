package flag

import (
	"fmt"
	"net/mail"
	"slices"
	"strings"
)

func NewEmailFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, *mail.Address:
			return v

		case mail.Address:
			return &v

		case string:
			if mail, err := mail.ParseAddress(v); err == nil {
				return mail
			}

			panic(fmt.Sprintf("%s Not A Valid Email Address", v))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A Valid Email Address", flag.name))
		}
	})

	return flag
}

func NewEmailListFlag(name string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	flag.SetConvert(func(value any) any {
		switch v := value.(type) {
		case nil, []*mail.Address:
			return v

		case []mail.Address:
			mails := make([]*mail.Address, 0)

			for i := range v {
				mails = append(mails, &v[i])
			}

			return mails

		case []string:
			mails := make([]*mail.Address, 0)

			for i := range v {
				mail, err := mail.ParseAddress(v[i])

				if err != nil {
					panic(fmt.Sprintf("%s Not A Valid Email Address", v[i]))
				}

				mails = append(mails, mail)
			}

			return mails

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be A List Of Email Addresses", flag.name))
		}
	})

	flag.variadic = true

	return flag
}

func NewEmailEnumFlag(name string, enum []string) *Flag {
	flag := NewCustomFlag(name, func(value any) any {
		return nil
	})

	if len(enum) == 0 {
		panic(fmt.Sprintf("Flag '%s' Must Have At Least 1 Option", flag.name))
	}

	for i := range enum {
		mail, err := mail.ParseAddress(enum[i])

		if err != nil {
			panic(fmt.Sprintf("Flag '%s' Has An Invalid Option: %s", flag.name, enum[i]))
		}

		flag.enum = append(flag.enum, mail)
	}

	flag.SetConvert(func(value any) any {
		if value == nil {
			return nil
		}

		options := make([]string, 0)

		for i := range flag.enum {
			options = append(options, flag.enum[i].(*mail.Address).String())
		}

		switch v := value.(type) {
		case *mail.Address:
			if slices.Contains(options, v.String()) {
				return v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case mail.Address:
			if slices.Contains(options, v.String()) {
				return &v
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		case string:
			mail, err := mail.ParseAddress(v)

			if err != nil {
				panic(fmt.Sprintf("%s Not A Valid Email Address", v))
			}

			if slices.Contains(options, mail.String()) {
				return mail
			}

			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))

		default:
			panic(fmt.Sprintf("Flag '%s' Must Be One Of %s", flag.name, strings.Join(options, ", ")))
		}
	})

	return flag
}
