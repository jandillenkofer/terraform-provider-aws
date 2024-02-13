// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kafka

import (
	"context"

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
			Factory:  dataSourceBootstrapBrokers,
			TypeName: "aws_msk_bootstrap_brokers",
			Name:     "Bootstrap Brokers",
		},
		{
			Factory:  dataSourceBrokerNodes,
			TypeName: "aws_msk_broker_nodes",
			Name:     "Broker Nodes",
		},
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_msk_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  dataSourceConfiguration,
			TypeName: "aws_msk_configuration",
			Name:     "Configuration",
		},
		{
			Factory:  dataSourceKafkaVersion,
			TypeName: "aws_msk_kafka_version",
			Name:     "Kafka Version",
		},
		{
			Factory:  dataSourceVPCConnection,
			TypeName: "aws_msk_vpc_connection",
			Name:     "VPC Connection",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCluster,
			TypeName: "aws_msk_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  resourceClusterPolicy,
			TypeName: "aws_msk_cluster_policy",
			Name:     "Cluster Policy",
		},
		{
			Factory:  resourceConfiguration,
			TypeName: "aws_msk_configuration",
			Name:     "Configuration",
		},
		{
			Factory:  resourceReplicator,
			TypeName: "aws_msk_replicator",
			Name:     "Replicator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  resourceSCRAMSecretAssociation,
			TypeName: "aws_msk_scram_secret_association",
			Name:     "SCRAM Secret Association",
		},
		{
			Factory:  resourceSingleSCRAMSecretAssociation,
			TypeName: "aws_msk_single_scram_secret_association",
			Name:     "Single SCRAM Secret Association",
		},
		{
			Factory:  resourceServerlessCluster,
			TypeName: "aws_msk_serverless_cluster",
			Name:     "Serverless Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  resourceVPCConnection,
			TypeName: "aws_msk_vpc_connection",
			Name:     "VPC Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kafka
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
