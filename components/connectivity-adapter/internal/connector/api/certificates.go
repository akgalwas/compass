package api

import (
	"encoding/json"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/api/middlewares"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/graphql"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/apperrors"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/reqerror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type certRequest struct {
	CSR string `json:"csr"`
}

type certificatesHandler struct {
	client graphql.Client
	logger *log.Logger
}

func NewCertificatesHandler(client graphql.Client, logger *log.Logger) certificatesHandler {
	return certificatesHandler{
		client: client,
		logger: logger,
	}
}

func (ch *certificatesHandler) SignCSR(w http.ResponseWriter, r *http.Request) {
	authorizationHeaders, err := middlewares.GetAuthHeadersFromContext(r.Context(), middlewares.AuthorizationHeadersKey)
	if err != nil {
		reqerror.WriteErrorMessage(w, "Failed to read authorization context.", apperrors.CodeForbidden)

		return
	}

	contextLogger := contextLogger(ch.logger, authorizationHeaders.GetClientID())
	certRequest, err := readCertRequest(r)
	if err != nil {
		err = errors.Wrap(err, "Failed to read certificate request")
		contextLogger.Error(err.Error())
		reqerror.WriteError(w, err, apperrors.CodeInternal)

		return
	}

	contextLogger.Info("Generating certificate")

	certificationResult, err := ch.client.SignCSR(certRequest.CSR, authorizationHeaders)
	if err != nil {
		err = errors.Wrap(err, "Failed to generate certificate")
		contextLogger.Error(err.Error())
		reqerror.WriteError(w, err, apperrors.CodeInternal)

		return
	}

	certResponse := graphql.ToCertResponse(certificationResult)
	respondWithBody(w, http.StatusOK, certResponse)
}

func readCertRequest(r *http.Request) (*certRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error while reading request body: %s")
	}
	defer r.Body.Close()

	var tokenRequest certRequest
	err = json.Unmarshal(b, &tokenRequest)
	if err != nil {
		return nil, errors.Wrap(err, "error while unmarshalling request body: %s")
	}

	return &tokenRequest, nil
}

func respondWithBody(w http.ResponseWriter, statusCode int, responseBody interface{}) {
	respond(w, statusCode)
	json.NewEncoder(w).Encode(responseBody)
}

func respond(w http.ResponseWriter, statusCode int) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJson)
	w.WriteHeader(statusCode)
}

func contextLogger(logger *log.Logger, application string) *log.Entry {
	return logger.WithField("application", application)
}
