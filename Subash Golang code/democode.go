package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type GoAssignment struct {
}

type ListOfMap map[string]interface{}

func main() {
	fmt.Printf("#########################################################")
	fmt.Printf("\nStarted the execution @ %v \n", time.Now())
	var rawdata = GoAssignment{}.loadRawDataList()
	GoAssignment{}.doTheWork(rawdata)
	fmt.Printf(" \n ")
	fmt.Printf("Finished the execution @ %v ", time.Now())
}

func (ga GoAssignment) doTheWork(rawDataList []string) {
	fmt.Println("Coming inside doTheWork .... ")

	//fmt.Println("Input Data is :  \n ", rawDataList)

	// IMPLEMENT YOUR CODE LOGIC HERE .....
	ipData := make(map[string][2]float64)
	for _, data := range rawDataList {
		ipAdd, pageload := parsefields(data)
		if ipAdd != "" && pageload > 0 {
			ipstats := ipData[ipAdd]
			ipstats[0] += pageload
			ipstats[1]++
			ipData[ipAdd] = ipstats
		}
	}
	uniqueIps := getsorted(ipData)
	fmt.Printf("\nUnique Ip Address:%d\n", len(uniqueIps))
	for _, ip := range uniqueIps {
		stats := ipData[ip]
		if stats[1] > 0 {
			fmt.Printf("Ip Address: %s,Average PageLoadTime:%2f ms \n", ip, stats[0]/stats[1])
		}
	}
}

func parsefields(data string) (string, float64) {
	var ipAddress string
	var pageloadtime float64
	for _, field := range strings.Split(data, "~$~") {
		switch {
		case strings.Contains(field, "$#$ip#$#"):
			parts := strings.Split(field, "$#$ip#$#")
			if len(parts) > 1 {
				ipAddress = parts[1]
			}
		case strings.Contains(field, "PageLoadTime#$#"):
			parts := strings.Split(field, "PageLoadTime#$#")
			if len(parts) > 1 {
				pageloadtime = parseToFloat(parts[1])
			}
		}
	}
	return ipAddress, pageloadtime
}
func parseToFloat(value string) float64 {
	if result, err := strconv.ParseFloat(value, 64); err == nil {
		return result
	}
	return 0.0
}
func getsorted(ipData map[string][2]float64) []string {
	uniqueip := make([]string, 0, len(ipData))
	for ip := range ipData {
		uniqueip = append(uniqueip, ip)
	}
	sort.Strings(uniqueip)
	return uniqueip
}

func (ga GoAssignment) loadRawDataList() []string {
	var al []string
	al = append(al, "a121a4cb-8d51-4622-923c-3755c80b51b8$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:10~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.0; en-us; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#1748~$~FirstbyteTime#$#1500~$~ServerConnectionTime#$#169~$~ResponseAvailableTime#$#1331~$~FrontEndTime#$#248~$~DocumentReadyTime#$#143~$~DocumentDownloadTime#$#74~$~DocumentProcessingTime#$#69~$~PageRenderTime#$#105~$~DNSLookupTime#$#1~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#~$~pageType#$#IFrame")
	al = append(al, "bc940d03-ef43-4a43-b7d6-9834a49a30f5$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:10~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.0; en-us; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#1522~$~FirstbyteTime#$#1306~$~ServerConnectionTime#$#169~$~ResponseAvailableTime#$#1137~$~FrontEndTime#$#216~$~DocumentReadyTime#$#125~$~DocumentDownloadTime#$#65~$~DocumentProcessingTime#$#60~$~PageRenderTime#$#91~$~DNSLookupTime#$#1~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#~$~pageType#$#IFrame")
	al = append(al, "17440ba6-71d2-4107-98a8-08e16175d7db$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Android~$~device#$#Mobile~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 15:17:48~$~userAgent#$#Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36~$~referrer#$#http://192.168.11.121:8780/corebanking/retail/thirdpartytransfer.jsp~$~errorCount#$#~$~PageLoadTime#$#2938~$~FirstbyteTime#$#2521~$~ServerConnectionTime#$#170~$~ResponseAvailableTime#$#2351~$~FrontEndTime#$#417~$~DocumentReadyTime#$#240~$~DocumentDownloadTime#$#125~$~DocumentProcessingTime#$#115~$~PageRenderTime#$#177~$~DNSLookupTime#$#2~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/loginaction.jsp~$~error#$#~$~pageType#$#Page")
	al = append(al, "7d732744-a24c-4355-b634-68504af4010d$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:14~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#1350~$~FirstbyteTime#$#1158~$~ServerConnectionTime#$#169~$~ResponseAvailableTime#$#989~$~FrontEndTime#$#192~$~DocumentReadyTime#$#110~$~DocumentDownloadTime#$#57~$~DocumentProcessingTime#$#53~$~PageRenderTime#$#82~$~DNSLookupTime#$#1~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#~$~pageType#$#IFrame")
	al = append(al, "e57e7965-5aab-4631-8721-d58b8d6ea0b5$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:14~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#2611~$~FirstbyteTime#$#2240~$~ServerConnectionTime#$#170~$~ResponseAvailableTime#$#2070~$~FrontEndTime#$#371~$~DocumentReadyTime#$#213~$~DocumentDownloadTime#$#111~$~DocumentProcessingTime#$#102~$~PageRenderTime#$#158~$~DNSLookupTime#$#2~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#~$~pageType#$#IFrame")
	al = append(al, "703f7517-b39d-4c41-91c0-9a0dbd0484b7$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:15~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#1128~$~FirstbyteTime#$#968~$~ServerConnectionTime#$#169~$~ResponseAvailableTime#$#799~$~FrontEndTime#$#160~$~DocumentReadyTime#$#92~$~DocumentDownloadTime#$#48~$~DocumentProcessingTime#$#44~$~PageRenderTime#$#68~$~DNSLookupTime#$#1~$~TCPConnectTime#$#1~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#~$~pageType#$#IFrame")
	al = append(al, "f3e6c4dd-3689-4d72-a7af-c807b320cbfb$#$ip#$#192.168.11.31~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2015-01-19 17:56:33~$~userAgent#$#Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; WOW64; Trident/6.0)~$~referrer#$#http://192.168.11.121:8780/corebanking/retail/thirdpartytransfer.jsp~$~errorCount#$#~$~FirstbyteTime#$#1689~$~DocumentDownloadTime#$#84~$~DocumentProcessingTime#$#77~$~PageLoadTime#$#1968~$~ServerConnectionTime#$#169~$~ResponseAvailableTime#$#1520~$~FrontEndTime#$#279~$~DocumentReadyTime#$#161~$~PageRenderTime#$#118~$~pageType#$#Ajax~$~url#$#http://192.168.11.121:8780/corebanking/retail/feesandcharges.jsp~$~error#$#")
	al = append(al, "d8aa77be-769d-4755-b2d5-26d01be6ff72$#$ip#$#190.25.228.161~$~sessionId#$#~$~os#$#Windows~$~device#$#Desktop~$~browserDetails#$#Internet Explorer 10.0~$~country#$#Malaysia~$~country_code#$#MY~$~region#$#Selangor~$~city#$#Kuala Lumpur~$~latitude#$#80.257616~$~longitude#$#12.968093~$~timeStr#$#2014-12-15 17:02:15~$~userAgent#$#Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30~$~referrer#$#~$~errorCount#$#~$~PageLoadTime#$#4522~$~FirstbyteTime#$#3880~$~ServerConnectionTime#$#172~$~ResponseAvailableTime#$#3708~$~FrontEndTime#$#642~$~DocumentReadyTime#$#369~$~DocumentDownloadTime#$#192~$~DocumentProcessingTime#$#177~$~PageRenderTime#$#273~$~DNSLookupTime#$#3~$~TCPConnectTime#$#2~$~url#$#http://192.168.11.121:8780/corebanking/retail/interbanktransfer.jsp~$~error#$#[Uncaught TypeError: Cannot read property ´tostring´ of null,http://192.168.11.121:8780/corebanking/script/ek.validation.js,249]~$~pageType#$#IFrame")

	return al
}
