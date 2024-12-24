package apollo

type State string

const (
	UnknownS         State = ":UNKNOWN"
	Online           State = ":online"
	Offline          State = ":offline"
	PreInstall       State = ":pre_install"
	PreInstallFailed State = ":pre_install_failed"
	Inventory        State = ":inventory"
	Test             State = ":test"
	OnJob            State = ":onjob"
	Resigned         State = ":RESIGNED"
	Resigning        State = ":resigning"
)

type Raid string

const (
	UnknownR Raid = ":UNKNOWN"
	Raid0    Raid = ":RAID0"
	Raid1    Raid = ":RAID1"
	Raid2    Raid = ":RAID2"
	Raid3    Raid = ":RAID3"
	Raid5    Raid = ":RAID5"
	Raid6    Raid = ":RAID6"
	Raid7    Raid = ":RAID7"
	Raid53   Raid = ":RAID53"
	Raid10   Raid = ":RAID10"
)

type Priority string

const (
	UnknownP Priority = ":UNKNOWN"
	P0       Priority = ":P0"
	P1       Priority = ":P1"
	P2       Priority = ":P2"
	P3       Priority = ":P3"
	P4       Priority = ":P4"
)

type IP string

const (
	UnknownIP IP = ":UNKNOWN"
	V4        IP = ":V4"
	V6        IP = ":V6"
)

type Device string

const (
	UnknownD Device = ":UNKNOWN"
	Storage  Device = ":storage"
	Memory   Device = ":memory"
	Cpu      Device = ":cpu"
	Network  Device = ":network"
)

type Disk string

const (
	UnknownDi Disk = ":UNKNOWN"
	Sas       Disk = ":SAS"
	Ssd       Disk = ":SSD"
	Sata      Disk = ":SATA"
)

type Machine string

const (
	UnknownM Machine = ":UNKNOWN"
	Physical Machine = ":physical"
	Virtual  Machine = ":virtual"
)

type Manufacturer string

const (
	UnknownMan  Manufacturer = ":UNKNOWN"
	Dell        Manufacturer = ":Dell"
	HP          Manufacturer = ":HP"
	HW          Manufacturer = ":HW"
	Sugon       Manufacturer = ":SUGON"
	PowerLeader Manufacturer = ":POWERLEADER"
	Lenovo      Manufacturer = ":LENOVO"
	H3C         Manufacturer = ":H3C"
	ZTE         Manufacturer = ":ZTE"
	Inspur      Manufacturer = ":INSPUR"
	Huawei      Manufacturer = ":HUAWEI"
)
