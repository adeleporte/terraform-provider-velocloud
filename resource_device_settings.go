package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getCidrNetmask(cidrIP string, cidrPrefix int) string {
	_, ipv4Net, err := net.ParseCIDR(fmt.Sprintf("%s/%d", cidrIP, cidrPrefix))
	if err != nil {
		log.Fatal(err)
	}
	mask := ipv4Net.Mask

	if len(mask) != 4 {
		panic("ipv4Mask: length must be 4 bytes")
	}

	return fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
}

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
						"advertise": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"override": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"dhcp_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
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
							Optional: true,
							Default:  "",
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
						"nat_direct": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"wan_overlay": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
			"static_route": {
				Type:        schema.TypeList,
				Description: "Static route description",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet_cidr_ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"subnet_cidr_prefix": {
							Type:     schema.TypeString,
							Required: true,
						},
						"next_hop": {
							Type:     schema.TypeString,
							Required: true,
						},
						"interface": {
							Type:     schema.TypeString,
							Required: true,
						},
						"cost": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"preferred": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"advertise": {
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
	override := vlan["override"].(bool)
	advertise := vlan["advertise"].(bool)
	dhcp_enabled := vlan["dhcp_enabled"].(bool)

	routed_interfaces := d.Get("routed_interface").([]interface{})
	static_routes := d.Get("static_route").([]interface{})

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
	dhcp := network0["dhcp"].(map[string]interface{})
	interfaces := data["routedInterfaces"].([]interface{})

	// Update the module
	network0["cidrIp"] = cidr_ip
	network0["cidrPrefix"] = cidr_prefix
	network0["advertise"] = advertise
	network0["override"] = override
	network0["netmask"] = getCidrNetmask(cidr_ip, cidr_prefix)

	if dhcp_enabled == false {
		dhcp["enabled"] = false
		dhcp["override"] = true
	} else {
		dhcp["enabled"] = true
		dhcp["override"] = false
	}

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
	lan := data["lan"].(map[string]interface{})
	networks := lan["networks"].([]interface{})
	network0 := networks[0].(map[string]interface{})
	interfaces := data["routedInterfaces"].([]interface{})
	dhcp := network0["dhcp"].(map[string]interface{})

	// Update the module
	network0["cidrIp"] = nil
	network0["cidrPrefix"] = nil
	network0["advertise"] = true
	network0["override"] = false

	dhcp["leaseTimeSeconds"] = 3600
	dhcp["override"] = false

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

	data["segments"].([]interface{})[0].(map[string]interface{})["routes"].(map[string]interface{})["static"] = []int{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)

	log.Println(buf)

	_, err = velo.UpdateDeviceSettingsModule(client, enterprise_id, id, data)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags

}
