package certgraphapi

// PKIRegistryInfo holds information about TLS artifacts stored in etcd. This includes object location and metadata based on object annotations
type PKIRegistryInfo struct {
	// +mapType:=atomic
	CertificateAuthorityBundles []PKIRegistryInClusterCABundle `json:"certificateAuthorityBundles"`
	// +mapType:=atomic
	CertKeyPairs []PKIRegistryInClusterCertKeyPair `json:"certKeyPairs"`
}

// PKIRegistryInClusterCertKeyPair identifies certificate key pair and stores its metadata
type PKIRegistryInClusterCertKeyPair struct {
	// SecretLocation points to the secret location
	SecretLocation InClusterSecretLocation `json:"secretLocation"`
	// CertKeyInfo stores metadata for certificate key pair
	CertKeyInfo PKIRegistryCertKeyPairInfo `json:"certKeyInfo"`
}

// PKIRegistryCertKeyPairInfo holds information about certificate key pair
type PKIRegistryCertKeyPairInfo struct {
	// OwningJiraComponent is a component name when a new OCP issue is filed in Jira
	OwningJiraComponent string `json:"owningJiraComponent"`
	// Description is a one sentence description of the certificate pair purpose
	Description string `json:"description"`

	//CertificateData PKIRegistryCertKeyMetadata
}

// PKIRegistryInClusterCABundle holds information about certificate authority bundle
type PKIRegistryInClusterCABundle struct {
	// ConfigMapLocation points to the configmap location
	ConfigMapLocation InClusterConfigMapLocation `json:"configMapLocation"`
	// CABundleInfo stores metadata for the certificate authority bundle
	CABundleInfo PKIRegistryCertificateAuthorityInfo `json:"certificateAuthorityBundleInfo"`
}

// PKIRegistryCertificateAuthorityInfo holds information about certificate authority bundle
type PKIRegistryCertificateAuthorityInfo struct {
	// OwningJiraComponent is a component name when a new OCP issue is filed in Jira
	OwningJiraComponent string `json:"owningJiraComponent"`
	// Description is a one sentence description of the certificate pair purpose
	Description string `json:"description"`

	//CertificateData []PKIRegistryCertKeyMetadata
}
