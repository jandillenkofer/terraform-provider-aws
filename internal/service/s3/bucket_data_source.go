// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_s3_bucket", name="Bucket")
func dataSourceBucket() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceBucketRead,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bucket_regional_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hosted_zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"website_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"website_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	awsClient := meta.(*conns.AWSClient)
	conn := awsClient.S3Client(ctx)
	var optFns []func(*s3.Options)

	bucket := d.Get("bucket").(string)
	// Via S3 access point: "Invalid configuration: region from ARN `us-east-1` does not match client region `aws-global` and UseArnRegion is `false`".
	if arn.IsARN(bucket) && conn.Options().Region == names.GlobalRegionID {
		optFns = append(optFns, func(o *s3.Options) { o.UseARNRegion = true })
	}
	err := findBucket(ctx, conn, bucket, optFns...)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading S3 Bucket (%s): %s", bucket, err)
	}

	region, err := findBucketRegion(ctx, awsClient, bucket, optFns...)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading S3 Bucket (%s) Region: %s", bucket, err)
	}

	d.SetId(bucket)
	if arn.IsARN(bucket) {
		d.Set("arn", bucket)
	} else {
		arn := arn.ARN{
			Partition: awsClient.Partition,
			Service:   "s3",
			Resource:  bucket,
		}.String()
		d.Set("arn", arn)
	}
	d.Set("bucket_domain_name", awsClient.PartitionHostname(fmt.Sprintf("%s.s3", bucket)))
	d.Set("bucket_regional_domain_name", bucketRegionalDomainName(bucket, region))
	if hostedZoneID, err := hostedZoneIDForRegion(region); err == nil {
		d.Set("hosted_zone_id", hostedZoneID)
	} else {
		log.Printf("[WARN] HostedZoneIDForRegion: %s", err)
	}
	d.Set("region", region)
	if _, err := findBucketWebsite(ctx, conn, bucket, ""); err == nil {
		endpoint, domain := bucketWebsiteEndpointAndDomain(bucket, region)
		d.Set("website_domain", domain)
		d.Set("website_endpoint", endpoint)
	} else if !tfresource.NotFound(err) {
		log.Printf("[WARN] Reading S3 Bucket (%s) Website: %s", bucket, err)
	}

	return diags
}
