package slack

import (
	"context"
	"encoding/json"
)

type AddUser struct {
	SlackID     string `json:"slack_id,omitempty"`
	ExternalID  string `json:"external_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
}

type AddCallRequest struct {
	ID                string    `json:"id"`
	CreatedBy         string    `json:"created_by"`
	DateStart         int       `json:"date_start"`
	ExternalUniqueID  string    `json:"external_unique_id"`
	JoinURL           string    `json:"join_url"`
	DesktopAppJoinURL string    `json:"desktop_app_join_url"`
	ExternalDisplayID string    `json:"external_display_id"`
	Title             string    `json:"title"`
	Users             []AddUser `json:"users"`
}

type AddCallResponse struct {
	SlackResponse
	Call struct {
		ID                string `json:"id"`
		DateStart         int    `json:"date_start"`
		ExternalUniqueID  string `json:"external_unique_id"`
		JoinURL           string `json:"join_url"`
		ExternalDisplayID string `json:"external_display_id"`
		Title             string `json:"title"`
	} `json:"call"`
}

// AddCall will register a call and return the ID as a string
func (api *Client) AddCall(req AddCallRequest) (string, error) {
	encoded, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	endpoint := api.endpoint + "calls.add"
	resp := &AddCallResponse{}
	err = postJSON(context.Background(), api.httpclient, endpoint, api.token, encoded, resp, api)
	if err != nil {
		return "", err
	}
	return resp.Call.ID, resp.Err()
}

// EndCall will end a registered call
func (api *Client) EndCall(id string) error {
	encoded, err := json.Marshal(struct {
		ID string `json:"id"`
	}{id})
	if err != nil {
		return err
	}
	endpoint := api.endpoint + "calls.end"
	resp := &AddCallResponse{}
	err = postJSON(context.Background(), api.httpclient, endpoint, api.token, encoded, resp, api)
	if err != nil {
		return err
	}
	return resp.Err()
}
