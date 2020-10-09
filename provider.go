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
			"vco": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_URL", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("VCO_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"velocloud_business_policies": resourceBusinessPolicies(),
			"velocloud_address_group":     resourceAddressGroup(),
			"velocloud_port_group":        resourcePortGroup(),
			"velocloud_edge":              resourceEdge(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"velocloud_profile":       dataSourceProfile(),
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (vco != "") && (token != "") {
		c, err := velo.NewClient(&vco, &token)
		//c := &http.Client{Timeout: 10 * time.Second}

		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	return nil, diag.Errorf("Missing credentials")
}
