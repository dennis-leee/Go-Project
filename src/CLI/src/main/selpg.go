package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

type Args struct {
	startPage int
	endPage int
	pageLength int
	formFeed bool
	inputFile string
	printerDestination string
}

func main() {
	args := new(Args)
	flag.IntVar(&args.startPage, "s", -1, "The start page of the extracted page range")
	flag.IntVar(&args.endPage, "e", -1, "The end page of the extracted page range")
	flag.IntVar(&args.pageLength, "l", 72, "The number of rows per page")
	flag.BoolVar(&args.formFeed, "f", false, "Whether used form-feed character to delimit the pages")
	flag.StringVar(&args.printerDestination, "d", "", "The name of the printer to use")
	flag.Parse()
	//var file= flag.Args()

	//fmt.Println("startPage has value ", args.startPage)
	//fmt.Println("endPage has value ", args.endPage)
	//fmt.Println("length has value ", args.pageLength)
	//fmt.Println("form-feed has value ", args.formFeed)
	//fmt.Println("printerDestination has value ", args.printerDestination)
	//fmt.Println("the number of args ", flag.NArg())
	//fmt.Println("the number of args ", len(flag.Args()))
	//fmt.Println("the name or content of file ", file)

	parseArgs(args)
}

func parseArgs(args *Args) {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("selpg [-s start page(>= 1)] [-e end page(>= s)] [-l page length] [-f use form-feed (l and f can only choose one)] [-d printer destination] [input file]")
		flag.PrintDefaults()
	}

	if (args.startPage == -1 || args.endPage == -1 || args.startPage > args.endPage || args.startPage < 1 || args.endPage < 1) ||
		(args.pageLength != 72 && args.formFeed) || (flag.NArg() > 1) {
		flag.Usage()
		return
	}

	if flag.NArg() == 1 {
		args.inputFile = flag.Args()[0]
	}

	selpg(args)
}

func selpg(args *Args) {
	curPage := 1
	curLines := 0
	line := new(bufio.Scanner)

	//判断从文件读输入还是从标准输入读输入
	if args.inputFile != "" {
		fin, err := os.OpenFile(args.inputFile, os.O_RDONLY, os.ModeType)
		defer fin.Close()
		if err != nil {
			panic(err)
			return
		}
		line = bufio.NewScanner(fin)    //从文件读
	} else {
		line = bufio.NewScanner(os.Stdin)     //从标准输入读
	}

	//根据分页模式进行分页
	if !args.formFeed { //根据给定行数分页
		for line.Scan() {
			if curPage > args.endPage {
				break
			}
			if curPage >= args.startPage {
				os.Stdout.Write([]byte(line.Text() + "\n"))
			}
			curLines++
			if curLines %= args.pageLength; curLines == 0 {
				curPage++
			}
		}
	} else {      //根据换页符分页
		lineEnd := false
		for line.Scan() {
			for _,c := range line.Text() {
				if lineEnd {      //遇到换页符后页数加1
					curPage++
					lineEnd = false
				}

				if c == '\f' {
					lineEnd = true
					if curPage >= args.startPage {
						os.Stdout.Write([]byte("\n"))
					}
					//末页刚好填满，内容也刚好结束
					if curPage == args.endPage {
						break
					}
				} else {
					if curPage >= args.startPage {
						os.Stdout.Write([]byte(string(c)))
					}
				}
			}
		}
		//末页未填满，但是内容已结束
		if curPage == args.endPage && !lineEnd {
			os.Stdout.Write([]byte("\n"))
		}
	}

	//判断文件是否有endPage页
	if curPage < args.endPage {
		fmt.Fprint(os.Stderr, "This file doesn't have that many pages!\n")
	}
}