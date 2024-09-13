package utils

import (
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func GetOGImage(resp io.ReadCloser) (string, error) {
	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		return "", err
	}

	ogImage, ok := doc.Find("meta[property='og:image']").Attr("content")

	if !ok {
		return "", errors.New("cant find og:image")
	}

	return ogImage, nil
}

func GetTagWithAttribute(resp io.ReadCloser, tag string, attr string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(resp)
	if err != nil {
		return "", err
	}

	result, ok := doc.Find(tag).Attr(attr)

	if !ok {
		return "", errors.New("cant find tag")
	}

	return result, nil
}
