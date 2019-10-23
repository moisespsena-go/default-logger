package defaultlogger

import (
	"os"
	"github.com/moisespsena-go/logging"
)

var Format = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05.999 -07:00}%{color} %{pid} %{level:.4s} [%{module}]: %{message}%{color:reset}`,
)

func init() {
	logging.SetBackend(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), Format))
}

func GetOrCreateLogger(module string) logging.Logger {
	return logging.GetOrCreateLogger(module)
}
