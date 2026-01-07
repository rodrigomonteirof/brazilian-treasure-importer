package tesouro

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	Result struct {
		Resources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"resources"`
	} `json:"result"`
}

func GetCSVUrl(apiUrl string) (string, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("failed to get API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var data apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("failed to decode JSON: %w", err)
	}

	for _, resource := range data.Result.Resources {
		if resource.Format == "CSV" {
			return resource.URL, nil
		}
	}

	return "", fmt.Errorf("CSV resource not found")
}

