package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

var QoSRuleClassValues = []string{"realtime", "transactional", "bulk"}
var QoSServiceGroupValues = []string{"ALL", "PUBLIC_WIRED", "PRIVATE_WIRED"}
var networkserviceValues = []string{"auto", "fixed"}

func resourceBusinessPolicies() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBusinessPoliciesCreate,
		ReadContext:   resourceBusinessPoliciesRead,
		UpdateContext: resourceBusinessPoliciesUpdate,
		DeleteContext: resourceBusinessPoliciesDelete,
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
			"rule": getQoSRulesSchema(),
		},
	}
}

func getQoSRulesSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Description: "List of QoS rules",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"appid": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  -1,
					Optional: true,
				},
				"hostname": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "",
					Optional: true,
				},
				"dip": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "any",
					Optional: true,
				},
				"dsm": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "255.255.255.255",
					Optional: true,
				},
				"dport_low": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  "-1",
					Optional: true,
				},
				"dport_high": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  "-1",
					Optional: true,
				},
				"proto": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  "-1",
					Optional: true,
				},
				"priority": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "normal",
					Optional: true,
				},
				"rxbandwidthpct": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  "-1",
					Optional: true,
				},
				"txbandwidthpct": &schema.Schema{
					Type:     schema.TypeFloat,
					Default:  "-1",
					Optional: true,
				},
				"networkservice": &schema.Schema{
					Type:         schema.TypeString,
					Default:      "auto",
					Optional:     true,
					ValidateFunc: validation.StringInSlice(networkserviceValues, false),
				},
				"serviceclass": &schema.Schema{
					Type:         schema.TypeString,
					Default:      "bulk",
					Optional:     true,
					ValidateFunc: validation.StringInSlice(QoSRuleClassValues, false),
				},
				"linksteering": &schema.Schema{
					Type:         schema.TypeString,
					Default:      "ALL",
					Optional:     true,
					ValidateFunc: validation.StringInSlice(QoSServiceGroupValues, false),
				},
				"saddressgroup": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "",
					Optional: true,
				},
				"daddressgroup": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "",
					Optional: true,
				},
				"sportgroup": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "",
					Optional: true,
				},
				"dportgroup": &schema.Schema{
					Type:     schema.TypeString,
					Default:  "",
					Optional: true,
				},
			},
		},
	}
}

func resourceBusinessPoliciesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)

	qosmodule, err := velo.GetQosModule(client, profile_id)
	if err != nil {
		return diag.FromErr(err)
	}

	raw := qosmodule.(map[string]interface{})
	qos_module_id := int(raw["id"].(float64))
	data := raw["data"].(map[string]interface{})
	segments := data["segments"].([]interface{})
	segment := segments[segment_id].(map[string]interface{})
	defaults := segment["defaults"].([]interface{})

	rulesFromSchema := d.Get("rule").([]interface{})
	new_rules := make([]interface{}, len(rulesFromSchema))
	for i, ruleFromSchema := range rulesFromSchema {
		r := ruleFromSchema.(map[string]interface{})

		// Copy the last default rule
		new_rule, err := velo.DeepCopy(defaults[len(defaults)-1].(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}

		match := new_rule["match"].(map[string]interface{})
		action := new_rule["action"].(map[string]interface{})
		qos := action["QoS"].(map[string]interface{})
		txScheduler := qos["txScheduler"].(map[string]interface{})
		rxScheduler := qos["rxScheduler"].(map[string]interface{})
		edge2CloudRouteAction := action["edge2CloudRouteAction"].(map[string]interface{})
		edge2DataCenterRouteAction := action["edge2DataCenterRouteAction"].(map[string]interface{})
		edge2EdgeRouteAction := action["edge2EdgeRouteAction"].(map[string]interface{})

		// Update rule
		new_rule["name"] = r["name"].(string)

		// Match
		match["appid"] = r["appid"].(float64)
		match["hostname"] = r["hostname"].(string)
		match["dip"] = r["dip"].(string)
		match["dport_low"] = r["dport_low"].(float64)
		match["dport_high"] = r["dport_low"].(float64)
		match["dsm"] = r["dsm"].(string)
		match["proto"] = r["proto"].(float64)

		if daddressgroup, ok := r["daddressgroup"].(string); ok {
			match["dAddressGroup"] = daddressgroup
		}

		if saddressgroup, ok := r["saddressgroup"].(string); ok {
			match["sAddressGroup"] = saddressgroup
		}

		if dportgroup, ok := r["dportgroup"].(string); ok {
			match["dPortGroup"] = dportgroup
		}

		if sportgroup, ok := r["sportgroup"].(string); ok {
			match["sPortGroup"] = sportgroup
		}

		// Action
		qos["type"] = r["serviceclass"].(string)
		rxScheduler["bandwidthCapPct"] = r["rxbandwidthpct"].(float64)
		txScheduler["bandwidthCapPct"] = r["txbandwidthpct"].(float64)

		rxScheduler["priority"] = r["priority"].(string)
		txScheduler["priority"] = r["priority"].(string)

		edge2CloudRouteAction["serviceGroup"] = r["linksteering"].(string)
		edge2DataCenterRouteAction["serviceGroup"] = r["linksteering"].(string)
		edge2EdgeRouteAction["serviceGroup"] = r["linksteering"].(string)

		if (r["networkservice"].(string)) == "auto" {
			edge2CloudRouteAction["routePolicy"] = "gateway"
			edge2DataCenterRouteAction["routePolicy"] = "auto"
			edge2EdgeRouteAction["routePolicy"] = "gateway"
		} else {
			edge2CloudRouteAction["routePolicy"] = "auto"
			edge2DataCenterRouteAction["routePolicy"] = "auto"
			edge2EdgeRouteAction["routePolicy"] = "auto"
		}

		if (r["linksteering"].(string)) == "ALL" {
			edge2CloudRouteAction["linkPolicy"] = "auto"
			edge2DataCenterRouteAction["linkPolicy"] = "auto"
			edge2EdgeRouteAction["linkPolicy"] = "auto"
		} else {
			edge2CloudRouteAction["linkPolicy"] = "fixed"
			edge2DataCenterRouteAction["linkPolicy"] = "fixed"
			edge2EdgeRouteAction["linkPolicy"] = "fixed"
		}

		new_rules[i] = new_rule
	}

	segment["rules"] = new_rules

	// Update QoS Configuration module
	velo.UpdateConfigurationModule(client, qos_module_id, data)

	d.SetId(fmt.Sprint(qos_module_id))
	resourceBusinessPoliciesRead(ctx, d, m)

	return diags
}

func resourceBusinessPoliciesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)
	rules, err := velo.GetQosRules(client, profile_id, segment_id)

	if err != nil {
		return diag.FromErr(err)
	}

	rs := make([]interface{}, len(rules))

	for i, rule := range rules {
		r := rule.(map[string]interface{})

		rid := make(map[string]interface{})

		match := r["match"].(map[string]interface{})
		action := r["action"].(map[string]interface{})
		qos := action["QoS"].(map[string]interface{})
		txScheduler := qos["txScheduler"].(map[string]interface{})
		rxScheduler := qos["txScheduler"].(map[string]interface{})
		edge2CloudRouteAction := action["edge2CloudRouteAction"].(map[string]interface{})
		edge2DataCenterRouteAction := action["edge2DataCenterRouteAction"].(map[string]interface{})
		edge2EdgeRouteAction := action["edge2EdgeRouteAction"].(map[string]interface{})

		rid["name"] = r["name"].(string)

		rid["appid"] = match["appid"].(float64)
		rid["hostname"] = match["hostname"].(string)
		rid["dip"] = match["dip"].(string)
		rid["dsm"] = match["dsm"].(string)
		rid["dport_low"] = match["dport_low"].(float64)
		rid["dport_high"] = match["dport_high"].(float64)
		rid["proto"] = match["proto"].(float64)

		if daddressgroup, ok := match["dAddressGroup"].(string); ok {
			rid["daddressgroup"] = daddressgroup
		}

		if saddressgroup, ok := match["sAddressGroup"].(string); ok {
			rid["saddressgroup"] = saddressgroup
		}

		if dportgroup, ok := match["dPortGroup"].(string); ok {
			rid["dportgroup"] = dportgroup
		}

		if sportgroup, ok := match["sPortGroup"].(string); ok {
			rid["sportgroup"] = sportgroup
		}

		rid["serviceclass"] = qos["type"].(string)
		rid["rxbandwidthpct"] = rxScheduler["bandwidthCapPct"].(float64)
		rid["txbandwidthpct"] = txScheduler["bandwidthCapPct"].(float64)
		rid["priority"] = rxScheduler["priority"].(string)
		rid["priority"] = txScheduler["priority"].(string)

		rid["linksteering"] = edge2CloudRouteAction["serviceGroup"].(string)
		rid["linksteering"] = edge2DataCenterRouteAction["serviceGroup"].(string)
		rid["linksteering"] = edge2EdgeRouteAction["serviceGroup"].(string)

		rid["networkservice"] = edge2CloudRouteAction["routePolicy"].(string)
		rid["networkservice"] = edge2DataCenterRouteAction["routePolicy"].(string)
		rid["networkservice"] = edge2EdgeRouteAction["routePolicy"].(string)

		rs[i] = rid
	}

	d.Set("rule", rs)

	return diags
}

func resourceBusinessPoliciesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceBusinessPoliciesCreate(ctx, d, m)

	return diags
}

func resourceBusinessPoliciesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(*velo.Client)
	profile_id := d.Get("profile").(int)
	segment_id := d.Get("segment").(int)

	qosmodule, err := velo.GetQosModule(client, profile_id)
	if err != nil {
		return diag.FromErr(err)
	}
	raw := qosmodule.(map[string]interface{})
	qos_module_id := int(raw["id"].(float64))
	data := raw["data"].(map[string]interface{})
	segments := data["segments"].([]interface{})
	segment := segments[segment_id].(map[string]interface{})

	segment["rules"] = make([]interface{}, 0)

	// Update QoS Configuration module
	_, err = velo.UpdateConfigurationModule(client, qos_module_id, data)

	if err != nil {
		d.SetId(fmt.Sprint(qos_module_id))
	}

	return diags
}
