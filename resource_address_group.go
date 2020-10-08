package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func resourceAddressGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAddressGroupCreate,
		ReadContext:   resourceAddressGroupRead,
		UpdateContext: resourceAddressGroupUpdate,
		DeleteContext: resourceAddressGroupDelete,
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
				Description: "List of IP ranges",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"mask": &schema.Schema{
							Type:     schema.TypeString,
							Default:  "255.255.255.255",
							Optional: true,
						},
						"rule_type": &schema.Schema{
							Type:     schema.TypeString,
							Default:  "exact",
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAddressGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	ip_ranges := d.Get("range").([]interface{})
	data := make([]velo.Address_group_data, len(ip_ranges))

	for i, ip_range := range ip_ranges {
		ipr := ip_range.(map[string]interface{})
		data[i].IP = ipr["ip"].(string)
		data[i].Mask = ipr["mask"].(string)
		data[i].RuleType = ipr["rule_type"].(string)
	}

	address_group := velo.Enterprise_insert_address_group{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        "address_group",
		Data:        data,
	}

	resp, err := velo.InsertAddressGroup(client, address_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))

	resourceAddressGroupRead(ctx, d, m)

	return diags
}

func resourceAddressGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	address_group := velo.Enterprise_get_address_group{
		Type: "address_group",
	}

	resp, err := velo.GetAddressGroup(client, address_group)
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

func resourceAddressGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	ip_ranges := d.Get("range").([]interface{})
	data := make([]velo.Address_group_data, len(ip_ranges))

	for i, ip_range := range ip_ranges {
		ipr := ip_range.(map[string]interface{})
		data[i].IP = ipr["ip"].(string)
		data[i].Mask = ipr["mask"].(string)
		data[i].RuleType = ipr["rule_type"].(string)
	}

	address_group := velo.Enterprise_update_address_group{
		ID:          id,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Data:        data,
	}

	resp, err := velo.UpdateAddressGroup(client, address_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))

	resourceAddressGroupRead(ctx, d, m)

	return diags
}

func resourceAddressGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	address_group := velo.Enterprise_delete_address_group{
		ID: id,
	}

	_, err := velo.DeleteAddressGroup(client, address_group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
