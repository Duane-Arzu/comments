<<<<<<< HEAD
// Filename: internal/validator/validator.go
package validator

import (
	"regexp"
	"slices"
)

// We will create a new type named Validator
=======
package validator

// new type named Validator
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
type Validator struct {
	Errors map[string]string
}

<<<<<<< HEAD
var EmailRX = regexp.MustCompile(
	"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Construct a new Validator and return a pointer to it
// All validation errors go into this one Validator instance
=======
// construct new validator and return a pointer to it
// all validation errors go into thie one Validator instance
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

<<<<<<< HEAD
// Let's check  to see if the Validator's map contains any entries
=======
// checking to see if the Validator's map contains any entries
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
func (v *Validator) IsEmpty() bool {
	return len(v.Errors) == 0
}

// Add a new error entry to the Validator's error map
<<<<<<< HEAD
// Check first if an entry with the same key does not already exist
=======
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
func (v *Validator) AddError(key string, message string) {
	_, exists := v.Errors[key]
	if !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(acceptable bool, key string, message string) {
	if !acceptable {
		v.AddError(key, message)
	}
}
<<<<<<< HEAD

func PermittedValue(value string, permittedValues ...string) bool {
	return slices.Contains(permittedValues, value)
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
=======
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
