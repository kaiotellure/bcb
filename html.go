package main

import (
	"strconv"
	"strings"
)

type ToHTML struct {
	lines string
}

func (i *ToHTML) handleCommand(line string) string {
	splitted := strings.Split(line, ":")
	command := splitted[0][1:]
	args := splitted[1:]

	switch command {
	case "CHAPTER":
		panicif(len(args) < 1, "chapter requires 1 arguments: "+line)
		return i.uiChapter(args[0])
	}

	panic("unknown command: " + line)
}

var odd bool
var linecount int

func (i *ToHTML) FeedLine(line string) {

	if strings.HasPrefix(line, "$") {
		i.lines += i.handleCommand(line)
		return
	}

	linecount++
	id := strconv.Itoa(linecount)
	index := f("<a id=\"%s\" href=\"javascript:save(%s)\" class=\"index\">%s</a>", id, id, id)

	if odd = !odd; odd {
		i.lines += f("<div class=\"line muted\">%s%s</div>\n", index, line)
		return
	}

	i.lines += f("<div class=\"line\">%s%s</div>\n", index, line)
}

const CHAPTER_MODEL = `
<div class="chapter" style="margin: 5vh 0; gap: 1vw; display: flex; justify-content: space-around; align-items: center;">
	<img style="width: 5vw;" src="$$ORNAMENT$$">
	<span style="text-transform: uppercase; font-family: serif; font-weight: bold;" class="text">$$TITLE$$</span>
	<div style="height: 1px; width: 100%; background: khaki;"></div>
	<img style="width: 5vw; rotate: 180deg;" src="$$ORNAMENT$$">
</div>`

func (i *ToHTML) uiChapter(title string) string {
	return remap(map[string]string{
		"$$TITLE$$":    title,
		"$$ORNAMENT$$": dataUri("./assets/images/ornament.svg", "image/svg+xml"),
	}, CHAPTER_MODEL)
}
