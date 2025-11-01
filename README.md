# go-render-delta-json

Package based in [https://github.com/dchenk/go-render-quill](https://github.com/dchenk/go-render-quill)

Package `quill` takes a Quill-based Delta (https://github.com/quilljs/delta) as a JSON array of `insert` operations
and renders the defined HTML document.

Complete documentation at GoDoc: [https://godoc.org/github.com/chocobits/go-delta-json-to-html](https://godoc.org/github.com/chocobits/go-delta-json-to-html)

## Usage

```go
package main

import (
	"fmt"

	dj "github.com/chocobits/go-delta-json-to-html"
)

func main() {
	var delta = `[{"insert":"This "},{"attributes":{"italic":true},"insert":"is"},
    {"insert":" "},{"attributes":{"bold":true},"insert":"Awesome!"},{"insert":"\n"}]`

	html, err := dj.Render([]byte(delta))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(html))
}
// Output: <p>This <em>is</em> <strong>Awesome!</strong></p>
```

## Supported Formats

### Inline
 - Background color
 - Bold
 - Text color
 - Italic
 - Link
 - Size
 - Strikethrough
 - Superscript/Subscript
 - Underline

### Block
 - Blockquote
 - Header
 - Indent
 - List (ul and ol, including nested lists)
 - List (ul and ol, including nested lists)
	 - Additionally, ordered lists may include subtypes via the Delta `list` attribute. For example,
		 `{"list":"roman"}` will be rendered as an ordered list with a `data-list-type` attribute:
		 `<ol data-list-type="roman">...</ol>`. Alignment attributes on list items will be rendered as classes
		 on the `<li>` elements (for example `align-justify`).
 - Text alignment
 - Code block

### Embeds
 - Image (an inline format)

## Extending

The simple `Formatter` interface is all you need to implement for most block and inline formats. Instead of `Render` use `RenderExtended`
and provide a function that returns a `Formatter` for inserts that have the format you need.

For more control, you can also implement `FormatWriter` or `FormatWrapper`.
