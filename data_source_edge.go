package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourceEdge() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEdgeRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	edgeame := d.Get("name").(string)
	id, err := velo.GetEdges(client, edgeame)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", int(id)))

	return diags
}
