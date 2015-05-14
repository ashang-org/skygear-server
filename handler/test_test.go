package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/oursky/ourd/router"
)

type SingleRouteRouter router.Router

func newSingleRouteRouter(handler router.Handler, prepareFunc func(*router.Payload)) *SingleRouteRouter {
	r := router.NewRouter()
	r.Map("", handler, func(p *router.Payload, _ *router.Response) int {
		prepareFunc(p)
		return 200
	})
	return (*SingleRouteRouter)(r)
}

func (r *SingleRouteRouter) POST(body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "", strings.NewReader(body))
	resp := httptest.NewRecorder()

	(*router.Router)(r).ServeHTTP(resp, req)
	return resp
}

// shouldEqualJSON asserts eqaulity of two JSON bytes or strings by
// their key / value, regardless of the actual position of those
// key-value pairs
func shouldEqualJSON(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return fmt.Sprintf("shouldEqualJSON receives only one expected argument")
	}

	actualBytes, err := interfaceToByteSlice(actual)
	if err != nil {
		return fmt.Sprintf("%[1]v is %[1]T, not []byte or string", actual)
	}

	expectedBytes, err := interfaceToByteSlice(expected[0])
	if err != nil {
		return fmt.Sprintf("%[1]v is %[1]T, not []byte or string", expected[0])
	}

	actualMap, expectedMap := map[string]interface{}{}, map[string]interface{}{}

	if err := json.Unmarshal(actualBytes, &actualMap); err != nil {
		return fmt.Sprintf("invalid JSON of L.H.S.: %v; actual = \n%v", err, actual)
	}

	if err := json.Unmarshal(expectedBytes, &expectedMap); err != nil {
		return fmt.Sprintf("invalid JSON of R.H.S.: %v; expected = \n%v", err, expected[0])
	}

	if !reflect.DeepEqual(actualMap, expectedMap) {
		return fmt.Sprintf(`Expected: '%s'
Actual:   '%s'`, prettyPrintJSONMap(expectedMap), prettyPrintJSONMap(actualMap))
	}

	return ""
}

func interfaceToByteSlice(i interface{}) ([]byte, error) {
	if b, ok := i.([]byte); ok {
		return b, nil
	}
	if s, ok := i.(string); ok {
		return []byte(s), nil
	}

	return nil, errors.New("cannot convert")
}

func prettyPrintJSONMap(m map[string]interface{}) []byte {
	b, _ := json.Marshal(&m)
	return b
}