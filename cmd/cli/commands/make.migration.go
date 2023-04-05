/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// pingCmd represents the ping command
var makeMigration = &cobra.Command{
	Use:   "make:migration",
	Short: "Command for generate file migration",
	Long:  `Command for generate file migration in database/migrations folder`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := os.Getwd()

		fmt.Println(path)

		fmt.Print("Name of migration: ")

		var response string

		_, err := fmt.Scanln(&response)
		if err != nil {
			log.Fatal(err)
		}

		if response == "" {
			fmt.Println("Empty name. Close")
			return
		}

		now := time.Now().UnixNano() / int64(time.Millisecond)

		fileName := "database/migrations/" + response + "_" + strconv.Itoa(int(now)) + ".go"

		data := `
package migrations

import (
 "gorm.io/gorm"
)

func M_{response}_{time}(db *gorm.DB) bool {

 return true
}
 `
		data = strings.Replace(data, "{time}", strconv.Itoa(int(now)), -1)
		data = strings.Replace(data, "{response}", response, -1)
		fmt.Println(data)
		err = ioutil.WriteFile(fileName, []byte(data), 0666)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(makeMigration)
}
