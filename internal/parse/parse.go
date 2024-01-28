/*
Copyright © 2024 Mike Borozdin mikebz@
*/

package parse

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
