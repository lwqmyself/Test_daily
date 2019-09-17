package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {

	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	//标准的稀疏数组应该还有一个记录元素的二维数组的规模
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}

	}

	for i, valNode := range sparseArr {
		fmt.Printf("%d : %d %d %d\n", i, valNode.row, valNode.col, valNode.val)

	}

	var chessMap2 [11][11]int
	for i, valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}

	}
	f, err := os.OpenFile("ww2.txt", os.O_CREATE, os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	for i, valNode := range sparseArr {
		str := fmt.Sprintf("%d : %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
		f.WriteString(str)

	}
	/*for i := 1; i<=10000000;i++  {
		str := fmt.Sprintf("这是第%d行数据\n",i)
		f.WriteString(str)
	}*/

	fmt.Println(f.Name())
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadBytes('\n')
		//line = strings.TrimSpace(string(line))

		fmt.Print(string(line))
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}

	name := "1234百万红包"
	words := ([]rune)(name)
	var inStr string
	for i := 0; i < len(words); i++ {
		fmt.Println(string(words[i]))
		inStr = inStr + string(words[i:i+1]) + ","
	}
	fmt.Println(inStr)

}
