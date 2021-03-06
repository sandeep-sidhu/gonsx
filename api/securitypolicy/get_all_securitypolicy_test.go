package securitypolicy

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var getAllSecurityPoliciesAPI *GetAllSecurityPoliciesAPI

func setupGetAll() {
	getAllSecurityPoliciesAPI = NewGetAll()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllSecurityPoliciesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/2.0/services/policy/securitypolicy/all", getAllSecurityPoliciesAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()

	xmlContent, err := ioutil.ReadFile("test_data/get_all_securitypolicy.xml")
	assert.Nil(t, err)

	xmlerr := xml.Unmarshal(xmlContent, getAllSecurityPoliciesAPI.ResponseObject())
	assert.Nil(t, xmlerr)

	assert.Len(t, getAllSecurityPoliciesAPI.GetResponse().SecurityPolicies, 8)
	assert.Equal(t, "OVP_sp_test2", getAllSecurityPoliciesAPI.GetResponse().SecurityPolicies[0].Name)
	assert.Equal(t, "OVP_firewall1", getAllSecurityPoliciesAPI.GetResponse().SecurityPolicies[0].ActionsByCategory.Actions[0].Name)
	assert.Equal(t, "allow", getAllSecurityPoliciesAPI.GetResponse().SecurityPolicies[0].ActionsByCategory.Actions[0].Action)
}
