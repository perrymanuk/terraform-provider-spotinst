package elastigroup_gcp_network_interface

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst/commons"
)

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//            Setup
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {

	fieldsMap[NetworkInterface] = commons.NewGenericField(
		commons.ElastigroupGCPNetworkInterface,
		NetworkInterface,
		&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					string(Network): {
						Type:     schema.TypeString,
						Required: true,
					},
					string(ProjectID): {
						Type:     schema.TypeString,
						Required: false,
						Optional: true,
					},

					string(AccessConfigs): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(Name): {
									Type:     schema.TypeString,
									Optional: true,
								},

								string(Type): {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},

					string(AliasIPRanges): {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								string(IPCIDRRange): {
									Type:     schema.TypeString,
									Required: true,
								},

								string(SubnetworkRangeName): {
									Type:     schema.TypeString,
									Required: true,
								},
							},
						},
					},
				},
			},
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(NetworkInterface)); ok {
				if networks, err := expandNetworkInterface(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetNetworkInterfaces(networks)
				}
			}
			return nil
		},
		func(resourceObject interface{}, resourceData *schema.ResourceData, meta interface{}) error {
			egWrapper := resourceObject.(*commons.ElastigroupGCPWrapper)
			elastigroup := egWrapper.GetElastigroup()
			if v, ok := resourceData.GetOk(string(NetworkInterface)); ok {
				if networks, err := expandNetworkInterface(v); err != nil {
					return err
				} else {
					elastigroup.Compute.LaunchSpecification.SetNetworkInterfaces(networks)
				}
			}
			return nil
		},
		nil,
	)

}

//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
//         Expand Fields
//-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

// expandNetworkInterface sets the values from the plan as objects
func expandNetworkInterface(data interface{}) ([]*gcp.NetworkInterface, error) {
	list := data.([]interface{})

	if list != nil && list[0] != nil {
		ifaces := make([]*gcp.NetworkInterface, 0, len(list))
		for _, item := range list {
			m := item.(map[string]interface{})
			iface := &gcp.NetworkInterface{}

			if v, ok := m[string(Network)].(string); ok && v != "" {
				iface.SetNetwork(spotinst.String(v))
			}

			if v, ok := m[string(ProjectID)].(string); ok && v != "" {
				iface.SetProjectID(spotinst.String(v))
			}

			if v, ok := m[string(AccessConfigs)]; ok {
				accessConfigs, err := expandAccessConfigs(v)
				if err != nil {
					return nil, err
				}

				if accessConfigs != nil {
					iface.SetAccessConfigs(accessConfigs)
				}
			} else {
				iface.AccessConfigs = nil
			}

			if v, ok := m[string(AliasIPRanges)]; ok {
				aliasRange, err := expandAliasIPRanges(v)
				if err != nil {
					return nil, err
				}
				if aliasRange != nil {
					iface.SetAliasIPRanges(aliasRange)
				}
			} else {
				iface.AliasIPRanges = nil
			}

			ifaces = append(ifaces, iface)
		}
		return ifaces, nil
	}
	return nil, nil
}

// expandAccessConfigs sets the values from the plan as objects
func expandAccessConfigs(data interface{}) ([]*gcp.AccessConfig, error) {
	list := data.(*schema.Set).List()
	accessConfigs := make([]*gcp.AccessConfig, 0, len(list))

	for _, item := range list {
		attr := item.(map[string]interface{})

		accessConfig := &gcp.AccessConfig{}

		if v, ok := attr[string(Name)].(string); ok && v != "" {
			accessConfig.SetName(spotinst.String(v))
		}

		if v, ok := attr[string(Type)].(string); ok && v != "" {
			accessConfig.SetType(spotinst.String(v))
		}

		accessConfigs = append(accessConfigs, accessConfig)
	}
	return accessConfigs, nil
}

// expandAccessConfigs sets the values from the plan as objects
func expandAliasIPRanges(data interface{}) ([]*gcp.AliasIPRange, error) {
	list := data.(*schema.Set).List()
	aliasRanges := make([]*gcp.AliasIPRange, 0, len(list))

	for _, item := range list {
		m := item.(map[string]interface{})
		aliasRange := &gcp.AliasIPRange{}

		if v, ok := m[string(SubnetworkRangeName)].(string); ok && v != "" {
			aliasRange.SetSubnetworkRangeName(spotinst.String(v))
		}

		if v, ok := m[string(IPCIDRRange)].(string); ok && v != "" {
			aliasRange.SetIPCIDRRange(spotinst.String(v))
		}

		aliasRanges = append(aliasRanges, aliasRange)
	}
	return aliasRanges, nil
}
