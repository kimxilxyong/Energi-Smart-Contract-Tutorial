// Package webserver static HTTP server
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/common/hexutil"
	"energi.world/core/gen3/params"
	"energi.world/core/gen3/rpc"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

const version = 105
const serverIPC = "/home/kim/.energicore3/testnet/energi3.ipc"
const gasDonorAddress = "0x09ae1a5ddfd481cfd3cc4390b0e08a0832709a06"
const weiGasDonation = 500432 * 20000000000  // 0.5 million gas * gasPrice for gas to send
/*
	Returns the most probable filepath and mime type
*/
func getSemanticFile(basedir string, subdir string, filename string) (string, string, error) {
	var fullfilepath string
	var mime string
	var mimedir string
	var err error

	ext := filepath.Ext(filename)
	if ext != "" {
		if ext[0] == '.' {
			ext = ext[1:]
		}
	} else {
		mime = ""
		fmt.Println("ERROR Unknown request ", filename)
		err = errors.New("Unknown request " + filename)
		return "", mime, err
	}

	if filename == "js/app.js.LICENSE.txt" {
		fullfilepath = filepath.Join(basedir, filename)
		if _, err := os.Stat(fullfilepath); err == nil {
			mime = "text/plain"
			return fullfilepath, mime, nil
		}
	}

	if ext == "css" {
		mime = "text/css"
	} else if ext == "js" {
		mime = "application/javascript"
	} else if ext == "xml" {
		mime = "application/xml"
	} else if ext == "htm" || ext == "html" {
		mime = "text/html"
	} else if ext == "png" {
		mime = "image/png"
	} else if ext == "jpg" || ext == "jpeg" || ext == "jpe" {
		mime = "image/jpeg"
	} else if ext == "gif" {
		mime = "image/gif"
	} else if ext == "ico" {
		mime = "image/x-icon"
	} else if ext == "svg" {
		mime = "image/svg+xml"
	} else if ext == "json" {
		mime = "application/json"
	} else if ext == "map" {
		mime = "application/json"
	} else if ext == "ttf" {
		mime = "font/ttf"
	} else if ext == "woff" {
		mime = "font/woff"
	} else if ext == "woff2" {
		mime = "font/woff2"
	} else if ext == "eot" {
		mime = "application/vnd.ms-fontobject"
	} else {
		mime = "message/http"
		fmt.Println("ERROR: Unknown file extension", filename)
		err = errors.New("Unknown file extension " + ext)
		return "", mime, err
	}

	fullfilepath = filepath.Join(filename)
	if _, err := os.Stat(fullfilepath); err == nil {
		return fullfilepath, mime, nil
	}
	fullfilepath = filepath.Join(basedir, filename)
	if _, err := os.Stat(fullfilepath); err == nil {
		return fullfilepath, mime, nil
	}
	fullfilepath = filepath.Join(basedir, subdir, filename)
	if _, err := os.Stat(fullfilepath); err == nil {
		return fullfilepath, mime, nil
	}

	if strings.Split(mime, "/")[0] == "image" {
		mimedir = "images"
	} else if strings.Split(mime, "/")[1] == "css" {
		mimedir = "css"
	} else if strings.Split(mime, "/")[1] == "javascript" {
		mimedir = "js"
	} else if strings.Split(mime, "/")[1] == "html" {
		mimedir = "html"
	} else if strings.Split(mime, "/")[1] == "json" {
		mimedir = "json"
	} else if strings.Split(mime, "/")[0] == "font" {
		mimedir = "fonts"
	} else if strings.Split(mime, "/")[1] == "vnd.ms-fontobject" {
		mimedir = "fonts"
	}

	fullfilepath = filepath.Join(basedir, subdir, mimedir, filename)
	if _, err := os.Stat(fullfilepath); err == nil {
		return fullfilepath, mime, nil
	}
	fullfilepath = filepath.Join(basedir, mimedir, filename)
	if _, err := os.Stat(fullfilepath); err == nil {
		return fullfilepath, mime, nil
	}

	fmt.Println("ERROR: File not found ", filename)
	err = errors.New("File not found " + mime + " " + filename)
	return "", mime, err
}

func faucetGUIHandler(w http.ResponseWriter, r *http.Request) {
	var subdir string
	var accept []string
	var accepts []string
	var atype string
	var adetail string

	//fmt.Println("Incoming request: ", r.URL.Path)
	//fmt.Println("Header: ", r.Header.Get("Accept"))

	accepts = strings.Split(r.Header.Get("Accept"), ",")
	accept = strings.Split(accepts[0], "/")
	if len(accept) > 1 {
		atype = accept[0]
		adetail = accept[1]
	}

	if atype == "image" {
		subdir = "images"
	} else if atype == "text" {
		if adetail == "javascript" {
			subdir = "js"
		} else if adetail == "html" {
			subdir = ""
		} else if adetail == "css" {
			subdir = "css"
		}
	} else {
		subdir = ""
	}

	filename := strings.Split(r.URL.Path[len("/"):], "?")[0]
	if filename == "" {
		filename = "index.html"
		log.Println("New client:", r.RemoteAddr)
	}

	filenamePath, mime, err := getSemanticFile("www", subdir, filename)
	if err != nil {
		filenamePath, mime, err = getSemanticFile("www", "static", filename)
		if err != nil {
			log.Println("ERROR:", err)
			w.WriteHeader(http.StatusNotFound)
			//http.Error(w, "Not Found: "+mime+" "+filenamePath, http.StatusNotFound)
			return
		}
	}

	body, err := ioutil.ReadFile(filenamePath)
	if err != nil {
		log.Println("ERROR:", err)
		w.WriteHeader(http.StatusNotFound)
		//http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	//fmt.Println("Serving file ", filenamePath)

	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")
	w.Write(body)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	var subdir string
	var accept []string
	var accepts []string
	var atype string
	var adetail string

	//fmt.Println("Incoming request: ", r.URL.Path)
	//fmt.Println("Header: ", r.Header.Get("Accept"))

	accepts = strings.Split(r.Header.Get("Accept"), ",")
	accept = strings.Split(accepts[0], "/")
	if len(accept) > 1 {
		atype = accept[0]
		adetail = accept[1]
	}

	if atype == "image" {
		subdir = "images"
	} else if atype == "text" {
		if adetail == "javascript" {
			subdir = "js"
		} else if adetail == "html" {
			subdir = ""
		} else if adetail == "css" {
			subdir = "css"
		}
	} else {
		subdir = ""
	}

	filename := strings.Split(r.URL.Path[len("/"):], "?")[0]
	if filename == "" {
		filename = "index.html"
		log.Println("New HTML client:", r.RemoteAddr)
	}

	filenamePath, mime, err := getSemanticFile("html", subdir, filename)
	if err != nil {
		filenamePath, mime, err = getSemanticFile("html", "static", filename)
		if err != nil {
			log.Println("ERROR:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	body, err := ioutil.ReadFile(filenamePath)
	if err != nil {
		log.Println("ERROR:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("HTML Serving file ", filenamePath)

	w.Header().Set("Content-Type", mime+"; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")
	w.Write(body)
}

func markdownHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Incoming MD request:", r)

	filename := r.URL.Path[len("/md/"):]

	md, err := ioutil.ReadFile("md/" + filename)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Serving file md/", filename)
	}

	//extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.Mmark | SuperSubscript
	//parser := parser.NewWithExtensions(extensions)

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

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")
	w.Write(buf.Bytes())
}

// AjaxResponse Json Response
type AjaxResponse struct {
	Version         uint16 `json:"version"`
	Status          bool   `json:"status"`
	TransactionHash string `json:"transactionHash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Wei             uint   `json:"value"`
	Err             error  `json:"error"`
}

// MsgError Our error object which is returned in an ajax response
type MsgError struct {
	Code 		int 		`json:"code"`
	Message string 	`json:"message"`
}
// Implement the error interface for our own error
func (e *MsgError) Error() string {
  return "Error " + strconv.Itoa(e.Code) + ", " + e.Message
}

// AJAX Request Handler
func ajaxHandlerRequestGas(w http.ResponseWriter, r *http.Request) {
	var sendGasTo string
	var tx hexutil.Bytes
	var from string
	var to string
	var wei uint

	log.Println("Gas request from: ", r.RemoteAddr)

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Content-Language", "en-US")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendGasTo = r.Form["toAdr"][0]

	var response AjaxResponse
	response.Version = version

	//  TODO check if sendGasTo is a valid address
	if sendGasTo == "" {
		log.Println("ERROR: Ajax data ", r.Form)
		log.Println("ERROR: Adr ", r.Form["toAdr"])
		//http.Error(w, r.Form.Encode(), http.StatusBadRequest)
		//return
		response.Status = false
		response.Err = &MsgError{Code: 1, Message: "address not set" }
	} else {
		// TODO security

		log.Println("Incoming AJAX request, sendGasTo ", sendGasTo)
		balance, err := getBalance(gasDonorAddress);
		if err != nil {
			log.Println("ERROR: getBalance(gas) failed:", err)
			response.Status = false
			response.Err = &MsgError{Code: 2, Message: "getBalance(gasDonorAddress) failed: " + err.Error() }
		} else {
			log.Println("Current gasFrom Balance: ",  weiToEther(&balance), "NRG")

			// Ceck the bablance of the requester
			balance, err = getBalance(sendGasTo);
			if err != nil {
				log.Println("ERROR: getBalance(to) failed:", err)
			} else {
				log.Println("Current gasTo Balance: ",  weiToEther(&balance), "NRG")

				// Allow him to have double the gas donation
				weiInt := big.NewInt(weiGasDonation*2)
				if weiInt.Cmp(&balance) > 0  {
					tx, from, to, wei, err = sendGas(sendGasTo)
					response.Err = err
					if err != nil {
						response.Status = false
					} else {
						response.Status = true
					}
					response.TransactionHash = tx.String()
					response.To = to
					response.From = from
					response.Wei = wei
				} else {
					// requester balance is more than double weiGasDonation
					response.To = sendGasTo
					response.Wei = weiGasDonation
					response.Status = false
					response.Err = &MsgError{Code: 3, Message: "You have enough gas (" + balance.String() + ")"}
					log.Println("You have enough gas",  response)
				}
			}
	  }
	}

	// create json from response struct
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("ERROR json.Marshal: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("jsonResponse ", string(jsonResponse))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jsonResponse)
}

var lastGasToAddress string
var lastGasCount time.Duration

// returns the transaction hash
func sendGas(gasToAddress string) (hexutil.Bytes, string, string, uint, error) {
	var err error
	var ec *rpc.Client // NOTE: rpc NOT ethclient

	if (lastGasToAddress == gasToAddress) {
		lastGasCount++
		time.Sleep((50*lastGasCount) * time.Millisecond)
	} else {
		lastGasCount = 0
		lastGasToAddress = gasToAddress
	}

	ctx := context.Background()

	ec, err = rpc.DialContext(ctx, serverIPC)
	if err != nil {
		log.Println("ERROR: could not connect to Ethereum gateway:", err)
	}
	defer ec.Close()

	// eth_sendTransaction
	var hex hexutil.Bytes
	var msg ethereum.CallMsg

	// Gas donor
	msg.From = common.HexToAddress(gasDonorAddress)
	// Gas beggar
	t := common.HexToAddress(gasToAddress)
	msg.To = &t
	// how much wei we send
	msg.Value = big.NewInt(weiGasDonation)

	err = ec.CallContext(ctx, &hex, "eth_sendTransaction", toCallArg(msg))
	if err != nil {
		log.Println("ERROR: eth_sendTransaction failed:", err)
	}
	return hex, gasDonorAddress, gasToAddress, weiGasDonation, err
}

// returns the balance
func getBalance(address string) (big.Int, error) {
	var err error
	var ec *rpc.Client // NOTE: rpc NOT ethclient

	ctx := context.Background()

	ec, err = rpc.DialContext(ctx, serverIPC)
	if err != nil {
		fmt.Println("ERROR: could not connect to Ethereum gateway:", err)
	}
	defer ec.Close()

	// eth_getBalance
	var response string
	err = ec.CallContext(ctx, &response, "eth_getBalance", address, "latest")
	if err != nil {
		fmt.Println("ERROR: eth_getBalance failed:", err)
		return big.Int{}, err
	}
	return ParseBigInt(response)
}

func main() {

	lastGasCount = 0
	log.SetFlags(log.LstdFlags)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	} else {

		balance, err := getBalance(gasDonorAddress);
		if err != nil {
			log.Fatal("Communication with node failed!")
			log.Fatal(err)
			os.Exit(1)
		}

		log.Println("Server version", version)
		log.Println("IPC: ", serverIPC)
		log.Println("Gas Donor: ", gasDonorAddress)
		log.Println("Gas Donation amount: ", weiGasDonation, "wei")
		log.Println("Gas Balance: ",  weiToEther(&balance), "NRG")
		log.Println("Start serving files in", filepath.Join(path, "www"))
		log.Println("Start serving AJAX gas donation reguests under /ajax/requestGas")
	}

	http.HandleFunc("/ajax/requestGas", ajaxHandlerRequestGas)
	http.HandleFunc("www.servht.ml/md/", markdownHandler)
	http.HandleFunc("www.servht.ml/", htmlHandler)

	http.HandleFunc("faucet.servht.ml/ajax/requestGas", ajaxHandlerRequestGas)
	http.HandleFunc("faucet.servht.ml/", faucetGUIHandler)

	http.HandleFunc("localhost/ajax/requestGas", ajaxHandlerRequestGas)
	http.HandleFunc("localhost/md/", markdownHandler)
	http.HandleFunc("localhost/", faucetGUIHandler)

	//log.Fatal(http.ListenAndServeTLS(":443", "ssl_server.crt", "ssl_server.key", nil))

	log.Println("Listening on :443 with ssl_server.crt, ssl_server.key")
	log.Fatal(http.ListenAndServeTLS(":443", "ssl_server.crt", "ssl_server.key", nil))

	print("Exiting server")
	os.Exit(0)
}

// Web3 Section ******************************************
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

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)

	return i, err
}

func etherToWei(val *big.Int) *big.Int {
	return new(big.Int).Mul(val, big.NewInt(params.Ether))
}

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

/*
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

	http.HandleFunc("/", generalHandler)
	http.HandleFunc("/md/", markdownHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	print("Exiting server")
	os.Exit(0)
}
*/
