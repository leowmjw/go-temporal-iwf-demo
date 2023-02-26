package main

import (
	"crypto"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/revocation"
	"github.com/digitorus/pdfsign/sign"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("Welcome to esign!!")
	run()
}

func run() error {

	input_file, err := os.Open("/tmp/input.pdf")
	if err != nil {
		return err
	}
	defer input_file.Close()

	output_file, err := os.Create("/tmp/output.pdf")
	if err != nil {
		return err
	}
	defer output_file.Close()

	finfo, err := input_file.Stat()
	if err != nil {
		return err
	}
	size := finfo.Size()

	rdr, err := pdf.NewReader(input_file, size)
	if err != nil {
		return err
	}

	fmt.Println("PAGES:", rdr.NumPage())
	sd := sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        "John Doe",
				Location:    "Somewhere on the globe",
				Reason:      "My season for siging this document",
				ContactInfo: "How you like",
				Date:        time.Now().Local(),
			},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		Signer:          nil,         // crypto.Signer
		DigestAlgorithm: crypto.SHA1, // hash algorithm for the digest creation
		Certificate:     nil,         // x509.Certificate
		//CertificateChains: certificate_chains, // x509.Certificate.Verify()
		TSA: sign.TSA{
			URL:      "https://freetsa.org/tsr",
			Username: "",
			Password: "",
		},

		// The follow options are likely to change in a future release
		//
		// cache revocation data when bulk signing
		RevocationData: revocation.InfoArchival{},
		// custom revocation lookup
		RevocationFunction: sign.DefaultEmbedRevocationStatusFunction,
	}
	spew.Dump(sd)
	err = sign.Sign(input_file, output_file, rdr, size, sd)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Signed PDF written to /tmp/output.pdf")
	}

	return nil
}
