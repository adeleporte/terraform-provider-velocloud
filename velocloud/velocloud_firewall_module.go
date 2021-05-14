package velocloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type FirewallRuleMatch struct {
	AppID         int    `json:"appid"`
	ClassID       int    `json:"classid"`
	Dscp          int    `json:"dscp"`
	SIP           string `json:"sip"`
	SPortHigh     int    `json:"sport_high"`
	SPortLow      int    `json:"sport_low"`
	SAddressGroup string `json:"sAddressGroup"`
	SPortGroup    string `json:"sPortGroup"`
	SSM           string `json:"ssm"`
	SMac          string `json:"smac"`
	SVlan         int    `json:"svlan"`
	OSVersion     int    `json:"os_version"`
	Hostname      string `json:"hostname"`
	DIP           string `json:"dip"`
	DPortLow      int    `json:"dport_low"`
	DPortHigh     int    `json:"dport_high"`
	DAddressGroup string `json:"dAddressGroup"`
	DPortGroup    string `json:"dPortGroup"`
	DSM           string `json:"dsm"`
	DMac          string `json:"dmac"`
	DVlan         int    `json:"dvlan"`
	Proto         int    `json:"proto"`
	SRuleType     string `json:"s_rule_type"`
	DRuleType     string `json:"d_rule_type"`
}

type FirewallInboundNat struct {
	LanIP    string `json:"lan_ip"`
	LanPort  int    `json:"lan_port"`
	Outbound bool   `json:"outbound"`
}

type FirewallInboundAction struct {
	Type           string             `json:"type"`
	Interface      string             `json:"interface"`
	SubInterfaceID int                `json:"subinterfaceId"`
	Nat            FirewallInboundNat `json:"nat"`
	DRuleType      string             `json:"d_rule_type"`
}

type FirewallInboundRule struct {
	Name          string                `json:"name"`
	Match         FirewallRuleMatch     `json:"match"`
	Action        FirewallInboundAction `json:"action"`
	RuleLogicalId string                `json:"ruleLogicalId,omitempty"`
}

type FirewallOutboundAction struct {
	AllowOrDeny string `json:"allow_or_deny"`
}

type FirewallOutboundRule struct {
	Name          string                 `json:"name"`
	Match         FirewallRuleMatch      `json:"match"`
	Action        FirewallOutboundAction `json:"action"`
	RuleLogicalId string                 `json:"ruleLogicalId,omitempty"`
}

type ModuleSegmentMetaData struct {
	Name             string `json:"name,omitempty"`
	SegmentID        int    `json:"segmentId"`
	SegmentLogicalID string `json:"segmentLogicalId,omitempty"`
	Type             string `json:"type,omitempty"`
}

type FirewallSegment struct {
	FirewallLoggingEnabled  bool                   `json:"firewall_logging_enabled"`
	StatefulFirewallEnabled bool                   `json:"stateful_firewall_enabled,omitempty"`
	Outbound                []FirewallOutboundRule `json:"outbound"`
	Segment                 ModuleSegmentMetaData  `json:"segment"`
}

type FirewallSSH struct {
	Enabled         bool     `json:"enabled"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallLocalUI struct {
	Enabled         bool     `json:"enabled"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	PortNumber      int      `json:"portNumber,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallSNMP struct {
	Enabled         bool     `json:"enabled"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallICMP struct {
	Enabled         bool     `json:"enabled"`
	AllowSelectedIp []string `json:"allowSelectedIp,omitempty"`
	RuleLogicalId   string   `json:"ruleLogicalId,omitempty"`
}

type FirewallServices struct {
	LoggingEnabled bool            `json:"loggingEnabled"`
	SSH            FirewallSSH     `json:"ssh,omitempty"`
	LocalUI        FirewallLocalUI `json:"localUi,omitempty"`
	SNMP           FirewallSNMP    `json:"snmp,omitempty"`
	ICMP           FirewallICMP    `json:"icmp,omitempty"`
}

type FirewallData struct {
	FirewallEnabled         bool                  `json:"firewall_enabled"`
	InboundLoggingEnabled   bool                  `json:"inboundLoggingEnabled,omitempty"`
	StatefulFirewallEnabled bool                  `json:"stateful_firewall_enabled,omitempty"`
	FirewallLoggingEnabled  bool                  `json:"firewall_logging_enabled,omitempty"`
	SyslogForwarding        bool                  `json:"syslog_forwarding,omitempty"`
	Inbound                 []FirewallInboundRule `json:"inbound"`
	Segments                []FirewallSegment     `json:"segments"`
}

type ConfigurationFirewallModule struct {
	Name string       `json:"name"`
	Data FirewallData `json:"data"`
}

type UpdateConfigurationFirewallModuleBody struct {
	ID           int                         `json:"id"`
	EnterpriseID int                         `json:"enterpriseId,omitempty"`
	Update       ConfigurationFirewallModule `json:"_update"`
}

type InsertConfigurationFirewallModuleBody struct {
	ConfigurationID int          `json:"configurationId,omitempty"`
	EnterpriseID    int          `json:"enterpriseId,omitempty"`
	Name            string       `json:"name"`
	Data            FirewallData `json:"data"`
}

type UpdateConfigurationFirewallModule_result struct {
	Error string `json:"error"`
	Rows  int    `json:"rows"`
}

type InsertConfigurationFirewallModule_result struct {
	Error string `json:"error"`
	Rows  int    `json:"rows"`
	ID    string `json:"id"`
}

type GetConfigurationModulesBody struct {
	ConfigurationID int      `json:"configurationId"`
	EnterpriseID    int      `json:"enterpriseId,omitempty"`
	Modules         []string `json:"modules"`
}

type GetConfigurationModulesBody_result struct {
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	Type            string       `json:"type"`
	Description     string       `json:"description"`
	ConfigurationID int          `json:"configurationId"`
	Data            FirewallData `json:"data"`
	Refs            interface{}  `json:"refs"`
}

// GetConfiguration ...
func GetFirewallModule(c *Client, enterpriseid int, profileid int) (interface{}, error) {

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
		if module["name"] == "firewall" {
			return module, nil
		}
	}

	return nil, errors.New("cannot find firewall module")

}

// UpdateFirewallModule ...
func UpdateFirewallModule(c *Client, body UpdateConfigurationFirewallModuleBody) (UpdateConfigurationFirewallModule_result, error) {

	resp := UpdateConfigurationFirewallModule_result{}

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

// GetFirewallModules ...
func GetFirewallModules(c *Client, configurationId int, enterpriseId int) (FirewallData, error) {

	resp := []GetConfigurationModulesBody_result{}

	body := GetConfigurationModulesBody{
		ConfigurationID: configurationId,
		EnterpriseID:    enterpriseId,
		Modules:         []string{"firewall"},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/getConfigurationModules", c.HostURL), buf)
	if err != nil {
		fmt.Println(err)
		return FirewallData{}, err
	}

	// Send the request
	r, err := c.doRequest(req)
	if err != nil {
		fmt.Println(err)
		return FirewallData{}, err
	}

	// parse response body
	err = json.Unmarshal(r, &resp)
	if err != nil {
		fmt.Println(err)
		return FirewallData{}, err
	}

	fw := resp[0].Data

	return fw, nil
}

// InsertFirewallModule ...
func InsertFirewallModule(c *Client, body InsertConfigurationFirewallModuleBody) (InsertConfigurationFirewallModule_result, error) {

	resp := InsertConfigurationFirewallModule_result{}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/configuration/insertConfigurationModule", c.HostURL), buf)
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
