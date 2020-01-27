package api

import (
	"net/http"

	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/api/middlewares"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/graphql"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/apperrors"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/reqerror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type revocationsHandler struct {
	gqlClient graphql.Client
	logger    *log.Logger
}

func NewRevocationsHandler(client graphql.Client, logger *log.Logger) revocationsHandler {
	return revocationsHandler{
		gqlClient: client,
		logger:    logger,
	}
}

func (rh *revocationsHandler) RevokeCertificate(w http.ResponseWriter, r *http.Request) {
	authorizationHeaders, err := middlewares.GetAuthHeadersFromContext(r.Context(), middlewares.AuthorizationHeadersKey)
	if err != nil {
		rh.logger.Errorf("Failed to read authorization context: %s.", err)
		reqerror.WriteErrorMessage(w, "Failed to read authorization context.", apperrors.CodeForbidden)

		return
	}

	application := authorizationHeaders.GetClientID()
	contextLogger := contextLogger(rh.logger, application)

	contextLogger.Info("Revoke certificate")

	err = rh.gqlClient.Revoke(authorizationHeaders)
	if err != nil {
		err = errors.Wrap(err, "Failed to revoke certificate")
		contextLogger.Error(err.Error())
		reqerror.WriteError(w, err, apperrors.CodeInternal)
	}

	respond(w, http.StatusCreated)
}
