package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func resourcePortGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortGroupCreate,
		ReadContext:   resourcePortGroupRead,
		UpdateContext: resourcePortGroupUpdate,
		DeleteContext: resourcePortGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"logicalid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:        schema.TypeList,
				Description: "List of Port ranges",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"port_low": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"port_high": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePortGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	port_ranges := d.Get("range").([]interface{})
	data := make([]velo.Port_group_data, len(port_ranges))

	for i, port_range := range port_ranges {
		ppr := port_range.(map[string]interface{})
		data[i].Proto = ppr["proto"].(int)
		data[i].PortLow = ppr["port_low"].(int)
		data[i].PortHigh = ppr["port_high"].(int)

	}

	port_group := velo.Enterprise_insert_port_group{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        "port_group",
		Data:        data,
	}

	resp, err := velo.InsertPortGroup(client, port_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))

	resourcePortGroupRead(ctx, d, m)

	return diags
}

func resourcePortGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	Port_group := velo.Enterprise_get_port_group{
		Type: "port_group",
	}

	resp, err := velo.GetPortGroup(client, Port_group)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, v := range resp {
		if v.ID == id {
			d.Set("name", v.Name)
			d.Set("logicalid", v.LogicalID)
		}
	}

	return diags
}

func resourcePortGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	port_ranges := d.Get("range").([]interface{})
	data := make([]velo.Port_group_data, len(port_ranges))

	for i, ip_range := range port_ranges {
		ppr := ip_range.(map[string]interface{})
		data[i].Proto = ppr["proto"].(int)
		data[i].PortLow = ppr["port_low"].(int)
		data[i].PortHigh = ppr["port_high"].(int)
	}

	Port_group := velo.Enterprise_update_port_group{
		ID:          id,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Data:        data,
	}

	resp, err := velo.UpdatePortGroup(client, Port_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))

	resourcePortGroupRead(ctx, d, m)

	return diags
}

func resourcePortGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	port_group := velo.Enterprise_delete_port_group{
		ID: id,
	}

	_, err := velo.DeletePortGroup(client, port_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
