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
			"enterpriseid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"activationkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"activationstate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"edgestate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hastate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"islive": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"servicestate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	edgename := d.Get("name").(string)
	id, err := velo.GetEdges(client, edgename)

	edge := velo.Enterprise_get_edge{
		ID: id,
	}

	resp, err := velo.ReadEdge(client, edge)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("activationkey", resp.ActivationKey)
	d.Set("activationstate", resp.ActivationState)
	d.Set("edgestate", resp.EdgeState)
	d.Set("hastate", resp.HaState)
	d.Set("islive", resp.IsLive)
	d.Set("servicestate", resp.ServiceState)

	d.SetId(fmt.Sprintf("%d", id))

	return diags
}
