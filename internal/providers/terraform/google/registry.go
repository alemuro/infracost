package google

import "github.com/infracost/infracost/internal/schema"

var ResourceRegistry []*schema.RegistryItem = []*schema.RegistryItem{
	GetComputeDiskRegistryItem(),
	GetComputeImageRegistryItem(),
	GetComputeSnapshotRegistryItem(),
	GetComputeInstanceRegistryItem(),
	GetComputeAddressRegistryItem(),
	GetComputeGlobalAddressRegistryItem(),
}

// FreeResources grouped alphabetically
var FreeResources []string = []string{
	"google_compute_attached_disk",
	"google_compute_backend_bucket",
	"google_compute_backend_bucket_signed_url_key",
	"google_compute_backend_service",
	"google_compute_backend_service_signed_url_key",
	"google_compute_disk_iam",
	"google_compute_firewall",
	"google_compute_global_network_endpoint",
	"google_compute_global_network_endpoint_group",
	"google_compute_health_check",
	"google_compute_http_health_check",
	"google_compute_https_health_check",
	"google_compute_image_iam",
	"google_compute_instance_group",
	"google_compute_instance_group_named_port",
	"google_compute_instance_iam",
	"google_compute_machine_image_iam",
	"google_compute_managed_ssl_certificate",
	"google_compute_network",
	"google_compute_network_endpoint",
	"google_compute_network_endpoint_group",
	"google_compute_network_peering",
	"google_compute_network_peering_routes_config",
	"google_compute_project_default_network_tier",
	"google_compute_project_metadata",
	"google_compute_project_metadata_item",
	"google_compute_region_backend_service",
	"google_compute_region_disk_iam",
	"google_compute_region_health_check",
	"google_compute_region_network_endpoint_group",
	"google_compute_region_per_instance_config",
	"google_compute_region_url_map",
	"google_compute_route",
	"google_compute_router",
	"google_compute_router_bgp_peer",
	"google_compute_router_interface",
	"google_compute_shared_vpc_host_project",
	"google_compute_shared_vpc_service_project",
	"google_compute_ssl_certificate",
	"google_compute_ssl_policy",
	"google_compute_subnetwork",
	"google_compute_subnetwork_iam",
	"google_compute_url_map",
}

var UsageOnlyResources []string = []string{}

// TODO: This is a list of all the google_compute* resources that may have prices:
// compute_instance scratch_disk
// VM instance (https://cloud.google.com/compute/vm-instance-pricing):
// google_compute_instance_from_machine_image
// google_compute_instance_from_template
//
// Node groups and autoscaling:
// google_compute_autoscaler
// google_compute_instance_template
// google_compute_target_pool
// google_compute_instance_group_manager
// google_compute_per_instance_config
// google_compute_region_autoscaler
// google_compute_node_group
// google_compute_node_template
// google_compute_region_instance_group_manager
// google_compute_region_per_instance_config
//
// Disk and images (https://cloud.google.com/compute/disks-image-pricing):
// google_compute_image
// google_compute_machine_image
// google_compute_region_disk
// google_compute_snapshot
//
// Cloud VPN (https://cloud.google.com/vpc/network-pricing#vpn-pricing):
// google_compute_vpn_gateway
// google_compute_vpn_tunnel
// google_compute_external_vpn_gateway
// google_compute_ha_vpn_gateway
//
// Load balancers (https://cloud.google.com/vpc/network-pricing#lb):
// google_compute_forwarding_rule
// google_compute_global_forwarding_rule
// google_compute_target_grpc_proxy
// google_compute_target_http_proxy
// google_compute_target_https_proxy
// google_compute_target_ssl_proxy
// google_compute_target_tcp_proxy
// google_compute_region_target_http_proxy
// google_compute_region_target_https_proxy
//
// Cloud NAT (https://cloud.google.com/vpc/network-pricing#nat-pricing):
// google_compute_router_nat
//
// Packet mirroring (https://cloud.google.com/vpc/network-pricing#packet-mirroring):
// google_compute_packet_mirroring
//
// Cloud interconnect (https://cloud.google.com/vpc/network-pricing#interconnect-pricing):
// google_compute_interconnect_attachment
//
// Others:
// google_compute_organization_security_policy
// google_compute_organization_security_policy_association
// google_compute_organization_security_policy_rule
// google_compute_region_disk_resource_policy_attachment
// google_compute_reservation
// google_compute_resource_policy
// google_compute_security_policy
// google_usage_export_bucket
