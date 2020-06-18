module tcpinfo

go 1.13

require (
	github.com/crutcha/netlink v1.0.0
	github.com/vishvananda/netlink v1.1.0
	github.com/vishvananda/netns v0.0.0-20200520041808-52d707b772fe // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1
)

replace github.com/vishvananda/netlink => github.com/crutcha/netlink v1.0.1-0.20200618161744-9f7fe3e78ae7
