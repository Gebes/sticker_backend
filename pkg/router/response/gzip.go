package response

import (
	"compress/gzip"
	. "github.com/Gebes/there/v2"
	"io"
	"net/http"
	"strings"
)

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

//Gzip if returned will continue to the next middleware or the response
func Gzip(response HttpResponse) *gzipMiddleware {
	r := &gzipMiddleware{response}
	return r
}

type gzipMiddleware struct {
	response HttpResponse
}

func (j gzipMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get(RequestHeaderAcceptEncoding), "gzip") {
		j.response.ServeHTTP(rw, r)
		return
	}

	gz := gzip.NewWriter(rw)
	defer gz.Close()

	rw.Header().Set("Content-Encoding", "gzip")

	var responseWriter http.ResponseWriter = gzipResponseWriter{Writer: gz, ResponseWriter: rw}
	j.response.ServeHTTP(responseWriter, r)
}
