package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"log"

	"golang.org/x/net/http2"
)

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	/* https/http2 client */
	certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal("unable to load server cert and key")
	}

	rootCACert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal("unable to load CA cert for client")
	}

	rootCACertPool := x509.NewCertPool()
	rootCACertPool.AppendCertsFromPEM(rootCACert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{certificate},
		InsecureSkipVerify: false,
		RootCAs:            rootCACertPool,
	}
	tlsConfig.BuildNameToCertificate()
	req, err := http.NewRequest("GET", "https://Server:8085/getAllArticles", nil)
	if err != nil {
		log.Fatal("Error while creating https request")
	}
	clientTLS := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	resp, err := clientTLS.Do(req)
	if err != nil {
		log.Fatal("Error while making https request : ", err.Error())
	}
	log.Println("resp status : " + resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response body")
	}
	log.Println("resp body : " + string(respBody))

	///////////////////////////////////////////////////////////////

	/* h2c client POST*/

	article := Article{ID: "3", Title: "Hello3", Desc: "Article Description3", Content: "Article Content3"}
	articleBytes, _ := json.Marshal(article)
	req, err = http.NewRequest(http.MethodPost, "http://Server:8084/setArticle", bytes.NewBuffer(articleBytes))
	if err != nil {
		log.Fatal("Error while creating https request")
	}
	req.Header.Set("Content-Type", "application/json")
	clientH2C := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	resp, err = clientH2C.Do(req)
	if err != nil {
		log.Fatal("Error while making https request : ", err.Error())
	}
	log.Println("resp status : " + resp.Status)
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response body")
	}
	log.Println("resp body : " + string(respBody))

	///////////////////////////////////////////////////////////////

	/* http client GET*/

	req, err = http.NewRequest(http.MethodGet, "http://Server:8084/getAllArticles", nil)
	if err != nil {
		log.Fatal("Error while creating https request")
	}
	clientHTTP := http.DefaultClient

	resp, err = clientHTTP.Do(req)
	if err != nil {
		log.Fatal("Error while making https request : ", err.Error())
	}
	log.Println("resp status : " + resp.Status)
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response body")
	}
	log.Println("resp body : " + string(respBody))
}
