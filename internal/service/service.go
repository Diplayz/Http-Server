package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	if data == "" {
		return "", fmt.Errorf("ошибка: пустая строка")
	}

	clean := strings.TrimSpace(data)
	clean = strings.Trim(clean, "\n\r\t")

	converter := morse.NewConverter(
		morse.DefaultMorse,
		morse.WithCharSeparator(" "),
		morse.WithWordSeparator(" "),
		morse.WithLowercaseHandling(true),
		morse.WithHandler(morse.IgnoreHandler),
		morse.WithTrailingSeparator(false),
	)

	textResult := converter.ToText(clean)
	if textResult != "" && textResult != clean {
		return textResult, nil
	}

	morseResult := converter.ToMorse(clean)
	if morseResult != "" {
		return morseResult, nil
	}

	return "", fmt.Errorf("ошибка конвертации")
}
