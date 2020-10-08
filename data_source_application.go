package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourceApplication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	applicationame := d.Get("name").(string)
	id, err := velo.GetApplications(client, applicationame)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", int(id)))

	return diags
}
