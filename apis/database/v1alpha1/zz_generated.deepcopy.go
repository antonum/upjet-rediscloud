//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertObservation) DeepCopyInto(out *AlertObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertObservation.
func (in *AlertObservation) DeepCopy() *AlertObservation {
	if in == nil {
		return nil
	}
	out := new(AlertObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertParameters) DeepCopyInto(out *AlertParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertParameters.
func (in *AlertParameters) DeepCopy() *AlertParameters {
	if in == nil {
		return nil
	}
	out := new(AlertParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseObservation) DeepCopyInto(out *DatabaseObservation) {
	*out = *in
	if in.DBID != nil {
		in, out := &in.DBID, &out.DBID
		*out = new(float64)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(string)
		**out = **in
	}
	if in.PublicEndpoint != nil {
		in, out := &in.PublicEndpoint, &out.PublicEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseObservation.
func (in *DatabaseObservation) DeepCopy() *DatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(DatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseParameters) DeepCopyInto(out *DatabaseParameters) {
	*out = *in
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = make([]AlertParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AverageItemSizeInBytes != nil {
		in, out := &in.AverageItemSizeInBytes, &out.AverageItemSizeInBytes
		*out = new(float64)
		**out = **in
	}
	if in.ClientSSLCertificate != nil {
		in, out := &in.ClientSSLCertificate, &out.ClientSSLCertificate
		*out = new(string)
		**out = **in
	}
	if in.DataEviction != nil {
		in, out := &in.DataEviction, &out.DataEviction
		*out = new(string)
		**out = **in
	}
	if in.DataPersistence != nil {
		in, out := &in.DataPersistence, &out.DataPersistence
		*out = new(string)
		**out = **in
	}
	if in.EnableTLS != nil {
		in, out := &in.EnableTLS, &out.EnableTLS
		*out = new(bool)
		**out = **in
	}
	if in.ExternalEndpointForOssClusterAPI != nil {
		in, out := &in.ExternalEndpointForOssClusterAPI, &out.ExternalEndpointForOssClusterAPI
		*out = new(bool)
		**out = **in
	}
	if in.HashingPolicy != nil {
		in, out := &in.HashingPolicy, &out.HashingPolicy
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MemoryLimitInGb != nil {
		in, out := &in.MemoryLimitInGb, &out.MemoryLimitInGb
		*out = new(float64)
		**out = **in
	}
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]ModulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PeriodicBackupPath != nil {
		in, out := &in.PeriodicBackupPath, &out.PeriodicBackupPath
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ReplicaOf != nil {
		in, out := &in.ReplicaOf, &out.ReplicaOf
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Replication != nil {
		in, out := &in.Replication, &out.Replication
		*out = new(bool)
		**out = **in
	}
	if in.SourceIps != nil {
		in, out := &in.SourceIps, &out.SourceIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(float64)
		**out = **in
	}
	if in.SupportOssClusterAPI != nil {
		in, out := &in.SupportOssClusterAPI, &out.SupportOssClusterAPI
		*out = new(bool)
		**out = **in
	}
	if in.ThroughputMeasurementBy != nil {
		in, out := &in.ThroughputMeasurementBy, &out.ThroughputMeasurementBy
		*out = new(string)
		**out = **in
	}
	if in.ThroughputMeasurementValue != nil {
		in, out := &in.ThroughputMeasurementValue, &out.ThroughputMeasurementValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseParameters.
func (in *DatabaseParameters) DeepCopy() *DatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(DatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulesObservation) DeepCopyInto(out *ModulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulesObservation.
func (in *ModulesObservation) DeepCopy() *ModulesObservation {
	if in == nil {
		return nil
	}
	out := new(ModulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModulesParameters) DeepCopyInto(out *ModulesParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModulesParameters.
func (in *ModulesParameters) DeepCopy() *ModulesParameters {
	if in == nil {
		return nil
	}
	out := new(ModulesParameters)
	in.DeepCopyInto(out)
	return out
}
