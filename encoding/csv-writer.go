/***

Writing Files into a CSV.

***/

package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/csv"
)

type Country struct {
	Code string;
	Name string;
	Population int;
}

func main() {
	const million = 1000000;

	records := []Country {
								{Code: "IT", Name: "Italy", Population: 65*million},
								{Code: "ES", Name: "Spain", Population: 70*million},
								{Code: "FR", Name: "France", Population: 110*million},
								{Code: "PO", Name: "Portugal", Population: 40*million},
								{Code: "BE", Name: "Belize", Population: 20*million},
						 }

	w := csv.NewWriter(os.Stdout);
	defer w.Flush();

	for _, r := range records {
		if err := w.Write([]string{r.Code, r.Name, strconv.Itoa(r.Population)}); err != nil {
			fmt.Println("error: ", err);
			return;
		}
	}
}


