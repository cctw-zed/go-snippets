package designpattern

type ReportBuilder interface {
	SetTitle(title string) ReportBuilder
	SetBody(body string) ReportBuilder
	SetFooter(footer string) ReportBuilder
	Build() Report
}

type Report struct {
	Title  string
	Body   string
	Footer string
}

type PDFReportBuilder struct {
	report Report
}

func (b *PDFReportBuilder) SetTitle(title string) ReportBuilder {
	b.report.Title = "PDF: " + title
	return b
}

func (b *PDFReportBuilder) SetBody(body string) ReportBuilder {
	b.report.Body = "PDF: " + body
	return b
}

func (b *PDFReportBuilder) SetFooter(footer string) ReportBuilder {
	b.report.Footer = "PDF: " + footer
	return b
}

func (b *PDFReportBuilder) Build() Report {
	return b.report
}


type HTMLReportBuilder struct {
	report Report
}

func (b *HTMLReportBuilder) SetTitle(title string) ReportBuilder {
	b.report.Title = "HTML: " + title
	return b
}

func (b *HTMLReportBuilder) SetBody(body string) ReportBuilder {
	b.report.Body = "HTML: " + body
	return b
}

func (b *HTMLReportBuilder) SetFooter(footer string) ReportBuilder {
	b.report.Footer = "HTML: " + footer
	return b
}

func (b *HTMLReportBuilder) Build() Report {
	return b.report
}
