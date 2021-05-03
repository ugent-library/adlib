package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugent-library/go-adlib/adlib"
)

func main() {
	rootCmd.AddCommand(convertCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert Adlib XML records to JSON",
	Run: func(cmd *cobra.Command, args []string) {
		dec := adlib.NewGroupedXMLDecoder(os.Stdin)
		enc := json.NewEncoder(os.Stdout)

		for {
			rec, err := dec.Decode()
			if err != nil {
				log.Panic(err)
			}
			if rec == nil {
				break
			}
			err = enc.Encode(rec)
			if err != nil {
				log.Panic(err)
			}
		}
	},
}
