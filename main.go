package main

import (
	"go_concurrent/process"
	"time"
)

func main() {
	processer := &process.LogProcesser{
		Rc:        make(chan []byte),
		Wc:        make(chan *process.Message),
		MReader:   process.Reader{LogPath: "./access.log"},
		MWriter:   process.Writer{InfluxdbDsn: "mysql"},
		MAnalyzer: process.Analyzer{},
	}

	go processer.ReadLog()
	go processer.AnalyzeLog()
	go processer.WriteLog()

	time.Sleep(20 * time.Second)
}
