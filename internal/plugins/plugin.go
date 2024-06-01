package plugins

type Plugin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Version     string `json:"version"`
}

var Plugins = []Plugin{
	{
		Name:        "GitHub",
		Description: "Scans GitHub repositories and integrates data with the service catalog.",
		Icon:        "/static/icons/github.png",
		Version:     "1.0.0",
	},
	{
		Name:        "Gitlab",
		Description: "Scans Gitlab repositories and integrates data with the service catalog.",
		Icon:        "/static/icons/gitlab.png",
		Version:     "1.0.0",
	},
}
