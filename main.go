package main

import (
	"bufio"
	"fmt"
	"github.com/adrianmo/go-nmea"
	"log"
	"net"
	"os"
	"strings"
)

const (
	serverport	string	= "192.168.31.5:5555"
)

func main() {
	// connect to gps server
	conn, err := net.Dial("tcp", serverport)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		// wait for reply
		gpstext := bufio.NewScanner(conn)
		gpstext.Split(bufio.ScanLines)
		for gpstext.Scan() {
			gpstextline := gpstext.Text()
			//fmt.Println(gpstextline)
			if strings.HasPrefix(gpstextline, "$") {
				gpsa(gpstextline)
			}
		}
	}
}

func gpsa(gpstextline string) {
	// gps data
	s, err := nmea.Parse(gpstextline)
	if err != nil {
		log.Fatalln(err)
	}
	if s.DataType() == nmea.TypeGGA {
		m := s.(nmea.GGA)
		fmt.Printf("\033[0;0H")
		fmt.Printf(" \n")
		fmt.Printf("  UTC Time: %s\n", m.Time)
		fmt.Printf("  Sate num: %d\n", m.NumSatellites)
		fmt.Printf("  Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
		fmt.Printf("  Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))

	}
}