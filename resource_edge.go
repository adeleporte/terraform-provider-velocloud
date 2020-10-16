package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/adeleporte/terraform-provider-velocloud/velocloud"
	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

var EdgeModels = []string{"edge500", "edge5X0", "edge510", "edge510lte", "edge6X0", "edge840", "edge1000", "edge1000qat", "edge3X00", "edge3X10", "virtual"}

func resourceEdge() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEdgeCreate,
		ReadContext:   resourceEdgeRead,
		UpdateContext: resourceEdgeUpdate,
		DeleteContext: resourceEdgeDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"configurationid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"enterpriseid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"modelnumber": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(EdgeModels, false),
			},
			"serialnumber": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"haenabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"generatecertificate": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"subjectcn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"subjecto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"subjectou": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"challengepassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"privatekeypassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"custominfo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"site": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Site description",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"contactname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"contactphone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"contactmobile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"contactemail": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"streetaddress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"streetaddress2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"city": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"state": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"country": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"postalcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"lat": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"lon": &schema.Schema{
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"timezone": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"locale": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingsameaslocation": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"shippingcontactname": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingstreetaddress": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingstreetaddress2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingcity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingcountry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"shippingpostalcode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"edgeprofileid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceEdgeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	site_schema := d.Get("site").([]interface{})
	site_map := site_schema[0].(map[string]interface{})

	site_name, _ := site_map["name"].(string)

	site_contactname, _ := site_map["contactname"].(string)
	site_contactphone, _ := site_map["contactphone"].(string)
	site_contactmobile, _ := site_map["contactmobile"].(string)
	site_contactemail, _ := site_map["contactemail"].(string)
	site_streetaddress, _ := site_map["streetaddress"].(string)
	site_streetaddress2, _ := site_map["streetaddress2"].(string)
	site_city, _ := site_map["city"].(string)
	site_state, _ := site_map["state"].(string)
	site_country, _ := site_map["country"].(string)
	site_postalcode, _ := site_map["postalcode"].(string)

	site_shippingcontactname, _ := site_map["shippingcontactname"].(string)
	site_shippingstreetaddress, _ := site_map["shippingstreetaddress"].(string)
	site_shippingstreetaddress2, _ := site_map["shippingstreetaddress2"].(string)
	site_shippingcity, _ := site_map["shippingcontactcity"].(string)
	site_shippingcountry, _ := site_map["shippingcontactcountry"].(string)
	site_shippingpostalcode, _ := site_map["shippingcontactpostalcode"].(string)

	site_lat, _ := site_map["lat"].(float64)
	site_lon, _ := site_map["lon"].(float64)

	site_timezone, _ := site_map["timezone"].(string)
	site_locale, _ := site_map["locale"].(string)

	site_shippingsameaslocation, _ := site_map["shippingsameaslocation"].(bool)

	site := velocloud.Site{
		Name:                   site_name,
		ContactName:            site_contactname,
		ContactPhone:           site_contactphone,
		ContactMobile:          site_contactmobile,
		ContactEmail:           site_contactemail,
		StreetAddress:          site_streetaddress,
		StreetAddress2:         site_streetaddress2,
		City:                   site_city,
		State:                  site_state,
		Country:                site_country,
		PostalCode:             site_postalcode,
		ShippingContactName:    site_shippingcontactname,
		ShippingAddress:        site_shippingstreetaddress,
		ShippingAddress2:       site_shippingstreetaddress2,
		ShippingCity:           site_shippingcity,
		ShippingCountry:        site_shippingcountry,
		ShippingPostalCode:     site_shippingpostalcode,
		Lat:                    site_lat,
		Lon:                    site_lon,
		Timezone:               site_timezone,
		Locale:                 site_locale,
		ShippingSameAsLocation: site_shippingsameaslocation,
	}

	edge := velo.Enterprise_provision_edge{
		Name:                d.Get("name").(string),
		EnterpriseID:        d.Get("enterpriseid").(int),
		ConfigurationID:     d.Get("configurationid").(int),
		ModelNumber:         d.Get("modelnumber").(string),
		SerialNumber:        d.Get("serialnumber").(string),
		Description:         d.Get("description").(string),
		HaEnabled:           d.Get("haenabled").(bool),
		GenerateCertificate: d.Get("generatecertificate").(bool),
		SubjectCN:           d.Get("subjectcn").(string),
		SubjectO:            d.Get("subjecto").(string),
		SubjectOU:           d.Get("subjectou").(string),
		ChallengePassword:   d.Get("challengepassword").(string),
		PrivateKeyPassword:  d.Get("privatekeypassword").(string),
		CustomInfo:          d.Get("custominfo").(string),
		Site:                site,
	}

	resp, err := velo.InsertEdge(client, edge)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))
	d.Set("activationkey", resp.ActivationKey)

	resourceEdgeRead(ctx, d, m)

	return diags
}

func resourceEdgeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	edge_id, _ := strconv.Atoi(d.Id())
	enterprise_id := d.Get("enterpriseid").(int)

	edge := velo.Enterprise_get_edge{
		ID: edge_id,
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

	// Read Edge Specific Configuration Profile
	edge_profile_id, err := velo.GetEdgeSpecificProfile(client, edge_id, enterprise_id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("edgeprofileid", edge_profile_id)

	return diags
}

func resourceEdgeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	site_schema := d.Get("site").([]interface{})
	site_map := site_schema[0].(map[string]interface{})

	site_name, _ := site_map["name"].(string)

	site_contactname, _ := site_map["contactname"].(string)
	site_contactphone, _ := site_map["contactphone"].(string)
	site_contactmobile, _ := site_map["contactmobile"].(string)
	site_contactemail, _ := site_map["contactemail"].(string)
	site_streetaddress, _ := site_map["streetaddress"].(string)
	site_streetaddress2, _ := site_map["streetaddress2"].(string)
	site_city, _ := site_map["city"].(string)
	site_state, _ := site_map["state"].(string)
	site_country, _ := site_map["country"].(string)
	site_postalcode, _ := site_map["postalcode"].(string)

	site_shippingcontactname, _ := site_map["shippingcontactname"].(string)
	site_shippingstreetaddress, _ := site_map["shippingstreetaddress"].(string)
	site_shippingstreetaddress2, _ := site_map["shippingstreetaddress2"].(string)
	site_shippingcity, _ := site_map["shippingcontactcity"].(string)
	site_shippingcountry, _ := site_map["shippingcontactcountry"].(string)
	site_shippingpostalcode, _ := site_map["shippingcontactpostalcode"].(string)

	site_lat, _ := site_map["lat"].(float64)
	site_lon, _ := site_map["lon"].(float64)

	site_timezone, _ := site_map["timezone"].(string)
	site_locale, _ := site_map["locale"].(string)

	site_shippingsameaslocation, _ := site_map["shippingsameaslocation"].(bool)

	site := velocloud.Site{
		Name:                   site_name,
		ContactName:            site_contactname,
		ContactPhone:           site_contactphone,
		ContactMobile:          site_contactmobile,
		ContactEmail:           site_contactemail,
		StreetAddress:          site_streetaddress,
		StreetAddress2:         site_streetaddress2,
		City:                   site_city,
		State:                  site_state,
		Country:                site_country,
		PostalCode:             site_postalcode,
		ShippingContactName:    site_shippingcontactname,
		ShippingAddress:        site_shippingstreetaddress,
		ShippingAddress2:       site_shippingstreetaddress2,
		ShippingCity:           site_shippingcity,
		ShippingCountry:        site_shippingcountry,
		ShippingPostalCode:     site_shippingpostalcode,
		Lat:                    site_lat,
		Lon:                    site_lon,
		Timezone:               site_timezone,
		Locale:                 site_locale,
		ShippingSameAsLocation: site_shippingsameaslocation,
	}

	edge := velo.Enterprise_update_edge{
		ID:           id,
		EnterpriseID: d.Get("enterpriseid").(int),
		Update: velocloud.Enterprise_update_edge_data{
			Name:         d.Get("name").(string),
			Description:  d.Get("description").(string),
			SerialNumber: d.Get("serialnumber").(string),
			Site:         site,
		},
	}

	_, err := velo.UpdateEdge(client, edge)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceEdgeRead(ctx, d, m)

	return diags
}

func resourceEdgeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	delete := velo.Edge_delete_edge{
		ID:           id,
		EnterpriseID: d.Get("enterpriseid").(int),
	}

	_, err := velo.DeleteEdge(client, delete)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
