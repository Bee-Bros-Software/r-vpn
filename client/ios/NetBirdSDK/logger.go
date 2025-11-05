package R-VPNSDK

import (
	"github.com/Bee-Bros-Software/r-vpn/util"
)

// InitializeLog initializes the log file.
func InitializeLog(logLevel string, filePath string) error {
	return util.InitLog(logLevel, filePath)
}
