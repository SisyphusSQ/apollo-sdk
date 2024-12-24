package apollo

import "errors"

var (
	JsonMarshalFailed = errors.New("json marshal is failed")
	BadGateway        = errors.New("bad gateway")
)
