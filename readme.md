# Search text engine
A implementation of a search text engine from [this tutorial](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/).

### Steps to run
1. Download this [file](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz) and extract it in the repo directory, It is a kind of dictionary xml file from wikipedia.
2. run `go mod install` to install the only dependence we need.
3. run `go build` to build the binary
4. execute the binary `./text-search-engine`
5. You can search as many sentences you want.
6. Type `exit` to... exit