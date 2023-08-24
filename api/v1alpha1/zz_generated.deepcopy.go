//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1alpha1

import (
	apiv1alpha1 "github.com/spectrocloud-labs/valid8or/api/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsAuth) DeepCopyInto(out *AwsAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsAuth.
func (in *AwsAuth) DeepCopy() *AwsAuth {
	if in == nil {
		return nil
	}
	out := new(AwsAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsValidator) DeepCopyInto(out *AwsValidator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsValidator.
func (in *AwsValidator) DeepCopy() *AwsValidator {
	if in == nil {
		return nil
	}
	out := new(AwsValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsValidator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsValidatorList) DeepCopyInto(out *AwsValidatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AwsValidator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsValidatorList.
func (in *AwsValidatorList) DeepCopy() *AwsValidatorList {
	if in == nil {
		return nil
	}
	out := new(AwsValidatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AwsValidatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsValidatorSpec) DeepCopyInto(out *AwsValidatorSpec) {
	*out = *in
	out.Auth = in.Auth
	if in.IamRules != nil {
		in, out := &in.IamRules, &out.IamRules
		*out = make([]IamRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceQuotaRules != nil {
		in, out := &in.ServiceQuotaRules, &out.ServiceQuotaRules
		*out = make([]ServiceQuotaRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TagRules != nil {
		in, out := &in.TagRules, &out.TagRules
		*out = make([]TagRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsValidatorSpec.
func (in *AwsValidatorSpec) DeepCopy() *AwsValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(AwsValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsValidatorStatus) DeepCopyInto(out *AwsValidatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1alpha1.ValidationCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsValidatorStatus.
func (in *AwsValidatorStatus) DeepCopy() *AwsValidatorStatus {
	if in == nil {
		return nil
	}
	out := new(AwsValidatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamRule) DeepCopyInto(out *IamRule) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]PolicyDocument, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamRule.
func (in *IamRule) DeepCopy() *IamRule {
	if in == nil {
		return nil
	}
	out := new(IamRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDocument) DeepCopyInto(out *PolicyDocument) {
	*out = *in
	if in.Statements != nil {
		in, out := &in.Statements, &out.Statements
		*out = make([]StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDocument.
func (in *PolicyDocument) DeepCopy() *PolicyDocument {
	if in == nil {
		return nil
	}
	out := new(PolicyDocument)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuota) DeepCopyInto(out *ServiceQuota) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuota.
func (in *ServiceQuota) DeepCopy() *ServiceQuota {
	if in == nil {
		return nil
	}
	out := new(ServiceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaRule) DeepCopyInto(out *ServiceQuotaRule) {
	*out = *in
	if in.ServiceQuotas != nil {
		in, out := &in.ServiceQuotas, &out.ServiceQuotas
		*out = make([]ServiceQuota, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaRule.
func (in *ServiceQuotaRule) DeepCopy() *ServiceQuotaRule {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatementEntry) DeepCopyInto(out *StatementEntry) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(Condition)
		(*in).DeepCopyInto(*out)
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatementEntry.
func (in *StatementEntry) DeepCopy() *StatementEntry {
	if in == nil {
		return nil
	}
	out := new(StatementEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagRule) DeepCopyInto(out *TagRule) {
	*out = *in
	if in.ARNs != nil {
		in, out := &in.ARNs, &out.ARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagRule.
func (in *TagRule) DeepCopy() *TagRule {
	if in == nil {
		return nil
	}
	out := new(TagRule)
	in.DeepCopyInto(out)
	return out
}
