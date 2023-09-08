package thenextleg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const host = "https://api.thenextleg.io/v2/"

type Client struct {
	authToken string
}

func NewClient(authToken string) *Client {
	return &Client{authToken: authToken}
}

type ImagineRequest struct {
	Msg             string `json:"msg"`
	Ref             string `json:"ref"`
	WebhookOverride string `json:"webhookOverride"`
	IgnorePrefilter bool   `json:"ignorePrefilter"`
}
type ImagineResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"createdAt"`
}

func (self *Client) Imagine(r ImagineRequest) (ImagineResponse, error) {
	var result ImagineResponse
	err := self.postJson("/imagine", r, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}

type ImagineResultResponse struct {
	CreatedAt            time.Time `json:"createdAt"`
	Buttons              []string  `json:"buttons"`
	Type                 string    `json:"type"`
	ImageUrl             string    `json:"imageUrl"`
	ImageUrls            []string  `json:"imageUrls"`
	ButtonMessageId      string    `json:"buttonMessageId"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              string    `json:"content"`
	Ref                  string    `json:"ref"`
	ResponseAt           time.Time `json:"responseAt"`
}

func (self *Client) ImagineResult(messageId string) (ImagineResultResponse, error) {
	var response ImagineResultResponse
	if err := self.get("/message/"+messageId, nil, response); err != nil {
		return response, err
	}
	return response, nil
}

type DescribeRequest struct {
	Url             string `json:"url"`
	Ref             string `json:"ref"`
	WebhookOverride string `json:"webhookOverride"`
	IgnorePrefilter bool   `json:"ignorePrefilter"`
}

type DescribeResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (self *Client) Describe(r DescribeRequest) (DescribeResponse, error) {
	var response DescribeResponse
	if err := self.postJson("/describe", r, response); err != nil {
		return response, err
	}
	return response, nil
}

type DescribeResultResponse struct {
	CreatedAt            time.Time `json:"createdAt"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              []string  `json:"content"`
	Ref                  string    `json:"ref"`
	Type                 string    `json:"type"`
	ResponseAt           time.Time `json:"responseAt"`
}

func (self *Client) DescribeResult(messageId string) (DescribeResultResponse, error) {
	var response DescribeResultResponse
	if err := self.get("/message/"+messageId, nil, response); err != nil {
		return response, err
	}
	return response, nil
}

func (self *Client) get(path string, request url.Values, response interface{}) error {
	u := host + path
	if len(request) > 0 {
		u += "?" + request.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+self.authToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}
	return nil
}

func (self *Client) postJson(path string, request interface{}, response interface{}) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(request); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, host+path, &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+self.authToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
