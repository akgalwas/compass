package main

import (
	"github.com/kyma-incubator/compass/components/provisioner/internal/api"
	"github.com/kyma-incubator/compass/components/provisioner/internal/hydroform"
	"github.com/kyma-incubator/compass/components/provisioner/internal/persistence"
	"github.com/kyma-incubator/compass/components/provisioner/internal/persistence/database"
	"github.com/kyma-incubator/compass/components/provisioner/internal/persistence/dbsession"
	"github.com/kyma-incubator/compass/components/provisioner/internal/provisioning"
	"github.com/pkg/errors"

	"path/filepath"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newPersistenceService(connectionString, schemaPath string) (persistence.RuntimeService, persistence.OperationService, error) {
	connection, err := database.InitializeDatabase(connectionString, schemaPath)
	if err != nil {
		return nil, nil, err
	}

	dbSessionFactory := dbsession.NewFactory(connection)
	uuidGenerator := persistence.NewUUIDGenerator()
	runtimeService := persistence.NewRuntimeService(dbSessionFactory, uuidGenerator)
	operationService := persistence.NewOperationService(dbSessionFactory)

	return runtimeService, operationService, nil
}

func newProvisioningService(runtimeService persistence.RuntimeService, operationService persistence.OperationService, secrets v1.SecretInterface) provisioning.ProvisioningService {
	hydroformClient := hydroform.NewHydroformClient(secrets)

	return provisioning.NewProvisioningService(operationService, runtimeService, hydroformClient)
}

func newSecretsInterface(namespace string) (v1.SecretInterface, error) {
	k8sConfig, err := restclient.InClusterConfig()
	if err != nil {
		logrus.Warnf("Failed to read in cluster config: %s", err.Error())
		logrus.Info("Trying to initialize with local config")
		home := homedir.HomeDir()
		k8sConfPath := filepath.Join(home, ".kube", "config")
		k8sConfig, err = clientcmd.BuildConfigFromFlags("", k8sConfPath)
		if err != nil {
			return nil, errors.Errorf("failed to read k8s in-cluster configuration, %s", err.Error())
		}
	}

	coreClientset, err := kubernetes.NewForConfig(k8sConfig)
	if err != nil {
		return nil, errors.Errorf("failed to create k8s core client, %s", err.Error())
	}

	return coreClientset.CoreV1().Secrets(namespace), nil
}

func newResolver(connectionString string, schemaFilePath string, namespace string) *api.Resolver {
	runtimeService, operationService, err := newPersistenceService(connectionString, schemaFilePath)

	exitOnError(err, "Error while initializing persistence")

	secretInterface, err := newSecretsInterface(namespace)

	exitOnError(err, "Cannot create secrets interface")

	service := newProvisioningService(runtimeService, operationService, secretInterface)

	return api.NewResolver(service)
}
