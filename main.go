package main

import (
	"HackerNewsBot"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// step 1 get file name
	fileName := os.Args[1]
	log.Println(fileName)
	// step 2 get story
	lines := GetNewsLetterDatas()
	// step 3 write in to file
	f, err := os.OpenFile(fileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for i := 0; i < len(lines); i++ {
		if _, err := f.WriteString("\n" + lines[i] + "\n"); err != nil {
			log.Println(err)
		}
	}
	fmt.Print("job finish")
}

func GetNewsLetterDatas() []string {
	c := new(HackerNewsBot.BotClient)
	/*
		var newItems []HackerNewsBot.Item
		var showItems []HackerNewsBot.Item
		var askItems []HackerNewsBot.Item
	*/
	var err error
	var topItems []HackerNewsBot.Item
	var lines []string
	topItems, err = getItemsByType(HackerNewsBot.Top, c)
	if err != nil {
		log.Println(err)
	}
	toplines := prepareNewsLetterText(topItems)
	for _, val := range toplines {
		lines = append(lines, val)
	}
	/*
		showItems, err = getItemsByType(HackerNewsBot.Show, c)
		if err != nil {
			log.Println(err)
		}
		showlines := prepareNewsLetterText(showItems)
		for _, val := range showlines {
			lines = append(lines, val)
		}
		askItems, err = getItemsByType(HackerNewsBot.Ask, c)
		if err != nil {
			log.Println(err)
		}
		asklines := prepareNewsLetterText(askItems)
		for _, val := range asklines {
			lines = append(lines, val)
		}
	*/
	return lines
}

func getItemsByType(typeCode string, client *HackerNewsBot.BotClient) ([]HackerNewsBot.Item, error) {
	collection, err := client.GetStors(HackerNewsBot.Top)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var itemList []HackerNewsBot.Item
	for i := 0; i < 30; i++ {
		itemId := (*collection)[i]
		story, err := client.GetItem(itemId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		itemList = append(itemList, *story)
	}
	return itemList, nil
}

func prepareNewsLetterText(items []HackerNewsBot.Item) []string {
	var l []string
	for i := 0; i < len(items); i++ {
		l = append(l, strings.Join([]string{"##", strconv.Itoa(i+1) + ".", items[i].Title}, " "))
		l = append(l, items[i].By+":"+items[i].Text)
		l = append(l, "[原文链接]("+items[i].URL+")")
	}
	return l
}
