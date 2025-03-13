/***
XML parsing examples,

***/

package main

import (
	"fmt"
	"os"
	"io"
	"strings" // to parse to io.Reader
	"encoding/xml"
	"encoding/json"
)

// for ex1.xml,
type Address struct {
	XMLName xml.Name `xml:"address"`
	Address string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`	
}

func ex1() {
	file, _ := os.Open("ex1.xml");
	body, _ := io.ReadAll(file);
	file.Close();
	decoder := xml.NewDecoder(strings.NewReader(string(body)));
	var address Address;
	err := decoder.Decode(&address); if err != nil {
		fmt.Println(err);
		return
	}
	fmt.Println("address : ", address.Address);
	fmt.Println("Type : ", address.Type);
}

// for ex2.xml,
type Cpe struct {
	XMLName xml.Name `xml:"cpe"`
	Value string `xml:",chardata"`
}

func ex2() {
	file, _ := os.Open("ex2.xml");
	body, _ := io.ReadAll(file);
	file.Close();                                              
	decoder := xml.NewDecoder(strings.NewReader(string(body)));
	var cpe Cpe;
	err := decoder.Decode(&cpe); if err != nil {
		fmt.Println(err);
		return
	}
	fmt.Println("cbe : ",cpe.Value); 
}

type Host struct {
	XMLName xml.Name `xml:"host"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Address Address `xml:"address"`
}

// for ex3.xml,
type Hostname struct {
	XMLName xml.Name `xml:"hostname"`
	Name []string `xml:"name,attr"`
}

func main() {
	file, _ := os.Open("ex3.xml");
	body, _ := io.ReadAll(file);
	file.Close();                                              
	decoder := xml.NewDecoder(strings.NewReader(string(body)));
	for {
		var host Host;
		err := decoder.Decode(&host); if err != nil {
			fmt.Println(err);
			break	
		}
		jsonData, _ := json.Marshal(host);
		fmt.Println(string(jsonData));
	}



}


