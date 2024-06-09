/***
Exercise 12.4 : Write an article with "Title" and save it.
Then refer to an article by "Title" and read it .
***/

package main

import (
	"bufio"
	"strings"
	"log"
	"fmt"
	"os"
)

type Page struct {
	Title string;
	Body []byte;
}

func saveArticle(page *Page) {

// first, read from user input
	input := bufio.NewReader(os.Stdin);

	var lines string;
	for {
		line := input.ReadString('\n');

		if strings.TrimSpace(line) == "quit" {
			break;
		}

		lines = lines + line + "\n";
	}

// then write user input into a file,
	file, err := os.OpenFile(page.Title, os.O_WRONLY| os.O_CREATE, 0666);
	if err != nil {
		fmt.Println("cannot create file in path");
		return
	}

	writer := bufio.NewWriter(file);

	writer.WriteString(lines);
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

/*
	fmt.Println("Choose your option [A] Read an article [B] Write an article\n\n");
	var input string;
	fmt.Scan(&input);

	for {
		switch input
		case A:
	}
*/
}
