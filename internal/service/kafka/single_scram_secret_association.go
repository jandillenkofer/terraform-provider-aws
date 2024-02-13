// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"golang.org/x/exp/slices"
)

// @SDKResource("aws_msk_single_scram_secret_association", name="Single SCRAM Secret Association)
func resourceSingleSCRAMSecretAssociation() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceSingleSCRAMSecretAssociationCreate,
		ReadWithoutTimeout:   resourceSingleSCRAMSecretAssociationRead,
		DeleteWithoutTimeout: resourceSingleSCRAMSecretAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"cluster_arn": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidARN,
			},
			"secret_arn": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidARN,
			},
		},
	}
}

func resourceSingleSCRAMSecretAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).KafkaClient(ctx)

	clusterARN := d.Get("cluster_arn").(string)
	secretARN := d.Get("secret_arn").(string)

	if err := associateSRAMSecret(ctx, conn, clusterARN, secretARN); err != nil {
		return sdkdiag.AppendErrorf(diags, "creating Single MSK SCRAM Secret Association (%s): %s", clusterARN, err)
	}

	d.SetId(fmt.Sprintf("%s:%s", clusterARN, secretARN))

	return append(diags, resourceSingleSCRAMSecretAssociationRead(ctx, d, meta)...)
}

func resourceSingleSCRAMSecretAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).KafkaClient(ctx)

	clusterARN := d.Get("cluster_arn").(string)
	secretARN := d.Get("secret_arn").(string)

	id, err := findSCRAMSecretsByClusterARNAndSecretARN(ctx, conn, clusterARN, secretARN)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] MSK Single SCRAM Secret Association (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading MSK SCRAM Single Secret Association (%s): %s", d.Id(), err)
	}

	if id != nil {
		d.SetId(*id)
		d.Set("cluster_arn", clusterARN)
		d.Set("secret_arn", secretARN)
	} else {
		d.SetId("")
	}

	return diags
}

func resourceSingleSCRAMSecretAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).KafkaClient(ctx)

	clusterARN := d.Get("cluster_arn").(string)
	secretARN := d.Get("secret_arn").(string)

	err := disassociateSRAMSecret(ctx, conn, clusterARN, secretARN)

	if errs.IsA[*types.NotFoundException](err) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting MSK Single SCRAM Secret Association (%s): %s", d.Id(), err)
	}

	return diags
}

func findSCRAMSecretsByClusterARNAndSecretARN(ctx context.Context, conn *kafka.Client, clusterARN string, secretARN string) (*string, error) {
	input := &kafka.ListScramSecretsInput{
		ClusterArn: aws.String(clusterARN),
	}

	pages := kafka.NewListScramSecretsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if errs.IsA[*types.NotFoundException](err) {
			return nil, &retry.NotFoundError{
				LastError:   err,
				LastRequest: input,
			}
		}

		if err != nil {
			return nil, err
		}

		if slices.Contains(page.SecretArnList, secretARN) {
			id := fmt.Sprintf("%s:%s", clusterARN, secretARN)
			return &id, nil
		}
	}

	return nil, nil
}

func associateSRAMSecret(ctx context.Context, conn *kafka.Client, clusterARN string, secretARN string) error {
	input := &kafka.BatchAssociateScramSecretInput{
		ClusterArn:    aws.String(clusterARN),
		SecretArnList: []string{secretARN},
	}

	output, err := conn.BatchAssociateScramSecret(ctx, input)

	if err == nil {
		err = unprocessedScramSecretsError(output.UnprocessedScramSecrets, false)
	}

	if err != nil {
		return err
	}

	return nil
}

func disassociateSRAMSecret(ctx context.Context, conn *kafka.Client, clusterARN string, secretARN string) error {
	input := &kafka.BatchDisassociateScramSecretInput{
		ClusterArn:    aws.String(clusterARN),
		SecretArnList: []string{secretARN},
	}

	output, err := conn.BatchDisassociateScramSecret(ctx, input)

	if err == nil {
		err = unprocessedScramSecretsError(output.UnprocessedScramSecrets, true)
	}

	if err != nil {
		return err
	}

	return nil
}
