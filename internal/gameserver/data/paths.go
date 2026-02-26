package data

import "path/filepath"

type Paths struct {
	Root string
}

func (p Paths) ExcelOutputDir() string { return filepath.Join(p.Root, "ExcelOutput") }
func (p Paths) ConfigDir() string      { return filepath.Join(p.Root, "Config") }
func (p Paths) TextMapDir() string     { return filepath.Join(p.Root, "TextMap") }
