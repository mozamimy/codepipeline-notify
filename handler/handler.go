package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type CodePipelineEventDetail struct {
	Pipeline    string      `json:"pipeline"`
	State       string      `json:"state"`
	ExecutionID string      `json:"execution-id"`
	Version     json.Number `json:"version"`
}

type CodePipelineEvent struct {
	Detail CodePipelineEventDetail `json:"detail"`
}

type SlackPayload struct {
	Text        string            `json:"text"`
	Username    string            `json:"username"`
	Icon_emoji  string            `json:"icon_emoji"`
	Icon_url    string            `json:"icon_url"`
	Channel     string            `json:"channel"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color  string       `json:"color"`
	Fields []SlackField `json:"fields"`
}

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func HandleRequest(codePipelineEvent CodePipelineEvent) {
	field := SlackField{
		Value: fmt.Sprintf("The state of pipeline `%s` is changed to `%s` (execution_id: %s)", codePipelineEvent.Detail.Pipeline, codePipelineEvent.Detail.State, codePipelineEvent.Detail.ExecutionID),
		Short: false,
	}
	colorMap := map[string]string{
		"CANCELED":   "warning",
		"FAILED":     "danger",
		"RESUMED":    "warning",
		"STARTED":    "good",
		"SUCCEEDED":  "good",
		"SUPERSEDED": "warning",
	}
	attachment := SlackAttachment{
		Color:  colorMap[codePipelineEvent.Detail.State],
		Fields: []SlackField{field},
	}
	params, _ := json.Marshal(SlackPayload{
		Username:    "CodePipeline",
		Icon_emoji:  os.Getenv("SLACK_EMOJI_ICON"),
		Channel:     os.Getenv("SLACK_CHANNEL"),
		Attachments: []SlackAttachment{attachment},
	})

	resp, _ := http.PostForm(
		os.Getenv("SLACK_WEBHOOK_URL"),
		url.Values{
			"payload": {string(params)},
		},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf(string(body))
}
