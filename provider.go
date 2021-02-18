package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"vco": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_URL", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_TOKEN", nil),
			},
			"operator": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				DefaultFunc: schema.EnvDefaultFunc("VCO_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				DefaultFunc: schema.EnvDefaultFunc("VCO_PASSWORD", nil),
			},
			"skip_ssl_verification": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"velocloud_business_policies": resourceBusinessPolicies(),
			"velocloud_address_group":     resourceAddressGroup(),
			"velocloud_port_group":        resourcePortGroup(),
			"velocloud_edge":              resourceEdge(),
			"velocloud_firewall_rules":    resourceFirewallRules(),
			"velocloud_device_settings":   resourceDeviceSettings(),
			"velocloud_enterprise":        resourceEnterprise(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"velocloud_profile":       dataSourceProfile(),
			"velocloud_enterprise":    dataSourceEnterprise(),
			"velocloud_address_group": dataSourceAddressGroup(),
			"velocloud_port_group":    dataSourcePortGroup(),
			"velocloud_application":   dataSourceApplication(),
			"velocloud_edge":          dataSourceEdge(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	vco := d.Get("vco").(string)
	token := d.Get("token").(string)
	operator := d.Get("operator").(bool)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	ssl := d.Get("skip_ssl_verification").(bool)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (vco != "") && (token != "") && !operator {
		c, err := velo.NewTokenClient(&vco, &token, &ssl)
		//c := &http.Client{Timeout: 10 * time.Second}

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	if (vco != "") && (username != "") && (password != "") {
		c, err := velo.NewUsernamePasswordClient(&vco, &username, &password, &ssl, &operator)
		//c := &http.Client{Timeout: 10 * time.Second}

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	return nil, diag.Errorf("Missing credentials")
}
