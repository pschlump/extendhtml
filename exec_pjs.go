package template

import (
	template "github.com/pschlump/extend"
)

// AvailableTemplates returns a array string listing the defined templates by name,
// If there are none, it returns an empty array.  (PJS-Wed Jun 15 12:47:01 MDT 2016)
// PJS - Ported to go v1.12 - Wed Mar 20 07:04:09 MDT 2019
func (t *Template) AvailableTemplates() (rv []string) {
	return t.text.AvailableTemplates()
}

func SetNoValue(s string) {
	template.NoValue = s
}
