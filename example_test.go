package fonts

import (
	"code.google.com/p/freetype-go/freetype"
	"fmt"
	"io/ioutil"
)

func ExampleLoad() {
	fontReader, _ := Load("Helvetica", "DejaVuSansMono")
	fontBytes, _ := ioutil.ReadAll(fontReader)
	_, err := freetype.ParseFont(fontBytes) // "code.google.com/p/freetype-go/freetype"
	fmt.Println(err)
	// Output: <nil>
}
