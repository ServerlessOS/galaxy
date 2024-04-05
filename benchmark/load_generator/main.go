package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// small test bed 10% ~ 90% mem，由于最大内存为4GB，将内存缩小到10%
// light functions 10 ~ 25 %
// heavy function 75% ~ 90 %
func main() {
	// 打开文件
	file, err := os.Open("./data/memory.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	memorys := make([]int, 0)
	// 使用 bufio.Scanner 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m, _ := strconv.Atoi(line)
		memorys = append(memorys, m)
	}
	// 检查是否有读取错误
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
	}
	println(len(memorys))
	origin_func := Percentile(memorys, 0.1, 0.9)
	_ = WriteArrayToFile("./data/origin", origin_func, 0.1)
	light_func := Percentile(memorys, 0.1, 0.25)
	_ = WriteArrayToFile("./data/light_func", light_func, 1)
	heavy_func := Percentile(memorys, 0.75, 0.9)
	_ = WriteArrayToFile("./data/heavy_func", heavy_func, 1)

	println(len(Percentile(memorys, 0.1, 0.9)))
	println(len(Percentile(memorys, 0.1, 0.25)))
	println(len(Percentile(memorys, 0.75, 0.9)))

}

// Percentile 返回数组在指定百分位范围内的切片
func Percentile(arr []int, lowerPercent float64, upperPercent float64) []int {
	// 对数组进行排序
	sort.Ints(arr)

	// 计算百分位位置
	lowerIndex := int(float64(len(arr)-1) * lowerPercent)
	upperIndex := int(float64(len(arr)-1) * upperPercent)

	// 截取百分位范围内的切片
	result := arr[lowerIndex : upperIndex+1]

	return result
}

func WriteArrayToFile(filename string, arr []int, factor float64) error {
	// 打开文件，如果文件不存在则创建
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将数组中的值逐行写入文件
	for _, value := range arr {
		_, err := fmt.Fprintln(file, int(float64(value)*factor))
		if err != nil {
			return err
		}
	}

	return nil
}
