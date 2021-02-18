package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/adeleporte/terraform-provider-velocloud/velocloud"
	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func resourceEnterprise() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnterpriseCreate,
		ReadContext:   resourceEnterpriseRead,
		UpdateContext: resourceEnterpriseUpdate,
		DeleteContext: resourceEnterpriseDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configurationid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
			"gatewaypoolid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"networkid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"enableenterprisedelegationtooperator": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enableenterpriseusermanagementdelegationtooperator": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"user": {
				Type:        schema.TypeList,
				Description: "User description",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password": {
							Type:     schema.TypeString,
							Required: true,
						},
						"password2": {
							Type:     schema.TypeString,
							Required: true,
						},
						"email": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceEnterpriseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	user_schema := d.Get("user").([]interface{})
	user_map := user_schema[0].(map[string]interface{})

	user_username, _ := user_map["username"].(string)
	user_password, _ := user_map["password"].(string)
	user_password2, _ := user_map["password2"].(string)
	user_email, _ := user_map["email"].(string)

	user := velocloud.User{
		Username:  user_username,
		Password:  user_password,
		Password2: user_password2,
		Email:     user_email,
	}

	Enterprise := velo.Enterprise_insert_enterprise{
		Name:                                 d.Get("name").(string),
		ConfigurationID:                      d.Get("configurationid").(int),
		NetworkID:                            d.Get("networkid").(int),
		GatewayPoolID:                        d.Get("gatewaypoolid").(int),
		EnableEnterpriseDelegationToOperator: d.Get("enableenterprisedelegationtooperator").(bool),
		EnableEnterpriseUserManagementDelegationToOperator: d.Get("enableenterpriseusermanagementdelegationtooperator").(bool),
		User: user,
	}

	resp, err := velo.InsertEnterprise(client, Enterprise)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", resp.ID))

	resourceEnterpriseRead(ctx, d, m)

	return diags
}

func resourceEnterpriseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	/*
		client := m.(*velo.Client)
		Enterprise_id, _ := strconv.Atoi(d.Id())
		enterprise_id := d.Get("enterpriseid").(int)

		Enterprise := velo.Enterprise_get_Enterprise{
			ID: Enterprise_id,
		}

		resp, err := velo.ReadEnterprise(client, Enterprise)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("activationkey", resp.ActivationKey)
		d.Set("activationstate", resp.ActivationState)
		d.Set("Enterprisestate", resp.EnterpriseState)
		d.Set("hastate", resp.HaState)
		d.Set("islive", resp.IsLive)
		d.Set("servicestate", resp.ServiceState)

		// Read Enterprise Specific Configuration Profile
		Enterprise_profile_id, err := velo.GetEnterpriseSpecificProfile(client, Enterprise_id, enterprise_id)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("Enterpriseprofileid", Enterprise_profile_id)
	*/
	return diags
}

func resourceEnterpriseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	/*
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

		Enterprise := velo.Enterprise_update_Enterprise{
			ID:           id,
			EnterpriseID: d.Get("enterpriseid").(int),
			Update: velocloud.Enterprise_update_Enterprise_data{
				Name:         d.Get("name").(string),
				Description:  d.Get("description").(string),
				SerialNumber: d.Get("serialnumber").(string),
				Site:         site,
			},
		}

		_, err := velo.UpdateEnterprise(client, Enterprise)
		if err != nil {
			return diag.FromErr(err)
		}

		resourceEnterpriseRead(ctx, d, m)
	*/
	return diags
}

func resourceEnterpriseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := m.(*velo.Client)
	id, _ := strconv.Atoi(d.Id())

	delete := velo.Enterprise_delete_enterprise{
		EnterpriseID: id,
	}

	_, err := velo.DeleteEnterprise(client, delete)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
