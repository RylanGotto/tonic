package openai

type NewsFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Type       string `json:"type"`
		Properties struct {
			Query struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"query"`
			Category struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"category"`
			Source struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Default     any    `json:"default"`
			} `json:"source"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}

type GoogleFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Type       string `json:"type"`
		Properties struct {
			Query struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"query"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}

type WikipediaFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Type       string `json:"type"`
		Properties struct {
			Query struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"query"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}

type FetchFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  []struct {
		Type       string `json:"type"`
		Properties struct {
			URL struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"url"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}

type OmniSearchFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  struct {
		Type       string `json:"type"`
		Properties struct {
			Query struct {
				Type        string `json:"type"`
				Description string `json:"description"`
			} `json:"query"`
		} `json:"properties"`
		Required []string `json:"required"`
	} `json:"parameters"`
}
