package designpattern

import "fmt"

// 组件接口：表示文件系统中的元素
type FileSystemElement interface {
	GetName() string
	Print(indent string)
}

// 叶子节点：文件
type File struct {
	name string
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) Print(indent string) {
	fmt.Printf("%s- %s\n", indent, f.GetName())
}

// 复合节点：文件夹
type Directory struct {
	name     string
	elements []FileSystemElement
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) Print(indent string) {
	fmt.Printf("%s+ %s\n", indent, d.GetName())
	for _, element := range d.elements {
		element.Print(indent + "  ")
	}
}

func (d *Directory) Add(element FileSystemElement) {
	d.elements = append(d.elements, element)
}

// Region 行政区，作为组合模式component接口
type Region interface {
	Name() string    // 名称
	Population() int //人口
	GDP() float64    // gdp
}

// town 区镇，组合模式中相当于叶子节点
type town struct {
	name       string
	population int
	gdp        float64
}

// NewTown 创建区镇，根据名称、人口、GDP
func NewTown(name string, population int, gdp float64) *town {
	return &town{
		name:       name,
		population: population,
		gdp:        gdp,
	}
}

func (c *town) Name() string {
	return c.name
}

func (c *town) Population() int {
	return c.population
}

func (c *town) GDP() float64 {
	return c.gdp
}

// cities 市，包括县市或者地市，组合模式中相当于composite
type cities struct {
	name    string
	regions map[string]Region
}

// NewCities 创建一个市
func NewCities(name string) *cities {
	return &cities{
		name:    name,
		regions: make(map[string]Region),
	}
}

func (c *cities) Name() string {
	return c.name
}

func (c *cities) Population() int {
	sum := 0
	for _, r := range c.regions {
		sum += r.Population()
	}
	return sum
}

func (c *cities) GDP() float64 {
	sum := 0.0
	for _, r := range c.regions {
		sum += r.GDP()
	}
	return sum
}

// Add 添加多个行政区
func (c *cities) Add(regions ...Region) {
	for _, r := range regions {
		c.regions[r.Name()] = r
	}
}

// Remove 递归删除行政区
func (c *cities) Remove(name string) {
	for n, r := range c.regions {
		if n == name {
			delete(c.regions, name)
			return
		}
		if city, ok := r.(*cities); ok {
			city.Remove(name)
		}
	}
}

func (c *cities) Regions() map[string]Region {
	return c.regions
}
