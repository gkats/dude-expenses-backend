package log

import (
	"dude_expenses/app"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

type logger interface {
	Log()
}

type httpLogger struct {
	w      io.Writer
	userId string
	ip     string
	method string
	path   string
	params string
	status int
}

func HttpLogger(env *app.Env, r *http.Request) *httpLogger {
	return &httpLogger{
		w:      env.GetLogStream(),
		userId: env.GetUserId(),
		ip:     getIP(r),
		method: r.Method,
		path:   r.URL.EscapedPath(),
		params: getParams(r),
	}
}

func (l httpLogger) Log() {
	entry := l.buildLogEntry()
	entry = append(entry, '\n')
	l.w.Write(entry)
}

func (l *httpLogger) SetStatus(status int) {
	l.status = status
}

func (l httpLogger) buildLogEntry() []byte {
	buf := make([]byte, 0)
	buf = append(buf, "level=I"...)
	buf = append(buf, " time="+time.Now().UTC().Format("2006-01-02T15:04:05MST")...)
	buf = append(buf, " user="+l.userId...)
	buf = append(buf, " ip="+l.ip...)
	buf = append(buf, " method="+l.method...)
	buf = append(buf, " path="+l.path...)
	buf = append(buf, " params="+l.params...)
	buf = append(buf, " status="+strconv.Itoa(l.status)...)
	return buf
}

func getIP(r *http.Request) string {
	if forwarded := r.Header.Get("X-Forwarded-For"); len(forwarded) > 0 {
		return forwarded
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}
	return ip
}

func getParams(r *http.Request) string {
	if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
		params, _ := ioutil.ReadAll(r.Body)
		return string(params)
	} else if r.Method == "GET" {
		return r.URL.Query().Encode()
	}
	return ""
}
