package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numberOfMonitoring = 2
const delayTime = 2

func main() {

	showIntro()

	for {
		selectedOption := showMenuAndReadSelectedOption()

		switchCaseFunc(selectedOption)
	}

}

func showIntro() {

	fmt.Println("Enter with your name")
	var name string
	fmt.Scan(&name)
	fmt.Println()
	var version float32 = 1.1
	fmt.Println("Welcome Mrs.", name)
	fmt.Println("This program is running on version", version)
	fmt.Println()

}

func showMenuAndReadSelectedOption() int8 {

	fmt.Println("1 - Init Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Leave")
	fmt.Println()

	var selectedOption int8
	// 8 bits (signed integer, from -127 to +127).
	//fmt.Scanf("%d", &selectedOption) // its one way to do it

	fmt.Scan(&selectedOption)
	// another way of doing the same thing as above but with a pointer

	fmt.Println("The address of my variable is", &selectedOption)
	fmt.Println("The selected option was", selectedOption)

	return selectedOption
}

func switchCaseFunc(selectedOption int8) {

	switch selectedOption {
	case 1:
		fmt.Println()
		fmt.Println("Starting Monitoring ...")
		initMonitoring()
	case 2:
		fmt.Println()
		fmt.Println("Showing Logs ...")
		readLogs()
	case 0:
		fmt.Println("Leaving Monitoring. See You Later.")
		os.Exit(0)
	default:
		fmt.Println("This Command Was Not Recongnized")
		os.Exit(-1)
	}
}

func initMonitoring() {
	fmt.Println("Monitoring Started. Please wait.")
	fmt.Println()

	urls := readUrlsFromFile()
	// Testing n times the websites in the slice
	for i := 0; i < numberOfMonitoring; i++ {

		for i, url := range urls {
			fmt.Println("Testing site", (i + 1), ":", url)
			testingUrls(url)
		}

		time.Sleep(delayTime * time.Second)
		fmt.Println()
	}

}

func testingUrls(urlWeb string) {
	response, err := http.Get(urlWeb)

	if err != nil {
		fmt.Print("The website", urlWeb, "had some errors:", err.Error())
		registerLogs(urlWeb, false)
	} else if response.StatusCode == 200 {
		fmt.Println("The website:", urlWeb, "was loaded successfully!")
		registerLogs(urlWeb, true)
	} else {
		fmt.Println("The website:", urlWeb, "has some issues. Status:", response.Status)
		registerLogs(urlWeb, false)
	}
}

func readUrlsFromFile() []string {
	var urls []string

	//file, err := os.ReadFile("websites.txt")

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("An error occured", string(err.Error()))
	}

	reader := bufio.NewReader(file)

	for {

		url, errReader := reader.ReadString('\n')

		if errReader != nil {
			fmt.Println("An error occured", err)
		}

		url = strings.TrimSpace(url)

		urls = append(urls, url)

		if errReader == io.EOF {
			fmt.Println()
			break
		}

	}

	file.Close()

	return urls
}

func registerLogs(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error register logs", err.Error())
	}

	//Time function uses some numbers to specify the regex of time formatting
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func readLogs() {
	file, err := os.OpenFile("logs.txt", os.O_RDWR, 0666)

	if err != nil {
		fmt.Println("Error reading logs", err.Error())
	}

	reader := bufio.NewReader(file)

	for {
		url, rErr := reader.ReadString('\n')

		if rErr == io.EOF {
			fmt.Println()
			break
		}

		url = strings.TrimSpace(url)

		fmt.Println(url)

	}

	file.Close()
}

func testDoubleReturnFunction() (string, int) {
	name := "Raveline"
	age := 18

	return name, age
}

func ifCaseFunc(selectedOption int8) {
	if selectedOption == 1 {
		fmt.Println("Starting Monitoring ...")
	} else if selectedOption == 2 {
		fmt.Println("Showing Logs ...")
	} else if selectedOption == 0 {
		fmt.Println("Leaving Monitoring. See You Later.")
		os.Exit(0)
	} else {
		fmt.Println("This Command Was Not Recongnized")
		os.Exit(-1)
	}
}
