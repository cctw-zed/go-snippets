package designpattern

import "testing"

func TestAbstractFactory(t *testing.T) {

	var factory GUIFactory

	// 使用 Windows GUI 组件
	factory = &WindowsFactory{}
	button := factory.CreateButton()
	mouse := factory.CreateMouse()
	button.Render()
	mouse.Click()

	// 使用 macOS GUI 组件
	factory = &MacOSFactory{}
	button = factory.CreateButton()
	mouse = factory.CreateMouse()
	button.Render()
	mouse.Click()
}
