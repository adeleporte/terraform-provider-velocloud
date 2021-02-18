package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourceAddressGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAddressGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"logicalid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterpriseid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func dataSourceAddressGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	enterprise_id := d.Get("enterpriseid").(int)

	if client.Operator && enterprise_id == 0 {
		return diag.Errorf("Enterprise ID is missing (logged as an operator)")
	}

	name := d.Get("name").(string)

	address_group := velo.Enterprise_get_address_group{
		Type:         "address_group",
		EnterpriseID: enterprise_id,
	}

	resp, err := velo.GetAddressGroup(client, address_group)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range resp {
		if v.Name == name {
			d.SetId(fmt.Sprintf("%d", v.ID))
			d.Set("name", v.Name)
			d.Set("logicalid", v.LogicalID)
		}
	}

	return diags
}
