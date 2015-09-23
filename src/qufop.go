package main

import (
	"fmt"
	"github.com/qiniu/log"
	"os"
	"ufop"
	"ufop/amerge"
	"ufop/html2image"
	"ufop/html2pdf"
	"ufop/mkzip"
	"ufop/unzip"
)

func help() {
	fmt.Println(`qufop v1.0

qufop [<UfopConfig>]`)
}

func main() {
	log.SetOutput(os.Stdout)

	args := os.Args
	argc := len(args)

	var configFilePath string

	switch argc {
	case 2:
		configFilePath = args[1]
	default:
		help()
		return
	}

	//load config
	ufopConf := &ufop.UfopConfig{}
	confErr := ufopConf.LoadFromFile(configFilePath)
	if confErr != nil {
		log.Error("load config file error,", confErr)
		return
	}

	ufopServ := ufop.NewServer(ufopConf)

	//register job handlers
	if err := ufopServ.RegisterJobHandler("amerge.conf", &amerge.AudioMerger{}); err != nil {
		log.Error(err)
	}

	if err := ufopServ.RegisterJobHandler("html2image.conf", &html2image.Html2Imager{}); err != nil {
		log.Error(err)
	}

	if err := ufopServ.RegisterJobHandler("html2pdf.conf", &html2pdf.Html2Pdfer{}); err != nil {
		log.Error(err)
	}

	if err := ufopServ.RegisterJobHandler("mkzip.conf", &mkzip.Mkzipper{}); err != nil {
		log.Error(err)
	}

	if err := ufopServ.RegisterJobHandler("unzip.conf", &unzip.Unzipper{}); err != nil {
		log.Error(err)
	}

	//listen
	ufopServ.Listen()
}
