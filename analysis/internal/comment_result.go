package internal

import "strings"

const (
	nameTag          = "name"
	countTag         = "count"
	sep              = "="
	defaultCountType = "uint8"
)

type Comment struct {
	CountType  string
	StructName string
}

func NewComment(comment, tag string) *Comment {
	tags := strings.Split(strings.Replace(comment, tag, "", 1),
		";")

	c := &Comment{}

	for _, a := range tags {
		if strings.Contains(a, nameTag) {
			c.StructName = strings.Replace(a, nameTag+sep, "", 1)
		}

		if strings.Contains(a, countTag) {
			c.CountType = strings.Replace(a, countTag+sep, "", 1)
		} else {
			c.CountType = defaultCountType
		}
	}

	return c
}
