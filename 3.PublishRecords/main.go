package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bitspill/flod/chaincfg"
	"github.com/bitspill/flod/rpcclient"
	"github.com/bitspill/floutil"
	"github.com/oipwg/oip/publishing"
	"github.com/oipwg/proto/go/pb_oip5"

	"github.com/oipwg/OipDemoPublisher/plant"
)

var floAddress = "F..."
var floWifKey = "R..."

var skipPlants = 40000
var limitPlants = 2

var (
	addr      floutil.Address
	wif       *floutil.WIF
	txSent    = 0
	floClient *rpcclient.Client
)

func init() {
	var err error

	addr, err = floutil.DecodeAddress(floAddress, &chaincfg.MainNetParams)
	check(err)

	wif, err = floutil.DecodeWIF(floWifKey)
	check(err)
}

func main() {
	mainBody()
}

func mainBody() {
	var err error
	start := time.Now()

	defer func() {
		fmt.Printf("Sent %d tx in %s\n", txSent, time.Since(start))
	}()

	cfg := &rpcclient.ConnConfig{
		Host:         "127.0.0.1:8334",
		Endpoint:     "ws",
		User:         "user",
		Pass:         "pass",
		DisableTLS:   true,
		Certificates: nil,
	}

	floClient, err = rpcclient.New(cfg, nil)
	check(err)

	pub, err := publishing.NewAddress(floClient, addr, wif, &chaincfg.MainNetParams)
	check(err)

	pub.(*publishing.Address).WaitForAncestorConfirmations = true

	err = pub.UpdateUtxoSet()
	check(err)

	fmt.Printf("obtained utxo in %s\n", time.Since(start))
	start = time.Now()

	f, err := os.Open("./data/plants.csv")
	check(err)

	r := csv.NewReader(f)
	// Read the header
	plantData, err := r.Read()
	if !(plantData[0] == "Symbol" &&
		plantData[1] == "Synonym Symbol" &&
		plantData[2] == "Scientific Name with Author" &&
		plantData[3] == "Common Name" &&
		plantData[4] == "Family") {
		panic("unexpected header")
	}

	i := 0
	// Read and process data records
	for ; err != io.EOF; plantData, err = r.Read() {
		check(err)

		i++
		if i <= skipPlants {
			continue
		}
		if i > skipPlants+limitPlants {
			break
		}

		o5 := createPlantPublish(plantData[0], plantData[1], plantData[2], plantData[3], plantData[4])
		res, err := pub.Publish(o5)
		if res != nil {
			for _, sbr := range res.Sbr {
				txSent += len(sbr.TxHash)
				for _, h := range sbr.TxHash {
					fmt.Println(h.String())
				}
			}
		}
		check(err)
	}
}

func createPlantPublish(symbol, synonym, scientificName, name, family string) *pb_oip5.OipFive {
	plantInfo := &templates.Tmpl_025412C9{
		Symbol:                   symbol,
		SynonymSymbol:            synonym,
		ScientificNameWithAuthor: scientificName,
		CommonName:               name,
		Family:                   family,
	}
	details, err := pb_oip5.CreateOipDetails(plantInfo)
	check(err)

	recordPublish := &pb_oip5.OipFive{
		Record: &pb_oip5.RecordProto{
			Details: details,
		},
	}

	return recordPublish
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
