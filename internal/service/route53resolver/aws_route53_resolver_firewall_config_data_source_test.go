package route53resolver_test

import (
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccRoute53ResolverFirewallDataSource_basic(t *testing.T) {
	dataSourceName := "data.aws_route53_resolver_firewall_config.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheck(t) },
		ErrorCheck:        acctest.ErrorCheck(t, route53resolver.EndpointsID),
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFirewallDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceName, "firewall_fail_open", regexp.MustCompile(`ENABLED|DISABLED`)),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "owner_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "resource_id"),
				),
			},
		},
	})
}

func testAccFirewallDataSourceConfig_basic() string {
	return `
resource "aws_vpc" "test" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
}

resource "aws_route53_resolver_firewall_config" "test" {
  resource_id        = aws_vpc.test.id
  firewall_fail_open = "ENABLED"
}

data "aws_route53_resolver_firewall_config" "test" {
  resource_id = aws_vpc.test.id
}

`
}
