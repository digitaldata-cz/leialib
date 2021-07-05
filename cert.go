package leialib

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

func LoadCertFile(s string) (crt *x509.Certificate) {
	f, err := ioutil.ReadFile(s)
	if err != nil {
		PrintlnStderr(err.Error())
		os.Exit(1)
	}
	p, _ := pem.Decode(f)
	if p == nil {
		PrintlnStderr("Can not decode PEM block in file " + s)
		os.Exit(1)
	}
	crt, err = x509.ParseCertificate(p.Bytes)
	if err != nil {
		PrintlnStderr(err.Error())
		os.Exit(1)
	}
	return crt
}
