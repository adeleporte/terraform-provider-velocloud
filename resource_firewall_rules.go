package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func resourceFirewallRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallRulesCreate,
		ReadContext:   resourceFirewallRulesRead,
		UpdateContext: resourceFirewallRulesUpdate,
		DeleteContext: resourceFirewallRulesDelete,
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
			"firewall_status": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"firewall_stateful": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"firewall_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"firewall_syslog": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Rules description",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"dip": {
							Type:     schema.TypeString,
							Default:  "any",
							Optional: true,
						},
						"sip": {
							Type:     schema.TypeString,
							Default:  "any",
							Optional: true,
						},
						"dsm": {
							Type:     schema.TypeString,
							Default:  "any",
							Optional: true,
						},
						"ssm": {
							Type:     schema.TypeString,
							Default:  "any",
							Optional: true,
						},
						"d_port_low": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"d_port_high": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"s_port_low": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"s_port_high": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"s_address_group": {
							Type:     schema.TypeString,
							Default:  "",
							Optional: true,
						},
						"d_address_group": {
							Type:     schema.TypeString,
							Default:  "",
							Optional: true,
						},
						"s_port_group": {
							Type:     schema.TypeString,
							Default:  "",
							Optional: true,
						},
						"d_port_group": {
							Type:     schema.TypeString,
							Default:  "",
							Optional: true,
						},
						"appid": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"classid": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"dscp": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"svlan": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"dvlan": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"hostname": {
							Type:     schema.TypeString,
							Default:  "",
							Optional: true,
						},
						"os_version": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"proto": {
							Type:     schema.TypeInt,
							Default:  -1,
							Optional: true,
						},
						"smac": {
							Type:     schema.TypeString,
							Default:  "any",
							Optional: true,
						},
						"d_rule_type": {
							Type:     schema.TypeString,
							Default:  "prefix",
							Optional: true,
						},
						"s_rule_type": {
							Type:     schema.TypeString,
							Default:  "prefix",
							Optional: true,
						},
						"action": {
							Type:     schema.TypeString,
							Default:  "allow",
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceFirewallRulesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)
	rulesFromSchema := d.Get("rule").([]interface{})
	enterprise_id := d.Get("enterpriseid").(int)

	fwmodule, err := velo.GetFirewallModule(client, enterprise_id, profile_id)
	if err != nil {
		return diag.FromErr(err)
	}

	raw := fwmodule.(map[string]interface{})
	fw_module_id := int(raw["id"].(float64))

	outbound_rules := make([]velo.FirewallOutboundRule, len(rulesFromSchema))
	for i, ruleFromSchema := range rulesFromSchema {
		rule := ruleFromSchema.(map[string]interface{})
		outbound_rules[i] = velo.FirewallOutboundRule{
			Name: rule["name"].(string),
			Match: velo.FirewallRuleMatch{
				DIP:           rule["dip"].(string),
				DSM:           rule["dsm"].(string),
				SIP:           rule["sip"].(string),
				SSM:           rule["ssm"].(string),
				SAddressGroup: rule["s_address_group"].(string),
				DAddressGroup: rule["d_address_group"].(string),
				SPortGroup:    rule["s_port_group"].(string),
				DPortGroup:    rule["d_port_group"].(string),
				DPortLow:      rule["d_port_low"].(int),
				DPortHigh:     rule["d_port_high"].(int),
				SPortLow:      rule["s_port_low"].(int),
				SPortHigh:     rule["s_port_high"].(int),
				AppID:         rule["appid"].(int),
				ClassID:       rule["classid"].(int),
				DRuleType:     rule["d_rule_type"].(string),
				SRuleType:     rule["s_rule_type"].(string),
				Dscp:          rule["dscp"].(int),
				DVlan:         rule["dvlan"].(int),
				SVlan:         rule["svlan"].(int),
				Hostname:      rule["hostname"].(string),
				OSVersion:     rule["os_version"].(int),
				Proto:         rule["proto"].(int),
				SMac:          rule["smac"].(string),
			},
			Action: velo.FirewallOutboundAction{
				AllowOrDeny: rule["action"].(string),
			},
		}
	}

	fw_data := velo.FirewallData{
		FirewallEnabled:         d.Get("firewall_status").(bool),
		StatefulFirewallEnabled: d.Get("firewall_stateful").(bool),
		FirewallLoggingEnabled:  d.Get("firewall_logging").(bool),
		SyslogForwarding:        d.Get("firewall_syslog").(bool),
		Inbound:                 []velo.FirewallInboundRule{},
		Segments: []velo.FirewallSegment{{
			Outbound: outbound_rules,
			Segment: velo.ModuleSegmentMetaData{
				SegmentID: segment_id,
			},
		},
		},
	}

	update := velo.ConfigurationFirewallModule{
		Name: "firewall",
		Data: fw_data,
	}

	firewall := velo.UpdateConfigurationFirewallModuleBody{
		ID:           int(fw_module_id),
		EnterpriseID: d.Get("enterpriseid").(int),
		Update:       update,
	}

	_, err = velo.UpdateFirewallModule(client, firewall)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(fw_module_id))

	resourceFirewallRulesRead(ctx, d, m)

	return diags
}

func resourceFirewallRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)
	enterprise_id := d.Get("enterpriseid").(int)

	fwmodule, err := velo.GetFirewallModules(client, profile_id, enterprise_id)
	if err != nil {
		return diag.FromErr(err)
	}

	rulestoschema := []map[string]interface{}{}

	for _, v := range fwmodule.Segments[segment_id].Outbound {
		item := map[string]interface{}{}

		item["name"] = v.Name
		item["action"] = v.Action
		item["dip"] = v.Match.DIP
		item["dsm"] = v.Match.DSM
		item["sip"] = v.Match.SIP
		item["ssm"] = v.Match.SSM
		item["s_address_group"] = v.Match.SAddressGroup
		item["d_address_group"] = v.Match.DAddressGroup
		item["s_port_group"] = v.Match.SPortGroup
		item["d_port_group"] = v.Match.DPortGroup
		item["d_port_low"] = v.Match.DPortLow
		item["d_port_high"] = v.Match.DPortHigh
		item["s_port_low"] = v.Match.SPortLow
		item["d_port_high"] = v.Match.DPortHigh
		item["appid"] = v.Match.AppID
		item["classid"] = v.Match.ClassID
		item["d_rule_type"] = v.Match.DRuleType
		item["s_rule_type"] = v.Match.SRuleType
		item["dscp"] = v.Match.Dscp
		item["dvlan"] = v.Match.DVlan
		item["svlan"] = v.Match.SVlan
		item["hostname"] = v.Match.Hostname
		item["os_version"] = v.Match.OSVersion
		item["proto"] = v.Match.Proto
		item["smac"] = v.Match.SMac

		rulestoschema = append(rulestoschema, item)
	}
	d.Set("firewall_status", fwmodule.FirewallEnabled)
	d.Set("firewall_stateful", fwmodule.StatefulFirewallEnabled)
	d.Set("firewall_logging", fwmodule.FirewallLoggingEnabled)
	d.Set("firewall_syslog", fwmodule.SyslogForwarding)
	d.Set("rule", rulestoschema)

	return diags
}

func resourceFirewallRulesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceFirewallRulesCreate(ctx, d, m)

	return diags
}

func resourceFirewallRulesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)
	enterprise_id := d.Get("enterpriseid").(int)

	if client.Operator && enterprise_id == 0 {
		return diag.Errorf("Enterprise ID is missing (logged as an operator)")
	}

	fwmodule, err := velo.GetFirewallModule(client, enterprise_id, profile_id)
	if err != nil {
		return diag.FromErr(err)
	}

	raw := fwmodule.(map[string]interface{})
	fw_module_id := int(raw["id"].(float64))

	outbound_rules := []velo.FirewallOutboundRule{}

	fw_data := velo.FirewallData{
		FirewallEnabled: true,
		Inbound:         []velo.FirewallInboundRule{},
		Segments: []velo.FirewallSegment{{
			Outbound: outbound_rules,
			Segment: velo.ModuleSegmentMetaData{
				SegmentID: segment_id,
			},
		},
		},
	}

	update := velo.ConfigurationFirewallModule{
		Name: "firewall",
		Data: fw_data,
	}

	firewall := velo.UpdateConfigurationFirewallModuleBody{
		ID:           int(fw_module_id),
		EnterpriseID: d.Get("enterpriseid").(int),
		Update:       update,
	}

	_, err = velo.UpdateFirewallModule(client, firewall)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
