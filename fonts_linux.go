package fonts

// The list of font paths on Linux
// this could be improved to use fontconfig
var FontPaths = []string{
	"$HOME/.fonts/**/*",
	"/usr/local/share/fonts/**/*",
	"/usr/share/fonts/**/*",
	"/usr/X11/lib/X11/fonts/**/*",
}
