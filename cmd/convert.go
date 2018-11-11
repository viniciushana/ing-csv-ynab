package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertRecord(record []string) ([]string, error) {
	if record[0] == "Datum" {
		return []string{"Date", "Payee", "Memo", "Outflow", "Inflow"}, nil
	}

	date := fmt.Sprintf("%s-%s-%s", record[0][0:4], record[0][4:6], record[0][6:8])
	payee := record[1]

	amount, err := strconv.ParseFloat(strings.Replace(record[6], ",", ".", -1), 32)
	if err != nil {
		return nil, fmt.Errorf("unable to parse amount: %s", err)
	}
	inAmount, outAmount := "", ""

	if inOut := record[5]; inOut == "Bij" {
		inAmount = fmt.Sprintf("%.2f", amount)
	} else {
		outAmount = fmt.Sprintf("%.2f", amount)
	}

	return []string{date, payee, payee, outAmount, inAmount}, nil
}

func convert(input io.Reader, output io.Writer) error {
	inCsv := csv.NewReader(input)
	outCsv := csv.NewWriter(output)
	defer outCsv.Flush()
	total := 0

	for {
		rec, err := inCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		conv, err := convertRecord(rec)
		if err != nil {
			return fmt.Errorf("error converting entry: %s", err)
		}

		total++
		outCsv.Write(conv)
	}

	if total < 2 {
		return fmt.Errorf("empty input file")
	}

	return nil
}

var convertCmd = &cobra.Command{
	Use:   "convert [ing csv file] [desired output file name]",
	Short: "Converts an ING CSV file to a YNAB-compliant one",
	Long: `Takes a CSV bank statement from ING, reads its data and outputs it in a format and disposition
that YNAB is able to read.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Converting file", args[0], "to", args[1])

		input, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer input.Close()

		output, err := os.Create(args[1])
		if err != nil {
			panic(err)
		}
		defer output.Close()

		err = convert(input, output)
		if err != nil {
			fmt.Println("Error converting", err)
		} else {
			fmt.Println("Done!")
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
}
