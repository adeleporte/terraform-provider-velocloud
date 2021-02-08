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
			"profile": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"enterpriseid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"segment": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"vlan": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Vlan description",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  24,
						},
					},
				},
			},
			"routed_interface": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Interface description",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cidr_prefix": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  24,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "DHCP",
						},
						"override": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"nat_direct": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"wan_overlay": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
			"static_route": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Static route description",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_cidr_ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_cidr_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"next_hop": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"preferred": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"advertise": &schema.Schema{
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
	//enterprise_id := d.Get("enterpriseid").(int)
	vlan := (d.Get("vlan").([]interface{}))[0].(map[string]interface{})
	cidr_ip := vlan["cidr_ip"].(string)
	cidr_prefix := vlan["cidr_prefix"].(int)

	routed_interfaces := d.Get("routed_interface").([]interface{})
	static_routes := d.Get("static_route").([]interface{})

	dmodule, err := velo.GetDeviceSettingsModule(client, edgeprofile_id)
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
	// routes := data["segments"].([]interface{})[0].(map[string]interface{})["routes"].(map[string]interface{})["static"].([]interface{})

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
				intf["natDirect"] = item["nat_direct"]
				addressing["cidrIp"] = item["cidr_ip"]
				addressing["cidrPrefix"] = item["cidr_prefix"]
				addressing["gateway"] = item["gateway"]
				addressing["netmask"] = item["netmask"]
				addressing["type"] = item["type"]
				if item["gateway"] == "" {
					addressing["gateway"] = nil
				} else {
					addressing["gateway"] = item["gateway"]
				}
				if item["wan_overlay"] == true {
					addressing["wanOverlay"] = "AUTO_DISCOVERED"
				} else {
					addressing["wanOverlay"] = "DISABLED"
				}
			}
		}

	}
	type Route struct {
		Destination    string `json:"destination"`
		Gateway        string `json:"gateway"`
		WanInterface   string `json:"wanInterface"`
		CidrPrefix     string `json:"cidrPrefix"`
		Cost           int    `json:"cost"`
		SubInterfaceID int    `json:"subinterfaceId"`
		Preferred      bool   `json:"preferred"`
		Advertise      bool   `json:"advertise"`
	}

	var routes []interface{}

	for _, w := range static_routes {
		item := w.(map[string]interface{})
		route := Route{
			Destination:    item["subnet_cidr_ip"].(string),
			CidrPrefix:     item["subnet_cidr_prefix"].(string),
			Gateway:        item["next_hop"].(string),
			WanInterface:   item["interface"].(string),
			Cost:           item["cost"].(int),
			Preferred:      item["preferred"].(bool),
			Advertise:      item["advertise"].(bool),
			SubInterfaceID: -1,
		}
		routes = append(routes, route)
	}

	data["segments"].([]interface{})[0].(map[string]interface{})["routes"].(map[string]interface{})["static"] = routes

	log.Println(data)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	log.Println(buf)

	_, err = velo.UpdateDeviceSettingsModule(client, id, data)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return resourceDeviceSettingsRead(ctx, d, m)
}

func resourceDeviceSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags

}

func resourceDeviceSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics

	return resourceDeviceSettingsCreate(ctx, d, m)

}

func resourceDeviceSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	// Get info from the schema
	edgeprofile_id, _ := d.Get("profile").(int)
	//enterprise_id := d.Get("enterpriseid").(int)

	routed_interfaces := d.Get("routed_interface").([]interface{})

	dmodule, err := velo.GetDeviceSettingsModule(client, edgeprofile_id)
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
				intf["natDirect"] = true
				addressing["cidrIp"] = nil
				addressing["cidrPrefix"] = nil
				addressing["gateway"] = nil
				addressing["netmask"] = nil
				addressing["type"] = "DHCP"
				addressing["wanOverlay"] = "AUTO_DISCOVERED"
			}
		}

	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	log.Println(buf)

	_, err = velo.UpdateDeviceSettingsModule(client, id, data)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
