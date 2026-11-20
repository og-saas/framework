package shorturl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
	Domain  string `json:"domain"`
}

type ShortenResponse struct {
	Link string `json:"link"`
}

func ShortURL(token, longURL string) (string, error) {
	body, _ := json.Marshal(ShortenRequest{
		LongURL: longURL,
		Domain:  "bit.ly",
	})

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api-ssl.bitly.com/v4/shorten",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated {
		data, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("bitly error: %s", data)
	}

	var result ShortenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Link, nil
}
