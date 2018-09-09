package zadara

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func initHttpMock(url string, method string, json string) {
	httpmock.Activate()

	//url e.g. "https://api.mybiz.com/articles.json"
	httpmock.RegisterResponder(method, url,
		httpmock.NewStringResponder(200, json))
}

const (
	VpsaHost     = "my.vpsa.com"
	VpsaUser     = "u"
	VpsaPassword = "p"
	VpsaToken    = "token"
)

func getClientAndApiUrl(apiPath string) (*Client, string) {
	return NewClient(VpsaHost, VpsaUser, VpsaPassword, VpsaToken), fmt.Sprintf("http://%s/%s", VpsaHost, apiPath)
}

//API Tests
func TestGetVolumesOK(t *testing.T) {

	c, url := getClientAndApiUrl(VolumesPath)

	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewStringResponder(200, RootVolumeResponseJson))

	defer httpmock.DeactivateAndReset()

	vols, err := c.GetVolumes()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 2, len(vols), "volume count should be equal")

	assert.Equal(t, mockedVol1.Name, vols[0].Name, "volume name should be equal")
	assert.Equal(t, mockedVol2.Name, vols[1].Name, "volume name should be equal")

	assert.Equal(t, mockedVol1.PoolName, vols[0].PoolName, "pool name should be equal")
	assert.Equal(t, mockedVol2.PoolName, vols[1].PoolName, "pool name should be equal")

	assert.Equal(t, mockedVol1.Encryption, vols[0].Encryption, "vol encryption should be equal")
	assert.Equal(t, mockedVol1.VirtualCapacity, vols[0].VirtualCapacity, "vol virtual capacity should be equal")
}

func TestGetMirrorJobsOk(t *testing.T) {
	c, url := getClientAndApiUrl(MirrorsPath)

	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodGet, url,
		httpmock.NewStringResponder(200, RootMirrorResponseJson))

	defer httpmock.DeactivateAndReset()

	mirrors, err := c.GetMirrors()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, 1, len(mirrors), "volume count should be equal")

	assert.Equal(t, mockedMirror1.Name, mirrors[0].Name, "mirror name should be equal")
	assert.Equal(t, mockedMirror1.Source.VpsaName, mirrors[0].Source.VpsaName, "source vpsa name should be equal")
	assert.Equal(t, mockedMirror1.Source.Provider, mirrors[0].Source.Provider, "source provider name should be equal")

	assert.Equal(t, mockedMirror1.Destination.VpsaName, mirrors[0].Destination.VpsaName, "dest vpsa name should be equal")
	assert.Equal(t, mockedMirror1.Destination.Provider, mirrors[0].Destination.Provider, "dest provider name should be equal")

	//assert.Equal(t, mockedMirror1.Created, mirrors[0].Created.Day(), "Mirror's day created must be equal")
	//assert.Equal(t, time.Time(mockedMirror1.Created).Month().String(), time.Time(mirrors[0].Created).Month().String(), "Mirror's month created must be equal")
	//assert.Equal(t, mockedMirror1.Created.Year(), mirrors[0].Created.Year(), "Mirror's year created must be equal")
}
