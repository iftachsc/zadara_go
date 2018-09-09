package zadara

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Host string
	//User     string
	//Password string
	Token string
}

func NewClient(host string, user string, password string, token string) *Client {
	zc := new(Client)
	zc.Host = host
	zc.Token = token
	return zc
}

func (c *Client) invokeGet(path string) ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	url := fmt.Sprintf("http://%s/%s", c.Host, path)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Access-Key", c.Token)
	//req.Header.Add("Content-Type","application/json")
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *Client) GetVolumes() ([]Volume, error) {

	jsonBytes, err := c.invokeGet(VolumesPath) //This reads raw request body
	if err != nil {
		return nil, err
	}

	var resp RootVolumeResponse
	json.Unmarshal(jsonBytes, &resp)
	//println("volume:",zadara.Response.Volumes[0].Vol_name)
	return resp.Response.Volumes, nil
}

func (c *Client) GetMirrors() ([]Mirror, error) {

	jsonBytes, err := c.invokeGet(MirrorsPath) //This reads raw request body
	if err != nil {
		return nil, err
	}

	var resp RootMirrorResponse
	json.Unmarshal(jsonBytes, &resp)
	return resp.Response.Mirrors, nil
}
