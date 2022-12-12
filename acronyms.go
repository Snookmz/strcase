package strcase

import "sync"

var uppercaseAcronym = sync.Map{}

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	uppercaseAcronym.Store("ID", "Id")
	uppercaseAcronym.Store(key, val)
}
