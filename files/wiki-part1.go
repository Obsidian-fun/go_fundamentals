/***
Exercise 12.4 : Write an article with "Title" and save it.
Then refer to an article by "Title" and read it .
***/

package main

import (
	"bufio"
	"log"
	"fmt"
)

type Page struct {
	Title string;
	Body []byte;
}

func (p *Page) saveArticle(title *Page) {

	var input string;

	file, err := os.OpenFile(title, os.O_WRONLY| os.O_CREATE, 0666);
	if err != nil {
		fmt.Println("cannot create file in path");
		return
	}

	writer := bufio.NewWriter(file);

}

func main() {]
	if len(os.Args) != 2 {
		log.Fatal("please specify the name of the article");
	}

	var article = Page {
									Title: os.Args[1];
								}

	saveArticle(&article);
}
