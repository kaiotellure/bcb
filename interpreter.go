package main

import (
	"strconv"
	"strings"
)

var alternate bool

type HTMLInterpreter struct {
	lines      string
	line_count int
}

func NewHTMLInterpreter() *HTMLInterpreter {
	return &HTMLInterpreter{}
}

func (i *HTMLInterpreter) GetDocument() string {
	return remap(map[string]string{
		"$$LOGO$$":        dataUri("./assets/images/logo.png", "image/jpg"),
		"$$FONTREGULAR$$": dataUri("./assets/fonts/main_regular.woff2", "font/woff2"),
		"$$FONTBOLD$$":    dataUri("./assets/fonts/main_bold.woff2", "font/woff2"),
		"$$LINES$$":       i.lines,
	}, readFile("./assets/template.html"))
}

func (i *HTMLInterpreter) insert(line string) {
	i.lines += line
}

func (i *HTMLInterpreter) on_command(line string) bool {
	command, args := parse_command(line)
	switch command {
	case "CHAPTER":
		panicif(len(args) < 1, "chapter requires 1 arguments: "+line)
		i.insert(html_chapter(args[0]))
		return true // should return
	default:
		panic("unknown command: " + line)
	}
}

func (i *HTMLInterpreter) on_line(line string) bool {
	i.line_count++
	index := html_index(strconv.Itoa(i.line_count))

	var muted string
	if alternate = !alternate; alternate {
		muted = " muted"
	}

	i.insert(f("<div class=\"line%s\">%s%s</div>\n", muted, index, line))
	return false
}

func (i *HTMLInterpreter) report(kind, value string) {
	switch kind {
	case "cut":
		if strings.HasPrefix(value, "$") {
			if i.on_command(value) {
				return
			}
		}
		if i.on_line(value) {
			return
		}
	}
}

func parse_command(line string) (string, []string) {
	splitted := strings.Split(line, ":")
	return splitted[0][1:], splitted[1:]
}

const CHAPTER_MODEL = `
<div class="chapter">
	$$ORNAMENT$$
	<span class="uppercase bold">$$TITLE$$</span>
	<div id="separator"></div>
	$$ORNAMENT$$
</div>`

func html_chapter(title string) string {
	return remap(map[string]string{
		"$$TITLE$$":    title,
		"$$ORNAMENT$$": readFile("./assets/images/ornament.svg"),
	}, CHAPTER_MODEL)
}

func html_index(id string) string {
	return f(
		"<a id=\"%s\" href=\"javascript:save(%s)\" class=\"bold muted mr\">%s</a>",
		id, id, id,
	)
}
