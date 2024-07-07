
package forms

// import (
// 	"crypto/internal/edwards25519/field"

// 	"golang.org/x/text/message"
// )

type errors map[string][]string
//add the error message for the given field
func (e errors) Add(field, message string){
	e[field]=append(e[field],message)
}

//get returns the first errror message 
func(e errors) Get(field string)string{
	es :=e[field]
	if len(es) == 0{
		return ""
	}
	return es[0]
}