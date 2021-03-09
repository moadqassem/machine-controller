package tinkerbell

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/provider/baremetal/plugins"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 10 * time.Second

type driver struct {
	tinkServerAddress string

	client   *http.Client
	username string
	password string
}

// NewTinkerbellDriver returns a new TinkerBell driver with a configured tinkserver address and a client timeout.
func NewTinkerbellDriver(tinkServerAddress, username, password string, timeout time.Duration) (*driver, error) {
	if tinkServerAddress == "" || username == "" || password == "" {
		return nil, errors.New("tink-server address, username or server cannot be empty")
	}

	if timeout == 0 {
		timeout = defaultTimeout
	}

	d := &driver{
		tinkServerAddress: tinkServerAddress,
		client: &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout: timeout,
			},
			Timeout: timeout,
		},
		username: username,
		password: password,
	}

	return d, nil
}
func (d *driver) GetServer(serverId string) (plugins.Server, error) {
	if serverId == "" {
		return nil, errors.New("server id cannot be empty")
	}

	requestUrl := url.URL{
		Host:   d.tinkServerAddress,
		Path:   fmt.Sprintf("/v1/hardware/%s", serverId),
		Scheme: "http",
		User:   url.UserPassword(d.username, d.password),
	}
	getRequest, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build get servers request: %v", err)
	}

	res, err := d.client.Do(getRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get server request: %v", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the get server response: %v", err)
	}

	server := &Hardware{}
	if err := json.Unmarshal(data, server); err != nil {
		return nil, fmt.Errorf("failed to unmrashal server data: %v", err)
	}

	return server, nil
}

func (d *driver) ListServers() ([]plugins.Server, error) {
	panic("implement me")
}

func (d *driver) ProvisionServer() (string, error) {
	panic("implement me")
}

func (d *driver) DeprovisionServer() error {
	panic("implement me")
}
