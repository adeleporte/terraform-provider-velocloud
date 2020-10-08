package velocloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetEdges ...
func GetEdges(c *Client, edge_name string) (float64, error) {

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

			return edge["id"].(float64), nil
		}
	}

	return -1, errors.New("cant find edge")

}
