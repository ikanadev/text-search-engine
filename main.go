package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func checkErr(e error) {
	if e != nil {
		fmt.Printf("Error: %s \n", e.Error())
		os.Exit(1)
	}
}
func main() {
	fmt.Println("Loading documents...")
	start := time.Now()
	docs, err := loadDocuments("enwiki-latest-abstract1.xml")
	if err != nil {
		checkErr(err)
	}
	fmt.Printf("Loaded %d documents in %v\n", len(docs), time.Since(start))

	fmt.Println("Indexing documents...")
	start = time.Now()
	idx := make(index)
	idx.add(docs)
	fmt.Printf("Indexed %d documents in %v\n", len(docs), time.Since(start))

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Search something:")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			fmt.Println("Bye...")
			os.Exit(0)
		}
		start = time.Now()
		mathchedIDs := idx.search(text)
		fmt.Printf("%d results found in %v\n", len(mathchedIDs), time.Since(start))

		for _, id := range mathchedIDs {
			doc := docs[id]
			fmt.Printf("%d\n%s\n%s\n\n", id, doc.Title, doc.Text)
		}
		fmt.Println("#########################################################")
		fmt.Println("Search or exit:")
	}

}
