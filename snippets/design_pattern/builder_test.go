package designpattern

import (
	"fmt"
	"testing"
)

/*
建造者模式适用于以下场景：

当一个对象的构建过程复杂且涉及多个步骤时。
当一个对象的构建过程需要产生不同的表示或配置时。
当需要将对象的构建过程与表示分离，以便同样的构建过程可以创建不同的对象表示时。
以下是一个使用建造者模式的示例：

假设我们正在开发一个应用程序，需要创建一些复杂的报告。报告可以有多种格式（例如，PDF、HTML、Markdown 等），并且可以包含多种部分（例如，标题、正文、页脚等）。我们可以使用建造者模式来实现一个灵活且可扩展的报告构建系统。
*/
func TestBuilder(t *testing.T) {
	pdf := &PDFReportBuilder{}
	pdfReport := pdf.SetTitle("My Title").SetBody("My Body").SetFooter("My Footer").Build()
	t.Logf("pdfReport: %+v", pdfReport)

	html := &HTMLReportBuilder{}
	htmlReport := html.SetTitle("My Title").SetBody("My Body").SetFooter("My Footer").Build()
	t.Logf("htmlReport: %+v", htmlReport)
}

func TestBuilder2(t *testing.T) {
	pancakeCook := NewPancakeCook(NewNormalPancakeBuilder())
	fmt.Printf("摊一个普通煎饼 %#v\n", pancakeCook.MakePancake())

	pancakeCook.SetPancakeBuilder(NewHealthyPancakeBuilder())
	fmt.Printf("摊一个健康的加量煎饼 %#v\n", pancakeCook.MakeBigPancake())
}
