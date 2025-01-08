package validator

import (
	"fmt"
	"net/http"
	"strings"
)

type Rules map[string][]string

type errors map[string]string

type validated map[string]string

type Validator struct {
	Rules     Rules
	errors    errors
	validated validated
	failed    bool
}

func (v *Validator) Validate(req *http.Request) bool {
	v.validated = make(validated)
	v.errors = make(errors)

	for name, rules := range v.Rules {

		for k := range rules {
			rule := strings.Split(rules[k], ":")

			switch rule[0] {
			case "required":
				if req.FormValue(name) == "" {
					v.errors[name] = "The " + name + " field is required"
				}
				break
				// case "min":
				// 	if utf8.RuneCountInString(req.FormValue(name)) < int(rule[1]) {
				// 		v.errors[name] = "The " + name + " must not be less then " + rule[1] + "chars"
				// 	}
				// 	break
				// case "max":
				// 	if utf8.RuneCountInString(req.FormValue(name)) > int(rule[1]) {
				// 		v.errors[name] = "The " + name + " must not be greater then " + rule[1] + "chars"
				// 	}
				// 	break
			}

			fmt.Println()

		}

		v.validated[name] = req.FormValue(name)
	}

	v.failed = len(v.errors) != 0

	return v.failed == false
}

func (v *Validator) Validated() validated {
	return v.validated
}

func (v *Validator) Errors() errors {
	return v.errors
}

func (v *Validator) Get(name string) string {
	value, ok := v.validated[name]

	if !ok {
		return ""
	}

	return value
}
