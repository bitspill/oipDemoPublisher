package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bitspill/flod/chaincfg"
	"github.com/bitspill/flod/rpcclient"
	"github.com/bitspill/floutil"
	"github.com/oipwg/oip/publishing"
	"github.com/oipwg/proto/go/pb_oip5"
	"github.com/oipwg/proto/go/pb_oip5/pb_templates"
)

var floAddress = "F..."
var floWifKey = "R..."

func main() {
	b, err := ioutil.ReadFile("../plant/plant.dso")
	if err != nil {
		check(err)
	}

	rtp := &pb_templates.RecordTemplateProto{
		FriendlyName:       "Checklist Plant",
		Description:        "The PLANTS Database provides standardized information about the vascular plants, mosses, liverworts, hornworts, and lichens of the U.S. and its territories.",
		DescriptorSetProto: b,
	}
	template := &pb_oip5.OipFive{
		RecordTemplate: rtp,
	}

	addr, err := floutil.DecodeAddress(floAddress, &chaincfg.MainNetParams)
	check(err)

	cfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:8334",
		Endpoint:     "ws",
		User:         "user",
		Pass:         "pass",
		DisableTLS:   true,
		Certificates: nil,
	}
	floClient, err := rpcclient.New(cfg, nil)
	check(err)

	wif, err := floutil.DecodeWIF(floWifKey)
	check(err)

	pub, err := publishing.NewAddress(floClient, addr, wif, &chaincfg.MainNetParams)
	check(err)

	err = pub.UpdateUtxoSet()
	check(err)

	res, err := pub.Publish(template)
	if res != nil {
		for _, sbr := range res.Sbr {
			for _, h := range sbr.TxHash {
				fmt.Println(h.String())
			}
		}
	}
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
