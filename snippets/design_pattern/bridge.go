package designpattern

import "fmt"

// 实现部分：渲染引擎接口
type Renderer interface {
	RenderShape(shape string)
}

// 具体实现：不同的渲染引擎
type OpenGLRenderer struct{}

func (r *OpenGLRenderer) RenderShape(shape string) {
	fmt.Printf("Rendering %s using OpenGL\n", shape)
}

type VulkanRenderer struct{}

func (r *VulkanRenderer) RenderShape(shape string) {
	fmt.Printf("Rendering %s using Vulkan\n", shape)
}

// 抽象部分：形状接口
type Shape interface {
	Draw()
}

// 具体抽象：不同的形状
type Circle struct {
	renderer Renderer
}

func (c *Circle) Draw() {
	c.renderer.RenderShape("Circle")
}

type Rectangle struct {
	renderer Renderer
}

func (r *Rectangle) Draw() {
	r.renderer.RenderShape("Rectangle")
}
