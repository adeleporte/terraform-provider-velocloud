package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type EdgeDeviceSettingsData struct {
	Bgp              interface{} `json:"bgp"`
	Lan              interface{} `json:"lan"`
	RoutedInterfaces interface{} `json:"routedInterfaces"`
	Routes           interface{} `json:"routes"`
	Ha               interface{} `json:"ha"`
	Dns              interface{} `json:"dns"`
	Netflow          interface{} `json:"netflow"`
	Vqm              interface{} `json:"vqm"`
	Vrrp             interface{} `json:"vrrp"`
	Snmp             interface{} `json:"snmp"`
	MultiSourceQos   interface{} `json:"multiSourceQos"`
	Tacacs           interface{} `json:"tacacs"`
}

type ConfigurationDeviceSettingsModule struct {
	Name string                 `json:"name"`
	Data EdgeDeviceSettingsData `json:"data"`
}

type UpdateConfigurationDeviceSettingsModuleBody struct {
	ID     int                               `json:"id"`
	Update ConfigurationDeviceSettingsModule `json:"_update"`
}

type UpdateConfigurationDeviceSettingsModule_result struct {
	Error string `json:"error"`
	Rows  int    `json:"rows"`
}

type GetConfigurationDeviceSettingsModuleBody struct {
	ID           int      `json:"id"`
	EnterpriseID int      `json:"enterpriseId,omitempty"`
	With         []string `json:"with"`
}

// GetConfiguration ...
func GetDeviceSettingsModule(c *Client, enterpriseid int, profileid int) (map[string]interface{}, error) {

	var dd map[string]interface{}

	body := GetConfigurationDeviceSettingsModuleBody{
		ID:           profileid,
		EnterpriseID: enterpriseid,
		With:         []string{"modules"},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/getConfiguration", c.HostURL), buf)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return dd, err
	}

	// Unmarschal
	var resp map[string]interface{}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return resp, err
	}

	modules := resp["modules"].([]interface{})

	for _, v := range modules {
		item := v.(map[string]interface{})
		if item["name"] == "deviceSettings" {
			return item, nil
		}
	}

	return dd, errors.New("cannot find device settings module")

}

// UpdateDeviceSettingsModule ...
func UpdateDeviceSettingsModule(c *Client, enterpriseid int, devicemoduleid int, data map[string]interface{}) (UpdateConfigurationDeviceSettingsModule_result, error) {

	resp := UpdateConfigurationDeviceSettingsModule_result{}

	body := updateConfigurationModuleBody{
		ID:           devicemoduleid,
		EnterpriseID: enterpriseid,
		Update: updateConfigurationModule{
			Name: "deviceSettings",
			Data: data,
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/updateConfigurationModule", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	return resp, nil
}
