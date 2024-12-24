package apollo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) call(ctx context.Context, method string, params map[string]any) ([]byte, error) {
	body := map[string]any{
		"method":  method,
		"params":  params,
		"jsonrpc": "2.0",
		"id":      0,
	}
	reqBody, err := json.MarshalIndent(body, "", "	")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("token", c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadGateway {
		return nil, BadGateway
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func (c *Client) handleResp(r []byte) (*Resource, error) {
	var resp Resp[Resource]
	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}

	resource := resp.Result
	if resource.Rel == nil {
		resource.Rel = make(map[string][]Resource)
	}

	return &resource, nil
}

func (c *Client) handleResps(r []byte) ([]*Resource, error) {
	var (
		resp Resp[[]Resource]
		res  = make([]*Resource, 0)
	)

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}

	for _, re := range resp.Result {
		if re.Rel == nil {
			re.Rel = make(map[string][]Resource)
		}

		res = append(res, &re)
	}

	return res, nil
}

func (c *Client) handleListStr(r []byte) ([]string, error) {
	var resp Resp[[]string]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}
	return resp.Result, nil
}

func (c *Client) handleStr(r []byte) (string, error) {
	var resp Resp[string]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return "", JsonMarshalFailed
	}
	return resp.Result, nil
}

func (c *Client) handleAgg(r []byte) (*AggRes, error) {
	var resp Resp[AggRes]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}
	return &resp.Result, nil
}

func (c *Client) handleAggLeft(r []byte) (*AggResLeftJoin, error) {
	var resp Resp[AggResLeftJoin]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}
	return &resp.Result, nil
}

func (c *Client) handleBool(r []byte) (bool, error) {
	var resp Resp[bool]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return false, JsonMarshalFailed
	}
	return resp.Result, nil
}

func (c *Client) handleOpsGroup(r []byte) (*OpsGroup, error) {
	var resp Resp[OpsGroup]

	err := json.Unmarshal(r, &resp)
	if err != nil {
		c.log.Error(err, "apollo query resource failed")
		return nil, JsonMarshalFailed
	}
	return &resp.Result, nil
}
