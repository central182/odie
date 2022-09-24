package cli

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/fatih/color"

	"github.com/central182/odie/internal/domain/application"
	"github.com/central182/odie/internal/domain/dictionary/entry"
	"github.com/central182/odie/internal/domain/dictionary/headword"
)

type InitApplication func() application.Application

type Printer struct {
	i InitApplication
}

func NewPrinter(i InitApplication) Printer {
	return Printer{i: i}
}

func (p Printer) PrintEntriesOfHeadword(hw string) (string, error) {
	h, herr := headword.New(hw)
	if herr != nil {
		return "", herr
	}

	es, eerr := p.i().GetEntriesByHeadword(h)

	t, err := template.New("entries").Funcs(
		template.FuncMap{
			"spellHeadword": func(h headword.Headword) string {
				return color.New(color.Bold).Sprint(h.Spelling())
			},
			"extractPronunciations": func(e entry.Entry) []entry.Pronunciation {
				ps, _ := e.Pronunciations()
				return ps
			},
			"showIpa": func(p entry.Pronunciation) string {
				return fmt.Sprintf("/%s/", p.PhoneticSpelling())
			},
			"nameLexicalCategory": func(lc entry.LexicalCategory) string {
				return color.New(color.Italic, color.FgYellow).Sprint(lc.Name())
			},
			"numberSense": func(i int) string {
				return fmt.Sprintf("%d", i+1)
			},
			"numberSubsense": func(i, j int) string {
				return fmt.Sprintf("%d.%d", i+1, j+1)
			},
			"calculatePaddingForExamplesOfSense": func(i int) string {
				return strings.Repeat(" ", len(fmt.Sprintf("%d", i+1))+1)
			},
			"calculatePaddingForExamplesOfSubsense": func(i, j int) string {
				return strings.Repeat(" ", len(fmt.Sprintf("%d.%d", i+1, j+1))+1)
			},
			"showExample": func(example string) string {
				return color.New(color.Italic, color.Faint).Sprint(example)
			},
			"extractExamplesOfSense": func(s entry.Sense) []string {
				es, _ := s.Examples()
				return es
			},
			"extractSubsensesOfSense": func(s entry.Sense) []entry.Subsense {
				ss, _ := s.Subsenses()
				return ss
			},
			"extractExamplesOfSubsense": func(s entry.Subsense) []string {
				es, _ := s.Examples()
				return es
			},
		},
	).Parse(`{{- /* no-op */ -}}
{{ range . -}}
{{ spellHeadword .Headword }} {{ if extractPronunciations . }}{{ range $i, $pronunciation := extractPronunciations . }}{{ if ne $i 0 }}, {{ end }}{{ showIpa $pronunciation }}{{ end }} {{ end }}{{ nameLexicalCategory .LexicalCategory }}
{{ range $i, $sense := .Senses }}
{{ numberSense $i }} {{ $sense.Description }}
{{- range extractExamplesOfSense $sense }}
{{ calculatePaddingForExamplesOfSense $i }}{{ showExample . }}
{{- end }}
{{ range $j, $subsense := extractSubsensesOfSense $sense }}
{{ numberSubsense $i $j }} {{ $subsense.Description }}
{{- range $ej, $example := extractExamplesOfSubsense $subsense }}
{{ calculatePaddingForExamplesOfSubsense $i $j }}{{ showExample $example }}
{{- end }}
{{ end -}}
{{ end }}
{{ end -}}
`)
	if err != nil {
		return "", err
	}

	output := new(bytes.Buffer)
	if err := t.Execute(output, es); err != nil {
		return "", err
	}

	return output.String(), eerr
}
