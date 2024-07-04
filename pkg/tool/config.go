package tool

import (
	"fmt"
	"time"
)

type Config struct {
	StaticAddrMap map[string][]string `yaml:"staticAddrMap"`
	Node          *NodeConfig         `yaml:"node"`
	Listen        *ListenConfig       `yaml:"listen"`
	Device        *DeviceConfig       `yaml:"device"`
	Handshake     *HandshakeConfig    `yaml:"handshake"`
	Firewall      *FirewallConfig     `yaml:"firewall"`
	Relay         *RelayConfig        `yaml:"relay"`
}

type NodeConfig struct {
	AmNode bool     `yaml:"amNode"`
	Addrs  []string `yaml:"addrs"`
}

type ListenConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Batch       int    `yaml:"batch"`
	ReadBuffer  int    `yaml:"readBuffer"`
	WriteBuffer int    `yaml:"writeBuffer"`
	Routines    int    `yaml:"routines"`
}

type DeviceConfig struct {
	Type string `yaml:"type"`
	// Dev Name of the device. If not set, a default will be chosen by the OS.
	// For macOS: if set, must be in the form `utun[0-9]+`.
	// For NetBSD: Required to be set, must be in the form `tun[0-9]+`
	Name               string `yaml:"name"`
	IP                 string `yaml:"ip"`
	Mask               string `yaml:"mask"`
	DropLocalBroadcast bool   `yaml:"dropLocalBroadcast"`
	DropMulticast      bool   `yaml:"dropMulticast"`
	MTU                int    `yaml:"mtu"`
}

// HandshakeConfig 握手配置
type HandshakeConfig struct {
	SyncDevice    time.Duration `yaml:"syncDevice"`    // 设备同步时间
	SyncNode      time.Duration `yaml:"syncNode"`      // 节点同步时间
	TryInterval   time.Duration `yaml:"tryInterval"`   // 尝试间隔
	Retries       int           `yaml:"retries"`       // 尝试次数
	TriggerBuffer int           `yaml:"triggerBuffer"` // 触发缓冲
	UseRelay      bool          `yaml:"useRelay"`      // 是否使用中继
}

type FirewallConfig struct {
	OutboundDefaultAction string         `yaml:"outboundDefaultAction"`
	InboundDefaultAction  string         `yaml:"inboundDefaultAction"`
	Outbound              []OutboundRule `yaml:"outbound"`
	Inbound               []InboundRule  `yaml:"inbound"`
}

type RelayConfig struct {
	AmRelay  bool     `yaml:"amRelay"`
	UseRelay bool     `yaml:"useRelay"`
	Addrs    []string `yaml:"addrs"`
}

type OutboundRule struct {
	Port  string   `yaml:"port"`
	Proto string   `yaml:"proto"`
	Host  []string `yaml:"host"`
	// "allow" or "deny"
	Action string `yaml:"action"`
}

type InboundRule struct {
	Port  string   `yaml:"port"`
	Proto string   `yaml:"proto"`
	Host  []string `yaml:"host"`
	// "allow" or "deny"
	Action string `yaml:"action"`
}

func (r OutboundRule) String() string {
	return fmt.Sprintf("port=%v proto=%s hosts=%v action=%s", r.Port, r.Proto, r.Host, r.Action)
}

func (r InboundRule) String() string {
	return fmt.Sprintf("port=%v proto=%s hosts=%v action=%s", r.Port, r.Proto, r.Host, r.Action)
}
