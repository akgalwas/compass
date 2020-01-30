package api

import (
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/api/middlewares"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/director"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/graphql"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/internal/connector/model"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/apperrors"
	"github.com/kyma-incubator/compass/components/connectivity-adapter/pkg/reqerror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type managementInfoHandler struct {
	gqlClient                      graphql.Client
	logger                         *log.Logger
	connectivityAdapterMTLSBaseURL string
	directorClientProvider         director.DirectorClientProvider
}

func NewManagementInfoHandler(client graphql.Client, logger *log.Logger, connectivityAdapterMTLSBaseURL string, directorClientProvider director.DirectorClientProvider) managementInfoHandler {
	return managementInfoHandler{
		gqlClient:                      client,
		logger:                         logger,
		connectivityAdapterMTLSBaseURL: connectivityAdapterMTLSBaseURL,
		directorClientProvider:         directorClientProvider,
	}
}

func (mh *managementInfoHandler) GetManagementInfo(w http.ResponseWriter, r *http.Request) {

	authorizationHeaders, err := middlewares.GetAuthHeadersFromContext(r.Context(), middlewares.AuthorizationHeadersKey)
	if err != nil {
		mh.logger.Errorf("Failed to read authorization context: %s.", err)
		reqerror.WriteErrorMessage(w, "Failed to read authorization context.", apperrors.CodeForbidden)

		return
	}

	systemAuthID := authorizationHeaders.GetClientID()
	contextLogger := contextLogger(mh.logger, systemAuthID)

	// <AG>
	directorClient := mh.directorClientProvider.Client(r)

	application, err := directorClient.GetApplication(systemAuthID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get application")
		contextLogger.Error(err.Error())
		reqerror.WriteError(w, err, apperrors.CodeInternal)

		return
	}

	log.Info("Events base: " + application.EventingConfiguration.DefaultURL)

	// <AG>

	contextLogger.Info("Getting Management Info")

	configuration, err := mh.gqlClient.Configuration(authorizationHeaders)
	if err != nil {
		err = errors.Wrap(err, "Failed to get configuration")
		contextLogger.Error(err.Error())
		reqerror.WriteError(w, err, apperrors.CodeInternal)

		return
	}

	certInfo := graphql.ToCertInfo(configuration)

	//TODO: handle case when configuration.Token is nil
	managementInfoResponse := mh.makeManagementInfoResponse(
		application.Name,
		configuration.Token.Token,
		mh.connectivityAdapterMTLSBaseURL,
		application.EventingConfiguration.DefaultURL,
		certInfo)

	respondWithBody(w, http.StatusOK, managementInfoResponse, contextLogger)
}

func (m *managementInfoHandler) makeManagementInfoResponse(
	application,
	newToken,
	connectivityAdapterMTLSBaseURL,
	eventServiceBaseURL string,
	certInfo model.CertInfo) model.MgmtInfoReponse {

	return model.MgmtInfoReponse{
		ClientIdentity:  model.MakeClientIdentity(application, "", ""),
		URLs:            model.MakeManagementURLs(application, connectivityAdapterMTLSBaseURL, eventServiceBaseURL),
		CertificateInfo: certInfo,
	}
}
