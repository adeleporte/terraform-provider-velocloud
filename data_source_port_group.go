package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourcePortGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePortGroupRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePortGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	name := d.Get("name").(string)

	port_group := velo.Enterprise_get_port_group{
		Type: "port_group",
	}

	resp, err := velo.GetPortGroup(client, port_group)
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
