package types

type (
	FileInfo struct {
		Filename      string `json:"filename"`
		DaysAvailable int    `json:"days_available"`
	}
	UserFiles struct {
		Files []FileInfo `json:"files"`
	}
	Link struct {
		FileInfo
		URL string `json:"url"`
	}
)
