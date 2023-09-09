package thenextleg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const host = "https://api.thenextleg.io"

type Client struct {
	authToken string
}

func NewClient(authToken string) *Client {
	return &Client{authToken: authToken}
}

func (c *Client) Imagine(r ImagineRequest) (ImagineResponse, error) {
	var response ImagineResponse
	if err := c.postJson("/v2/imagine", r, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) ImagineProgress(messageId string) (ImagineProgress, error) {
	var response ImagineProgress
	if err := c.get("/v2/message/"+messageId, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) Describe(r DescribeRequest) (DescribeResponse, error) {
	var response DescribeResponse
	if err := c.postJson("/v2/describe", r, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) DescribeProgress(messageId string) (DescribeProgress, error) {
	var response DescribeProgress
	if err := c.get("/v2/message/"+messageId, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) Blend(r BlendRequest) (BlendResponse, error) {
	var response BlendResponse
	if err := c.postJson("/v2/blend", r, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) BlendProgress(messageId string) (BlendProgress, error) {
	var response BlendProgress
	if err := c.get("/v2/message/"+messageId, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) FaceSwap(r FaceSwapRequest) (io.Reader, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(r); err != nil {
		return nil, err
	}
	res, err := c.post(body, "/face-swap")
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (c *Client) GetImage(r GetImageRequest) (io.Reader, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(r); err != nil {
		return nil, err
	}
	res, err := c.post(body, "/getImage")
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (c *Client) IsThisNaughty(r IsThisNaughtyRequest) (IsThisNaughtyResponse, error) {
	var response IsThisNaughtyResponse
	err := c.postJson("/v2/is-this-naughty", r, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) Info(r ...InfoRequest) (InfoResponse, error) {
	var request InfoRequest
	if len(r) > 0 {
		request = r[0]
	}
	var response InfoResponse
	values := url.Values{}
	if request.WebhookOverride != "" {
		values.Set("webhookOverride", request.WebhookOverride)
	}
	if request.Ref != "" {
		values.Set("ref", request.Ref)
	}
	if err := c.get("/v2/info", values, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) InfoProgress(messageId string) (InfoProgress, error) {
	var response InfoProgress
	if err := c.get("/v2/message/"+messageId, nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) get(path string, request url.Values, response interface{}) error {
	u := host + path
	if len(request) > 0 {
		u += "?" + request.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.authToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

func (c *Client) postJson(path string, request interface{}, response interface{}) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(request); err != nil {
		return err
	}
	res, err := c.post(body, path)
	if err != nil {
		return err
	}
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}
	return nil
}

func (c *Client) post(body bytes.Buffer, path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, host+path, &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.authToken)
	return http.DefaultClient.Do(req)
}
