package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Site struct {
	ID                     int     `json:"id"`
	Created                string  `json:"created"`
	Name                   string  `json:"name"`
	ContactName            string  `json:"contactName"`
	ContactPhone           string  `json:"contactPhone"`
	ContactMobile          string  `json:"contactMobile"`
	ContactEmail           string  `json:"contactEmail"`
	StreetAddress          string  `json:"streetAddress"`
	StreetAddress2         string  `json:"streetAddress2"`
	City                   string  `json:"city"`
	State                  string  `json:"state"`
	PostalCode             string  `json:"postalCode"`
	Country                string  `json:"country"`
	Lat                    float64 `json:"lat"`
	Lon                    float64 `json:"lon"`
	Timezone               string  `json:"timezone"`
	Locale                 string  `json:"locale"`
	ShippingContactName    string  `json:"shippingContactName"`
	ShippingAddress        string  `json:"shippingAddress"`
	ShippingAddress2       string  `json:"shippingAddress2"`
	ShippingCity           string  `json:"shippingCity"`
	ShippingCountry        string  `json:"shippingCountry"`
	ShippingPostalCode     string  `json:"shippingPostalCode"`
	Modified               string  `json:"modified"`
	ShippingSameAsLocation bool    `json:"shippingSameAsLocation"`
}

type Enterprise_provision_edge struct {
	EnterpriseID        int    `json:"enterpriseId,omitempty"`
	ConfigurationID     int    `json:"configurationId"`
	Name                string `json:"name"`
	SerialNumber        string `json:"serialNumber"`
	ModelNumber         string `json:"modelNumber"`
	Description         string `json:"description"`
	HaEnabled           bool   `json:"haEnabled"`
	GenerateCertificate bool   `json:"generateCertificate"`
	SubjectCN           string `json:"subjectCN"`
	SubjectO            string `json:"subjectO"`
	SubjectOU           string `json:"subjectOU"`
	ChallengePassword   string `json:"challengePassword"`
	PrivateKeyPassword  string `json:"privateKeyPassword"`
	CustomInfo          string `json:"customInfo"`
	Site                Site   `json:"site"`
}

type Enterprise_provision_edge_result struct {
	ID            int    `json:"id"`
	ActivationKey string `json:"activationKey"`
}

type Enterprise_get_edge struct {
	EnterpriseID int      `json:"enterpriseId,omitempty"`
	ID           int      `json:"id"`
	With         []string `json:"with"`
}

type Enterprise_get_edge_result struct {
	ID              int    `json:"id"`
	ActivationKey   string `json:"activationKey"`
	ActivationState string `json:"activationState"`
	EdgeState       string `json:"edgeState"`
	HaState         string `json:"haState"`
	IsLive          int    `json:"isLive"`
	ServiceState    string `json:"serviceState"`
	Site            Site   `json:"site"`
}

type Enterprise_update_edge_data struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
	Description  string `json:"description"`
	CustomInfo   string `json:"customInfo"`
	Site         Site   `json:"site"`
}

type Enterprise_update_edge struct {
	EnterpriseID int                         `json:"enterpriseId,omitempty"`
	ID           int                         `json:"id"`
	Update       Enterprise_update_edge_data `json:"_update"`
}

type Enterprise_update_edge_result struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
	Rows  int    `json:"rows"`
}

type Edge_delete_edge struct {
	EnterpriseID int `json:"enterpriseId,omitempty"`
	ID           int `json:"id"`
}

type Edge_delete_edge_result struct {
	ID    int    `json:"id"`
	Error string `json:"error"`
	Rows  int    `json:"rows"`
}

type Edge_get_edge_configuration_stack struct {
	EdgeID       int `json:"edgeId"`
	EnterpriseID int `json:"enterpriseId,omitempty"`
}

type Edge_get_edge_configuration_stack_result struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Modules []interface{} `json:"modules"`
}

// InsertEdge ...
func InsertEdge(c *Client, body Enterprise_provision_edge) (Enterprise_provision_edge_result, error) {

	resp := Enterprise_provision_edge_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge/edgeProvision", c.HostURL), buf)
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

// GetEdges ...
func GetEdges(c *Client, edge_name string) (int, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getEnterpriseEdges", c.HostURL), nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return -1, err
	}

	// Unmarschal
	var results []interface{}
	err = json.Unmarshal(res, &results)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return -1, err
	}

	for _, v := range results {
		edge := v.(map[string]interface{})
		if edge["name"].(string) == edge_name {

			return int(edge["id"].(float64)), nil
		}
	}

	return -1, errors.New("cant find edge")

}

// UpdateEdge ...
func ReadEdge(c *Client, body Enterprise_get_edge) (Enterprise_get_edge_result, error) {

	resp := Enterprise_get_edge_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge/getEdge", c.HostURL), buf)
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

// UpdateEdge ...
func UpdateEdge(c *Client, body Enterprise_update_edge) (Enterprise_update_edge_result, error) {

	resp := Enterprise_update_edge_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge/updateEdgeAttributes", c.HostURL), buf)
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

// DeleteEdge ...
func DeleteEdge(c *Client, body Edge_delete_edge) (Edge_delete_edge_result, error) {

	resp := Edge_delete_edge_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge/deleteEdge", c.HostURL), buf)
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

// GetEdges ...
func GetEdgeSpecificProfile(c *Client, edge_id int, enterprise_id int) (int, error) {

	resp := []Edge_get_edge_configuration_stack_result{}

	body := Edge_get_edge_configuration_stack{
		EdgeID:       edge_id,
		EnterpriseID: enterprise_id,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/edge/getEdgeConfigurationStack", c.HostURL), buf)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return -1, err
	}

	// Unmarschal
	err = json.Unmarshal(res, &resp)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return -1, err
	}

	for _, v := range resp {
		if v.Name == "Edge Specific Profile" {
			return v.ID, nil
		}
	}

	return -1, errors.New("cant find edge specific profile")

}
