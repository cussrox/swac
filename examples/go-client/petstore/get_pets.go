// Code generated by github.com/swaggest/swac v0.0.1, DO NOT EDIT.

package petstore

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetPetsRequest is operation request value.
type GetPetsRequest struct {
	// Tags is an optional `tags` parameter in query.
	// tags to filter by
	Tags  []string
	// Limit is an optional `limit` parameter in query.
	// maximum number of results to return
	Limit *int64
}

// encode creates *http.Request for request data.
func (request *GetPetsRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestUri := baseURL + "/pets"
	query := make(url.Values, 2)
	query.Set("tags", strings.Join(request.Tags, ","))
	if request.Limit != nil {
		query.Set("limit", strconv.FormatInt(*request.Limit, 10))
	}
	if len(query) > 0 {
		requestUri += "?" + query.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, requestUri, nil)
	if err != nil {
	    return nil, err // add wrap with "failed to create request
	}
	req = req.WithContext(ctx)
	return req, err
}

// GetPetsResponse is operation response value.
type GetPetsResponse struct {
	StatusCode int
	OK         []Pet   // OK is a value of 200 OK response.
	Default    *Error  // Default is a default value of response.
}

// decode loads data from *http.Response.
func (result *GetPetsResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.OK)
	    if err != nil {
	        return err // add wrap with "failed to decode response"
	    }
	default:
	    err := json.NewDecoder(resp.Body).Decode(&result.Default)
	    if err != nil {
	        return err // add wrap with "failed to decode response"
	    }
	}
	return nil

}

// GetPets performs REST operation.
func (c *Client) GetPets(ctx context.Context, request GetPetsRequest) (GetPetsResponse, error) {
	result := GetPetsResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/pets")
	if c.Timeout != 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, c.Timeout)
		defer cancel()
	}
	req, err := request.encode(ctx, c.BaseURL)
	if err != nil {
	    return result, err // add wrap with "failed to create request"
	}
	resp, err := c.transport.RoundTrip(req)

	if err != nil {
	    return result, err // add wrap with "failed to send request"
	}
	defer resp.Body.Close()
	err = result.decode(resp)
	if err != nil {
	    return result, err // add wrap with "failed to decode response"
	}
	return result, nil
}
