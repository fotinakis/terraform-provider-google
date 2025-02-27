// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleGlobalInternallbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name                  = "tf-test-website-forwarding-rule%{random_suffix}"
  region                = "us-central1"
  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  allow_global_access   = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}
resource "google_compute_region_backend_service" "backend" {
  name                  = "tf-test-website-backend%{random_suffix}"
  region                = "us-central1"
  health_checks         = [google_compute_health_check.hc.id]
}
resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
resource "google_compute_network" "default" {
  name = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}
resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleBasicExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region", "port_range", "target"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_forwarding_rule" "default" {
  name       = "tf-test-website-forwarding-rule%{random_suffix}"
  target     = google_compute_target_pool.default.id
  port_range = "80"
}

resource "google_compute_target_pool" "default" {
  name = "tf-test-website-target-pool%{random_suffix}"
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleInternallbExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleInternallbExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region", "port_range", "target"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleInternallbExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name   = "tf-test-website-forwarding-rule%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
  ip_version            = "IPV4"
}

resource "google_compute_region_backend_service" "backend" {
  name          = "tf-test-website-backend%{random_suffix}"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-website-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-website-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-website-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleVpcPscExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleVpcPscExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region", "port_range", "target", "ip_address"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleVpcPscExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Forwarding rule for VPC private service connect
resource "google_compute_forwarding_rule" "default" {
  name                    = "tf-test-psc-endpoint%{random_suffix}"
  region                  = "us-central1"
  load_balancing_scheme   = ""
  target                  = google_compute_service_attachment.producer_service_attachment.id
  network                 = google_compute_network.consumer_net.name
  ip_address              = google_compute_address.consumer_address.id
  allow_psc_global_access = true
}

// Consumer service endpoint

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-consumer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-consumer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_address" "consumer_address" {
  name         = "tf-test-website-ip%{random_suffix}-1"
  region       = "us-central1"
  subnetwork   = google_compute_subnetwork.consumer_subnet.id
  address_type = "INTERNAL"
}


// Producer service attachment

resource "google_compute_network" "producer_net" {
  name                    = "tf-test-producer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "producer_subnet" {
  name          = "tf-test-producer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_subnetwork" "psc_producer_subnet" {
  name          = "tf-test-producer-psc-net%{random_suffix}"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"

  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_service_attachment" "producer_service_attachment" {
  name        = "tf-test-producer-service%{random_suffix}"
  region      = "us-central1"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = true
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_producer_subnet.name]
  target_service        = google_compute_forwarding_rule.producer_target_service.id
}

resource "google_compute_forwarding_rule" "producer_target_service" {
  name     = "tf-test-producer-forwarding-rule%{random_suffix}"
  region   = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.producer_net.name
  subnetwork            = google_compute_subnetwork.producer_subnet.name
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name     = "tf-test-producer-service-backend%{random_suffix}"
  region   = "us-central1"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  name     = "tf-test-producer-service-health-check%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleVpcPscNoAutomateDnsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleVpcPscNoAutomateDnsExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region", "port_range", "target", "ip_address"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleVpcPscNoAutomateDnsExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_forwarding_rule" "default" {
  name                    = "tf-test-psc-endpoint%{random_suffix}"
  region                  = "us-central1"
  load_balancing_scheme   = ""
  target                  = google_compute_service_attachment.producer_service_attachment.id
  network                 = google_compute_network.consumer_net.name
  ip_address              = google_compute_address.consumer_address.id
  allow_psc_global_access = true
  no_automate_dns_zone    = true
}

resource "google_compute_network" "consumer_net" {
  name                    = "tf-test-consumer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "consumer_subnet" {
  name          = "tf-test-consumer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.consumer_net.id
}

resource "google_compute_address" "consumer_address" {
  name         = "tf-test-website-ip%{random_suffix}-1"
  region       = "us-central1"
  subnetwork   = google_compute_subnetwork.consumer_subnet.id
  address_type = "INTERNAL"
}


resource "google_compute_network" "producer_net" {
  name                    = "tf-test-producer-net%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "producer_subnet" {
  name          = "tf-test-producer-net%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_subnetwork" "psc_producer_subnet" {
  name          = "tf-test-producer-psc-net%{random_suffix}"
  ip_cidr_range = "10.1.0.0/16"
  region        = "us-central1"

  purpose       = "PRIVATE_SERVICE_CONNECT"
  network       = google_compute_network.producer_net.id
}

resource "google_compute_service_attachment" "producer_service_attachment" {
  name        = "tf-test-producer-service%{random_suffix}"
  region      = "us-central1"
  description = "A service attachment configured with Terraform"

  enable_proxy_protocol = true
  connection_preference = "ACCEPT_AUTOMATIC"
  nat_subnets           = [google_compute_subnetwork.psc_producer_subnet.name]
  target_service        = google_compute_forwarding_rule.producer_target_service.id
}

resource "google_compute_forwarding_rule" "producer_target_service" {
  name     = "tf-test-producer-forwarding-rule%{random_suffix}"
  region   = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.producer_service_backend.id
  all_ports             = true
  network               = google_compute_network.producer_net.name
  subnetwork            = google_compute_subnetwork.producer_subnet.name
}

resource "google_compute_region_backend_service" "producer_service_backend" {
  name     = "tf-test-producer-service-backend%{random_suffix}"
  region   = "us-central1"

  health_checks = [google_compute_health_check.producer_service_health_check.id]
}

resource "google_compute_health_check" "producer_service_health_check" {
  name     = "tf-test-producer-service-health-check%{random_suffix}"

  check_interval_sec = 1
  timeout_sec        = 1
  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.steering",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleRegionalSteeringExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_compute_forwarding_rule" "steering" {
  name = "tf-test-steering-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
  source_ip_ranges = ["34.121.88.0/24", "35.187.239.137"]
  depends_on = [google_compute_forwarding_rule.external]
}

resource "google_compute_address" "basic" {
  name = "tf-test-website-ip%{random_suffix}"
  region = "us-central1"
}

resource "google_compute_region_backend_service" "external" {
  name = "tf-test-service-backend%{random_suffix}"
  region = "us-central1"
  load_balancing_scheme = "EXTERNAL"
}

resource "google_compute_forwarding_rule" "external" {
  name = "tf-test-external-forwarding-rule%{random_suffix}"
  region = "us-central1"
  ip_address = google_compute_address.basic.self_link
  backend_service = google_compute_region_backend_service.external.self_link
  load_balancing_scheme = "EXTERNAL"
}
`, context)
}

func TestAccComputeForwardingRule_forwardingRuleInternallbIpv6Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckComputeForwardingRuleDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_forwardingRuleInternallbIpv6Example(context),
			},
			{
				ResourceName:            "google_compute_forwarding_rule.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"backend_service", "network", "subnetwork", "no_automate_dns_zone", "region", "port_range", "target"},
			},
		},
	})
}

func testAccComputeForwardingRule_forwardingRuleInternallbIpv6Example(context map[string]interface{}) string {
	return acctest.Nprintf(`
// Forwarding rule for Internal Load Balancing
resource "google_compute_forwarding_rule" "default" {
  name   = "tf-test-ilb-ipv6-forwarding-rule%{random_suffix}"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
  ip_version            = "IPV6"
}

resource "google_compute_region_backend_service" "backend" {
  name          = "tf-test-ilb-ipv6-backend%{random_suffix}"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_health_check" "hc" {
  name               = "check-tf-test-ilb-ipv6-backend%{random_suffix}"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "default" {
  name                    = "tf-test-net-ipv6%{random_suffix}"
  auto_create_subnetworks = false
  enable_ula_internal_ipv6 = true
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-subnet-internal-ipv6%{random_suffix}"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  stack_type       = "IPV4_IPV6"
  ipv6_access_type = "INTERNAL"
  network       = google_compute_network.default.id
}
`, context)
}

func testAccCheckComputeForwardingRuleDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_forwarding_rule" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/forwardingRules/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("ComputeForwardingRule still exists at %s", url)
			}
		}

		return nil
	}
}
