// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package iam

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	iam_sdkv1 "github.com/aws/aws-sdk-go/service/iam"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAccessKeys,
			TypeName: "aws_iam_access_keys",
		},
		{
			Factory:  DataSourceAccountAlias,
			TypeName: "aws_iam_account_alias",
		},
		{
			Factory:  DataSourceGroup,
			TypeName: "aws_iam_group",
		},
		{
			Factory:  DataSourceInstanceProfile,
			TypeName: "aws_iam_instance_profile",
		},
		{
			Factory:  DataSourceInstanceProfiles,
			TypeName: "aws_iam_instance_profiles",
		},
		{
			Factory:  DataSourceOpenIDConnectProvider,
			TypeName: "aws_iam_openid_connect_provider",
		},
		{
			Factory:  DataSourcePolicy,
			TypeName: "aws_iam_policy",
		},
		{
			Factory:  DataSourcePolicyDocument,
			TypeName: "aws_iam_policy_document",
		},
		{
			Factory:  DataSourcePrincipalPolicySimulation,
			TypeName: "aws_iam_principal_policy_simulation",
		},
		{
			Factory:  DataSourceRole,
			TypeName: "aws_iam_role",
		},
		{
			Factory:  DataSourceRoles,
			TypeName: "aws_iam_roles",
		},
		{
			Factory:  DataSourceSAMLProvider,
			TypeName: "aws_iam_saml_provider",
		},
		{
			Factory:  DataSourceServerCertificate,
			TypeName: "aws_iam_server_certificate",
		},
		{
			Factory:  DataSourceSessionContext,
			TypeName: "aws_iam_session_context",
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_iam_user",
		},
		{
			Factory:  DataSourceUserSSHKey,
			TypeName: "aws_iam_user_ssh_key",
		},
		{
			Factory:  DataSourceUsers,
			TypeName: "aws_iam_users",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccessKey,
			TypeName: "aws_iam_access_key",
		},
		{
			Factory:  ResourceAccountAlias,
			TypeName: "aws_iam_account_alias",
		},
		{
			Factory:  resourceAccountPasswordPolicy,
			TypeName: "aws_iam_account_password_policy",
			Name:     "Account Password Policy",
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_iam_group",
			Name:     "Group",
		},
		{
			Factory:  ResourceGroupMembership,
			TypeName: "aws_iam_group_membership",
		},
		{
			Factory:  ResourceGroupPolicy,
			TypeName: "aws_iam_group_policy",
		},
		{
			Factory:  resourceGroupPolicyAttachment,
			TypeName: "aws_iam_group_policy_attachment",
			Name:     "Group Policy Attachment",
		},
		{
			Factory:  resourceInstanceProfile,
			TypeName: "aws_iam_instance_profile",
			Name:     "Instance Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "InstanceProfile",
			},
		},
		{
			Factory:  resourceOpenIDConnectProvider,
			TypeName: "aws_iam_openid_connect_provider",
			Name:     "OIDC Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "OIDCProvider",
			},
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_iam_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "Policy",
			},
		},
		{
			Factory:  resourcePolicyAttachment,
			TypeName: "aws_iam_policy_attachment",
			Name:     "Policy Attachment",
		},
		{
			Factory:  resourceRole,
			TypeName: "aws_iam_role",
			Name:     "Role",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "Role",
			},
		},
		{
			Factory:  ResourceRolePolicy,
			TypeName: "aws_iam_role_policy",
		},
		{
			Factory:  resourceRolePolicyAttachment,
			TypeName: "aws_iam_role_policy_attachment",
			Name:     "Role Policy Attachment",
		},
		{
			Factory:  resourceSAMLProvider,
			TypeName: "aws_iam_saml_provider",
			Name:     "SAML Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "SAMLProvider",
			},
		},
		{
			Factory:  ResourceSecurityTokenServicePreferences,
			TypeName: "aws_iam_security_token_service_preferences",
			Name:     "Security Token Service Preferences",
		},
		{
			Factory:  resourceServerCertificate,
			TypeName: "aws_iam_server_certificate",
			Name:     "Server Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "name",
				ResourceType:        "ServerCertificate",
			},
		},
		{
			Factory:  resourceServiceLinkedRole,
			TypeName: "aws_iam_service_linked_role",
			Name:     "Service Linked Role",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "ServiceLinkedRole",
			},
		},
		{
			Factory:  ResourceServiceSpecificCredential,
			TypeName: "aws_iam_service_specific_credential",
		},
		{
			Factory:  ResourceSigningCertificate,
			TypeName: "aws_iam_signing_certificate",
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_iam_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "User",
			},
		},
		{
			Factory:  ResourceUserGroupMembership,
			TypeName: "aws_iam_user_group_membership",
		},
		{
			Factory:  resourceUserLoginProfile,
			TypeName: "aws_iam_user_login_profile",
			Name:     "User Login Profile",
		},
		{
			Factory:  ResourceUserPolicy,
			TypeName: "aws_iam_user_policy",
		},
		{
			Factory:  resourceUserPolicyAttachment,
			TypeName: "aws_iam_user_policy_attachment",
			Name:     "User Policy Attachment",
		},
		{
			Factory:  resourceUserSSHKey,
			TypeName: "aws_iam_user_ssh_key",
			Name:     "User SSH Key",
		},
		{
			Factory:  resourceVirtualMFADevice,
			TypeName: "aws_iam_virtual_mfa_device",
			Name:     "Virtual MFA Device",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ResourceType:        "VirtualMFADevice",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IAM
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*iam_sdkv1.IAM, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return iam_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
