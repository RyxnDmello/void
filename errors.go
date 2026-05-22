package void

import (
	"fmt"
	"os"
)

type (
	ErrorKind string
	ErrorCode string
)

const (
	UsageError      ErrorKind = "USAGE_ERROR"
	RuntimeError    ErrorKind = "RUNTIME_ERROR"
	FrameworkError  ErrorKind = "FRAMEWORK_ERROR"
	DefinitionError ErrorKind = "DEFINITION_ERROR"
)

const (
	InvalidFlag                     ErrorCode = "INVALID_FLAG"
	UnknownCommand                  ErrorCode = "UNKNOWN_COMMAND"
	CommandRequired                 ErrorCode = "COMMAND_REQUIRED"
	InvalidFlagFormat               ErrorCode = "INVALID_FLAG_FORMAT"
	InvalidCommandRoute             ErrorCode = "INVALID_COMMAND_ROUTE"
	TooManyCommandFlags             ErrorCode = "TOO_MANY_COMMAND_FLAGS"
	TooManyCommandArguments         ErrorCode = "TOO_MANY_COMMAND_ARGUMENTS"
	TooManyCommandFlagValues        ErrorCode = "TOO_MANY_COMMAND_FLAG_VALUES"
	MissingRequiredCommandFlags     ErrorCode = "MISSING_REQUIRED_COMMAND_FLAGS"
	MissingRequiredCommandArguments ErrorCode = "MISSING_REQUIRED_COMMAND_ARGUMENTS"

	CommandPathRequired           ErrorCode = "COMMAND_PATH_REQUIRED"
	CommandNameRequired           ErrorCode = "COMMAND_NAME_REQUIRED"
	DuplicateCommandFlag          ErrorCode = "DUPLICATE_COMMAND_FLAG"
	DuplicateCommandRoute         ErrorCode = "DUPLICATE_COMMAND_ROUTE"
	CommandTaglineRequired        ErrorCode = "COMMAND_TAGLINE_REQUIRED"
	DuplicateCommandArgument      ErrorCode = "DUPLICATE_COMMAND_ARGUMENT"
	MultipleVariadicArguments     ErrorCode = "MULTIPLE_VARIADIC_ARGUMENTS"
	CommandDescriptionRequired    ErrorCode = "COMMAND_DESCRIPTION_REQUIRED"
	CommandExecuteHandlerRequired ErrorCode = "COMMAND_EXECUTE_HANDLER_REQUIRED"
)

type VoidError struct {
	kind       ErrorKind
	code       ErrorCode
	cause      error
	message    *string
	expected   *string
	received   *string
	suggestion *string
	command    *Command
	argument   *Argument
	flag       *Flag
	metadata   map[string]string
}

func NewError(kind ErrorKind, code ErrorCode, cause error) *VoidError {
	return &VoidError{
		kind:       kind,
		code:       code,
		cause:      cause,
		message:    nil,
		expected:   nil,
		received:   nil,
		suggestion: nil,
		command:    nil,
		argument:   nil,
		flag:       nil,
		metadata:   nil,
	}
}

func NewUsageError(code ErrorCode, cause error) *VoidError {
	return NewError(UsageError, code, cause)
}

func NewRuntimeError(code ErrorCode, cause error) *VoidError {
	return NewError(RuntimeError, code, cause)
}

func NewFrameworkError(code ErrorCode, cause error) *VoidError {
	return NewError(FrameworkError, code, cause)
}

func NewDefinitionError(code ErrorCode, cause error) *VoidError {
	return NewError(DefinitionError, code, cause)
}

func (voidError *VoidError) WithMessage(message string) *VoidError {
	voidError.message = &message
	return voidError
}

func (voidError *VoidError) WithExpected(expected string) *VoidError {
	voidError.expected = &expected
	return voidError
}

func (voidError *VoidError) WithReceived(received string) *VoidError {
	voidError.received = &received
	return voidError
}

func (voidError *VoidError) WithSuggestion(suggestion string) *VoidError {
	voidError.suggestion = &suggestion
	return voidError
}

func (voidError *VoidError) WithCommand(command *Command) *VoidError {
	voidError.command = command
	return voidError
}

func (voidError *VoidError) WithArgument(argument *Argument) *VoidError {
	voidError.argument = argument
	return voidError
}

func (voidError *VoidError) WithFlag(flag *Flag) *VoidError {
	voidError.flag = flag
	return voidError
}

func (voidError *VoidError) WithMetadata(metadata map[string]string) *VoidError {
	voidError.metadata = metadata
	return voidError
}

func (voidError *VoidError) Error() string {
	return fmt.Sprintf("ERROR: %s | %s", voidError.kind, voidError.code)
}

func (voidError *VoidError) Message() string {
	panic := fmt.Sprintf("%s\n", voidError.Error())

	if voidError.message != nil {
		panic += fmt.Sprintf("\nMessage:\n\t%s\n", *voidError.message)
	}

	if voidError.expected != nil {
		panic += fmt.Sprintf("\nExpected:\n\t%s\n", *voidError.expected)
	}

	if voidError.received != nil {
		panic += fmt.Sprintf("\nReceived:\n\t%s\n", *voidError.received)
	}

	if voidError.suggestion != nil {
		panic += fmt.Sprintf("\nSuggestion:\n\t%s\n", *voidError.suggestion)
	}

	return panic
}

func (voidError *VoidError) Fatal() {
	fmt.Print(voidError.Message())
	os.Exit(1)
}

func onError() {
	if rec := recover(); rec != nil {
		switch err := rec.(type) {
		case *VoidError:
			fmt.Print(err.Message())

		default:
			fmt.Print(err)
		}

		os.Exit(1)
	}
}
