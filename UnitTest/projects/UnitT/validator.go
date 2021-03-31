package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

type (
	RawPath struct {
		NasID    string `validate:"required,omitempty,gte=3,lte=10" json:"nasid"`
		SVCID    string `validate:"required,lte=2" json:"svc_id"`
		PJTID    string `validate:"required,gte=3,lte=10" json:"pjt_id"`
		FilePath string `validate:"required,gte=3,lte=500" json:"filepath"`
		FileName string `validate:"required,gte=5,lte=100" json:"filename"`
	}
)

func main() {

	getpath, _ := os.Getwd()
	fmt.Println(getpath)

	validate := validator.New()

	// rawpath
	rawpath := RawPath{
		NasID:    "nasid",
		SVCID:    "SV",
		PJTID:    "pjt.PJTID",
		FilePath: "filepath",
		FileName: "filename",
	}
	if e := validate.Struct(rawpath); e != nil {
		fmt.Printf("error opening file: %v", e.Error())
		//return nil
	}

}
