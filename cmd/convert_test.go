package cmd

import (
	"bytes"
	"reflect"
	"testing"
)

func TestConvertRecordIncomingAmount(t *testing.T) {
	result, err := convertRecord([]string{
		"20180903",
		"SUPERDRY BATAVIA 356 BATAVIA NLD",
		"NL64INGB0000000000",
		"",
		"BA",
		"Af",
		"46,50",
		"Betaalautomaat",
		"Pasvolgnr:001 02-09-2018 13:14 Transactie:K12345 Term:00000000",
	})

	expected := []string{
		"2018-09-03",
		"SUPERDRY BATAVIA 356 BATAVIA NLD",
		"SUPERDRY BATAVIA 356 BATAVIA NLD",
		"46.50",
		"",
	}

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %s, got: %s", expected, result)
	}
}

func TestConvertRecordOutgoingAmount(t *testing.T) {
	result, err := convertRecord([]string{
		"20180820",
		"AAB RETAIL INZ TIKKIE",
		"NL64INGB9999999999",
		"NL18ABNA0000000000",
		"OV", "Bij",
		"34,00",
		"Overschrijving",
		"Naam: AAB RETAIL INZ TIKKIE Omschrijving: Tikkie ID 000000000000, Sumo, Van SOMEONE, NL18ABNA0000000000 IBAN: NL61ABNA9999999999 Kenmerk: 666",
	})

	expected := []string{
		"2018-08-20",
		"AAB RETAIL INZ TIKKIE",
		"AAB RETAIL INZ TIKKIE",
		"",
		"34.00",
	}

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %s, got: %s", expected, result)
	}
}

func TestConvertWithEmptyFile(t *testing.T) {
	input, output := bytes.NewBufferString(""), bytes.NewBufferString("")
	err := convert(input, output)

	if err.Error() != "empty input file" {
		t.Error(err)
	}
}

func TestConvertWithUsualFile(t *testing.T) {
	input, output := bytes.NewBufferString("\"Datum\",\"Naam / Omschrijving\",\"Rekening\",\"Tegenrekening\",\"Code\",\"Af Bij\",\"Bedrag (EUR)\",\"MutatieSoort\",\"Mededelingen\"\n\"20180903\",\"SUPERDRY BATAVIA 356 BATAVIA NLD\",\"NL64INGB9999999999\",\"\",\"BA\",\"Af\",\"46,50\",\"Betaalautomaat\",\"Pasvolgnr:001 02-09-2018 13:14 Transactie:K12345 Term:00000000\"\n\"20180820\",\"AAB RETAIL INZ TIKKIE\",\"NL64INGB9999999999\",\"NL13ABNA0000000000\",\"OV\",\"Bij\",\"34,00\",\"Overschrijving\",\"Naam: AAB RETAIL INZ TIKKIE Omschrijving: Tikkie ID 000000000000, Sumo, Van SOMEONE, NL27INGB9999999999 IBAN: NL13ABNA9999999999 Kenmerk: 123\"\n"), bytes.NewBufferString("")
	err := convert(input, output)

	if err != nil {
		t.Error(err)
	}

	if res := output.String(); res != "Date,Payee,Memo,Outflow,Inflow\n2018-09-03,SUPERDRY BATAVIA 356 BATAVIA NLD,SUPERDRY BATAVIA 356 BATAVIA NLD,46.50,\n2018-08-20,AAB RETAIL INZ TIKKIE,AAB RETAIL INZ TIKKIE,,34.00\n" {
		t.Errorf("unexpected result: %s", res)
	}
}
