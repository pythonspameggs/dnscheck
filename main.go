package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(os.Stderr, "Usage: dnscheck <DOMAIN> <DNS RECORD TYPES>")
		os.Exit(1)
	}
	domain := os.Args[1]
	record := os.Args[2]
	fmt.Printf("### domain: %s. dns record types: %s ###\n", domain, record)

	switch record {
	case "cname", "CNAME":
		cname(domain)
	case "txt", "TXT":
		txt(domain)
	case "mx", "MX":
		mx(domain)
	case "ns", "NS":
		ns(domain)
	case "srv", "SRV":
		srv(domain)
	default:
		ip(domain)
	}
}

func cname(host string) {
	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Println("CNAME Record error: ", err.Error())
	} else {
		fmt.Printf("%s IN CNAME %s\n", host, cname)
	}
}

func txt(name string) {
	txts, err := net.LookupTXT(name)
	if err != nil {
		fmt.Println("TXT Record error: ", err.Error())
	} else if len(txts) == 0 {
		fmt.Printf("No TXT Record: %s\n", name)
	} else {
		for _, t := range txts {
			fmt.Printf("%s IN TXT %s\n", name, t)
		}
	}
}

func ip(host string) {
	ips, err := net.LookupIP(host)
	if err != nil {
		fmt.Println("A Record error: ", err.Error())
	} else if len(ips) == 0 {
		fmt.Printf("No A Record: %s\n", host)
	} else {
		for _, ip := range ips {
			fmt.Printf("%s IN A %s\n", host, ip)
		}
	}
}

func mx(name string) {
	mxs, err := net.LookupMX(name)
	if err != nil {
		fmt.Println("MX Record error: ", err.Error())
	} else {
		for _, mx := range mxs {
			fmt.Printf("%s IN MX %d %s\n", name, mx.Pref, mx.Host)
		}
	}
}

func ns(name string) {
	nss, err := net.LookupNS(name)
	if err != nil {
		fmt.Println("NS Record error: ", err.Error())
	} else if len(nss) == 0 {
		fmt.Printf("No NS Record: %s\n", name)
	} else {
		for _, ns := range nss {
			fmt.Printf("%s IN NS %s\n", name, ns.Host)
		}
	}
}

func srv(name string) {
	_, addrs, err := net.LookupSRV("", "", name)
	if err != nil {
		fmt.Printf("No SRV Record: %s\n", name)
	} else {
		for _, addr := range addrs {
			fmt.Printf("%s IN SRV %d %d %d %s\n", name, addr.Priority, addr.Weight, addr.Port, addr.Target)
		}
	}
}
