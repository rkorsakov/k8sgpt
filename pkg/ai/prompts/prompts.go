package prompts

import (
	"fmt"
	"strings"
)

var (
	promptMap       map[string]string
	currentLanguage string = "en"
)

func init() {
	// Load default prompts
	err := loadPrompts("en")
	if err != nil {
		fmt.Printf("Error loading default prompts: %v\n", err)
		return
	}
}

func loadPrompts(lang string) error {
	normalizedLang := strings.ToLower(lang)

	// Маппинг альтернативных названий языков на стандартные коды
	languageAliases := map[string]string{
		"russian":    "ru",
		"русский":    "ru",
		"english":    "en",
		"английский": "en",
	}

	if code, ok := languageAliases[normalizedLang]; ok {
		normalizedLang = code
	}

	prompts, ok := languagePrompts[normalizedLang]
	// Если языка нет в промтах, используется en
	if !ok {
		prompts = languagePrompts["en"]
		normalizedLang = "en"
	}

	promptMap = make(map[string]string)
	for key, prompt := range prompts {
		promptMap[key] = prompt.Template
	}

	currentLanguage = normalizedLang
	return nil
}

func SetLanguage(lang string) error {
	return loadPrompts(lang)
}

func GetCurrentLanguage() string {
	return currentLanguage
}

func GetPrompt(key string) string {
	if prompt, ok := promptMap[key]; ok {
		return prompt
	}
	return promptMap["default"]
}

var PromptMap = map[string]string{
	"raw":                           "raw",
	"default":                       "default",
	"PrometheusConfigValidate":      "prom_conf",
	"PrometheusConfigRelabelReport": "prom_relabel",
	"PolicyReport":                  "kyverno",
	"ClusterPolicyReport":           "kyverno",
}
