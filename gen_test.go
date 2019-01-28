package generate

import (
	"testing"
	//
	"github.com/jsonrouter/logging/testing"
	"github.com/jsonrouter/dashboard/gen"
	"github.com/jsonrouter/platforms/standard"
	"github.com/jsonrouter/core/openapi/v2"
)

func TestMain(t *testing.T) {

	Generate()

	router, err := jsonrouter.New(
		logs.NewClient().NewLogger(),
		openapiv2.New("localhost", "test"),
	)
	if err != nil {
		t.Fail()
	}

	staticContent := static.New()
	staticContent.Dashboard(router.Node)

}
