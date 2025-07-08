package prompts

import "fmt"

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
	prompts, ok := languagePrompts[lang]
	if !ok {
		return fmt.Errorf("language %s not supported (en, ru)", lang)
	}

	promptMap = make(map[string]string)
	for key, prompt := range prompts {
		promptMap[key] = prompt.Template
	}

	currentLanguage = lang
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
