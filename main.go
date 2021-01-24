package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/urfave/cli/v2"
)

type JisyoCSV struct {
	Yomi string `csv:"yomi"`
	Word string `csv:"word"`
	Note string `csv:"note"`
}

type GoogleContactsCSV struct {
	GivenName       string `csv:"Given Name"`
	FamilyName      string `csv:"Family Name"`
	GivenNameYomi   string `csv:"Given Name Yomi"`
	GroupMembership string `csv:"Group Membership"`
}

func main() {
	app := &cli.App{
		Name:    "syosyo",
		Usage:   "generate jisyo file",
		Version: "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Value:   "skk",
				Usage:   "jisyo file format (\"skk\", \"contacts\")",
				EnvVars: []string{"SYOSYO_FORMAT"},
			},
		},
		Action: generateJisyo,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func generateJisyo(c *cli.Context) error {
	// input: <SYOSYO_JISYO_NAME>.csv
	args := c.Args().Slice()
	var jisyoName string
	jisyoNameEnv := os.Getenv("SYOSYO_JISYO_NAME")
	if jisyoNameEnv != "" {
		jisyoName = jisyoNameEnv
	} else {
		if len(args) == 0 {
			return fmt.Errorf("no jisyo name")
		}
		jisyoName = args[0]
	}
	jisyoCSVFileName := jisyoName + ".csv"

	f, err := os.OpenFile(filepath.Join("csv", jisyoCSVFileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer f.Close()

	jisyoRows := []*JisyoCSV{}
	err = gocsv.UnmarshalFile(f, &jisyoRows)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	switch c.String("format") {
	case "skk":
		// output: SKK-JISYO-<SYOSYO_JISYO_NAME>.txt
		jisyoAll, err := convertCsvToSkk(jisyoRows)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		fileName := "SKK-JISYO-" + jisyoName + ".txt"
		err = export("skk", fileName, jisyoAll)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		log.Printf("Done: %s => %s", jisyoCSVFileName, fileName)
		return nil
	case "contacts":
		// output: GContacts-JISYO-<SYOSYO_JISYO_NAME>.csv
		jisyoAll, err := convertCsvToGoogleContacts(jisyoRows, jisyoName)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		fileName := "GContacts-JISYO-" + jisyoName + ".csv"
		err = export("contacts", fileName, jisyoAll)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		log.Printf("Done: %s => %s", jisyoCSVFileName, fileName)
		return nil
	default:
		return nil
	}
}

func export(baseDir, fileName, all string) error {
	err := checkBaseDir(baseDir)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	err = ioutil.WriteFile(filepath.Join(baseDir, fileName), []byte(all), 0644)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func checkBaseDir(baseDir string) error {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir(baseDir, 0755)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}
	return nil
}

func convertCsvToSkk(jisyoRows []*JisyoCSV) (string, error) {
	skkJisyoTmpl := "{{.yomi}} /{{.word}};{{.note}}/\n"
	skkJisyoAll := ";; okuri-nasi entries.\n"
	for _, row := range jisyoRows {
		// 雑に split
		yo := strings.Split(row.Yomi, ",")
		for _, y := range yo {
			t, err := template.New("SKKJisyo").Parse(skkJisyoTmpl)
			if err != nil {
				return "", fmt.Errorf("%w", err)
			}
			data := map[string]interface{}{
				"yomi": y,
				"word": row.Word,
				"note": row.Note,
			}
			var buf bytes.Buffer
			err = t.Execute(&buf, data)
			if err != nil {
				return "", fmt.Errorf("%w", err)
			}
			skkJisyoAll += buf.String()
		}
	}
	return skkJisyoAll, nil
}

func convertCsvToGoogleContacts(jisyoRows []*JisyoCSV, name string) (string, error) {
	var rows []GoogleContactsCSV
	for _, row := range jisyoRows {
		// 雑に split
		yo := strings.Split(row.Yomi, ",")
		for _, y := range yo {
			row := GoogleContactsCSV{
				GivenName:       row.Word,
				FamilyName:      "_",
				GivenNameYomi:   y,
				GroupMembership: name + " ::: * myContacts",
			}
			rows = append(rows, row)
		}
	}
	csv, err := gocsv.MarshalString(rows)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return csv, nil
}
