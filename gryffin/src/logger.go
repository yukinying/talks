
tcpout, err := net.Dial("tcp", "localhost:5000") // HL
logWriter = io.MultiWriter(os.Stdout, tcpout)

func (s *Scan) Logm(service, msg string) {
    m := &LogMessage{
        Service: service,
        Msg:     msg,
        Method: s.Request.Method,
        Url:    s.Request.URL.String(),
    }
    encoder := json.NewEncoder(s.LogWriter)
    encoder.Encode(m)
}