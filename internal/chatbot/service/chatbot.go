package service

import (
	dtos "bazar/internal/chatbot/DTOs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ChatbotService struct {
	info *dtos.Info
}

func NewChatbotService(jsonPath string) (*ChatbotService, error) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	var info dtos.Info
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &ChatbotService{info: &info}, nil
}

func (s *ChatbotService) Ask(question string) (string, error) {
	url := "https://api.ai21.com/studio/v1/j1-large/complete" // <-- endpoint correcto

	infoJSON, err := json.Marshal(s.info)
	if err != nil {
		return "", err
	}

	// Prompt que se usará
	prompt := fmt.Sprintf(
		"Eres un asistente que responde solo basado en esta información JSON: %s. "+
			"Si no puedes responder con la información disponible, di 'No sé'. "+
			"Pregunta: %s",
		string(infoJSON), question,
	)

	body := map[string]interface{}{
		"prompt":      prompt,
		"maxTokens":   200,
		"temperature": 0.7,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("AI21_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(respBody))

	var result struct {
		Completions []struct {
			Text string `json:"text"`
		} `json:"completions"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}
	if len(result.Completions) == 0 {
		return "", fmt.Errorf("no completions returned by AI21 API")
	}

	return result.Completions[0].Text, nil

}
