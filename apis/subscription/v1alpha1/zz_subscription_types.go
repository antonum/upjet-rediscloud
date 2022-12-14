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

type AllowlistObservation struct {
}

type AllowlistParameters struct {

	// Set of CIDR ranges that are allowed to access the databases associated with this subscription
	// +kubebuilder:validation:Optional
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// Set of security groups that are allowed to access the databases associated with this subscription
	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`
}

type CloudProviderObservation struct {

	// Cloud networking details, per region (single region or multiple regions for Active-Active cluster only)
	// +kubebuilder:validation:Required
	Region []RegionObservation `json:"region,omitempty" tf:"region,omitempty"`
}

type CloudProviderParameters struct {

	// Cloud account identifier. Default: Redis Labs internal cloud account (using Cloud Account Id = 1 implies using Redis Labs internal cloud account). Note that a GCP subscription can be created only with Redis Labs internal cloud account
	// +kubebuilder:validation:Optional
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// The cloud provider to use with the subscription, (either `AWS` or `GCP`)
	// +kubebuilder:validation:Optional
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// Cloud networking details, per region (single region or multiple regions for Active-Active cluster only)
	// +kubebuilder:validation:Required
	Region []RegionParameters `json:"region" tf:"region,omitempty"`
}

type CreationPlanObservation struct {
}

type CreationPlanParameters struct {

	// Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database
	// +kubebuilder:validation:Optional
	AverageItemSizeInBytes *float64 `json:"averageItemSizeInBytes,omitempty" tf:"average_item_size_in_bytes,omitempty"`

	// Maximum memory usage for each database
	// +kubebuilder:validation:Required
	MemoryLimitInGb *float64 `json:"memoryLimitInGb" tf:"memory_limit_in_gb,omitempty"`

	// Modules that will be used by the databases in this subscription.
	// +kubebuilder:validation:Required
	Modules []*string `json:"modules" tf:"modules,omitempty"`

	// The planned number of databases
	// +kubebuilder:validation:Required
	Quantity *float64 `json:"quantity" tf:"quantity,omitempty"`

	// Databases replication
	// +kubebuilder:validation:Required
	Replication *bool `json:"replication" tf:"replication,omitempty"`

	// Support Redis open-source (OSS) Cluster API
	// +kubebuilder:validation:Required
	SupportOssClusterAPI *bool `json:"supportOssClusterApi" tf:"support_oss_cluster_api,omitempty"`

	// Throughput measurement method, (either ???number-of-shards??? or ???operations-per-second???)
	// +kubebuilder:validation:Required
	ThroughputMeasurementBy *string `json:"throughputMeasurementBy" tf:"throughput_measurement_by,omitempty"`

	// Throughput value (as applies to selected measurement method)
	// +kubebuilder:validation:Required
	ThroughputMeasurementValue *float64 `json:"throughputMeasurementValue" tf:"throughput_measurement_value,omitempty"`
}

type NetworksObservation struct {
	NetworkingDeploymentCidr *string `json:"networkingDeploymentCidr,omitempty" tf:"networking_deployment_cidr,omitempty"`

	NetworkingSubnetID *string `json:"networkingSubnetId,omitempty" tf:"networking_subnet_id,omitempty"`

	NetworkingVPCID *string `json:"networkingVpcId,omitempty" tf:"networking_vpc_id,omitempty"`
}

type NetworksParameters struct {
}

type RegionObservation struct {

	// List of networks used
	Networks []NetworksObservation `json:"networks,omitempty" tf:"networks,omitempty"`
}

type RegionParameters struct {

	// Support deployment on multiple availability zones within the selected region
	// +kubebuilder:validation:Optional
	MultipleAvailabilityZones *bool `json:"multipleAvailabilityZones,omitempty" tf:"multiple_availability_zones,omitempty"`

	// Deployment CIDR mask
	// +kubebuilder:validation:Required
	NetworkingDeploymentCidr *string `json:"networkingDeploymentCidr" tf:"networking_deployment_cidr,omitempty"`

	// Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified)
	// +kubebuilder:validation:Optional
	NetworkingVPCID *string `json:"networkingVpcId,omitempty" tf:"networking_vpc_id,omitempty"`

	// List of availability zones used
	// +kubebuilder:validation:Required
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones" tf:"preferred_availability_zones,omitempty"`

	// Deployment region as defined by cloud provider
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type SubscriptionObservation struct {

	// A cloud provider object
	// +kubebuilder:validation:Required
	CloudProvider []CloudProviderObservation `json:"cloudProvider,omitempty" tf:"cloud_provider,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubscriptionParameters struct {

	// An allowlist object
	// +kubebuilder:validation:Optional
	Allowlist []AllowlistParameters `json:"allowlist,omitempty" tf:"allowlist,omitempty"`

	// A cloud provider object
	// +kubebuilder:validation:Required
	CloudProvider []CloudProviderParameters `json:"cloudProvider" tf:"cloud_provider,omitempty"`

	// Information about the planned databases used to optimise the database infrastructure. This information is only used when creating a new subscription and any changes will be ignored after this.
	// +kubebuilder:validation:Optional
	CreationPlan []CreationPlanParameters `json:"creationPlan,omitempty" tf:"creation_plan,omitempty"`

	// Memory storage preference: either ???ram??? or a combination of 'ram-and-flash???
	// +kubebuilder:validation:Optional
	MemoryStorage *string `json:"memoryStorage,omitempty" tf:"memory_storage,omitempty"`

	// A meaningful name to identify the subscription
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Payment method for the requested subscription. If credit card is specified, the payment method Id must be defined.
	// +kubebuilder:validation:Optional
	PaymentMethod *string `json:"paymentMethod,omitempty" tf:"payment_method,omitempty"`

	// A valid payment method pre-defined in the current account
	// +kubebuilder:validation:Optional
	PaymentMethodID *string `json:"paymentMethodId,omitempty" tf:"payment_method_id,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rediscloud}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
