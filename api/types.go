package api

type (
	APIResponse struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}
	FileUploadResponse struct {
		APIResponse
		UserToken string `json:"user_token"`
	}
	LinksResponse struct {
		APIResponse
		Links []Link `json:"links"`
	}
)

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

var (
	ResponseOK = APIResponse{
		Ok:      true,
		Message: "OK",
	}
)

func (r *APIResponse) Error() string {
	return r.Message
}
