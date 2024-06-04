package gotest

import (
	"fmt"

	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus"
)

func PrintRestful() {
	restfulParam := restful.Parameter{}
	fmt.Println(restfulParam)
}

func PrintK8s() {
	err := prometheus.AlreadyRegisteredError{}
	fmt.Println(err)
}