package process

type LogProcesser struct {
	Rc        chan []byte
	Wc        chan *Message
	MReader   Reader
	MWriter   Writer
	MAnalyzer Analyzer
}

type LogProcess interface {
	ReadLog()
	AnalyzeLog()
	WriteLog()
}

func (log *LogProcesser) ReadLog() {
	log.MReader.Read(log.Rc)
}

func (log *LogProcesser) AnalyzeLog() {
	log.MAnalyzer.Analyze(log.Rc, log.Wc)
}

func (log *LogProcesser) WriteLog() {
	log.MWriter.Write(log.Wc)
}
