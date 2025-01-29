package api_impl

import (
	"github.com/openziti/foundation/v2/errorz"
	"ztna-core/ztna/controller/api"
	"ztna-core/ztna/controller/rest_model"
	"net/http"
)

func RespondWithCreatedId(responder api.Responder, id string, link rest_model.Link) {
	createEnvelope := &rest_model.CreateEnvelope{
		Data: &rest_model.CreateLocation{
			Links: rest_model.Links{
				"self": link,
			},
			ID: id,
		},
		Meta: &rest_model.Meta{},
	}

	responder.Respond(createEnvelope, http.StatusCreated)
}

func RespondWithOk(responder api.Responder, data interface{}, meta *rest_model.Meta) {
	responder.Respond(&rest_model.Empty{
		Data: data,
		Meta: meta,
	}, http.StatusOK)
}

type FabricResponseMapper struct{}

func (self FabricResponseMapper) EmptyOkData() interface{} {
	return &rest_model.Empty{
		Data: map[string]interface{}{},
		Meta: &rest_model.Meta{},
	}
}

func (self FabricResponseMapper) MapApiError(requestId string, apiError *errorz.ApiError) interface{} {
	return &rest_model.APIErrorEnvelope{
		Error: ToRestModel(apiError, requestId),
		Meta: &rest_model.Meta{
			APIEnrollmentVersion: ApiVersion,
			APIVersion:           ApiVersion,
		},
	}
}
