// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_acmpca_certificate":           DataSourceCertificate,
		"aws_acmpca_certificate_authority": DataSourceCertificateAuthority,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_acmpca_certificate":                       ResourceCertificate,
		"aws_acmpca_certificate_authority":             ResourceCertificateAuthority,
		"aws_acmpca_certificate_authority_certificate": ResourceCertificateAuthorityCertificate,
		"aws_acmpca_permission":                        ResourcePermission,
		"aws_acmpca_policy":                            ResourcePolicy,
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ACMPCA
}

var ServicePackage = &servicePackage{}
