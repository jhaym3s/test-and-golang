package dictionary

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")
func (d Dictionary) Search(word string) (string,error) {
	definition, ok := d[word]
	if !ok {
		return "", errNotFound
	}

	return definition,nil
}