// Code generated by github.com/swaggest/swac v0.0.2, DO NOT EDIT.

package uber

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetMeRequest is operation request value.
type GetMeRequest struct {

}

// encode creates *http.Request for request data.
func (request *GetMeRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestUri := baseURL + "/me"
	req, err := http.NewRequest(http.MethodGet, requestUri, nil)
	if err != nil {
	    return nil, err
	}
	req = req.WithContext(ctx)
	return req, err
}

// GetMeResponse is operation response value.
type GetMeResponse struct {
	StatusCode int
	OK         *Profile  // OK is a value of 200 OK response.
	Default    *Error    // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetMeResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.OK)
	    if err != nil {
	        return err
	    }
	default:
	    err := json.NewDecoder(resp.Body).Decode(&result.Default)
	    if err != nil {
	        return err
	    }
	}
	return nil

}

// GetMe performs REST operation.
func (c *Client) GetMe(ctx context.Context, request GetMeRequest) (GetMeResponse, error) {
	result := GetMeResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/me")
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
