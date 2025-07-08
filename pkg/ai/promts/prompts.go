package promts

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Prompts struct {
	Prompts map[string]struct {
		Template string `yaml:"template"`
	} `yaml:"prompts"`
}

var (
	promptMap       map[string]string
	currentLanguage string = "en" // по умолчанию
)

func init() {
	// Загрузка промтов по умолчанию
	err := loadPrompts("en")
	if err != nil {
		return
	}
}

var localesFS embed.FS

func loadPrompts(lang string) error {
	data, err := localesFS.ReadFile(fmt.Sprintf("locales/%s.yaml", lang))
	if err != nil {
		return err
	}

	var prompts Prompts
	err = yaml.Unmarshal(data, &prompts)
	if err != nil {
		return err
	}

	promptMap = make(map[string]string)
	for key, prompt := range prompts.Prompts {
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
