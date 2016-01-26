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
		Links []string `json:"links"`
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
