package velocloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Address_group_data struct {
	IP       string `json:"ip"`
	Mask     string `json:"mask"`
	RuleType string `json:"rule_type"`
}

type Enterprise_get_address_group struct {
	EnterpriseID int    `json:"enterpriseId,omitempty"`
	ID           int    `json:"id"`
	Type         string `json:"type"`
}

type Enterprise_get_address_group_result struct {
	ID        int                  `json:"id"`
	Name      string               `json:"name"`
	Type      string               `json:"type"`
	LogicalID string               `json:"logicalId"`
	Data      []Address_group_data `json:"data"`
}

type Enterprise_insert_address_group struct {
	EnterpriseID int                  `json:"enterpriseId,omitempty"`
	Name         string               `json:"name"`
	Type         string               `json:"type"`
	Description  string               `json:"description"`
	Data         []Address_group_data `json:"data"`
}

type Enterprise_insert_address_group_result struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

type Enterprise_update_address_group struct {
	ID           int                  `json:"id"`
	EnterpriseID int                  `json:"enterpriseId,omitempty"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	Data         []Address_group_data `json:"data"`
}

type Enterprise_update_address_group_result struct {
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

type Enterprise_delete_address_group struct {
	EnterpriseID int `json:"enterpriseId,omitempty"`
	ID           int `json:"id"`
}

type Enterprise_delete_address_group_result struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

// InsertAddressGroup ...
func InsertAddressGroup(c *Client, body Enterprise_insert_address_group) (Enterprise_insert_address_group_result, error) {

	resp := Enterprise_insert_address_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/insertObjectGroup", c.HostURL), buf)
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

// GetAddressGroup ...
func GetAddressGroup(c *Client, body Enterprise_get_address_group) ([]Enterprise_get_address_group_result, error) {

	resp := []Enterprise_get_address_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getObjectGroups", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_address_group_result{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_address_group_result{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_address_group_result{}, err
	}

	return resp, nil
}

// UpdateAddressGroup ...
func UpdateAddressGroup(c *Client, body Enterprise_update_address_group) (Enterprise_update_address_group_result, error) {

	resp := Enterprise_update_address_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/updateObjectGroup", c.HostURL), buf)
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

// DeleteAddressGroup ...
func DeleteAddressGroup(c *Client, body Enterprise_delete_address_group) (Enterprise_delete_address_group_result, error) {

	resp := Enterprise_delete_address_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/deleteObjectGroup", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_address_group_result{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_address_group_result{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_address_group_result{}, err
	}

	return resp, nil
}
