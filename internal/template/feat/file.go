package feat

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FileContent(name string, template string) []byte {
	caser := cases.Title(language.AmericanEnglish)
	nameSingle, _ := strings.CutSuffix(name, "s")

	content := fmt.Sprintf(template, name, caser.String(name), caser.String(nameSingle))

	return []byte(content)
}
