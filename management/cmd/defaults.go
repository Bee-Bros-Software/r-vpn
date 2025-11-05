package cmd

const (
	defaultMgmtDataDir   = "/var/lib/rvpn/"
	defaultMgmtConfigDir = "/etc/rvpn"
	defaultLogDir        = "/var/log/rvpn"

	oldDefaultMgmtDataDir   = "/var/lib/wiretrustee/"
	oldDefaultMgmtConfigDir = "/etc/wiretrustee"
	oldDefaultLogDir        = "/var/log/wiretrustee"

	defaultMgmtConfig    = defaultMgmtConfigDir + "/management.json"
	defaultLogFile       = defaultLogDir + "/management.log"
	oldDefaultMgmtConfig = oldDefaultMgmtConfigDir + "/management.json"
	oldDefaultLogFile    = oldDefaultLogDir + "/management.log"

	defaultSingleAccModeDomain = "rvpn.selfhosted"
)
