package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/zcalusic/sysinfo"

	"skrepysh-agent/pkg/config"
)

type Client struct {
	IP   string
	OS   string
	conf *config.SkrepyshBackendConfig
}

func New(conf *config.SkrepyshBackendConfig) (*Client, error) {
	ip, err := getIPAddress()
	if err != nil {
		return nil, err
	}
	os := getOSName()
	return &Client{
		ip, os, conf,
	}, nil
}

func (c *Client) Init() error {
	initReq := initRequest{
		c.IP, c.OS,
	}
	data, err := json.Marshal(initReq)
	if err != nil {
		return err
	}

	u, err := url.JoinPath(fmt.Sprintf("http://%s:%d", c.conf.Host, c.conf.Port), "/init")
	if err != nil {
		return err
	}
	_, err = http.Post(u, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete() error {
	deleteReq := deleteRequest{
		c.IP,
	}
	data, err := json.Marshal(deleteReq)
	if err != nil {
		return err
	}

	u, err := url.JoinPath(fmt.Sprintf("http://%s:%d", c.conf.Host, c.conf.Port), "/delete")
	if err != nil {
		return err
	}
	_, err = http.Post(u, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}

func getIPAddress() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}
	return "", fmt.Errorf("unable to determine IP address")
}

func getOSName() string {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	return si.OS.Name
}

type initRequest struct {
	IP string `json:"ip"`
	OS string `json:"os"`
}

type deleteRequest struct {
	IP string `json:"ip"`
}
