// Package webserver static HTTP server
package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/accounts"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/accounts/keystore"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/rpc"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func setContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Path[len("/"):]
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Start serving files in", filename)
	}

	setContent(w)
	w.Write(body)
}

func markdownHandler(w http.ResponseWriter, r *http.Request) {

	filename := r.URL.Path[len("/md/"):]

	md, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Serving file", filename)
	}

	//extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Mmark | SuperSubscript
	//parser := parser.NewWithExtensions(extensions)

	setContent(w)
	//html := markdown.ToHTML(md, parser, nil)
	//html := blackfriday.Run(md)

	hl := highlighting.NewHighlighting(
		highlighting.WithStyle("monokai"),
		highlighting.WithGuessLanguage(true),
		highlighting.WithFormatOptions(
			chromahtml.WithLineNumbers(true),
			chromahtml.LineNumbersInTable(true),
			//chromahtml.PreventSurroundingPre(true),
			//chromahtml.WithClasses(true),
		),
	)

	gm := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Typographer, hl),

		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := gm.Convert(md, &buf); err != nil {
		panic(err)
	}

	w.Write(buf.Bytes())
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func main() {
	var adr common.Address
	var f *Faucet
	var err error
	var backend bind.ContractBackend
	var auth *bind.TransactOpts
	var ks *keystore.KeyStore
	var am *accounts.Manager
	var ec *rpc.Client // NOTE: rpc NOT ethclient

	ctx := context.Background()

	ec, err = rpc.DialContext(ctx, "/home/moe/.energicore3/testnet/energi3.ipc")
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer ec.Close()

	// eth_sendTransaction
	var hex hexutil.Bytes
	var msg ethereum.CallMsg

	msg.From = common.HexToAddress("0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30")
	To := common.HexToAddress("0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c")
	msg.To = &To
	// value := big.NewInt(1000000000000000000) // in wei (1 eth)
	msg.Value = big.NewInt(3000000) // 3 million wei to send

	fmt.Println("Sending from ", msg.From, " to ", msg.To, " Value ", msg.Value)
	fmt.Println("Msg: ", toCallArg(msg))
	err = ec.CallContext(ctx, &hex, "eth_sendTransaction", toCallArg(msg))
	if err != nil {
		log.Fatalf("eth_sendTransaction: %v\n", err)
	} else {
		fmt.Println("Result: ", hex)
	}
	os.Exit(123)

	//err = client.c.CallContext //(ctx, &hex, "eth_sendTransaction", toCallArg(msg), "1")
	//if err != nil {
	//		return nil, err
	//}

	ks = keystore.NewKeyStore(
		"/home/moe/.energicore3/testnet/keystore/",
		keystore.LightScryptN,
		keystore.LightScryptP)

	am = accounts.NewManager(ks)

	for i, w := range am.Wallets() {
		fmt.Println("Wallet ", i, ": ", w)
	}

	os.Exit(123)

	//keystore.NewKeyStore()
	//auth, _ = bind.SignerFn()

	adr = common.HexToAddress("0x7A5c0AAAa0F92106a59EbE2F8D26af289b0ad104")
	fmt.Println("Connecting to faucet Smart Contract")

	f, err = NewFaucet(adr, backend)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Contract ", f)
	}

	var trans *types.Transaction
	trans, err = f.RequestGas(auth, adr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Transaction ", trans)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Start serving files in", path)
	}

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/md/", markdownHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	print("Exiting server")
	os.Exit(0)
}
