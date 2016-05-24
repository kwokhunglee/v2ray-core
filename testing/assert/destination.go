package assert

import (
	v2net "github.com/v2ray/v2ray-core/common/net"
	"github.com/v2ray/v2ray-core/common/serial"
)

func Destination(value v2net.Destination) *DestinationSubject {
	return &DestinationSubject{value: value}
}

type DestinationSubject struct {
	Subject
	value v2net.Destination
}

func (this *DestinationSubject) Named(name string) *DestinationSubject {
	this.Subject.Named(name)
	return this
}

func (this *DestinationSubject) DisplayString() string {
	return this.Subject.DisplayString(this.value.String())
}

func (this *DestinationSubject) IsTCP() {
	if !this.value.IsTCP() {
		this.Fail(this.DisplayString(), "is", serial.StringT("a TCP destination"))
	}
}

func (this *DestinationSubject) IsNotTCP() {
	if this.value.IsTCP() {
		this.Fail(this.DisplayString(), "is not", serial.StringT("a TCP destination"))
	}
}

func (this *DestinationSubject) IsUDP() {
	if !this.value.IsUDP() {
		this.Fail(this.DisplayString(), "is", serial.StringT("a UDP destination"))
	}
}

func (this *DestinationSubject) IsNotUDP() {
	if this.value.IsUDP() {
		this.Fail(this.DisplayString(), "is not", serial.StringT("a UDP destination"))
	}
}