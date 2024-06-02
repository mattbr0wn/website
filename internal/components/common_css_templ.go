// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.709
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "strings"

func Base() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`height:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:REM, sans-serif;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-optical-sizing:auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-weight:200;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-style:normal;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Base`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Wrapper() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`max-width:75ch;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0 auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`min-height:80vh;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-top:30px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-bottom:10px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-left:5px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding-right:5px;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Wrapper`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func Main() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:Lora, serif;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-optical-sizing:auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:22;`)
	templ_7745c5c3_CSSBuilder.WriteString(`line-height:1.58;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex:1;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Main`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func BodyText() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:Lora, serif;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-optical-sizing:auto;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:22;`)
	templ_7745c5c3_CSSBuilder.WriteString(`line-height:1.58;`)
	templ_7745c5c3_CSSID := templ.CSSID(`BodyText`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func NoLinkStyles() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`text-decoration:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:inherit;`)
	templ_7745c5c3_CSSID := templ.CSSID(`NoLinkStyles`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}
