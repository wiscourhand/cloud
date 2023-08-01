package main

import (
	"bytes"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/wiscourhand/rpcsvr/kitex_gen/base"
	"github.com/wiscourhand/rpcsvr/kitex_gen/kitex/test/server"
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
	"time"
)

const (
	exampleMethodURL = "http://127.0.0.1:8888/AServiceName/ExampleMethod"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func BenchmarkAServiceNameParallel(b *testing.B) {
	runtime.GOMAXPROCS(16)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			newExample := genExample(1)
			resp, err := example_(newExample)
			Assert(b, err == nil, err)
			fmt.Println(*resp)
		}
	})
}

func example_(exm *server.ExampleReq) (rResp *string, err error) {
	reqBody, err := json.Marshal(exm)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	reader := bytes.NewReader(reqBody)
	req, err := http.NewRequest("POST", exampleMethodURL, reader)
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	var resp *http.Response
	resp, err = httpCli.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(body, &rResp); err != nil {
		return
	}
	return
}

func genExample(id int) *server.ExampleReq {
	return &server.ExampleReq{
		Msg: fmt.Sprintf("test-%d", id),
		Base: &base.Base{
			LogID:  fmt.Sprint(id),
			Caller: "czt",
			Addr:   "163",
			Client: "c",
			TrafficEnv: &base.TrafficEnv{
				Open: true,
				Env:  "env",
			},
			Extra: nil,
		},
	}
	//return &demo.Student{
	//	Id:   int32(id),
	//	Name: fmt.Sprintf("student-%d", id),
	//	College: &demo.College{
	//		Name:    "",
	//		Address: "",
	//	},
	//	Email: []string{fmt.Sprintf("student-%d@nju.com", id)},
	//}
}

// Assert asserts cond is true, otherwise fails the test.
func Assert(t testingTB, cond bool, val ...interface{}) {
	t.Helper()
	if !cond {
		if len(val) > 0 {
			val = append([]interface{}{"assertion failed:"}, val...)
			t.Fatal(val...)
		} else {
			t.Fatal("assertion failed")
		}
	}
}

// testingTB is a subset of common methods between *testing.T and *testing.B.
type testingTB interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
}
