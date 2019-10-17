// Code generated by github.com/swaggest/swac v0.0.3, DO NOT EDIT.

package uspto

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

// GetRequest is operation request value.
type GetRequest struct {

}

// encode creates *http.Request for request data.
func (request *GetRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestUri := baseURL + "/"
	req, err := http.NewRequest(http.MethodGet, requestUri, nil)
	if err != nil {
	    return nil, err
	}
	req = req.WithContext(ctx)
	return req, err
}

// GetResponse is operation response value.
type GetResponse struct {
	StatusCode int
	OK         *ComponentsSchemasDataSetList  // OK is a value of 200 OK response.
}

// decode loads data from *http.Response.
func (result *GetResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.OK)
	    if err != nil {
	        return err
	    }
	default:
	    return errors.New("unexpected response status: " + resp.Status)
	}
	return nil

}

// Get performs REST operation.
func (c *Client) Get(ctx context.Context, request GetRequest) (GetResponse, error) {
	result := GetResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/")
	if c.Timeout != 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, c.Timeout)
		defer cancel()
	}
	req, err := request.encode(ctx, c.BaseURL)
	if err != nil {
	    return result, err
	}
	resp, err := c.transport.RoundTrip(req)

	if err != nil {
	    return result, err
	}
	defer resp.Body.Close()
	err = result.decode(resp)
	if err != nil {
	    return result, err
	}
	return result, nil
}
