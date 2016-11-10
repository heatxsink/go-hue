package hue

type ApiResponse struct {
	Success map[string]interface{} `json:"success"`
	Error   *ApiResponseError      `json:"error"`
}

type ApiResponseError struct {
	Type        uint   `json:"type"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
