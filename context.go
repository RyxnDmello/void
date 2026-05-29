package void

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	ContextFlagNotFound         ErrorCode = "CONTEXT_FLAG_NOT_FOUND"
	ContextFlagTypeMismatch     ErrorCode = "CONTEXT_FLAG_TYPE_MISMATCH"
	ContextArgumentNotFound     ErrorCode = "CONTEXT_ARGUMENT_NOT_FOUND"
	ContextArgumentTypeMismatch ErrorCode = "CONTEXT_ARGUMENT_TYPE_MISMATCH"
)

type Context struct {
	command   commandView
	arguments map[string]any
	flags     map[string]any
}

func newContext(command *Command, arguments map[string]any, flags map[string]any) *Context {
	return &Context{
		command:   command,
		arguments: arguments,
		flags:     flags,
	}
}

func (context *Context) Command() commandView {
	return context.command
}

func (context *Context) Path() string {
	return strings.Join(context.command.Path(), " ")
}

func (context *Context) Arguments() map[string]any {
	return context.arguments
}

func (context *Context) GetArgument(name string) any {
	value, _ := getArgument(context, name)
	return value
}

func (context *Context) MustGetArgument(name string) any {
	return mustGetArgumentAs[any](context, name)
}

func (context *Context) MustGetStringArgument(name string) string {
	return mustGetArgumentAs[string](context, name)
}

func (context *Context) MustGetStringListArgument(name string) []string {
	return mustGetArgumentAs[[]string](context, name)
}

func (context *Context) MustGetIntArgument(name string) int {
	return mustGetArgumentAs[int](context, name)
}

func (context *Context) MustGetIntListArgument(name string) []int {
	return mustGetArgumentAs[[]int](context, name)
}

func (context *Context) MustGetFloatArgument(name string) float64 {
	return mustGetArgumentAs[float64](context, name)
}

func (context *Context) MustGetFloatListArgument(name string) []float64 {
	return mustGetArgumentAs[[]float64](context, name)
}

func (context *Context) MustGetBoolArgument(name string) bool {
	return mustGetArgumentAs[bool](context, name)
}

func (context *Context) MustGetBoolListArgument(name string) []bool {
	return mustGetArgumentAs[[]bool](context, name)
}

func (context *Context) Flags() map[string]any {
	return context.flags
}

func (context *Context) GetFlag(name string) any {
	value, _ := getFlag(context, name)
	return value
}

func (context *Context) MustGetFlag(name string) any {
	return mustGetFlagAs[any](context, name)
}

func (context *Context) MustGetStringFlag(name string) string {
	return mustGetFlagAs[string](context, name)
}

func (context *Context) MustGetStringListFlag(name string) []string {
	return mustGetFlagAs[[]string](context, name)
}

func (context *Context) MustGetIntFlag(name string) int {
	return mustGetFlagAs[int](context, name)
}

func (context *Context) MustGetIntListFlag(name string) []int {
	return mustGetFlagAs[[]int](context, name)
}

func (context *Context) MustGetFloatFlag(name string) float64 {
	return mustGetFlagAs[float64](context, name)
}

func (context *Context) MustGetFloatListFlag(name string) []float64 {
	return mustGetFlagAs[[]float64](context, name)
}

func (context *Context) MustGetBoolFlag(name string) bool {
	return mustGetFlagAs[bool](context, name)
}

func (context *Context) MustGetBoolListFlag(name string) []bool {
	return mustGetFlagAs[[]bool](context, name)
}

func (context *Context) MustGetStringMapFlag(name string) map[string]string {
	return mustGetFlagAs[map[string]string](context, name)
}

func (context *Context) MustGetIntMapFlag(name string) map[string]int {
	return mustGetFlagAs[map[string]int](context, name)
}

func (context *Context) MustGetFloatMapFlag(name string) map[string]float64 {
	return mustGetFlagAs[map[string]float64](context, name)
}

func (context *Context) MustGetBoolMapFlag(name string) map[string]bool {
	return mustGetFlagAs[map[string]bool](context, name)
}

func (context *Context) MustGetComplexMapFlag(name string) map[string]any {
	return mustGetFlagAs[map[string]any](context, name)
}

func (context *Context) MustGetTimeFlag(name string) time.Time {
	return mustGetFlagAs[time.Time](context, name)
}

func (context *Context) MustGetTimeListFlag(name string) []time.Time {
	return mustGetFlagAs[[]time.Time](context, name)
}

func (context *Context) MustGetDurationFlag(name string) time.Duration {
	return mustGetFlagAs[time.Duration](context, name)
}

func (context *Context) MustGetDurationListFlag(name string) []time.Duration {
	return mustGetFlagAs[[]time.Duration](context, name)
}

func (context *Context) MustGetRegexpFlag(name string) *regexp.Regexp {
	return mustGetFlagAs[*regexp.Regexp](context, name)
}

func (context *Context) MustGetRegexpListFlag(name string) []*regexp.Regexp {
	return mustGetFlagAs[[]*regexp.Regexp](context, name)
}

func (context *Context) MustGetURLFlag(name string) *url.URL {
	return mustGetFlagAs[*url.URL](context, name)
}

func (context *Context) MustGetURLListFlag(name string) []*url.URL {
	return mustGetFlagAs[[]*url.URL](context, name)
}

func (context *Context) MustGetIPFlag(name string) net.IP {
	return mustGetFlagAs[net.IP](context, name)
}

func (context *Context) MustGetIPListFlag(name string) []net.IP {
	return mustGetFlagAs[[]net.IP](context, name)
}

func (context *Context) MustGetIPNetFlag(name string) *net.IPNet {
	return mustGetFlagAs[*net.IPNet](context, name)
}

func (context *Context) MustGetIPNetListFlag(name string) []*net.IPNet {
	return mustGetFlagAs[[]*net.IPNet](context, name)
}

func (context *Context) MustGetEmailFlag(name string) *mail.Address {
	return mustGetFlagAs[*mail.Address](context, name)
}

func (context *Context) MustGetEmailListFlag(name string) []*mail.Address {
	return mustGetFlagAs[[]*mail.Address](context, name)
}

func getArgument(context *Context, name string) (any, string) {
	name = mustParseArgumentName(name)

	if len(context.arguments) == 0 {
		return nil, name
	}

	return context.arguments[name], name
}

func mustGetArgumentAs[T any](context *Context, name string) T {
	value, name := getArgument(context, name)

	if value == nil {
		panic(
			NewUsageError(ContextArgumentNotFound, nil).
				WithCommand(context.command).
				WithMessage(
					"Argument value could not be found.",
				).
				WithExpected(
					"A resolved argument value to exist in the command context.",
				).
				WithReceived(
					fmt.Sprintf("Argument '%s' was not found.", name),
				).
				WithSuggestion(
					"Ensure the argument exists and has been parsed before accessing it.",
				),
		)
	}

	result, ok := value.(T)

	if !ok {
		panic(
			NewUsageError(ContextArgumentTypeMismatch, nil).
				WithCommand(context.command).
				WithMessage(
					"Argument value type mismatch detected.",
				).
				WithExpected(
					"The argument value to match the requested type.",
				).
				WithReceived(
					fmt.Sprintf("Argument '%s' could not be converted to the requested type.", name),
				).
				WithSuggestion(
					"Use the correct argument accessor for the underlying argument type.",
				),
		)
	}

	return result
}

func getFlag(context *Context, name string) (any, string) {
	name = mustParseFlagName(name)

	if len(context.flags) == 0 {
		return nil, name
	}

	return context.flags[name], name
}

func mustGetFlagAs[T any](context *Context, name string) T {
	value, name := getFlag(context, name)

	if value == nil {
		panic(
			NewUsageError(ContextFlagNotFound, nil).
				WithCommand(context.command).
				WithMessage(
					"Flag value could not be found.",
				).
				WithExpected(
					"A resolved flag value to exist in the command context.",
				).
				WithReceived(
					fmt.Sprintf("Flag '%s' was not found.", name),
				).
				WithSuggestion(
					"Ensure the flag exists and has been parsed before accessing it.",
				),
		)
	}

	result, ok := value.(T)

	if !ok {
		panic(
			NewUsageError(ContextFlagTypeMismatch, nil).
				WithCommand(context.command).
				WithMessage(
					"Flag value type mismatch detected.",
				).
				WithExpected(
					"The flag value to match the requested type.",
				).
				WithReceived(
					fmt.Sprintf("Flag '%s' could not be converted to the requested type.", name),
				).
				WithSuggestion(
					"Use the correct flag accessor for the underlying flag type.",
				),
		)
	}

	return result
}
