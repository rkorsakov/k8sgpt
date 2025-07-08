package prompts

import (
	"embed"
	"gopkg.in/yaml.v2"
)

//go:embed locales/*.yaml
var localesFS embed.FS

type Prompts struct {
	Prompts map[string]struct {
		Template string `yaml:"template"`
	} `yaml:"prompts"`
}

var (
	promptMap       map[string]string
	currentLanguage string = "en" // default
)

func init() {
	// Load default prompts
	err := loadPrompts("en")
	if err != nil {
		return
	}
}

func loadPrompts(lang string) error {
	data, err := localesFS.ReadFile("locales/" + lang + ".yaml")
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
