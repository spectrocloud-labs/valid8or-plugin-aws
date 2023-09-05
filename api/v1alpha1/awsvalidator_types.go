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

package v1alpha1

import (
	"fmt"

	valid8orv1alpha1 "github.com/spectrocloud-labs/valid8or/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AwsValidatorSpec defines the desired state of AwsValidator
type AwsValidatorSpec struct {
	Auth              AwsAuth            `json:"auth"`
	DefaultRegion     string             `json:"defaultRegion"`
	IamRoleRules      []IamRoleRule      `json:"iamRoleRules,omitempty"`
	IamUserRules      []IamUserRule      `json:"iamUserRules,omitempty"`
	IamGroupRules     []IamGroupRule     `json:"iamGroupRules,omitempty"`
	IamPolicyRules    []IamPolicyRule    `json:"iamPolicyRules,omitempty"`
	ServiceQuotaRules []ServiceQuotaRule `json:"serviceQuotaRules,omitempty"`
	TagRules          []TagRule          `json:"tagRules,omitempty"`
}

type AwsAuth struct {
	// Option 1: lookup AWS creds from a secret
	SecretName string `json:"secretName,omitempty"`
	// Option 2: specify a service account (EKS)
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}

type IamRoleRule struct {
	IamRoleName string           `json:"iamRoleName"`
	Policies    []PolicyDocument `json:"iamPolicies"`
}

func (r IamRoleRule) Name() string {
	return r.IamRoleName
}

func (r IamRoleRule) IAMPolicies() []PolicyDocument {
	return r.Policies
}

type IamUserRule struct {
	IamUserName string           `json:"iamUserName"`
	Policies    []PolicyDocument `json:"iamPolicies"`
}

func (r IamUserRule) Name() string {
	return r.IamUserName
}

func (r IamUserRule) IAMPolicies() []PolicyDocument {
	return r.Policies
}

type IamGroupRule struct {
	IamGroupName string           `json:"iamGroupName"`
	Policies     []PolicyDocument `json:"iamPolicies"`
}

func (r IamGroupRule) Name() string {
	return r.IamGroupName
}

func (r IamGroupRule) IAMPolicies() []PolicyDocument {
	return r.Policies
}

type IamPolicyRule struct {
	IamPolicyARN string           `json:"iamPolicyArn"`
	Policies     []PolicyDocument `json:"iamPolicies"`
}

func (r IamPolicyRule) Name() string {
	return r.IamPolicyARN
}

func (r IamPolicyRule) IAMPolicies() []PolicyDocument {
	return r.Policies
}

type PolicyDocument struct {
	Name       string           `json:"name"`
	Version    string           `json:"version"`
	Statements []StatementEntry `json:"statements"`
}

type StatementEntry struct {
	Condition *Condition `json:"condition,omitempty"`
	Effect    string     `json:"effect"`
	Actions   []string   `json:"actions"`
	Resources []string   `json:"resources"`
}

type Condition struct {
	Type   string   `json:"type"`
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

func (c *Condition) String() string {
	return fmt.Sprintf("%s: %s=%s", c.Type, c.Key, c.Values)
}

type ServiceQuotaRule struct {
	Region        string         `json:"region"`
	ServiceCode   string         `json:"serviceCode"`
	ServiceQuotas []ServiceQuota `json:"serviceQuotas"`
}

type ServiceQuota struct {
	Name   string `json:"name"`
	Buffer int    `json:"buffer"`
}

type TagRule struct {
	Key           string   `json:"key"`
	ExpectedValue string   `json:"expectedValue"`
	Region        string   `json:"region"`
	ResourceType  string   `json:"resourceType"`
	ARNs          []string `json:"arns"`
}

// AwsValidatorStatus defines the observed state of AwsValidator
type AwsValidatorStatus struct {
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []valid8orv1alpha1.ValidationCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AwsValidator is the Schema for the awsvalidators API
type AwsValidator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AwsValidatorSpec   `json:"spec,omitempty"`
	Status AwsValidatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AwsValidatorList contains a list of AwsValidator
type AwsValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AwsValidator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AwsValidator{}, &AwsValidatorList{})
}
