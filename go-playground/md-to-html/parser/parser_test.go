package parser

import (
	"strings"
	"testing"
)

func TestParseLine_SingleLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Heading 1",
			input:    "# Hello World",
			expected: "<h1>Hello World</h1>\n",
		},
		{
			name:     "Heading 3",
			input:    "### Subheading",
			expected: "<h3>Subheading</h3>\n",
		},
		{
			name:     "Simple Paragraph",
			input:    "Just some text.",
			expected: "<p>Just some text.</p>\n",
		},
		{
			name:     "Bold Line",
			input:    "**Bold Text**",
			expected: "<b>Bold Text</b>\n",
		},
		{
			name:     "Italic Line",
			input:    "*Italic Text*",
			expected: "<i>Italic Text</i>\n",
		},
		{
			name:     "Link Detection",
			input:    "[Google](#google)",
			expected: "<a href=\"#google\">Google</a>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			p := NewParser(&output)

			p.ParseLine(tt.input)

			result := output.String()
			if result != tt.expected {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, result)
			}
		})
	}
}

func TestParseLine_UnorderedList(t *testing.T) {
	inputLines := []string{
		"Paragraph before.",
		"* Item 1",
		"* Item 2",
		"Paragraph after.",
	}

	expectedOutput :=
		`<p>Paragraph before.</p>
<ul>
 <li>Item 1</li>
 <li>Item 2</li>
</ul>
<p>Paragraph after.</p>
`

	var output strings.Builder
	p := NewParser(&output)

	for _, line := range inputLines {
		p.ParseLine(line)
	}

	p.Finish()

	if output.String() != expectedOutput {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestParseLine_OrderedList(t *testing.T) {
	inputLines := []string{
		"1. First",
		"2. Second",
		"Not a number",
	}

	expectedOutput :=
		`<ol>
 <li>First</li>
 <li>Second</li>
</ol>
<p>Not a number</p>
`

	var output strings.Builder
	p := NewParser(&output)

	for _, line := range inputLines {
		p.ParseLine(line)
	}

	p.Finish()

	if output.String() != expectedOutput {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
