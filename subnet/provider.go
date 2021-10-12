package subnet

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"subnet_single": dataSubnetSet(),
			"subnet_list":   dataSubnetListSet(),
		},
	}
}

func dataSubnetSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSubnetRead,
		Schema: map[string]*schema.Schema{
			"cidr": {
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"included": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSubnetListSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSubnetListRead,
		Schema: map[string]*schema.Schema{
			"cidr_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"included": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"included_subnet_cidr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"included_subnet_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
