// Package fonts provides a cross-platform interface to list and load font data
package fonts

import "io"

// Font holds the name of a font and an optional string indicating the font's source
type Font struct {
	Name   string
	Source string
}

// List the available fonts
func List() ([]Font, error) {
	return listFonts()
}

// Load attempts to load font file data given font names
func Load(fontNames ...string) (io.Reader, error) {
	for _, fontName := range fontNames {
		b, e := loadFont(fontName)
		if e == nil {
			return b, e
		}
	}
	if len(fontNames) > 1 {
		return nil, errNoneFoundPlural
	}
	return nil, errNoneFound
}

// Load attempts to load the font's data from its Name
func (f Font) Load() (io.Reader, error) {
	return Load(f.Name)
}

// ValidExtensions is the list of file extensions to consider font files
var ValidExtensions = []string{"ttf", "otf"}
