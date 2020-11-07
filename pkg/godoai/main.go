package godoai

import (
	"github.com/godot-go/godot-go/pkg/gdnative"
)

type Main struct {
	gdnative.NodeImpl
	gdnative.UserDataIdentifiableImpl

	score int64
}

func (p *Main) ClassName() string {
	return "Main"
}

func (p *Main) BaseClass() string {
	return "Node"
}

func (p *Main) Init() {
}

func (p *Main) OnClassRegistered(e gdnative.ClassRegisteredEvent) {
	// methods
	e.RegisterMethod("_ready", "Ready")
}

func (p *Main) Ready() {
}

func NewMainWithOwner(owner *gdnative.GodotObject) Main {
	inst := gdnative.GetCustomClassInstanceWithOwner(owner).(*Main)
	return *inst
}

func init() {
	gdnative.RegisterInitCallback(func() {
		gdnative.RegisterClass(&Main{})
	})
}
