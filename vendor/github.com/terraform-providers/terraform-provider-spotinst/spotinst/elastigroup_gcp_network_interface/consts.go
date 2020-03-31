package elastigroup_gcp_network_interface

import "github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"

const (
	NetworkInterface commons.FieldName = "network_interface"
	Network          commons.FieldName = "network"
	ProjectID        commons.FieldName = "subnetwork_project"
	AccessConfigs    commons.FieldName = "access_configs"
	AliasIPRanges    commons.FieldName = "alias_ip_ranges"

	IPCIDRRange         commons.FieldName = "ip_cidr_range"
	SubnetworkRangeName commons.FieldName = "subnetwork_range_name"
	Name                commons.FieldName = "name"
	Type                commons.FieldName = "type"
)
