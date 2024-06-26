// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfrds "github.com/hashicorp/terraform-provider-aws/internal/service/rds"
)

func TestAccRDSOrderableInstanceDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"
	engine := "mysql"
	licenseModel := "general-public-license"
	storageType := "standard"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_basic(engine, licenseModel, storageType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "engine", engine),
					resource.TestCheckResourceAttr(dataSourceName, "license_model", licenseModel),
					resource.TestCheckResourceAttr(dataSourceName, "storage_type", storageType),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_version", "data.aws_rds_engine_version.default", "version"),
					resource.TestCheckResourceAttrPair(dataSourceName, "instance_class", "data.aws_rds_orderable_db_instance.dynamic", "instance_class"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_preferredClass(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_preferredClass(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "instance_class"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_preferredVersion(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_preferredVersion(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_version", "data.aws_rds_engine_version.default", "version"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_preferredClassAndVersion(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_preferredClassAndVersion(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "instance_class", "data.aws_rds_orderable_db_instance.dynamic", "instance_class"),
					resource.TestCheckResourceAttrPair(dataSourceName, "engine_version", "data.aws_rds_engine_version.default", "version"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsEnhancedMonitoring(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsEnhancedMonitoring(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_enhanced_monitoring", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_latestVersion(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_latestVersion(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineAuroraMySQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "false"),
					resource.TestMatchResourceAttr(dataSourceName, "engine_version", regexache.MustCompile(`^5\.7\.mysql_aurora\..*`)),
				),
			},
			{
				Config: testAccOrderableInstanceDataSourceConfig_latestVersion(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineAuroraMySQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestMatchResourceAttr(dataSourceName, "engine_version", regexache.MustCompile(`^5\.7\.mysql_aurora\..*`)),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsGlobalDatabases(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsGlobalDatabases(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_global_databases", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineAuroraMySQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestMatchResourceAttr(dataSourceName, "engine_version", regexache.MustCompile(`^8\.0\.mysql_aurora\..*`)),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsClusters(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsClusters(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_clusters", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineMySQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestMatchResourceAttr(dataSourceName, "instance_class", regexache.MustCompile(`^db\..*large$`)),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_readReplicaCapable(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_readReplicaCapable(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "read_replica_capable", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.InstanceEngineOracleEnterprise),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestCheckResourceAttrSet(dataSourceName, "instance_class"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsMultiAZ(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsMultiAZ(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_multi_az", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineMySQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportedEngineModes(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportedEngineModes(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineAuroraPostgreSQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "supported_engine_modes.0", "provisioned"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportedNetworkTypes(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportedNetworkTypes(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "engine", tfrds.ClusterEngineAuroraPostgreSQL),
					resource.TestCheckResourceAttr(dataSourceName, "engine_latest_version", "true"),
					resource.TestCheckTypeSetElemAttr(dataSourceName, "supported_network_types.*", "DUAL"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsIAMDatabaseAuthentication(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsIAMDatabaseAuthentication(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_iam_database_authentication", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsIops(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsIops(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_iops", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsKerberosAuthentication(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsKerberosAuthentication(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_kerberos_authentication", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsPerformanceInsights(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccOrderableInstancePreCheck(ctx, t)
			testAccPerformanceInsightsDefaultVersionPreCheck(ctx, t, "mysql")
		},
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsPerformanceInsights(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_performance_insights", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsStorageAutoScaling(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsStorageAutoScaling(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_storage_autoscaling", "true"),
				),
			},
		},
	})
}

func TestAccRDSOrderableInstanceDataSource_supportsStorageEncryption(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	dataSourceName := "data.aws_rds_orderable_db_instance.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccOrderableInstancePreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccOrderableInstanceDataSourceConfig_supportsStorageEncryption(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "supports_storage_encryption", "true"),
				),
			},
		},
	})
}

func testAccOrderableInstancePreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).RDSConn(ctx)

	input := &rds.DescribeOrderableDBInstanceOptionsInput{
		Engine:          aws.String("mysql"),
		DBInstanceClass: aws.String("db.m5.xlarge"),
	}

	_, err := conn.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)

	if acctest.PreCheckSkipError(err) {
		t.Skipf("skipping acceptance testing: %s", err)
	}

	if err != nil {
		t.Fatalf("unexpected PreCheck error: %s", err)
	}
}

func testAccOrderableInstanceDataSourceConfig_basic(engine, license, storage string) string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = %[1]q
}

data "aws_rds_orderable_db_instance" "dynamic" {
  engine                     = data.aws_rds_engine_version.default.engine
  engine_version             = data.aws_rds_engine_version.default.version
  preferred_instance_classes = [%[2]s]
}

data "aws_rds_orderable_db_instance" "test" {
  instance_class = data.aws_rds_orderable_db_instance.dynamic.instance_class
  engine         = data.aws_rds_engine_version.default.engine
  engine_version = data.aws_rds_engine_version.default.version
  license_model  = %[3]q
  storage_type   = %[4]q
}
`, engine, mainInstanceClasses, license, storage)
}

func testAccOrderableInstanceDataSourceConfig_preferredClass() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine         = data.aws_rds_engine_version.default.engine
  engine_version = data.aws_rds_engine_version.default.version
  license_model  = "general-public-license"

  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_preferredVersion() string {
	return `
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine        = data.aws_rds_engine_version.default.engine
  license_model = "general-public-license"
  storage_type  = "standard"

  preferred_engine_versions = [
    "18.42.32",
    data.aws_rds_engine_version.default.version,
    "not.a.version",
  ]
}
`
}

func testAccOrderableInstanceDataSourceConfig_preferredClassAndVersion() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "dynamic" {
  engine                     = data.aws_rds_engine_version.default.engine
  engine_version             = data.aws_rds_engine_version.default.version
  preferred_instance_classes = [%[1]s]
}

data "aws_rds_orderable_db_instance" "test" {
  engine        = data.aws_rds_engine_version.default.engine
  license_model = "general-public-license"

  preferred_instance_classes = [
    "db.xyz.xlarge",
    data.aws_rds_orderable_db_instance.dynamic.instance_class,
    "db.t3.small",
  ]
  preferred_engine_versions = [
    "18.42.32",
    data.aws_rds_engine_version.default.version,
    "not.a.version",
  ]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsEnhancedMonitoring() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                       = data.aws_rds_engine_version.default.engine
  license_model                = "general-public-license"
  storage_type                 = "standard"
  supports_enhanced_monitoring = true

  preferred_engine_versions  = ["8.0.25", "8.0.26", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsGlobalDatabases() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  supports_global_databases  = true
}
`, tfrds.ClusterEngineAuroraMySQL, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsClusters() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  storage_type               = "io1"
  supports_iops              = true
  supports_clusters          = true
}
`, tfrds.ClusterEngineMySQL, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_readReplicaCapable() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  read_replica_capable       = true
  storage_type               = "gp3"
}
`, tfrds.InstanceEngineOracleEnterprise, strings.Replace(mainInstanceClasses, "db.t3.small", "frodo", 1))
}

func testAccOrderableInstanceDataSourceConfig_supportsMultiAZ() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  supports_multi_az          = true
}
`, tfrds.ClusterEngineMySQL, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportedEngineModes() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  supports_clusters          = true
  supported_engine_modes     = ["provisioned", "serverless"]
}
`, tfrds.ClusterEngineAuroraPostgreSQL, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportedNetworkTypes() string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = true
  preferred_instance_classes = [%[2]s]
  supports_clusters          = true
  supported_network_types    = ["DUAL"]
}
`, tfrds.ClusterEngineAuroraPostgreSQL, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_latestVersion(latestVersion bool) string {
	return fmt.Sprintf(`
data "aws_rds_orderable_db_instance" "test" {
  engine                     = %[1]q
  engine_latest_version      = %[2]t
  preferred_instance_classes = [%[3]s]
}
`, tfrds.ClusterEngineAuroraMySQL, latestVersion, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsIAMDatabaseAuthentication() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                               = data.aws_rds_engine_version.default.engine
  license_model                        = "general-public-license"
  storage_type                         = "standard"
  supports_iam_database_authentication = true

  preferred_engine_versions  = ["8.0.25", "8.0.26", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsIops() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine        = data.aws_rds_engine_version.default.engine
  license_model = "general-public-license"
  supports_iops = true

  preferred_engine_versions  = ["8.0.20", "8.0.19", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsKerberosAuthentication() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "postgres"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                           = data.aws_rds_engine_version.default.engine
  license_model                    = "postgresql-license"
  storage_type                     = "standard"
  supports_kerberos_authentication = true

  preferred_engine_versions  = ["14.1", "13.5", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsPerformanceInsights() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                        = data.aws_rds_engine_version.default.engine
  license_model                 = "general-public-license"
  supports_performance_insights = true

  preferred_engine_versions  = ["8.0.25", "8.0.26", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsStorageAutoScaling() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                       = data.aws_rds_engine_version.default.engine
  license_model                = "general-public-license"
  supports_storage_autoscaling = true

  preferred_engine_versions  = ["8.0.20", "8.0.19", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}

func testAccOrderableInstanceDataSourceConfig_supportsStorageEncryption() string {
	return fmt.Sprintf(`
data "aws_rds_engine_version" "default" {
  engine = "mysql"
}

data "aws_rds_orderable_db_instance" "test" {
  engine                      = data.aws_rds_engine_version.default.engine
  license_model               = "general-public-license"
  storage_type                = "standard"
  supports_storage_encryption = true

  preferred_engine_versions  = ["8.0.25", "8.0.26", data.aws_rds_engine_version.default.version]
  preferred_instance_classes = [%[1]s]
}
`, mainInstanceClasses)
}
