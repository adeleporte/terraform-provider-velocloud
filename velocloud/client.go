package velocloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type updateConfigurationModule struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}

type updateConfigurationModuleBody struct {
	ID     int                       `json:"id"`
	Update updateConfigurationModule `json:"_update"`
}

type enterprise_get_object_groups struct {
	Type string `json:"type"`
}

// NewClient -
func NewClient(vco, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		// Default Hashicups URL
		HostURL: *vco,
	}

	if (vco != nil) && (token != nil) {
		c.HostURL = *vco
		c.Token = *token
	}

	if (vco != nil) && (token != nil) {
		/*
				// form request body
				rb, err := json.Marshal(AuthStruct{
					Username: *username,
					Password: *password,
				})
				if err != nil {
					return nil, err
				}

				// authenticate
				req, err := http.NewRequest("POST", fmt.Sprintf("%s/signin", c.HostURL), strings.NewReader(string(rb)))
				if err != nil {
					return nil, err
				}

				body, err := c.doRequest(req)

				// parse response body
				ar := AuthResponse{}
				err = json.Unmarshal(body, &ar)
				if err != nil {
					return nil, err
				}

			c.Token = ar.Token
		*/
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Token))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// DeepCopy ...
func DeepCopy(src map[string]interface{}) (map[string]interface{}, error) {

	var dst map[string]interface{}

	if src == nil {
		fmt.Println("Error src")
		return nil, fmt.Errorf("src cannot be nil")
	}

	bytes, err := json.Marshal(src)

	if err != nil {
		return nil, fmt.Errorf("Unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, &dst)

	if err != nil {
		fmt.Println("Error unmarshal")
		fmt.Println(err)
		return nil, fmt.Errorf("Unable to unmarshal into dst: %s", err)
	}
	return dst, nil

}

// GetProfile ...
func GetProfile(c *Client, profilename string) (int, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getEnterpriseConfigurationsPolicies", c.HostURL), nil)

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
		if v["name"] == profilename {

			return int(v["id"].(float64)), nil
		}
	}

	return 0, errors.New("cant find profile")

}

/*
// GetAddressGroups ...
func GetAddressGroups(c *Client, groupname string) (string, error) {

	body := enterprise_get_object_groups{
		Type: "address_group",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getObjectGroups", c.HostURL), buf)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return "", err
	}

	// Unmarschal
	var list []map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return "", err
	}

	for _, v := range list {
		if v["name"] == groupname {

			return v["logicalId"].(string), nil
		}
	}

	return "", errors.New("cant find groupname")

}

// GetAddressGroups ...
func GetPortGroups(c *Client, groupname string) (string, error) {

	body := enterprise_get_object_groups{
		Type: "port_group",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getObjectGroups", c.HostURL), buf)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Send the request
	res, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err.Error())
		//return ConfigurationResults{}
		return "", err
	}

	// Unmarschal
	var list []map[string]interface{}
	err = json.Unmarshal(res, &list)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return "", err
	}

	for _, v := range list {
		if v["name"] == groupname {

			return v["logicalId"].(string), nil
		}
	}

	return "", errors.New("cant find groupname")

}
*/
// GetApplications ...
func GetApplications(c *Client, application_name string) (float64, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/getIdentifiableApplications", c.HostURL), nil)

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
	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println("Error with unmarshal")
		fmt.Println(err.Error())
		return -1, err
	}

	applications := result["applications"].([]interface{})
	for _, v := range applications {
		application := v.(map[string]interface{})
		if application["displayName"].(string) == application_name {

			return application["id"].(float64), nil
		}
	}

	return -1, errors.New("cant find application")

}
