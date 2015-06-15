/*************************************************************************
	> File Name: ShortProcessor.go
	> Author: Wu Yinghao
	> Mail: wyh817@gmail.com
	> Created Time: 日  6/14 16:00:54 2015
 ************************************************************************/

package main

import (
	"net/http"
	"shortlib"
)

type ShortProcessor struct {
	*shortlib.BaseProcessor
}

func (this *ShortProcessor) ProcessRequest(method string, params map[string]string, body []byte, w http.ResponseWriter) error {

	return nil
}
