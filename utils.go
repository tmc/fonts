package fonts

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	errNoneFound       = errors.New("Font not found.")
	errNoneFoundPlural = errors.New("None of the specified fonts were found.")
)

// validFontFile determines if a file path appears to be a valid font
func validFontFile(fontFile string) bool {
	valid := false
	for _, ext := range ValidExtensions {
		valid = valid || ("."+ext == filepath.Ext(fontFile))
	}
	return valid
}

// listFonts produces a slice of `Font` based on "FontPaths" var
func listFonts() ([]Font, error) {
	result := make([]Font, 0)
	found := make(map[string]bool, 0)
	for _, p := range FontPaths {
		for _, ext := range ValidExtensions {
			matches, err := filepath.Glob(os.ExpandEnv(p) + ext)
			if err != nil {
				return nil, err
			}
			for _, fullPath := range matches {
				base := filepath.Base(fullPath)
				baseName := base[:len(base)-len(filepath.Ext(base))] // chop off ext
				if validFontFile(fullPath) && !found[baseName] {
					result = append(result, Font{baseName, p})
					found[baseName] = true
				}
			}
		}
	}
	return result, nil
}

// loadFont attempts to load the font with the given base name (without extension)
func loadFont(fontName string) (io.Reader, error) {
	for _, p := range FontPaths {
		for _, ext := range ValidExtensions {
			matches, err := filepath.Glob(os.ExpandEnv(p) + fontName + "." + ext)
			if err != nil {
				return nil, err
			}
			for _, match := range matches {
				return os.Open(match)
			}
		}
	}
	return nil, fmt.Errorf("Font '%s' not found.", fontName)
}
