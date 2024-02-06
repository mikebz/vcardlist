/*
Copyright © 2024 Mike Borozdin mikebz@
*/

package parse

import (
	"os"

	vcard "github.com/emersion/go-vcard"
)

type NameEmail struct {
	Name  string
	Email string
}

// Parse takes the directory, finds all the vCard files
// in it and then returns the resulting structs
func Parse(dir string) ([]NameEmail, error) {
	var nameEmails []NameEmail

	return nameEmails, nil
}

// Parse an individual vCard file.  Return an error if
// the file cannot be opened or parsed.
func ParseFile(filename string) (*NameEmail, error) {
	var nameEmail NameEmail

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := vcard.NewDecoder(file)
	card, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	// retrieve the values.  Name and email.
	// Try the preferred value first in both cases.
	nameEmail.Name = card.PreferredValue(vcard.FieldFormattedName)
	if nameEmail.Name == "" {
		nameEmail.Name = card.Value(vcard.FieldFormattedName)
	}
	nameEmail.Email = card.PreferredValue(vcard.FieldEmail)
	if nameEmail.Email == "" {
		nameEmail.Email = card.Value(vcard.FieldEmail)
	}
	return &nameEmail, nil
}
