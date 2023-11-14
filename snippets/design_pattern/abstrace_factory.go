package designpattern

import "fmt"

// 产品接口
type Button interface {
	Render()
}

type Mouse interface {
	Click()
}

// 具体产品
type WindowsButton struct{}

func (b *WindowsButton) Render() {
	fmt.Println("Rendering Windows button")
}

type WindowsMouse struct{}

func (b *WindowsMouse) Click() {
	fmt.Println("Clicking Windows mouse")
}

type MacOSButton struct{}

func (b *MacOSButton) Render() {
	fmt.Println("Rendering macOS button")
}

type MacOSMouse struct{}

func (b *MacOSMouse) Click() {
	fmt.Println("Clicking macOS mouse")
}

type GUIFactory interface {
	CreateButton() Button
	CreateMouse() Mouse
}

type WindowsFactory struct{}

func (f *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (f *WindowsFactory) CreateMouse() Mouse {
	return &WindowsMouse{}
}

type MacOSFactory struct{}

func (f *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (f *MacOSFactory) CreateMouse() Mouse {
	return &MacOSMouse{}
}