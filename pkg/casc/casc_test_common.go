package casc

import (
	"fmt"
	"github.com/jenkins-zh/jenkins-client/pkg/core"
	"net/http"

	"github.com/jenkins-zh/jenkins-client/pkg/mock/mhttp"
)

// PrepareForSASCReload only for test
func PrepareForSASCReload(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string) {
	request, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/configuration-as-code/reload", rootURL), nil)
	core.PrepareCommonPost(request, "", roundTripper, user, password, rootURL)
}

// PrepareForSASCApply only for test
func PrepareForSASCApply(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string) {
	request, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/configuration-as-code/apply", rootURL), nil)
	core.PrepareCommonPost(request, "", roundTripper, user, password, rootURL)
}

// PrepareForSASCExport only for test
func PrepareForSASCExport(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string) (
	response *http.Response) {
	request, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/configuration-as-code/export", rootURL), nil)
	response = core.PrepareCommonPost(request, "sample", roundTripper, user, password, rootURL)
	return
}

// PrepareForSASCExportWithCode only for test
func PrepareForSASCExportWithCode(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string, code int) {
	response := PrepareForSASCExport(roundTripper, rootURL, user, password)
	response.StatusCode = code
}

// PrepareForSASCSchema only for test
func PrepareForSASCSchema(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string) (
	response *http.Response) {
	request, _ := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/configuration-as-code/schema", rootURL), nil)
	response = core.PrepareCommonPost(request, "sample", roundTripper, user, password, rootURL)
	return
}

// PrepareForSASCSchemaWithCode only for test
func PrepareForSASCSchemaWithCode(roundTripper *mhttp.MockRoundTripper, rootURL, user, password string, code int) {
	response := PrepareForSASCSchema(roundTripper, rootURL, user, password)
	response.StatusCode = code
}
