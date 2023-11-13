package designpattern

import "testing"

/*
桥接模式主要用于将抽象部分与实现部分分离，使它们可以独立地变化。它主要用于以下场景：

当一个类存在多个独立变化的维度，并且这些维度都需要扩展时。

当需要避免多层继承结构，以减少代码的复杂性和提高可维护性时。

当需要在运行时切换实现部分或者抽象部分时。

示例：

假设我们正在开发一个跨平台的图形绘制库，该库需要支持多种形状（如圆形、矩形等）和多种渲染引擎（如OpenGL、Vulkan等）。我们可以使用桥接模式将形状（抽象部分）与渲染引擎（实现部分）分离，以便它们可以独立地变化。

在这个示例中，我们定义了一个名为 Renderer 的接口，它表示渲染引擎（实现部分）。我们还定义了两个具体的渲染引擎：OpenGLRenderer 和 VulkanRenderer。接着，我们创建了一个名为 Shape 的接口，它表示形状（抽象部分）。我们为每种形状定义了一个具体的类：Circle 和 Rectangle。这些形状包含一个 Renderer 类型的字段，用于将形状与渲染引擎连接起来。

通过使用桥接模式，我们可以将形状与渲染引擎分离，使它们可以独立地变化。这种模式有助于减少代码的复杂性，提高可维护性和可扩展性。
*/
func TestBridge(t *testing.T) {

	// 创建具体实现（渲染引擎）
	openglRenderer := &OpenGLRenderer{}
	vulkanRenderer := &VulkanRenderer{}

	// 创建具体抽象（形状）
	circle := &Circle{renderer: openglRenderer}
	circle2 := &Circle{renderer: vulkanRenderer}
	rectangle := &Rectangle{renderer: vulkanRenderer}
	rectangle2 := &Rectangle{renderer: openglRenderer}

	// 使用桥接模式绘制形状
	circle.Draw()
	circle2.Draw()
	rectangle.Draw()
	rectangle2.Draw()
}
