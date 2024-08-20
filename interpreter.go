package main

import (
	"strconv"
	"strings"
)

type HTMLInterpreter struct {
	lines      string
	line_count int
	odd        bool
}

func (i *HTMLInterpreter) GetDocument() string {
	return remap(map[string]string{
		"$$LOGO$$":       dataUri("./assets/images/logo.png", "image/jpg"),
		"$$BACKGROUND$$": dataUri("./assets/images/background.jpg", "image/jpeg"),
		"$$FONTREGULAR$$": dataUri("./assets/fonts/main_regular.woff2", "font/woff2"),
		"$$FONTBOLD$$": dataUri("./assets/fonts/main_bold.woff2", "font/woff2"),
		"$$LINES$$":      i.lines,
	}, readFile("./assets/template.html"))
}

func (i *HTMLInterpreter) parseCommand(line string) (string, []string) {
	splitted := strings.Split(line, ":")
	return splitted[0][1:], splitted[1:]
}

func (i *HTMLInterpreter) insert(line string) {
	i.lines += line
}

func (i *HTMLInterpreter) Feed(line string) {

	if strings.HasPrefix(line, "$") {
		command, args := i.parseCommand(line)
		switch command {
		case "CHAPTER":
			panicif(len(args) < 1, "chapter requires 1 arguments: "+line)
			i.insert(i.uiChapter(args[0]))
			return
		default:
			panic("unknown command: " + line)
		}
	}

	i.line_count++
	id := strconv.Itoa(i.line_count)

	index := f("<a id=\"%s\" href=\"javascript:save(%s)\" class=\"index\">%s</a>", id, id, id)

	if i.odd = !i.odd; i.odd {
		i.insert(f("<div class=\"line muted\">%s%s</div>\n", index, line))
		return
	}
	i.insert(f("<div class=\"line\">%s%s</div>\n", index, line))
}

const CHAPTER_MODEL = `
<div class="chapter" style="margin: 5vh 0; gap: 1vw; display: flex; justify-content: space-around; align-items: center;">
	<img style="width: 5vw;" src="$$ORNAMENT$$">
	<span class="text">$$TITLE$$</span>
	<div style="height: 1px; width: 100%; background: khaki;"></div>
	<img style="width: 5vw; rotate: 180deg;" src="$$ORNAMENT$$">
</div>`

func (i *HTMLInterpreter) uiChapter(title string) string {
	return remap(map[string]string{
		"$$TITLE$$":    title,
		"$$ORNAMENT$$": dataUri("./assets/images/ornament.svg", "image/svg+xml"),
	}, CHAPTER_MODEL)
}
