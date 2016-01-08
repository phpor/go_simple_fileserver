package controller
import "net/http"


func Upload(w http.ResponseWriter, r *http.Request) {

}

/*
 http/server.go 中有如下关于expect的逻辑，使得即使请求头中含有Expect，req.Header中也看不到
 
		// Expect 100 Continue support
		req := w.req
		if req.expectsContinue() {
			if req.ProtoAtLeast(1, 1) && req.ContentLength != 0 {
				// Wrap the Body reader with one that replies on the connection
				req.Body = &expectContinueReader{readCloser: req.Body, resp: w}
			}
			req.Header.Del("Expect")
		} else if req.Header.get("Expect") != "" {
			w.sendExpectationFailed()
			break
		}
		*/
