//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/mongodb/mongodb-kubernetes-operator/pkg/automationconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	if in.Modes != nil {
		in, out := &in.Modes, &out.Modes
		*out = make([]AuthMode, len(*in))
		copy(*out, *in)
	}
	if in.IgnoreUnknownUsers != nil {
		in, out := &in.IgnoreUnknownUsers, &out.IgnoreUnknownUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationRestriction) DeepCopyInto(out *AuthenticationRestriction) {
	*out = *in
	if in.ClientSource != nil {
		in, out := &in.ClientSource, &out.ClientSource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServerAddress != nil {
		in, out := &in.ServerAddress, &out.ServerAddress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationRestriction.
func (in *AuthenticationRestriction) DeepCopy() *AuthenticationRestriction {
	if in == nil {
		return nil
	}
	out := new(AuthenticationRestriction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutomationConfigOverride) DeepCopyInto(out *AutomationConfigOverride) {
	*out = *in
	if in.Processes != nil {
		in, out := &in.Processes, &out.Processes
		*out = make([]OverrideProcess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutomationConfigOverride.
func (in *AutomationConfigOverride) DeepCopy() *AutomationConfigOverride {
	if in == nil {
		return nil
	}
	out := new(AutomationConfigOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRole) DeepCopyInto(out *CustomRole) {
	*out = *in
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]Privilege, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]Role, len(*in))
		copy(*out, *in)
	}
	if in.AuthenticationRestrictions != nil {
		in, out := &in.AuthenticationRestrictions, &out.AuthenticationRestrictions
		*out = make([]AuthenticationRestriction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRole.
func (in *CustomRole) DeepCopy() *CustomRole {
	if in == nil {
		return nil
	}
	out := new(CustomRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCommunity) DeepCopyInto(out *MongoDBCommunity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCommunity.
func (in *MongoDBCommunity) DeepCopy() *MongoDBCommunity {
	if in == nil {
		return nil
	}
	out := new(MongoDBCommunity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBCommunity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCommunityList) DeepCopyInto(out *MongoDBCommunityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBCommunity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCommunityList.
func (in *MongoDBCommunityList) DeepCopy() *MongoDBCommunityList {
	if in == nil {
		return nil
	}
	out := new(MongoDBCommunityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBCommunityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCommunitySpec) DeepCopyInto(out *MongoDBCommunitySpec) {
	*out = *in
	if in.ReplicaSetHorizons != nil {
		in, out := &in.ReplicaSetHorizons, &out.ReplicaSetHorizons
		*out = make(ReplicaSetHorizonConfiguration, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(automationconfig.ReplicaSetHorizons, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	in.Security.DeepCopyInto(&out.Security)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]MongoDBUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StatefulSetConfiguration.DeepCopyInto(&out.StatefulSetConfiguration)
	in.AdditionalMongodConfig.DeepCopyInto(&out.AdditionalMongodConfig)
	if in.AutomationConfigOverride != nil {
		in, out := &in.AutomationConfigOverride, &out.AutomationConfigOverride
		*out = new(AutomationConfigOverride)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(Prometheus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCommunitySpec.
func (in *MongoDBCommunitySpec) DeepCopy() *MongoDBCommunitySpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBCommunitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCommunityStatus) DeepCopyInto(out *MongoDBCommunityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCommunityStatus.
func (in *MongoDBCommunityStatus) DeepCopy() *MongoDBCommunityStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBCommunityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBUser) DeepCopyInto(out *MongoDBUser) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]Role, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBUser.
func (in *MongoDBUser) DeepCopy() *MongoDBUser {
	if in == nil {
		return nil
	}
	out := new(MongoDBUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodConfiguration) DeepCopyInto(out *MongodConfiguration) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverrideProcess) DeepCopyInto(out *OverrideProcess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverrideProcess.
func (in *OverrideProcess) DeepCopy() *OverrideProcess {
	if in == nil {
		return nil
	}
	out := new(OverrideProcess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Privilege) DeepCopyInto(out *Privilege) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Privilege.
func (in *Privilege) DeepCopy() *Privilege {
	if in == nil {
		return nil
	}
	out := new(Privilege)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	out.TLSSecretRef = in.TLSSecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ReplicaSetHorizonConfiguration) DeepCopyInto(out *ReplicaSetHorizonConfiguration) {
	{
		in := &in
		*out = make(ReplicaSetHorizonConfiguration, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(automationconfig.ReplicaSetHorizons, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetHorizonConfiguration.
func (in ReplicaSetHorizonConfiguration) DeepCopy() ReplicaSetHorizonConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetHorizonConfiguration)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.DB != nil {
		in, out := &in.DB, &out.DB
		*out = new(string)
		**out = **in
	}
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Role) DeepCopyInto(out *Role) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Role.
func (in *Role) DeepCopy() *Role {
	if in == nil {
		return nil
	}
	out := new(Role)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Security) DeepCopyInto(out *Security) {
	*out = *in
	in.Authentication.DeepCopyInto(&out.Authentication)
	in.TLS.DeepCopyInto(&out.TLS)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]CustomRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Security.
func (in *Security) DeepCopy() *Security {
	if in == nil {
		return nil
	}
	out := new(Security)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetConfiguration) DeepCopyInto(out *StatefulSetConfiguration) {
	*out = *in
	in.SpecWrapper.DeepCopyInto(&out.SpecWrapper)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetConfiguration.
func (in *StatefulSetConfiguration) DeepCopy() *StatefulSetConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSpecWrapper) DeepCopyInto(out *StatefulSetSpecWrapper) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	out.CertificateKeySecret = in.CertificateKeySecret
	if in.CaCertificateSecret != nil {
		in, out := &in.CaCertificateSecret, &out.CaCertificateSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.CaConfigMap != nil {
		in, out := &in.CaConfigMap, &out.CaConfigMap
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}
