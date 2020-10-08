package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourceProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProfilesRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profilename := d.Get("name").(string)
	id, err := velo.GetProfile(client, profilename)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", id))

	return diags
}
