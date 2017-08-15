package util

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"os/signal"
	"strings"

	"xsec-dns-server/models"

	"github.com/miekg/dns"
	"github.com/urfave/cli"
)

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	doProxy(DNS_SERVER, w, r)
}

func doProxy(addr string, w dns.ResponseWriter, req *dns.Msg) {
	transport := "udp"
	if _, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		transport = "tcp"
	}
	if isTransfer(req) {
		if transport != "tcp" {
			dns.HandleFailed(w, req)
			return
		}
		t := new(dns.Transfer)
		c, err := t.In(req, addr)
		if err != nil {
			dns.HandleFailed(w, req)
			return
		}
		if err = t.Out(w, req, c); err != nil {
			dns.HandleFailed(w, req)
			return
		}
		return
	}
	c := &dns.Client{Net: transport}
	resp, _, err := c.Exchange(req, addr)
	if err != nil {
		dns.HandleFailed(w, req)
		return
	}

	targets := req.Question
	target := targets[0].Name
	client := w.RemoteAddr().String()
	clientIpPort := strings.Split(client, ":")
	clientIp := clientIpPort[0]
	retReplys := resp.Answer
	replys := make([]string, 0)
	for _, t := range retReplys {
		replys = append(replys, t.String())
	}

	if DEBUG_MODE {
		log.Printf("DB type: %v, Target:%v, Client: %v, Reply: %v\n", models.DATA_TYPE, target, clientIp, strings.Join(replys, ","))
	}

	switch models.DATA_TYPE {
	case "mongodb":
		dnsInfo := models.NewMgoDnsInfo(target, client, replys)
		go dnsInfo.Insert()
	default:
		dnsInfo := models.NewDnsInfo(target, clientIp, replys)
		go dnsInfo.Insert()
	}

	w.WriteMsg(resp)
}

func isTransfer(req *dns.Msg) bool {
	for _, q := range req.Question {
		switch q.Qtype {
		case dns.TypeIXFR, dns.TypeAXFR:
			return true
		}
	}
	return false
}

func Run(ctx *cli.Context) (err error) {
	// attach request handler func
	dns.HandleFunc(".", handleDnsRequest)

	// start server
	server := LISTEN_HOST
	port := LISTEN_PORT

	tcpServer := &dns.Server{Addr: fmt.Sprintf("%v:%v", server, port), Net: "tcp"}
	udpServer := &dns.Server{Addr: fmt.Sprintf("%v:%v", server, port), Net: "udp"}

	log.Printf("Starting at udp://%v:%v\n", server, port)
	go func() {
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}()

	log.Printf("Starting at tcp://%v:%v\n", server, port)
	go func() {
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}()

	// Wait for SIGINT or SIGTERM
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	udpServer.Shutdown()
	tcpServer.Shutdown()
	return err
}
