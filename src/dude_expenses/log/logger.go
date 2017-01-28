package log

import (
	"bufio"
	"dude_expenses/app"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
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
	ua     string
	params string
	status int
	reqRaw []byte
}

func HttpLogger(env *app.Env, r *http.Request) *httpLogger {
	logger := &httpLogger{
		w:      env.GetLogStream(),
		userId: env.GetUserId(),
		ip:     getIP(r),
	}
	logger.setRequestInfo(r)
	return logger
}

func (l httpLogger) Log() {
	entry := l.buildLogEntry()
	entry = append(entry, '\n')
	l.w.Write(entry)
}

func (l *httpLogger) SetStatus(status int) {
	l.status = status
}

func (l *httpLogger) setRequestInfo(r *http.Request) {
	// Get a request dump
	l.reqRaw = reqDump(r)

	var line string
	pathRegexp, _ := regexp.Compile("(.+)\\s(.+)\\sHTTP")
	userAgentRegexp, _ := regexp.Compile("User-Agent:\\s(.+)")
	getParamsRegexp, _ := regexp.Compile("(.+)\\?(.+)")

	// The raw request comes in lines, separated by \r\n
	s := bufio.NewScanner(strings.NewReader(string(l.reqRaw)))
	for s.Scan() {
		line = s.Text()
		l.setPath(line, pathRegexp, getParamsRegexp)
		l.setUserAgent(line, userAgentRegexp)
	}
	// Last line contains the request parameters
	if len(l.params) == 0 {
		l.params = line
	}
}

func (l *httpLogger) setPath(path string, pathRegexp *regexp.Regexp, getParamsRegexp *regexp.Regexp) {
	// Check for the request path portion
	// example POST /path HTTP/1.1
	matches := pathRegexp.FindStringSubmatch(path)
	if len(matches) > 0 {
		l.method = matches[1]
		l.path = matches[2]
		// Check for query string params (GET request)
		// example GET /path?param1=value&param2=value
		matches = getParamsRegexp.FindStringSubmatch(matches[2])
		if len(matches) > 0 {
			l.path = matches[1]
			l.params = toJson(matches[2])
		}
	}
}

func (l *httpLogger) setUserAgent(uaHeader string, uaRegexp *regexp.Regexp) {
	// Check for user agent header
	// example User-Agent: <ua>
	if matches := uaRegexp.FindStringSubmatch(uaHeader); len(matches) > 0 {
		l.ua = matches[1]
	}
}

func (l httpLogger) buildLogEntry() []byte {
	buf := make([]byte, 0)
	buf = append(buf, "level=I"...)
	buf = append(buf, " time="+time.Now().UTC().Format("2006-01-02T15:04:05MST")...)
	buf = append(buf, " user="+l.userId...)
	buf = append(buf, " ip="+l.ip...)
	buf = append(buf, " method="+l.method...)
	buf = append(buf, " path="+l.path...)
	buf = append(buf, " ua="+l.ua...)
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

func reqDump(r *http.Request) []byte {
	if dump, err := httputil.DumpRequest(r, true); err != nil {
		return []byte("")
	} else {
		return dump
	}
}

func toJson(queryString string) string {
	// Poor man's JSON encoding
	r := strings.NewReplacer("=", "\": \"", "&", "\", \"")
	return "{ \"" + r.Replace(queryString) + "\" }"
}
