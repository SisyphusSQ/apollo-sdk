package apollo

// ------- Resource -------

type (
	Attr map[string]any
	Rel  map[string][]Resource
)

type RType struct {
	Name string `json:"name"`
}

type ResBase struct {
	ID   int64 `json:"id,omitempty"`
	Type RType `json:"type"`
}

type Resource struct {
	ResBase

	Attrs Attr `json:"attributes"`
	Rel   Rel  `json:"relations"`
}

// ------- ResourceVO -------

type AggRes struct {
	Data     map[string]any `json:"data,omitempty"`
	Children []AggRes       `json:"children,omitempty"`
}

type AggResLeftJoin struct {
	Data []AggResLeftJoinItem `json:"data,omitempty"`
}

type AggResLeftJoinItem struct {
	Data     map[string]any                  `json:"data,omitempty"`
	Children map[string][]AggResLeftJoinItem `json:"children,omitempty"`
}

type OpsGroup struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	ProxyId     int64     `json:"proxyId"`
	Description string    `json:"description"`
	Template    string    `json:"template"`
	DutyId      int       `json:"dutyId"`
	Views       []View    `json:"views"`
	Users       []OpsUser `json:"users"`
	Owner       OpsUser   `json:"owner"`
}

type View struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}

type OpsUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

// ------- Response -------

type R interface {
	Resource | AggRes | AggResLeftJoin | OpsGroup | string | bool | []Resource | []string
}

type RespBase struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
}

type Resp[T R] struct {
	RespBase
	Result T `json:"result"`
}

// ------- used for user -------

type AttrBase struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

type EmptyRel struct{}
