
package main

import (
	"fmt"
	"encoding/xml"
	"io"
	"net/http"
)

type Status struct{
	Text string;
}

type User struct{
	XMLName xml.Name;
	Status Status;
}

func main() {
	resp, err := http.Get("http://twitter.com/users/Googland.xml");
	if err != nil {
		fmt.Println("error fetching", err);
		return
	}

	data, err := io.ReadAll(resp.Body);
	if err != nil {
		fmt.Println("cannot read file");
		return
	}
	defer resp.Body.Close();

	user := User{
						xml.Name{"","user"},
						Status{""},
					}

	err = xml.Unmarshal(data, &user); if err != nil {
			fmt.Println("can't retreive, xml structure changed");
			return
	}
	fmt.Printf("%s",user.Status.Text);

}

