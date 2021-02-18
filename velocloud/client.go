package velocloud

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Cookies    []string
	SSL        bool
	Operator   bool
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
	ID           int                       `json:"id"`
	EnterpriseID int                       `json:"enterpriseId,omitempty"`
	Update       updateConfigurationModule `json:"_update"`
}

type enterprise_get_object_groups struct {
	Type string `json:"type"`
}

type enterprise_get_configurations struct {
	EnterpriseId int `json:"enterpriseId,omitempty"`
}

// NewTokenClient
func NewTokenClient(vco, token *string, ssl *bool) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		HostURL:    *vco,
		SSL:        *ssl,
	}

	if (vco != nil) && (token != nil) {
		c.HostURL = *vco
		c.Token = *token
	}

	return &c, nil
}

// NewTokenClient
func NewUsernamePasswordClient(vco, username *string, password *string, ssl *bool, operator *bool) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
		HostURL:    *vco,
		SSL:        *ssl,
	}

	c.HostURL = *vco
	c.Operator = *operator

	if (username != nil) && (password != nil) {

		// form request body
		rb, err := json.Marshal(AuthStruct{
			Username: *username,
			Password: *password,
		})
		if err != nil {
			return nil, err
		}

		// Debug
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println(strings.NewReader(string(rb)))

		// authenticate
		if c.SSL {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		}

		// Operator Login
		var req *http.Request
		if c.Operator {
			req, err = http.NewRequest("POST", fmt.Sprintf("%s/login/operatorLogin", c.HostURL), strings.NewReader(string(rb)))
			if err != nil {
				return nil, err
			}
		} else {
			// Enterprise Login
			req, err = http.NewRequest("POST", fmt.Sprintf("%s/login/enterpriseLogin", c.HostURL), strings.NewReader(string(rb)))
			if err != nil {
				return nil, err
			}
		}

		res, err := c.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		authenticated := false
		cookies := []string{}
		for _, c := range res.Cookies() {
			cookies = append(cookies, c.Raw)

			if c.Name == "velocloud.session" {
				authenticated = true
			}
		}

		if !authenticated {
			return &c, fmt.Errorf("Not Authenticated")
		}

		c.Cookies = cookies

		// Debug
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Println("*********************************")
		log.Printf("%v \n", cookies)

		return &c, nil

	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Token))
	}

	if c.SSL {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if len(c.Cookies) > 0 {
		for _, c := range c.Cookies {
			req.Header.Set("Cookie", c)
			log.Printf("cookie: %s\n", c)
		}
	}

	// Debug
	log.Println("REQUEST DUMP")
	requestDump, errdebug := httputil.DumpRequest(req, true)
	if errdebug != nil {
		log.Println(errdebug)
	}
	log.Println(string(requestDump))
	log.Println("###################################################")
	//

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Debug
	log.Println("REPONSE DUMP")
	responseDump, errdebug := httputil.DumpResponse(res, true)
	if errdebug != nil {
		log.Println(errdebug)
	}
	log.Println(string(responseDump))
	log.Println("###################################################")
	//

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
func GetProfile(c *Client, profilename string, enterpriseID int) (int, error) {

	body := enterprise_get_configurations{
		EnterpriseId: enterpriseID,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/enterprise/getEnterpriseConfigurations", c.HostURL), buf)

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
