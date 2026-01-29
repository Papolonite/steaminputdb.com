package models

type ControllerType struct {
	Base Base `bun:",embed"`

	Name         string  `bun:",unique,notnull"`
	ReadableName *string `bun:","`
}
