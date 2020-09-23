package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	/*
			if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
				fmt.Println("Log file does not exist")

			} else {
				fmt.Println("Log file exists")
			}

			contentBytes, err := ioutil.ReadFile("names.txt")
			if err == nil {
				var contentStr string = string(contentBytes)
				fmt.Println(contentStr)
			}

			stmt := "Welcome IBM"
			err1 := ioutil.WriteFile("ibm.txt", []byte(stmt), 0644)

			if err1 != nil {
				fmt.Println(err1)
			}

			fmt.Println("------------------------------------")

			file, _ := os.Open("names.txt")

			fileScanner := bufio.NewScanner(file)
			lineCount := 0

			for fileScanner.Scan() {
				lineCount++
			}

			defer file.Close()
			fmt.Println(lineCount)

			fmt.Println("---------------reading a line---------------------")

			fmt.Println(ReadLine(1))

			fmt.Println(ReadLine(3))

			fmt.Println("---------------compare files---------------------")

			one, err := ioutil.ReadFile("names.txt")

			if err != nil {
				panic(err)
			}

			two, err2 := ioutil.ReadFile("names1.txt")

			if err2 != nil {
				panic(err2)
			}

			same := bytes.Equal(one, two)

			fmt.Println(same)

			fmt.Println("-------------delete file-----------------------")

			err3 := os.Remove("log.txt")

			if err3 != nil {
				panic(err3)
			}

		fmt.Println("--------------COPY AND MOVE ----------------------")

		original, err2 := os.Open("names1.txt")

		if err2 != nil {
			panic(err2)
		}


		os.Mkdir("target", 0755)
		original_copy, err1 := os.Create("target/names1.txt")

		if err1 != nil {
			panic(err1)
		}

		defer original_copy.Close()

		_, err3 := io.Copy(original_copy, original)

		if err3 != nil {
			panic(err3)
		}

		original.Close()
		os.Remove("names1.txt")
	*/
	fmt.Println("-------------renaming file-----------------------")

	os.Rename("ibm.txt", "ibmIndia.txt")

	fmt.Println("------------reading dir------------------------")

	files, err1 := ioutil.ReadDir("simpleshape")

	if err1 != nil {
		panic(err1)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	fmt.Println("------------delete folder------------------------")

	err2 := os.RemoveAll("delfolder")

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("------------------------------------")
}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("names.txt")

	fileScanner := bufio.NewScanner(file)
	lineCount := 1

	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++

	}

	defer file.Close()
	return ""
}
