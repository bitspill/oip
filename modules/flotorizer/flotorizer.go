package flotorizer

import (
	"strings"

	"github.com/bitspill/oip/config"
	"github.com/bitspill/oip/datastore"
	"github.com/bitspill/oip/events"
	"gopkg.in/olivere/elastic.v6"
)

func init() {
	log.Info("init flotorizer")
	if !config.IsTestnet() {
		events.SubscribeAsync("flo:floData", onFloData, false)
		events.SubscribeAsync("modules:flotorizer:flotorized", onFlotorized, false)
		datastore.RegisterMapping("flotorizer", "flotorizer.json")
	}
}

func onFloData(floData string, tx *datastore.TransactionData) {
	if tx.Block < 1500000 {
		return
	}
	prefix := "This document has been flotorized: "
	if strings.HasPrefix(floData, prefix) {
		events.Publish("modules:flotorizer:flotorized", strings.TrimPrefix(floData, prefix))
		return
	}
}

func onFlotorized(floData string) {
	f := Flotorized{
		Hash: floData,
		// TxId: txid,
	}
	bir := elastic.NewBulkIndexRequest().Index(datastore.Index("flotorizer")).Type("_doc"). /*Id(txid).*/ Doc(f)
	datastore.AutoBulk.Add(bir)
}

type Flotorized struct {
	Hash string `json:"hash"`
	TxId string `json:"txId"`
}
