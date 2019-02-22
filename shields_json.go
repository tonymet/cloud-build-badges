package main

type BadgesStruct struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

func StatusToShieldJson(statusText string) BadgesStruct {
	return BadgesStruct{
		SchemaVersion: 1,
		Label:         "Cloud Builders",
		Message:       statusText,
		Color:         shieldColor(statusText),
	}
}

func shieldColor(status string) (color string) {
	switch status {
	case "SUCCESS":
		return "green"
	default:
		return "red"
	}
}
