package main 

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
	"strings"
)

func verifyemailDomain(emailDomain string) {

	var hasMX, hasDMARC, hasSPF bool
	var SPFRecord, DMARCRecord string

	MXRecords, error := net.LookupMX(emailDomain)

	if error != nil {
		log.Printf("Error: %v\n", error)
	}
	
	if len(MXRecords) > 0 {
		hasMX = true
	}

	TXTRecords, error := net.LookupTXT(emailDomain)

	if error != nil {
		log.Printf("Error: %v\n", error)
	}
	
	for _, record := range TXTRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	DMARCRecords, error := net.LookupTXT("_dmarc" + emailDomain)

	if error != nil {
		log.Printf("Error: %v\n", error)
	}
	
	for _, record := range DMARCRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
		    DMARCRecord = record
			break
		}
	}

	fmt.Println("Domain: ", emailDomain)
	fmt.Println("Does the email domain have MX: ", hasMX)
	fmt.Println("Does the email domain have SPF: ", hasSPF)
	fmt.Println("Does the email domain have DMARC: ", hasDMARC)
	fmt.Println("SPFRecord: ", SPFRecord)
	fmt.Println("DMARCRecord: ", DMARCRecord)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Checking for: emailDomain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		verifyemailDomain(scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		log.Fatal("Unable to read from given input.", error)
	}
}