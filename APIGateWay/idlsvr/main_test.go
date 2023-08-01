package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wiscourhand/idlsvr/kitex_gen/demo"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	addIdlURL = "http://127.0.0.1:8888/idl-info/add/CServiceName"
)

var httpCli = &http.Client{Timeout: 3 * time.Second}

func BenchmarkAServiceName(b *testing.B) {
	for i := 1; i < b.N; i++ {
		newIdl := genExample(i)
		resp, err := register_(newIdl)
		Assert(b, err == nil, err)
		fmt.Println(*resp)
	}
}

func register_(idl *demo.IdlReq) (rResp *string, err error) {
	reqBody, err := json.Marshal(idl)
	if err != nil {
		return nil, fmt.Errorf("marshal request failed: err=%v", err)
	}
	reader := bytes.NewReader(reqBody)
	req, err := http.NewRequest("POST", addIdlURL, reader)
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

func genExample(id int) *demo.IdlReq {
	return &demo.IdlReq{
		IdlName: fmt.Sprintf("idl-test-%d", id),
		IdlInfo: &demo.IdlInfo{
			Content:  fmt.Sprint(id),
			Includes: nil,
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
