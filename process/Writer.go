package process

import (
	"log"
	"github.com/influxdata/influxdb/client/v2"
)

type Writer struct {
	InfluxdbDsn string
}

type Write interface {
	Write(Wc chan string)
}

const (
	MyDB     = "mydb"
	username = "ysnows"
	password = "Ysnows"
)

func (writer *Writer) Write(Wc chan *Message) {

	for value := range Wc {
		// Create a new HTTPClient
		c, err := client.NewHTTPClient(client.HTTPConfig{
			Addr:     "http://localhost:8086",
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()

		// Create a new point batch
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  MyDB,
			Precision: "s",
		})
		if err != nil {
			log.Fatal(err)
		}

		// Create a point and add to batch
		tags := map[string]string{"PATH": value.Path, "SCHEMA": value.Scheme, "METHOD": value.Method, "STATUS": value.Status}
		fields := map[string]interface{}{
			"UpstreamTime": value.UpstreamTime,
			"RequestTime":  value.RequestTime,
			"BytesSent":    value.BytesSend,
		}

		pt, err := client.NewPoint("nginx_log", tags, fields, value.TimeLocal)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := c.Write(bp); err != nil {
			log.Fatal(err)
		}

		// Close client resources
		if err := c.Close(); err != nil {
			log.Fatal(err)
		}

		println("write success")
	}
}
