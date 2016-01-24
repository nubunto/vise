package api

import (
	"github.com/nubunto/vise/persistence/types"
)

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
		Links []types.Link `json:"links"`
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
