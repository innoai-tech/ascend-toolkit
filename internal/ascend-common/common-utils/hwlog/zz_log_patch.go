package hwlog

import "io"

func init() {
	RunLog = &logger{}
	RunLog.initLogWriter(io.Discard)
}
