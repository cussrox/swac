// Code generated by github.com/swaggest/swac v0.1.0, DO NOT EDIT.

package uspto

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PostDatasetVersionRecordsRequest is operation request value.
type PostDatasetVersionRecordsRequest struct {
	// Version is a required `version` parameter in path.
	// Version of the dataset.
	Version  string
	// Dataset is a required `dataset` parameter in path.
	// Name of the dataset. In this case, the default value is oa_citations
	Dataset  string
	Criteria string  // Criteria is a required `criteria` parameter in formData.
	Start    *int64  // Start is an optional `start` parameter in formData.
	Rows     *int64  // Rows is an optional `rows` parameter in formData.
}

// encode creates *http.Request for request data.
func (request *PostDatasetVersionRecordsRequest) encode(ctx context.Context, baseURL string) (*http.Request, error) {
	requestUri := baseURL + "/" + url.PathEscape(request.Dataset) + "/" + url.PathEscape(request.Version) + "/records"
	formData := make(url.Values, 3)
	formData.Set("criteria", request.Criteria)
	if request.Start != nil {
		formData.Set("start", strconv.FormatInt(*request.Start, 10))
	}
	if request.Rows != nil {
		formData.Set("rows", strconv.FormatInt(*request.Rows, 10))
	}
	var body io.Reader
	if len(formData) > 0 {
		body = strings.NewReader(formData.Encode())
	}
	req, err := http.NewRequest(http.MethodPost, requestUri, body)
	if err != nil {
	    return nil, err
	}
	req = req.WithContext(ctx)
	return req, err
}

// PostDatasetVersionRecordsResponse is operation response value.
type PostDatasetVersionRecordsResponse struct {
	StatusCode int
	OK         []map[string]map[string]interface{}  // OK is a value of 200 OK response.
	NotFound   interface{}                          // NotFound is a value of 404 Not Found response.
}

// decode loads data from *http.Response.
func (result *PostDatasetVersionRecordsResponse) decode(resp *http.Response) error {
	result.StatusCode = resp.StatusCode
	switch resp.StatusCode {
	case http.StatusOK:
	    err := json.NewDecoder(resp.Body).Decode(&result.OK)
	    if err != nil {
	        return err
	    }
	case http.StatusNotFound:
	    err := json.NewDecoder(resp.Body).Decode(&result.NotFound)
	    if err != nil {
	        return err
	    }
	default:
	    return errors.New("unexpected response status: " + resp.Status)
	}
	return nil

}

// PostDatasetVersionRecords performs REST operation.
func (c *Client) PostDatasetVersionRecords(ctx context.Context, request PostDatasetVersionRecordsRequest) (PostDatasetVersionRecordsResponse, error) {
	result := PostDatasetVersionRecordsResponse{}
	ctx = context.WithValue(ctx, "restOperationPath", "/{dataset}/{version}/records")
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
