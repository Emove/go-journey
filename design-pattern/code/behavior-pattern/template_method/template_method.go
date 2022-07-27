package template_method

import "fmt"

// 模板方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。
// 从而使得实现符合开闭原则
// 如实现代码中通用步骤再父类中实现（准备、下载、保存、收尾）下载和保存的具体实现留到子类中，并且提供保存方法的默认实现
// 因为go不提供继承机制，需要使用匿名组合模拟实现继承
// 此处需要注意：因为父类需要调用子类方法，所以子类需要匿名组合父类的同时，父类也需要持有子类的引用

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

type HttpDownloader struct {
	*template
}

type FTPDownloader struct {
	*template
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

func NewHTTPDownloader() Downloader {
	downloader := &HttpDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HttpDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HttpDownloader) save() {
	fmt.Printf("http save\n")
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
