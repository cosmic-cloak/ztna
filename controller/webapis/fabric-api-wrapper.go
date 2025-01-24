package webapis

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/edge-api/rest_management_api_server"
	"github.com/openziti/foundation/v2/errorz"
	"github.com/cosmic-cloak/ztna/controller/api"
	"github.com/cosmic-cloak/ztna/controller/api_impl"
	"github.com/cosmic-cloak/ztna/controller/apierror"
	"github.com/cosmic-cloak/ztna/controller/env"
	"github.com/cosmic-cloak/ztna/controller/internal/permissions"
	"github.com/cosmic-cloak/ztna/controller/response"
	"net/http"
	"time"
)

func NewFabricApiWrapper(ae *env.AppEnv) api_impl.RequestWrapper {
	return &fabricWrapper{ae: ae}
}

type fabricWrapper struct {
	ae *env.AppEnv
}

func (self *fabricWrapper) WrapRequest(handler api_impl.RequestHandler, request *http.Request, entityId, entitySubId string) middleware.Responder {
	return middleware.ResponderFunc(func(writer http.ResponseWriter, producer runtime.Producer) {
		rc, err := env.GetRequestContextFromHttpContext(request)

		if rc == nil {
			rc = self.ae.CreateRequestContext(writer, request)
		}

		rc.SetProducer(producer)
		rc.SetEntityId(entityId)
		rc.SetEntitySubId(entitySubId)

		if err != nil {
			pfxlog.Logger().WithError(err).Error("could not retrieve request context")
			rc.RespondWithError(err)
			return
		}

		if !permissions.IsAdmin().IsAllowed(rc.ActivePermissions...) {
			rc.RespondWithApiError(errorz.NewUnauthorized())
			return
		}

		handler(self.ae.GetHostController().GetNetwork(), rc)
	})
}

func (self *fabricWrapper) WrapHttpHandler(handler http.Handler) http.Handler {
	wrapped := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set(ZitiInstanceId, self.ae.InstanceId)

		if r.URL.Path == api_impl.FabricRestApiRootPath {
			rw.Header().Set("content-type", "application/json")
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write(rest_management_api_server.SwaggerJSON)
			return
		}

		rc := self.ae.CreateRequestContext(rw, r)

		api.AddRequestContextToHttpContext(r, rc)

		err := self.ae.FillRequestContext(rc)
		if err != nil {
			rc.RespondWithError(err)
			return
		}

		//after request context is filled so that api session is present for session expiration headers
		response.AddHeaders(rc)

		handler.ServeHTTP(rw, r)
	})

	return api.TimeoutHandler(api.WrapCorsHandler(wrapped), 10*time.Second, apierror.NewTimeoutError(), response.EdgeResponseMapper{})
}

func (self *fabricWrapper) WrapWsHandler(handler http.Handler) http.Handler {
	wrapped := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rc := self.ae.CreateRequestContext(rw, r)

		err := self.ae.FillRequestContext(rc)
		if err != nil {
			rc.RespondWithError(err)
			return
		}

		if !permissions.IsAdmin().IsAllowed(rc.ActivePermissions...) {
			rc.RespondWithApiError(errorz.NewUnauthorized())
			return
		}

		handler.ServeHTTP(rw, r)
	})

	return wrapped
}
