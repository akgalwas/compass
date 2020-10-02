package certificates

import (
	"crypto/x509"
	"encoding/base64"
	"github.com/sirupsen/logrus"

	"github.com/kyma-incubator/compass/components/connector/internal/apperrors"
	"github.com/kyma-incubator/compass/components/connector/internal/secrets"

	"k8s.io/apimachinery/pkg/types"
)

const (
	caCertificateSecretKey     = "ca.crt"
	caKeySecretKey             = "ca.key"
	rootCACertificateSecretKey = "cacert"
)

//go:generate mockery -name=Service
type Service interface {
	// SignCSR takes encoded CSR, validates subject and generates Certificate based on CA stored in secret
	// returns base64 encoded certificate chain
	SignCSR(encodedCSR []byte, subject CSRSubject, clientId string) (EncodedCertificateChain, apperrors.AppError)
}

type certificateService struct {
	secretsRepository           secrets.Repository
	certUtil                    CertificateUtility
	caSecretName                types.NamespacedName
	rootCACertificateSecretName types.NamespacedName
	logger                           *logrus.Entry
}

func NewCertificateService(secretRepository secrets.Repository, certUtil CertificateUtility, caSecretName, rootCACertificateSecretName types.NamespacedName) Service {
	return &certificateService{
		secretsRepository:           secretRepository,
		certUtil:                    certUtil,
		caSecretName:                caSecretName,
		rootCACertificateSecretName: rootCACertificateSecretName,
		logger : logrus.WithField("CertificateService", "Certificate"),
	}
}

func (svc *certificateService) SignCSR(encodedCSR []byte, subject CSRSubject, clientId string) (EncodedCertificateChain, apperrors.AppError) {
	csr, err := svc.certUtil.LoadCSR(encodedCSR)
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	err = svc.checkCSR(csr, subject)
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	return svc.signCSR(csr, clientId)
}

func (svc *certificateService) signCSR(csr *x509.CertificateRequest, clientId string) (EncodedCertificateChain, apperrors.AppError) {
	svc.logger.Infof("signCSR for %s client started.", clientId)
	secretData, err := svc.secretsRepository.Get(svc.caSecretName)
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	svc.logger.Infof("secret read for %s client.", clientId)

	caCrt, err := svc.certUtil.LoadCert(secretData[caCertificateSecretKey])
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	svc.logger.Infof("cert loaded for %s client.", clientId)

	caKey, err := svc.certUtil.LoadKey(secretData[caKeySecretKey])
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	svc.logger.Infof("key loaded for %s client.", clientId)

	signedCrt, err := svc.certUtil.SignCSR(caCrt, csr, caKey)
	if err != nil {
		return EncodedCertificateChain{}, err
	}

	return svc.encodeCertificates(caCrt.Raw, signedCrt, clientId)
}

func (svc *certificateService) encodeCertificates(rawCaCertificate, rawClientCertificate []byte, clientId string) (EncodedCertificateChain, apperrors.AppError) {
	svc.logger.Infof("encodeCertificates for %s client started.", clientId)

	caCrtBytes := svc.certUtil.AddCertificateHeaderAndFooter(rawCaCertificate)

	svc.logger.Infof("certificate header and footer added for client %s. (ca certificate)", clientId)

	signedCrtBytes := svc.certUtil.AddCertificateHeaderAndFooter(rawClientCertificate)

	svc.logger.Infof("certificate header and footer added for client %s. (client certificate)", clientId)

	if svc.rootCACertificateSecretName.Name != "" {
		rootCABytes, err := svc.loadRootCACert(clientId)
		if err != nil {
			svc.logger.Errorf("failed to load root cert for client %s: %s", clientId, err.Error())
			return EncodedCertificateChain{}, err
		}
		svc.logger.Infof("root CA cert loaded for client %s.", clientId)

		caCrtBytes = svc.createCertChain(rootCABytes, caCrtBytes)
		svc.logger.Infof("cert chain created for client %s. (root chain)", clientId)
	}

	certChain := svc.createCertChain(signedCrtBytes, caCrtBytes)
	svc.logger.Infof("cert chain created for client %s. (client chain)", clientId)

	return encodeCertificateBase64(certChain, signedCrtBytes, caCrtBytes), nil
}

func (svc *certificateService) loadRootCACert(clientId string) ([]byte, apperrors.AppError) {
	svc.logger.Infof("loadRootCACert for %s client started.", clientId)
	secretData, err := svc.secretsRepository.Get(svc.rootCACertificateSecretName)
	if err != nil {
		return nil, err
	}

	svc.logger.Infof("secret for Root CA read for %s client.", clientId)

	rootCACrt, err := svc.certUtil.LoadCert(secretData[rootCACertificateSecretKey])
	if err != nil {
		return nil, err
	}

	svc.logger.Infof("cert loaded for Root CA for %s client.", clientId)

	return svc.certUtil.AddCertificateHeaderAndFooter(rootCACrt.Raw), nil
}

func (svc *certificateService) checkCSR(csr *x509.CertificateRequest, expectedSubject CSRSubject) apperrors.AppError {
	return svc.certUtil.CheckCSRValues(csr, expectedSubject)
}

func (svc *certificateService) createCertChain(clientCrt, caCrt []byte) []byte {
	return append(clientCrt, caCrt...)
}

func encodeCertificateBase64(certChain, clientCRT, caCRT []byte) EncodedCertificateChain {
	return EncodedCertificateChain{
		CertificateChain:  encodeStringBase64(certChain),
		ClientCertificate: encodeStringBase64(clientCRT),
		CaCertificate:     encodeStringBase64(caCRT),
	}
}

func encodeStringBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
