package process

import (
	"regexp"
	"time"
	"strconv"
	"strings"
	"net/url"
)

type Analyzer struct {
}

type Analyze interface {
	Analyze(Rc chan string, Wr chan string)
}

type Message struct {
	TimeLocal                    time.Time
	BytesSend                    int
	Path, Method, Scheme, Status string
	UpstreamTime, RequestTime    float64
}

func (analyzer *Analyzer) Analyze(Rc chan []byte, Wr chan *Message) {
	//127.0.0.1 - - [30/Jun/2018:23:58:16 +0800] "GET / HTTP/1.1" 200 12 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36"

	rep := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([\d\.-]+)`)
	loc, _ := time.LoadLocation("Asia/Shanghai")

	for v := range Rc {
		ret := rep.FindStringSubmatch(string(v))
		message := &Message{}

		//时间
		parseInLocation, _ := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
		message.TimeLocal = parseInLocation

		//流量
		bb, _ := strconv.Atoi(ret[8])
		message.BytesSend = bb

		split := strings.Split(ret[6], " ")

		message.Method = split[0]

		parse, _ := url.Parse(split[1])

		message.Path = parse.Path

		message.Scheme = ret[5]
		message.Status = ret[7]

		f, _ := strconv.ParseFloat(ret[12], 64)
		f2, _ := strconv.ParseFloat(ret[13], 64)

		message.UpstreamTime = f
		message.RequestTime = f2

		Wr <- message
	}
}
