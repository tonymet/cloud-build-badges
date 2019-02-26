package shields

const badgePrefix = "GCP Build | "

type BadgesStruct struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

func (b BadgesStruct) FromStatus(statusText string) BadgesStruct {
	b.Message = statusText
	b.Color = shieldColor(statusText)
	return b
}

func (b *BadgesStruct) SetLabel(label string) {
	b.Label = badgePrefix + label
}

func New() BadgesStruct {
	return BadgesStruct{
		SchemaVersion: 1,
		Color:         shieldColor(""),
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
