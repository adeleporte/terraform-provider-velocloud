package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDeviceSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeviceSettingsCreate,
		ReadContext:   resourceDeviceSettingsRead,
		UpdateContext: resourceDeviceSettingsUpdate,
		DeleteContext: resourceDeviceSettingsDelete,
		Schema: map[string]*schema.Schema{
			"profile": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"enterpriseid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"segment": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vlan": {
				Type:        schema.TypeList,
				Description: "Vlan description",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_prefix": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  24,
						},
					},
				},
			},
			"routed_interface": {
				Type:        schema.TypeList,
				Description: "Interface description",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_prefix": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  24,
						},
						"gateway": {
							Type:     schema.TypeString,
							Required: true,
						},
						"netmask": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "DHCP",
						},
						"override": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
		},
	}
}

func resourceDeviceSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics

	client := m.(*velo.Client)

	// Get info from the schema
	edgeprofile_id, _ := d.Get("profile").(int)
	enterprise_id := d.Get("enterpriseid").(int)
	vlan := (d.Get("vlan").([]interface{}))[0].(map[string]interface{})
	cidr_ip := vlan["cidr_ip"].(string)
	cidr_prefix := vlan["cidr_prefix"].(int)

	routed_interfaces := d.Get("routed_interface").([]interface{})

	dmodule, err := velo.GetDeviceSettingsModule(client, enterprise_id, edgeprofile_id)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get info from module
	id := int(dmodule["id"].(float64))
	data := dmodule["data"].(map[string]interface{})
	lan := data["lan"].(map[string]interface{})
	networks := lan["networks"].([]interface{})
	network0 := networks[0].(map[string]interface{})
	interfaces := data["routedInterfaces"].([]interface{})

	// Update the module
	network0["cidrIp"] = cidr_ip
	network0["cidrPrefix"] = cidr_prefix

	for _, v := range interfaces {
		intf := v.(map[string]interface{})
		addressing := intf["addressing"].(map[string]interface{})
		for _, w := range routed_interfaces {
			item := w.(map[string]interface{})
			if intf["name"] == item["name"] {
				intf["override"] = item["override"]
				addressing["cidrIp"] = item["cidr_ip"]
				addressing["cidrPrefix"] = item["cidr_prefix"]
				addressing["gateway"] = item["gateway"]
				addressing["netmask"] = item["netmask"]
				addressing["type"] = item["type"]
			}
		}

	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	log.Println(buf)

	_, err = velo.UpdateDeviceSettingsModule(client, enterprise_id, id, data)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceDeviceSettingsRead(ctx, d, m)
}

func resourceDeviceSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// To be implemented

	return diags

}

func resourceDeviceSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics

	// To be implemented

	return resourceDeviceSettingsCreate(ctx, d, m)

}

func resourceDeviceSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	// Get info from the schema
	edgeprofile_id, _ := d.Get("profile").(int)
	enterprise_id := d.Get("enterpriseid").(int)

	if client.Operator && enterprise_id == 0 {
		return diag.Errorf("Enterprise ID is missing (logged as an operator)")
	}

	routed_interfaces := d.Get("routed_interface").([]interface{})

	dmodule, err := velo.GetDeviceSettingsModule(client, enterprise_id, edgeprofile_id)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get info from module
	id := int(dmodule["id"].(float64))
	data := dmodule["data"].(map[string]interface{})
	//lan := data["lan"].(map[string]interface{})
	//networks := lan["networks"].([]interface{})
	//network0 := networks[0].(map[string]interface{})
	interfaces := data["routedInterfaces"].([]interface{})

	// Update the module
	//network0["cidrIp"] = nil
	//network0["cidrPrefix"] = nil
	// Need to find a way to reset vlan cidr and prefix

	for _, v := range interfaces {
		intf := v.(map[string]interface{})
		addressing := intf["addressing"].(map[string]interface{})
		for _, w := range routed_interfaces {
			item := w.(map[string]interface{})
			if intf["name"] == item["name"] {
				intf["override"] = false
				addressing["cidrIp"] = nil
				addressing["cidrPrefix"] = nil
				addressing["gateway"] = nil
				addressing["netmask"] = nil
				addressing["type"] = "DHCP"
			}
		}

	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	log.Println(buf)

	_, err = velo.UpdateDeviceSettingsModule(client, enterprise_id, id, data)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
