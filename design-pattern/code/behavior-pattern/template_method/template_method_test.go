package template_method

import "testing"

func TestNewHTTPDownloader(t *testing.T) {
	downloader := NewHTTPDownloader()
	downloader.Download("http://www.example.com/test.txt")
}

func TestNewFTPDownloader(t *testing.T) {
	downloader := NewFTPDownloader()
	downloader.Download("ftp://www.example.com/test.txt")
}
