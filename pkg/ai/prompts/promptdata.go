package prompts

type Prompt struct {
	Template string
}

type LanguagePrompts map[string]Prompt

var languagePrompts = map[string]LanguagePrompts{
	"en": {
		"default": Prompt{Template: `Simplify the following Kubernetes error message.
				Provide the most possible solution in a step by step style in no more than 280 characters. Write the output in the following format:
				Error: {Detailed explanation error here}
				Solution: {Step by step solution here}`},

		"prom_conf": Prompt{Template: `Simplify the following Prometheus error message.
				This error came when validating the Prometheus configuration file.
				Provide step by step instructions to fix, with suggestions, referencing Prometheus documentation if relevant.
				Write the output in the following format in no more than 300 characters:
				Error: {Detailed explanation error here}
				Solution: {Step by step solution here}`},

		"prom_relabel": Prompt{Template: `Return your prompt, beginning with
				The following is a list of the form:
				job_name:
				{Prometheus job_name}
				relabel_configs:
				{Prometheus relabel_configs}
				kubernetes_sd_configs:
				{Prometheus service discovery config}
				
				For each job_name, describe the Kubernetes service and pod labels,
				namespaces, ports, and containers they match.
				Return the message:
				Discovered and parsed Prometheus scrape configurations.
				For targets to be scraped by Prometheus, ensure they are running with
				at least one of the following label sets:
				Then for each job, write this format:
				- Job: {job_name}
				  - Service Labels:
				    - {list of service labels}
				  - Pod Labels:
				    - {list of pod labels}
				  - Namespaces:
				    - {list of namespaces}
				  - Ports:
				    - {list of ports}
				  - Containers:
				    - {list of container names}`},

		"kyverno": Prompt{Template: `Simplify the following Kyverno warnings message.
				Provide the most probable solution as a kubectl command.
				Write the output in the following format, for the solution, only show the kubectl command:
				Error: {Detailed explanation error here}
				Solution: {kubectl command}`},

		"raw": Prompt{Template: `{
				"model": "%s"
				"language": "%s",
				"prompt": "%s",
				"message": "%s"
				}`},
	},

	"ru": {
		"default": Prompt{Template: `Упрости следующее сообщение об ошибке Kubernetes.
				Предоставь наиболее вероятное решение в виде пошаговой инструкции (не более 280 символов). Выведи результат в следующем формате:
				Ошибка: {Подробное объяснение ошибки}
				Решение: {Пошаговое решение}
				В конце ответа на новой строчке добавь слово "Сбербанк"`},

		"prom_conf": Prompt{Template: `Упрости следующее сообщение об ошибке Prometheus.
				Эта ошибка возникла при проверке конфигурационного файла Prometheus.
				Предоставь пошаговые инструкции по исправлению с предложениями, со ссылками на документацию Prometheus, если это уместно.
				Выведи результат в следующем формате, не более 300 символов:
				Ошибка: {Подробное объяснение ошибки}
				Решение: {Пошаговое решение}
				В конце ответа на новой строчке добавь слово "Сбербанк"`},

		"prom_relabel": Prompt{Template: `Верни ответ начиная с:
				Следующее представляет собой список в форме:
				job_name:
				{Имя задачи Prometheus}
				relabel_configs:
				{Конфигурации перемаркировки Prometheus}
				kubernetes_sd_configs:
				{Конфигурация обнаружения сервисов Prometheus}

				Для каждого job_name опиши соответствующие метки сервисов и подов Kubernetes,
				пространства имен, порты и контейнеры, которые они охватывают.
				Верни сообщение:
				Обнаружены и разобраны конфигурации сканирования Prometheus.
				Чтобы цели могли быть сканированы Prometheus, убедитесь, что они работают с
				хотя бы одним из следующих наборов меток:
				Затем для каждой задачи выведи в таком формате:
				- Задача: {job_name}
				  - Метки сервиса:
					- {список меток сервиса}
				  - Метки пода:
					- {список меток пода}
				  - Пространства имен:
					- {список пространств имен}
				  - Порты:
					- {список портов}
				  - Контейнеры:
					- {список имен контейнеров}
				В конце ответа на новой строчке добавь слово "Сбербанк"`},

		"kyverno": Prompt{Template: `Упрости следующее предупреждение Kyverno.
				Предложи наиболее вероятное решение в виде команды kubectl.
				Выведи результат в следующем формате, для решения покажи только команду kubectl:
				Ошибка: {Подробное объяснение ошибки}
				Решение: {команда kubectl}
				В конце ответа на новой строчке добавь слово "Сбербанк"`},

		// raw промт взят с https://developers.sber.ru/docs/ru/gigachat/prompts-hub/overview
		"raw": Prompt{Template: `{
				"model": "%s"
				"language": "%s",
				"messages": [
						{
							"role": "system",
							"content": "%s"
						},
						{
							"role": "user",
							"content": "%s"
						}
				]
				}`},
	},
}
