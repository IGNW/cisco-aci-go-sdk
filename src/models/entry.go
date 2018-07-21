package models

const ENTRY_RESOURCE_PREFIX = "e"
const ENTRY_OBJECT_CLASS = "vzEntry"

// Represents an ACI Contract Filter Entry.
// https://pubhub.devnetcloud.com/media/apic-mim-ref-311/docs/MO-vzEntry.html
type Entry struct {
	ResourceAttributes
	Protocol            string `oneof=unspecified icmp igmp tcp egp igp udp icmpv6 eigrp ospfigp pim l2tp`
	ArpOpCodes          string `oneof=unspecified req reply`
	ApplyToFrag         bool
	EthernetType        string `oneof=unspecified ipv4 trill arp ipv6 mpls_ucast mac_security fcoe ip`
	ICMPv4Settings      string `oneof=echo-rep dst-unreach src-quench echo time-exceeded unspecified`
	ICMPv6Settings      string `oneof=unspecified dst-unreach time-exceeded echo-req echo-rep nbr-solicit nbr-advert redirect`
	DSCP                string `oneof=unspecified CS0 CS1 AF11 AF12 AF13 CS2 AF21 AF22 AF23 CS3 AF31 AF32 AF33 CS4 AF41 AF42 AF43 CS5 VA EF CS6 CS7`
	Stateful            bool
	TcpFlags            string `oneof=unspecified est syn ack fin rst`
	Source, Destination *ToFrom
}

// TODO: add validation rules
type ToFrom struct {
	To, From string
}

func (e *Entry) GetObjectClass() string {
	return ENTRY_OBJECT_CLASS
}

func (e *Entry) GetResourcePrefix() string {
	return ENTRY_RESOURCE_PREFIX
}

func (e *Entry) HasParent() bool {
	return true
}

func (e *Entry) ToMap() map[string]string {
	var model = e.ResourceAttributes.ToMap()

	model["applyToFrag"] = e.FormatBool(e.ApplyToFrag)
	model["arpOpc"] = e.ArpOpCodes
	model["dFromPort"] = ""
	model["dToPort"] = ""
	model["etherT"] = e.EthernetType
	model["icmpv4T"] = e.ICMPv4Settings
	model["icmpv6T"] = e.ICMPv6Settings
	model["matchDscp"] = e.DSCP
	model["prot"] = e.Protocol
	model["sFromPort"] = ""
	model["sToPort"] = ""
	model["stateful"] = e.FormatBool(e.Stateful)
	model["tcpRules"] = e.TcpFlags

	if e.Source != nil {
		model["sFromPort"] = e.Source.From
		model["sToPort"] = e.Source.To
	}

	if e.Destination != nil {
		model["dFromPort"] = e.Destination.From
		model["dToPort"] = e.Destination.To
	}

	return model
}

// NewFilterMap will construct a Filter from a string map.
func NewEntry(model map[string]string) *Entry {

	m := Entry{NewResourceAttributes(model),
		"",
		"",
		false,
		"",
		"",
		"",
		"",
		false,
		"",
		&ToFrom{To: "", From: ""},
		&ToFrom{To: "", From: ""},
	}

	m.ApplyToFrag = m.ParseBool(model["applyToFrag"])
	m.ArpOpCodes = model["arpOpc"]
	m.EthernetType = model["etherT"]
	m.ICMPv4Settings = model["icmpv4T"]
	m.ICMPv6Settings = model["icmpv6T"]
	m.DSCP = model["matchDscp"]
	m.Protocol = model["prot"]
	m.Stateful = m.ParseBool(model["stateful"])
	m.TcpFlags = model["tcpRules"]

	m.Source.From = model["sFromPort"]
	m.Source.To = model["sToPort"]

	m.Destination.From = model["dFromPort"]
	m.Destination.To = model["dToPort"]

	return &m
}

// NewEntryMap will construct a string map from reading Filter Entry values that can be converted to the type.
func NewEntryMap() map[string]string {

	m := NewResourceAttributesMap()

	m["applyToFrag"] = "no"
	m["arpOpc"] = "unspecified"
	m["dFromPort"] = ""
	m["dToPort"] = ""
	m["etherT"] = "unspecified"
	m["icmpv4T"] = "unspecified"
	m["icmpv6T"] = "unspecified"
	m["matchDscp"] = "unspecified"
	m["prot"] = "unspecified"
	m["sFromPort"] = ""
	m["sToPort"] = ""
	m["stateful"] = "no"
	m["tcpRules"] = "unspecified"

	return m
}
