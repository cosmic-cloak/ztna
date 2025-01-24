package env

import (
	"bytes"
	"github.com/cosmic-cloak/ztna/common/eid"
	"github.com/cosmic-cloak/ztna/controller/response"
	"io"
	"net/http"
)

func NewRequestContext(rw http.ResponseWriter, r *http.Request) *response.RequestContext {
	rid := eid.New()

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(body))

	requestContext := &response.RequestContext{
		Id:                rid,
		ResponseWriter:    rw,
		Request:           r,
		Body:              body,
		Identity:          nil,
		ApiSession:        nil,
		ActivePermissions: []string{},
	}

	requestContext.Responder = response.NewResponder(requestContext)

	return requestContext
}
