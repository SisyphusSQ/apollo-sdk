package apollo

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	url   string
	token string

	log     Logger
	timeout time.Duration
	client  *http.Client
}

func NewClient(c Config) (*Client, error) {
	if c.Url == "" || c.Token == "" {
		return nil, errors.New("url or token is empty")
	}

	cli := &Client{
		url:     c.Url,
		token:   c.Token,
		timeout: c.Timeout,
		log:     c.Logger,
		client:  &http.Client{Timeout: c.Timeout},
	}

	cli.log.Info("apollo client created", "url", cli.url)
	return cli, nil
}

// --------- QUERY ---------

func (c *Client) QueryResById(ctx context.Context, id int64) (*Resource, error) {
	var (
		params = map[string]any{
			"id": id,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by id", "id", id)
		return nil, err
	}
	return c.handleResp(r)
}

func (c *Client) QueryResByTypeAndName(ctx context.Context, rType, name string) (*Resource, error) {
	var (
		params = map[string]any{
			"type": rType,
			"name": name,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by type and name", "type", rType, "name", name)
		return nil, err
	}
	return c.handleResp(r)
}

func (c *Client) QueryResByType(ctx context.Context, rType string) ([]*Resource, error) {
	var (
		params = map[string]any{
			"type": rType,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by type", "type", rType)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByGraphAndTarget(ctx context.Context, graph, target string) ([]*Resource, error) {
	var (
		params = map[string]any{
			"graph":  graph,
			"target": target,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by graph and target", "graph", graph, "target", target)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByName(ctx context.Context, name string) ([]*Resource, error) {
	var (
		params = map[string]any{
			"name": name,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by name", "name", name)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByGroupAndType(ctx context.Context, rType, group string) ([]*Resource, error) {
	var (
		params = map[string]any{
			"type":       rType,
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by type and group", "type", rType, "group", group)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByTypeAndCondition(ctx context.Context, rType string, cond map[string]any) ([]*Resource, error) {
	var (
		params = map[string]any{
			"type":       rType,
			"conditions": cond,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by type and conditions", "type", rType, "conditions", cond)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByTypeAndRelationship(ctx context.Context, pType, relationship, sType string) ([]*Resource, error) {
	var (
		params = map[string]any{
			"primary_type":   pType,
			"relationship":   relationship,
			"secondary_type": sType,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by pType/relationship/sType",
			"type", pType, "relationship", relationship, "secondary_type", sType)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) QueryResByReferId(ctx context.Context, id int64) ([]*Resource, error) {
	var (
		params = map[string]any{
			"referenced_id": id,
		}
	)

	r, err := c.call(ctx, "query.resource", params)
	if err != nil {
		c.log.Error(err, "fail to query resource by referId", "ReferId", id)
		return nil, err
	}
	return c.handleResps(r)
}

func (c *Client) ListTypes(ctx context.Context) ([]string, error) {
	var (
		params = map[string]any{}
	)

	r, err := c.call(ctx, "query.ci.types", params)
	if err != nil {
		c.log.Error(err, "fail to list types")
		return nil, err
	}
	return c.handleListStr(r)
}

func (c *Client) ListOpsGroups(ctx context.Context) ([]string, error) {
	var (
		params = map[string]any{}
	)

	r, err := c.call(ctx, "query.ops.group", params)
	if err != nil {
		c.log.Error(err, "fail to list ops groups")
		return nil, err
	}
	return c.handleListStr(r)
}

func (c *Client) ListOpsGroupsWithUser(ctx context.Context, user string) ([]string, error) {
	var (
		params = map[string]any{
			"username": user,
		}
	)

	r, err := c.call(ctx, "query.ops.group", params)
	if err != nil {
		c.log.Error(err, "fail to list ops groups", "username", user)
		return nil, err
	}
	return c.handleListStr(r)
}

func (c *Client) ListUsers(ctx context.Context, group string) ([]string, error) {
	var (
		params = map[string]any{
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "query.ops.group.members", params)
	if err != nil {
		c.log.Error(err, "fail to list users", "group", group)
		return nil, err
	}
	return c.handleListStr(r)
}

func (c *Client) QueryOpsGroupOwner(ctx context.Context, group string) (string, error) {
	var (
		params = map[string]any{
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "query.ops.group.owner", params)
	if err != nil {
		c.log.Error(err, "fail to query ops group owner", "group", group)
		return "", err
	}
	return c.handleStr(r)
}

func (c *Client) QueryAggRes(ctx context.Context, graph string, fields [][]string) (*AggRes, error) {
	var (
		params = map[string]any{
			"graph":  graph,
			"fields": fields,
		}
	)

	r, err := c.call(ctx, "query.aggregate", params)
	if err != nil {
		c.log.Error(err, "fail to query aggregate resource", "graph", graph, "fields", fields)
		return nil, err
	}
	return c.handleAgg(r)
}

func (c *Client) QueryAggResWithGroup(ctx context.Context, graph, group string, fields [][]string) (*AggRes, error) {
	var (
		params = map[string]any{
			"graph":      graph,
			"fields":     fields,
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "query.aggregate", params)
	if err != nil {
		c.log.Error(err, "fail to query aggregate resource", "graph", graph, "fields", fields)
		return nil, err
	}
	return c.handleAgg(r)
}

func (c *Client) QueryAggResLeftJoin(ctx context.Context, graph, root string, fields []string) (*AggResLeftJoin, error) {
	var (
		params = map[string]any{
			"graph":  graph,
			"fields": fields,
			"root":   root,
		}
	)

	r, err := c.call(ctx, "query.aggregate", params)
	if err != nil {
		c.log.Error(err, "fail to query aggregate left join resource", "graph", graph, "fields", fields)
		return nil, err
	}
	return c.handleAggLeft(r)
}

func (c *Client) QueryResOpsGroupById(ctx context.Context, id int64) (*OpsGroup, error) {
	var (
		params = map[string]any{
			"id": id,
		}
	)

	r, err := c.call(ctx, "query.ci.ops.group", params)
	if err != nil {
		c.log.Error(err, "fail to query ops group", "id", id)
		return nil, err
	}
	return c.handleOpsGroup(r)
}

func (c *Client) QueryResOpsGroupByTypeAndName(ctx context.Context, rType, name string) (*OpsGroup, error) {
	var (
		params = map[string]any{
			"type": rType,
			"name": name,
		}
	)

	r, err := c.call(ctx, "query.ci.ops.group", params)
	if err != nil {
		c.log.Error(err, "fail to query resource ops group by type and name", "type", rType, "name", name)
		return nil, err
	}
	return c.handleOpsGroup(r)
}

// --------- CRUD ---------

func (c *Client) CreateRes(ctx context.Context, res Resource, group string) (*Resource, error) {
	var (
		params = map[string]any{
			"resource":   res,
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "create.resource", params)
	if err != nil {
		c.log.Error(err, "fail to create resource", "resource", res, "group", group)
		return nil, err
	}
	return c.handleResp(r)
}

func (c *Client) CreateResLst(ctx context.Context, resLst []Resource, group string) (bool, error) {
	var (
		params = map[string]any{
			"resources":  resLst,
			"group_name": group,
		}
	)

	r, err := c.call(ctx, "create.resource", params)
	if err != nil {
		c.log.Error(err, "fail to create resources", "resources", resLst, "group", group)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) UpdateRes(ctx context.Context, res Resource) (bool, error) {
	var (
		params = map[string]any{
			"resource": res,
		}
	)

	r, err := c.call(ctx, "update.resource", params)
	if err != nil {
		c.log.Error(err, "fail to update resource", "resource", res)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) UpdateResLst(ctx context.Context, resLst []Resource) (bool, error) {
	var (
		params = map[string]any{
			"resources": resLst,
		}
	)

	r, err := c.call(ctx, "update.resource", params)
	if err != nil {
		c.log.Error(err, "fail to update resources", "resources", resLst)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) UpdateResById(ctx context.Context, id int64, attr Attr) (bool, error) {
	var (
		params = map[string]any{
			"id":         id,
			"attributes": attr,
		}
	)

	r, err := c.call(ctx, "update.resource", params)
	if err != nil {
		c.log.Error(err, "fail to update resource", "id", id, "attributes", attr)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) UpdateResByTypeAndName(ctx context.Context, rtype, name string, attr Attr) (bool, error) {
	var (
		params = map[string]any{
			"type":       rtype,
			"name":       name,
			"attributes": attr,
		}
	)

	r, err := c.call(ctx, "update.resource", params)
	if err != nil {
		c.log.Error(err, "fail to update res by type and name", "type", rtype, "name", name)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) UpdateResRel(ctx context.Context, id int64, rels Rel, mode string) (bool, error) {
	var (
		params = map[string]any{
			"id":        id,
			"rels":      rels,
			"rels_mode": mode,
		}
	)

	r, err := c.call(ctx, "update.resource", params)
	if err != nil {
		c.log.Error(err, "fail to update resource relations", "id", id, "rels", rels, "mode", mode)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) DeliverRes(ctx context.Context, targetGroup string, id int64) (bool, error) {
	var (
		params = map[string]any{
			"id":                id,
			"target_group_name": targetGroup,
		}
	)

	r, err := c.call(ctx, "update.ci.ops.group", params)
	if err != nil {
		c.log.Error(err, "fail to deliver resource", "targetGroup", targetGroup, "id", id)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) DeleteById(ctx context.Context, id int64) (bool, error) {
	var (
		params = map[string]any{
			"id": id,
		}
	)

	r, err := c.call(ctx, "delete.resource", params)
	if err != nil {
		c.log.Error(err, "fail to delete resource", "id", id)
		return false, err
	}
	return c.handleBool(r)
}

func (c *Client) DeleteByTypeAndName(ctx context.Context, rtype, name string) (bool, error) {
	var (
		params = map[string]any{
			"type": rtype,
			"name": name,
		}
	)

	r, err := c.call(ctx, "delete.resource", params)
	if err != nil {
		c.log.Error(err, "fail to delete resource", "type", rtype, "name", name)
		return false, err
	}
	return c.handleBool(r)
}
