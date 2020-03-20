package controltransfer

import (
	"context"

	"github.com/google/gousb"
	"github.com/sirupsen/logrus"
)

type Command int

const (
	Down      = 0x01
	Up        = 0x02
	Left      = 0x04
	Right     = 0x08
	Fire      = 0x10
	Stop      = 0x20
	UpLeft    = Up | Left
	DownLeft  = Down | Left
	UpRight   = Up | Right
	DownRight = Down | Right
)

func (cmd Command) String() string {
	return map[Command]string{
		0x01:           "Down",
		0x02:           "Up",
		0x04:           "Left",
		0x0:            "Right",
		0x10:           "Fire",
		0x20:           "Stop",
		(Up | Left):    "UpLeft",
		(Down | Left):  "DownLeft",
		(Up | Right):   "UpRight",
		(Down | Right): "DownRight",
	}[cmd]
}

type Controller struct {
	Device *gousb.Device
}

func (c Controller) Send(ctx context.Context, command Command) (int, error) {
	logrus.Tracef("%s called\n", command)

	return c.Device.Control(0x21, 0x09, 0, 0, []byte{
		0x02,
		byte(command),
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00})
}
