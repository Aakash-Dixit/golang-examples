package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/soheilhy/cmux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var counter int
var mutex = &sync.Mutex{}
var articles = []Article{
	Article{ID: "1", Title: "Hello1", Desc: "Article Description1", Content: "Article Content1"},
	Article{ID: "2", Title: "Hello2", Desc: "Article Description2", Content: "Article Content2"},
}

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func echoString(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello")
}

func incrementCounter(rw http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(rw, strconv.Itoa(counter))
	mutex.Unlock()
}

func getArticles1(rw http.ResponseWriter, r *http.Request) {
	log.Println("getArticles 1")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(articles)
}

func getArticles2(rw http.ResponseWriter, r *http.Request) {
	log.Println("getArticles 2")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(articles)
}

func getArticle(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, article := range articles {
		if article.ID == id {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(article)
			return
		}
	}

	rw.WriteHeader(http.StatusNotFound)
	json.NewEncoder(rw).Encode("Article with id : " + id + " not found")
}

func getArticleWithParams(rw http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query()["id"]
	id := ids[0]
	titles := r.URL.Query()["title"]
	title := titles[0]
	for _, article := range articles {
		if article.ID == id && article.Title == title {
			rw.WriteHeader(http.StatusOK)
			json.NewEncoder(rw).Encode(article)
			return
		}
	}

	rw.WriteHeader(http.StatusNotFound)
	json.NewEncoder(rw).Encode("Article with id : " + id + " and title : " + title + " not found")
}

func setArticle(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var art Article
	err = json.Unmarshal(reqBody, &art)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	articles = append(articles, art)
	rw.WriteHeader(http.StatusCreated)
}

func updateArticle(rw http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var art Article
	err = json.Unmarshal(reqBody, &art)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	for index, article := range articles {
		if article.ID == art.ID {
			articles[index] = art
			return
		}
	}
	articles = append(articles, art)
	rw.WriteHeader(http.StatusCreated)
}

func getAllArticles(rw http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(bytes))
}

func deleteArticle(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range articles {
		if article.ID == id {
			articles = append(articles[:index], articles[index+1:]...)
			rw.WriteHeader(http.StatusOK)
			return
		}
	}
	rw.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hi")
	})

	http.HandleFunc("/getArticles", getArticles1)

	log.Println("Listening for http requests on 8081")
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	router := mux.NewRouter()
	router.HandleFunc("/getAllArticles", getArticles2).Methods("GET")
	router.HandleFunc("/getArticle/{id}", getArticle).Methods("GET")
	router.HandleFunc("/getArticleWithParams", getArticleWithParams).Methods("GET")
	router.HandleFunc("/getArticles", getArticles2).Methods("GET")
	router.HandleFunc("/setArticle", setArticle).Methods("POST")
	router.HandleFunc("/setArticle", updateArticle).Methods("PUT")
	router.HandleFunc("/deleteArticle/{id}", deleteArticle).Methods("DELETE")

	log.Println("Listening for http requests on 8082")
	go func() {
		log.Fatal(http.ListenAndServe(":8082", router))
	}()

	l, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("unable to create listener on port 8083")
	}

	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal("unable to load server cert and key")
	}

	clientCACert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal("unable to load CA cert for client")
	}

	clientCACertPool := x509.NewCertPool()
	clientCACertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCACertPool,
	}

	server := &http.Server{
		TLSConfig: tlsConfig,
		Handler:   router,
	}
	log.Println("Listening for https requests on 8083")
	go func() {
		log.Fatal(server.ServeTLS(l, "", ""))
	}()

	l1, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal("unable to create listener on port 8084")
	}
	cMux := cmux.New(l1)
	httpListener := cMux.Match(cmux.HTTP2(), cmux.HTTP1())
	httpsListener := cMux.Match(cmux.TLS())

	go serveHTTPAndH2C(httpListener, router)
	go serveHTTPSAndHTTP2(httpsListener, router)

	log.Println("Listening for http/h2c/https requests on 8084")
	go func() {
		// Start cmux serving.
		if err := cMux.Serve(); !strings.Contains(err.Error(),
			"use of closed network connection") {
			log.Fatal(err)
		}
	}()

	var httpServer = http.Server{
		Addr:      ":8085",
		TLSConfig: tlsConfig,
		Handler:   router,
	}
	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&httpServer, &http2Server)
	log.Println("Listening for http2 requests on 8085")
	log.Fatal(httpServer.ListenAndServeTLS("", ""))
}

func serveHTTPAndH2C(httpListener net.Listener, router http.Handler) error {
	http2Server := &http2.Server{}
	server := &http.Server{
		Handler: h2c.NewHandler(router, http2Server),
	}
	if err := server.Serve(httpListener); err != cmux.ErrListenerClosed {
		return err
	}
	return nil
}

func serveHTTPSAndHTTP2(httpsListener net.Listener, router http.Handler) error {
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal("unable to load server cert and key")
	}

	clientCACert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatal("unable to load CA cert for client")
	}

	clientCACertPool := x509.NewCertPool()
	clientCACertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCACertPool,
	}

	server := &http.Server{
		TLSConfig: tlsConfig,
		Handler:   router,
	}

	if err := server.ServeTLS(httpsListener, "", ""); err != cmux.ErrListenerClosed {
		return err
	}
	return nil
}
