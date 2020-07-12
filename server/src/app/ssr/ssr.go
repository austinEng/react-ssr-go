package ssr

import (
	"fmt"
  "log"
	"net/http"
	"runtime"
)

//#cgo CFLAGS: -I../../client_snap/gen/
//#cgo LDFLAGS: -L../../client_snap/gen/ -lclient_snap -lm -ldl
//#include <client_snap.h>
import "C"

type PageRenderRequest struct {
  Request 	*http.Request
  Reply chan string
}

func Init() {
	log.Println("page_renderer init")
	C.client_snap_init()
}

func NewPageRenderer(requests chan PageRenderRequest, worker_num int) {
	runtime.LockOSThread()
	log.Printf("worker/%d start\n", worker_num)
	page_renderer := C.client_snap_instance_new(C.CString("SSR"))

	for req := range requests {

		params := fmt.Sprintf("{ \"pageTitle\": \"%s\", \"items\": [\"A\", \"B\", \"C\", \"D\", \"E\", \"F\", \"G\", \"H\", \"I\", \"J\"] }", req.Request.URL.Path)

		var result_ptr *C.char = nil
		var result_len C.int = 0
		C.client_snap_instance_call(
			page_renderer,
			C.CString("Index"),
			C.CString(params),
			&result_ptr,
			&result_len,
		)
		
		defer close(req.Reply)
		req.Reply <- C.GoStringN(result_ptr, result_len)
	}

	log.Printf("worker/%d exit\n", worker_num)
	C.client_snap_instance_delete(page_renderer)
	runtime.UnlockOSThread()
}