package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetConfiguration ...
func GetQosModule(c *Client, enterpriseid int, profileid int) (interface{}, error) {

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
		return nil, err
	}

	// Unmarschal
	var list map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return nil, err
	}

	// Find modules
	modules, _ := list["modules"].([]interface{})

	for _, v := range modules {
		module := v.(map[string]interface{})
		if module["name"] == "QOS" {
			return module, nil
		}
	}

	return nil, errors.New("cannot find qos module")

}

// GetDefaultRule ...
func GetQosRules(c *Client, enterpriseid int, profile_id int, segment_id int) ([]interface{}, error) {

	qosmodule, err := GetQosModule(c, enterpriseid, profile_id)

	if err != nil {
		return nil, err
	}

	raw := qosmodule.(map[string]interface{})
	data := raw["data"].(map[string]interface{})
	segments := data["segments"].([]interface{})
	segment := segments[segment_id].(map[string]interface{})
	//defaults := segment["defaults"].([]interface{})
	rules := segment["rules"].([]interface{})

	return rules, nil

}

// GetDefaultQosRules ...
func GetDefaultQosRules(c *Client, enterpriseid int, profile_id int, segment_id int) ([]interface{}, error) {

	qosmodule, err := GetQosModule(c, enterpriseid, profile_id)

	if err != nil {
		return nil, err
	}

	raw := qosmodule.(map[string]interface{})
	data := raw["data"].(map[string]interface{})
	segments := data["segments"].([]interface{})
	segment := segments[segment_id].(map[string]interface{})
	defaults := segment["defaults"].([]interface{})
	//rules := segment["rules"].([]interface{})

	return defaults, nil

}

// UpdateConfigurationModule ...
func UpdateConfigurationModule(c *Client, enterpriseid int, qos_module_id int, data map[string]interface{}) (interface{}, error) {

	body := updateConfigurationModuleBody{
		ID:           qos_module_id,
		EnterpriseID: enterpriseid,
		Update: updateConfigurationModule{
			Name: "QOS",
			Data: data,
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/updateConfigurationModule", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Send the request
	_, err = c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil

}
