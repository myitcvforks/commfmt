package pkg

import (
	"strings"
	"unicode"
)

// justify  takes  an  input  and  a  desired  line  length and returns a
// fully-justified  version  of  the  text.  Taken  wholesale  from Malik
// Browne's   great   JavaScript   walkthrough   that   does   the   same
// https://www.malikbrowne.com/blog/text-justification-coding-question.
// This version adds support for godoc code examples.
func (c *Config) Justify(input string) (string, error) {
	paragraphs := strings.Split(input, "\n\n")

	builder := newErrWriter()
	for i, p := range paragraphs {
		if strings.HasPrefix(p, " ") || strings.HasPrefix(p, "\t") {
			builder.write(p)
		} else {
			builder.write(c.justifyParagraph(p))
		}

		if i < len(paragraphs)-1 {
			builder.write("\n\n")
		}
	}

	if err := builder.error(); err != nil {
		return "", err
	}

	return addComments(builder.string())
}

func addComments(input string) (string, error) {
	builder := newErrWriter()
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		builder.write("//")
		line := strings.TrimRightFunc(line, unicode.IsSpace)
		if len(line) > 0 {
			builder.write(" ")
			builder.write(line)
		}

		if i < len(lines)-1 {
			builder.write("\n")
		}
	}

	if err := builder.error(); err != nil {
		return "", err
	}

	return builder.string(), nil
}

func (c *Config) justifyParagraph(input string) string {
	words := strings.Fields(input)

	lines := []string{}
	index := 0

	for index < len(words) {
		count := len(words[index])
		last := index + 1

		for last < len(words) {
			if len(words[last])+count+1 > c.Width {
				break
			}
			count += len(words[last]) + 1
			last++
		}

		var line string
		difference := last - index - 1
		if last == len(words) || difference == 0 {
			for i := index; i < last; i++ {
				line += words[i] + " "
			}

			line = line[0 : len(line)-1]
			for i := len(line); i < c.Width; i++ {
				line += " "
			}
		} else {
			spaces := (c.Width - count) / difference
			remainder := (c.Width - count) % difference

			for i := index; i < last; i++ {
				line += words[i]

				if i < last-1 {
					limit := spaces
					if (i - index) < remainder {
						limit++
					}

					for j := 0; j <= limit; j++ {
						line += " "
					}
				}
			}
		}

		lines = append(lines, strings.TrimRight(line, " "))
		index = last
	}

	return strings.Join(lines, "\n")
}
