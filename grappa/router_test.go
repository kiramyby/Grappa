package grappa

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.routeRegister("GET", "/", nil)
	r.routeRegister("GET", "/hello/:name", nil)
	r.routeRegister("GET", "/hello/b/c", nil)
	r.routeRegister("GET", "/hi/:name", nil)
	r.routeRegister("GET", "/assets/*filepath", nil)

	return r
}

func TestDividePath(t *testing.T) {
	ok := reflect.DeepEqual(dividePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(dividePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(dividePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test dividePath failed")
	}
}

func TestParseRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.parseRoute("GET", "/hello/Kiracoon")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.path != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "Kiracoon" {
		t.Fatal("name should be equal to 'Kiracoon'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.path, ps["name"])

}

func TestParseRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.parseRoute("GET", "/assets/file1.txt")
	ok1 := n1.path == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("path should be /assets/*filepath & filepath should be file1.txt")
	}
	fmt.Printf("matched path1: %s, params1['filepath']: %s\n", n1.path, ps1["filepath"])

	n2, ps2 := r.parseRoute("GET", "/assets/css/test.css")
	ok2 := n2.path == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("path should be /assets/*filepath & filepath should be css/test.css")
	}
	fmt.Printf("matched path2: %s, params2['filepath']: %s\n", n2.path, ps2["filepath"])

}

func TestCheckPaths(t *testing.T) {
	r := newTestRouter()
	nodes := r.checkPaths("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of paths should be 4")
	}
}
