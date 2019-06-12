package tokenizers

import (
	"log"
	"strings"
	"testing"
	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis/analyzer/custom"
)

// FileIndexer is a data structure to hold the content of the file
type FileIndexer struct {
	FileName    string
	FileContent string
}

type FileIndexerArray struct {
	IndexerArray []FileIndexer
}

func panicMust(e error, logstring string) {
	if e != nil && !strings.Contains(e.Error(), "lost+found") {
		log.Printf(e.Error())
		panic(e)
	}
	log.Printf(logstring + "\n")
}

func indexedSearch(b *testing.B, indexFilename string, searchWord string) *bleve.SearchResult {
	b.StopTimer()
	//searchWord2 := parseSearch(searchWord)
	// opens the index file using bleve
	index, err := bleve.Open(indexFilename)
	if err != nil {
		panic(err)
	}
	// closes file after the function completes its execution
	defer func() {
		log.Printf("cleanup: index.Close()")
		index.Close()
	}()

	b.StartTimer()
	// makes query to search the string
	query := bleve.NewQueryStringQuery(searchWord)
	request := bleve.NewSearchRequestOptions(query, 50, 0, false)
	// matches the keyword if any from the index created and returns
	result, _ := index.Search(request)
	b.StopTimer()
	return result
}

func BenchmarkLoadDictionary1(b *testing.B)  {
	indexedSearch(b,"./root.bleve", "肉丝")
}

func BenchmarkLoadDictionary10(b *testing.B)  {
	keywords := []string{
		"肉丝",
		"黑丝",
		"灰丝",
		"白丝",
		"鱼香肉丝",
		"中国",
		"美国",
		"粉红",
		"旗袍",
		"写真",
		"美腿",
	}
	for _, k := range keywords {
		indexedSearch(b,"./root.bleve", k)
	}
}