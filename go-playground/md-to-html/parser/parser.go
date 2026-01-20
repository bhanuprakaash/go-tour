package parser

import (
	"fmt"
	"strings"
	"unicode"
)

type Parser struct {
	InList  bool
	InOlist bool
	Output  *strings.Builder
}

func NewParser(outputWriter *strings.Builder) *Parser {
	return &Parser{
		Output: outputWriter,
	}
}

func (p *Parser) ParseLine(line string) {

	trimmed := strings.TrimSpace(line)

	// parse unordered list
	if done := p.parseUnorderdList(trimmed); done {
		return
	}

	// parse ordered list
	if done := p.parseOrderedList(trimmed); done {
		return
	}
	// parse heading tags
	if tag := p.parseHeadings(line); tag != "" {
		p.Output.WriteString(tag + "\n")
		return
	}

	// parse italics and bold
	if tag := p.parseItalicsAndBold(line); tag != "" {
		p.Output.WriteString(tag + "\n")
		return
	}

	// parse a tag
	if tag := p.parseLinks(line); tag != "" {
		p.Output.WriteString(tag + "\n")
		return
	}

	// parse p tag
	if tag := p.parseParagraphs(line); tag != "" {
		p.Output.WriteString(tag + "\n")
		return
	}

}

func (p *Parser) Finish() {
	if p.InList {
		p.Output.WriteString("</ul>\n")
		p.InList = false
	}

	if p.InOlist {
		p.Output.WriteString("</ol>\n")
		p.InOlist = false
	}
}

func (p *Parser) parseHeadings(line string) string {
	h := 0
	for h < len(line) && line[h] == '#' {
		h++
	}
	if h > 0 && h <= 6 && h < len(line) && line[h] == ' ' {
		heading := strings.TrimSpace(line[h:])
		return fmt.Sprintf("<h%d>%s</h%d>", h, heading, h)
	}

	return ""
}

func (p *Parser) parseItalicsAndBold(line string) string {
	b := 0
	var tag string
	for b < len(line) && line[b] == '*' {
		b++
	}
	if b > 0 && b <= 2 && b < len(line) {
		text := line[b : len(line)-b]
		switch b {
		case 1:
			tag = fmt.Sprintf("<i>%s</i>", text)
		case 2:
			tag = fmt.Sprintf("<b>%s</b>", text)
		}
		return tag
	}

	return ""
}

func (p *Parser) parseUnorderdList(trimmed string) bool {
	isListItem := strings.HasPrefix(trimmed, "* ") || strings.HasPrefix(trimmed, "- ")

	if isListItem {
		content := trimmed[2:]
		if !p.InList {
			p.Output.WriteString("<ul>\n")
			p.InList = true
		}
		fmt.Fprintf(p.Output, " <li>%s</li>\n", p.parseContent(content))
		return true
	}

	if p.InList {
		p.Output.WriteString("</ul>\n")
		p.InList = false
		return false
	}

	return false
}

func (p *Parser) parseOrderedList(trimmed string) bool {
	dotIndex := strings.Index(trimmed, ".")
	isNumberListItem := false

	if dotIndex > 0 && len(trimmed) > dotIndex+1 && trimmed[dotIndex+1] == ' ' {
		numberPart := trimmed[:dotIndex]
		isAllDigits := true
		for _, r := range numberPart {
			if !unicode.IsDigit(r) {
				isAllDigits = false
				break
			}
		}
		isNumberListItem = isAllDigits
	}

	if isNumberListItem {
		content := strings.TrimSpace(trimmed[dotIndex+1:])

		if !p.InOlist {
			p.Output.WriteString("<ol>\n")
			p.InOlist = true
		}
		fmt.Fprintf(p.Output, " <li>%s</li>\n", p.parseContent(content))
		return true
	}

	if p.InOlist {
		p.Output.WriteString("</ol>\n")
		p.InOlist = false
		return false
	}

	return false
}

func (p *Parser) parseParagraphs(line string) string {
	if strings.TrimSpace(line) != "" {
		return fmt.Sprintf("<p>%s</p>", line)
	}
	return ""
}

func (p *Parser) parseLinks(text string) string {
	var sb strings.Builder
	i := 0

	for i < len(text) {
		if text[i] == '[' {
			rest := text[i:]
			closeBracketIdx := strings.Index(rest, "]")

			if closeBracketIdx != -1 {
				absCloseBracket := i + closeBracketIdx

				if absCloseBracket+1 < len(text) && text[absCloseBracket+1] == '(' {
					restAfterBracket := text[absCloseBracket+1:]
					closeParenIdx := strings.Index(restAfterBracket, ")")

					if closeParenIdx != -1 {
						absCloseParen := absCloseBracket + 1 + closeParenIdx

						linkText := text[i+1 : absCloseBracket]
						linkUrl := text[absCloseBracket+2 : absCloseParen]

						sb.WriteString(fmt.Sprintf(`<a href="%s">%s</a>`, linkUrl, linkText))
						i = absCloseParen + 1
						continue
					}
				}
			}
		}
		i++
	}
	return sb.String()
}

func (p *Parser) parseContent(text string) string {

	if link := p.parseLinks(text); link != "" {
		return link
	}

	if tag := p.parseItalicsAndBold(text); tag != "" {
		return tag
	}

	return text
}
