package notary

import "fmt"
import "regexp"

type RegexUtil struct {
	email *regexp.Regexp
}

type ErrRequired struct {
	error
	arg string
}

type ErrInvalidValue struct {
	error
	arg string
}

func (e *ErrRequired) Error() string {
	return fmt.Sprintf("%s is required!", e.arg)
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("%s is invalid!", e.arg)
}

func NewRegexUtil() *RegexUtil {
	r := RegexUtil{}
	if e, err := regexp.Compile("(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)"); err != nil {
		fmt.Errorf("Couldn't compile regex! %v", err.Error())
	} else {
		r.email = e
	}
	return &r
}

func (r *RegexUtil) MatchEmail(email string) bool {
	return r.email.Match([]byte(email))
}
