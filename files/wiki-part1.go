/***
Exercise 12.4 : Write an article with "Title" and save it.
Then refer to an article by "Title" and read it .
***/

package main

import (
	"bufio"
	"log"
	"fmt"
	"os"
)

type Page struct {
	Title string;
	Body []byte;
}

func saveArticle(title *Page) {

	var input string;

	file, err := os.OpenFile(title.Title, os.O_WRONLY| os.O_CREATE, 0666);
	if err != nil {
		fmt.Println("cannot create file in path");
		return
	}

	writer := bufio.NewWriter(file);

	writer.WriteString();
	writer.Flush();

}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please specify the name of the article");
	}

	var article = Page {
									Title: os.Args[1],
								}

	saveArticle(&article);
}
