package settings

type ServerSettings struct {
	Port               int
	Master             bool
	MasterReplId       string
	MasterReplIdOffset int
	MasterHost         string
	MasterPort         int
}

func GetRoleValue(serverSettings ServerSettings) string {
	if serverSettings.Master {
		return "master"
	} else {
		return "slave"
	}
}

const Role = "role"
const MasterReplId = "master_replid"
const MasterReplIdOffset = "master_repl_offset"
