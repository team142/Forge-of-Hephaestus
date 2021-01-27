package web

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func StartServer(listen string) {
	logrus.Infoln("> Starting webserver on", listen)
	http.HandleFunc("api/location", func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			logrus.Errorln(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("-----------")
		fmt.Println(string(b))
		fmt.Println("-----------")
		fmt.Println(" ")
	})
	http.ListenAndServe(listen, nil)
}
