/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AlertObservation struct {
}

type AlertParameters struct {

	// Alert name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Alert value
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type DatabaseObservation struct {

	// Identifier of the database created
	DBID *float64 `json:"dbId,omitempty" tf:"db_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Private endpoint to access the database
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// Public endpoint to access the database
	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`
}

type DatabaseParameters struct {

	// Set of alerts to enable on the database
	// +kubebuilder:validation:Optional
	Alert []AlertParameters `json:"alert,omitempty" tf:"alert,omitempty"`

	// Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database
	// +kubebuilder:validation:Optional
	AverageItemSizeInBytes *float64 `json:"averageItemSizeInBytes,omitempty" tf:"average_item_size_in_bytes,omitempty"`

	// SSL certificate to authenticate user connections
	// +kubebuilder:validation:Optional
	ClientSSLCertificate *string `json:"clientSslCertificate,omitempty" tf:"client_ssl_certificate,omitempty"`

	// (Optional) The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')
	// +kubebuilder:validation:Optional
	DataEviction *string `json:"dataEviction,omitempty" tf:"data_eviction,omitempty"`

	// Rate of database data persistence (in persistent storage)
	// +kubebuilder:validation:Optional
	DataPersistence *string `json:"dataPersistence,omitempty" tf:"data_persistence,omitempty"`

	// Use TLS for authentication
	// +kubebuilder:validation:Optional
	EnableTLS *bool `json:"enableTls,omitempty" tf:"enable_tls,omitempty"`

	// Should use the external endpoint for open-source (OSS) Cluster API
	// +kubebuilder:validation:Optional
	ExternalEndpointForOssClusterAPI *bool `json:"externalEndpointForOssClusterApi,omitempty" tf:"external_endpoint_for_oss_cluster_api,omitempty"`

	// List of regular expression rules to shard the database by. See the documentation on clustering for more information on the hashing policy - https://docs.redislabs.com/latest/rc/concepts/clustering/
	// +kubebuilder:validation:Optional
	HashingPolicy []*string `json:"hashingPolicy,omitempty" tf:"hashing_policy,omitempty"`

	// Maximum memory usage for this specific database
	// +kubebuilder:validation:Required
	MemoryLimitInGb *float64 `json:"memoryLimitInGb" tf:"memory_limit_in_gb,omitempty"`

	// Modules to be provisioned in the database
	// +kubebuilder:validation:Optional
	Modules []ModulesParameters `json:"modules,omitempty" tf:"modules,omitempty"`

	// A meaningful name to identify the database
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Password used to access the database. If left empty, the password will be generated automatically
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Path that will be used to store database backup files
	// +kubebuilder:validation:Optional
	PeriodicBackupPath *string `json:"periodicBackupPath,omitempty" tf:"periodic_backup_path,omitempty"`

	// The protocol that will be used to access the database, (either ‘redis’ or 'memcached’)
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Set of Redis database URIs, in the format `redis://user:password@host:port`, that this database will be a replica of. If the URI provided is Redis Labs Cloud instance, only host and port should be provided
	// +kubebuilder:validation:Optional
	ReplicaOf []*string `json:"replicaOf,omitempty" tf:"replica_of,omitempty"`

	// Databases replication
	// +kubebuilder:validation:Optional
	Replication *bool `json:"replication,omitempty" tf:"replication,omitempty"`

	// Set of CIDR addresses to allow access to the database
	// +kubebuilder:validation:Optional
	SourceIps []*string `json:"sourceIps,omitempty" tf:"source_ips,omitempty"`

	// Identifier of the subscription
	// +kubebuilder:validation:Required
	SubscriptionID *float64 `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// Support Redis open-source (OSS) Cluster API
	// +kubebuilder:validation:Optional
	SupportOssClusterAPI *bool `json:"supportOssClusterApi,omitempty" tf:"support_oss_cluster_api,omitempty"`

	// Throughput measurement method, (either ‘number-of-shards’ or ‘operations-per-second’)
	// +kubebuilder:validation:Required
	ThroughputMeasurementBy *string `json:"throughputMeasurementBy" tf:"throughput_measurement_by,omitempty"`

	// Throughput value (as applies to selected measurement method)
	// +kubebuilder:validation:Required
	ThroughputMeasurementValue *float64 `json:"throughputMeasurementValue" tf:"throughput_measurement_value,omitempty"`
}

type ModulesObservation struct {
}

type ModulesParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rediscloud}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
