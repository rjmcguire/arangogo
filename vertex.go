package arangogo

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CreateVertexResult struct {
	ID  string `json:"_id"`
	Key string `json:"_key"`
	Rev string `json:"_rev"`
}

type CreateVertexConfig struct {
	WaitForSync *bool
}

func (c *CreateVertexConfig) urlValues() url.Values {
	if c == nil {
		return nil
	}

	var params url.Values
	if c.WaitForSync != nil {
		params = make(url.Values)
		params.Set("waitForSync", strconv.FormatBool(*c.WaitForSync))
	}
	return params
}

func (c *Connection) CreateVertex(dbName, graphName, collName string, data interface{}, config *CreateVertexConfig) (r CreateVertexResult, rc int, err error) {
	path := buildPath(pathConfig{
		dbName:      dbName,
		pathFormat:  "/_api/gharial/%s/vertex/%s",
		pathParams:  []interface{}{graphName, collName},
		queryParams: config.urlValues(),
	})

	var body struct {
		Vertex CreateVertexResult `json:"vertex"`
	}
	rc, _, err = c.send(http.MethodPost, path, nil, data, &body)
	if err != nil {
		return body.Vertex, rc, fmt.Errorf("failed to create vertex: %v", err)
	}
	return body.Vertex, rc, nil
}

type GetVertexConfig struct {
	IfMatch string
}

func (c *GetVertexConfig) header() http.Header {
	if c == nil {
		return nil
	}
	var header http.Header
	if c.IfMatch != "" {
		header = make(http.Header)
		header.Set("if-match", c.IfMatch)
		return header
	}

	return nil
}

func (c *Connection) GetVertex(dbName, graphName, collName, vertexKey string, config *GetVertexConfig, vertexPtr interface{}) (rc int, err error) {
	path := buildPath(pathConfig{
		dbName:     dbName,
		pathFormat: "/_api/gharial/%s/vertex/%s/%s",
		pathParams: []interface{}{graphName, collName, vertexKey},
	})

	var body struct {
		Vertex interface{} `json:"vertex"`
	}
	if vertexPtr != nil {
		body.Vertex = vertexPtr
	}
	rc, _, err = c.send(http.MethodGet, path, config.header(), nil, &body)
	if err != nil {
		return rc, fmt.Errorf("failed to get vertex: %v", err)
	}
	return rc, nil
}

type ModifyVertexResult struct {
	ID     string `json:"_id"`
	Key    string `json:"_key"`
	Rev    string `json:"_rev"`
	OldRev string `json:"_oldRev"`
}

type ModifyVertexConfig struct {
	WaitForSync *bool
	KeepNull    *bool
	IfMatch     string
}

func (c *ModifyVertexConfig) header() http.Header {
	if c == nil {
		return nil
	}
	var header http.Header
	if c.IfMatch != "" {
		header = make(http.Header)
		header.Set("if-match", c.IfMatch)
		return header
	}

	return nil
}

func (c *ModifyVertexConfig) queryParams() url.Values {
	if c == nil {
		return nil
	}

	var params url.Values
	if c.WaitForSync != nil || c.KeepNull != nil {
		params = make(url.Values)
	}
	if c.WaitForSync != nil {
		params.Set("waitForSync", strconv.FormatBool(*c.WaitForSync))
	}
	if c.KeepNull != nil {
		params.Set("keepNull", strconv.FormatBool(*c.KeepNull))
	}
	return params
}

func (c *Connection) ModifyVertex(dbName, graphName, collName, vertexKey string, data interface{}, config *ModifyVertexConfig) (r ModifyVertexResult, rc int, err error) {
	path := buildPath(pathConfig{
		dbName:      dbName,
		pathFormat:  "/_api/gharial/%s/vertex/%s/%s",
		pathParams:  []interface{}{graphName, collName, vertexKey},
		queryParams: config.queryParams(),
	})

	var body struct {
		Vertex ModifyVertexResult `json:"vertex"`
	}
	rc, _, err = c.send(http.MethodPatch, path, config.header(), data, &body)
	if err != nil {
		return body.Vertex, rc, fmt.Errorf("failed to modify vertex: %v", err)
	}
	return body.Vertex, rc, nil
}

type ReplaceVertexResult struct {
	ID     string `json:"_id"`
	Key    string `json:"_key"`
	Rev    string `json:"_rev"`
	OldRev string `json:"_oldRev"`
}

type ReplaceVertexConfig struct {
	WaitForSync *bool
	IfMatch     string
}

func (c *ReplaceVertexConfig) header() http.Header {
	if c == nil {
		return nil
	}

	var header http.Header
	if c.IfMatch != "" {
		header = make(http.Header)
		header.Set("if-match", c.IfMatch)
		return header
	}
	return nil
}

func (c *ReplaceVertexConfig) queryParams() url.Values {
	if c == nil {
		return nil
	}

	var params url.Values
	if c.WaitForSync != nil {
		params = make(url.Values)
		params.Set("waitForSync", strconv.FormatBool(*c.WaitForSync))
	}
	return params
}

func (c *Connection) ReplaceVertex(dbName, graphName, collName, vertexKey string, data interface{}, config *ReplaceVertexConfig) (r ReplaceVertexResult, rc int, err error) {
	path := buildPath(pathConfig{
		dbName:      dbName,
		pathFormat:  "/_api/gharial/%s/vertex/%s/%s",
		pathParams:  []interface{}{graphName, collName, vertexKey},
		queryParams: config.queryParams(),
	})

	var body struct {
		Vertex ReplaceVertexResult `json:"vertex"`
	}
	rc, _, err = c.send(http.MethodPut, path, config.header(), data, &body)
	if err != nil {
		return body.Vertex, rc, fmt.Errorf("failed to replace vertex: %v", err)
	}
	return body.Vertex, rc, nil
}

type RemoveVertexConfig struct {
	WaitForSync *bool
	IfMatch     string
}

func (c *RemoveVertexConfig) header() http.Header {
	if c == nil {
		return nil
	}

	var header http.Header
	if c.IfMatch != "" {
		header = make(http.Header)
		header.Set("if-match", c.IfMatch)
		return header
	}
	return nil
}

func (c *RemoveVertexConfig) queryParams() url.Values {
	if c == nil {
		return nil
	}

	var params url.Values
	if c.WaitForSync != nil {
		params = make(url.Values)
		params.Set("waitForSync", strconv.FormatBool(*c.WaitForSync))
	}
	return params
}

func (c *Connection) RemoveVertex(dbName, graphName, collName, vertexKey string, config *RemoveVertexConfig) (removed bool, rc int, err error) {
	path := buildPath(pathConfig{
		dbName:      dbName,
		pathFormat:  "/_api/gharial/%s/vertex/%s/%s",
		pathParams:  []interface{}{graphName, collName, vertexKey},
		queryParams: config.queryParams(),
	})

	var body struct {
		Removed bool `json:"removed"`
	}
	rc, _, err = c.send(http.MethodDelete, path, config.header(), nil, &body)
	if err != nil {
		return body.Removed, rc, fmt.Errorf("failed to remove vertex: %v", err)
	}
	return body.Removed, rc, nil
}
