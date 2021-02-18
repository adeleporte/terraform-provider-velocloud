package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// AuthStruct -
type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type Enterprise_insert_enterprise struct {
	ConfigurationID                                    int    `json:"configurationId"`
	EnableEnterpriseDelegationToOperator               bool   `json:"enableEnterpriseDelegationToOperator"`
	EnableEnterpriseUserManagementDelegationToOperator bool   `json:"enableEnterpriseUserManagementDelegationToOperator"`
	GatewayPoolID                                      int    `json:"gatewayPoolId"`
	NetworkID                                          int    `json:"networkId"`
	Name                                               string `json:"name"`
	User                                               User   `json:"user"`
}

type Enterprise_insert_enterprise_result struct {
	ID   int `json:"id"`
	Rows int `json:"rows"`
}

type Enterprise_delete_enterprise struct {
	EnterpriseID int `json:"enterpriseId"`
}

type Enterprise_delete_enterprise_result struct {
	ID   int `json:"id"`
	Rows int `json:"rows"`
}

// GetEnterprise ...
func GetEnterprise(c *Client, enterprisename string) (int, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/network/getNetworkEnterprises", c.HostURL), nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return 0, err
	}

	// Unmarschal
	var list []map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return 0, err
	}

	for _, v := range list {
		if v["name"] == enterprisename {

			return int(v["id"].(float64)), nil
		}
	}

	return 0, errors.New("cant find enterprise")

}

// InsertEnterprise ...
func InsertEnterprise(c *Client, body Enterprise_insert_enterprise) (Enterprise_insert_enterprise_result, error) {

	resp := Enterprise_insert_enterprise_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/insertEnterprise", c.HostURL), buf)
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

// DeleteEnterprise ...
func DeleteEnterprise(c *Client, body Enterprise_delete_enterprise) (Enterprise_delete_enterprise_result, error) {

	resp := Enterprise_delete_enterprise_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/deleteEnterprise", c.HostURL), buf)
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
