package velocloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Port_group_data struct {
	Proto    int `json:"proto"`
	PortLow  int `json:"port_low"`
	PortHigh int `json:"port_high,omitempty"`
}

type Enterprise_get_port_group struct {
	EnterpriseID int    `json:"enterpriseId,omitempty"`
	ID           int    `json:"id"`
	Type         string `json:"type"`
}

type Enterprise_get_port_group_result struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	LogicalID string            `json:"logicalId"`
	Data      []Port_group_data `json:"data"`
}

type Enterprise_insert_port_group struct {
	EnterpriseID int               `json:"enterpriseId,omitempty"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Description  string            `json:"description"`
	Data         []Port_group_data `json:"data"`
}

type Enterprise_insert_port_group_result struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

type Enterprise_update_port_group struct {
	ID           int               `json:"id"`
	EnterpriseID int               `json:"enterpriseId,omitempty"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Data         []Port_group_data `json:"data"`
}

type Enterprise_update_port_group_result struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

type Enterprise_delete_port_group struct {
	EnterpriseID int    `json:"enterpriseId,omitempty"`
	ID           int    `json:"id"`
	Type         string `json:"type"`
}

type Enterprise_delete_port_group_result struct {
	ID    int    `json:"id"`
	Rows  int    `json:"rows"`
	Error string `json:"error"`
}

// InsertPortGroup ...
func InsertPortGroup(c *Client, body Enterprise_insert_port_group) (Enterprise_insert_port_group_result, error) {

	resp := Enterprise_insert_port_group_result{}

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

// GetPortGroup ...
func GetPortGroup(c *Client, body Enterprise_get_port_group) ([]Enterprise_get_port_group_result, error) {

	resp := []Enterprise_get_port_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getObjectGroups", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_port_group_result{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_port_group_result{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return []Enterprise_get_port_group_result{}, err
	}

	return resp, nil
}

// UpdatePortGroup ...
func UpdatePortGroup(c *Client, body Enterprise_update_port_group) (Enterprise_update_port_group_result, error) {

	resp := Enterprise_update_port_group_result{}

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

// DeletePortGroup ...
func DeletePortGroup(c *Client, body Enterprise_delete_port_group) (Enterprise_delete_port_group_result, error) {

	resp := Enterprise_delete_port_group_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/deleteObjectGroup", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_port_group_result{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_port_group_result{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return Enterprise_delete_port_group_result{}, err
	}

	return resp, nil
}
