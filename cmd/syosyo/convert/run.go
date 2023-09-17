package convert

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert JISYO_NAME",
		Short: "generate jisyo file",
		Long: `
$ ./syosyo convert inoriminase
2023/09/17 16:11:42 Done: inoriminase.csv => SKK-JISYO-inoriminase.txt
`,
		RunE: run,
	}

	cmd.Flags().String("format", "skk", "jisyo file format (skk, contacts)")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	// input: <SYOSYO_JISYO_NAME>.csv

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return errors.WithStack(err)
	}

	var jisyoName string
	jisyoNameEnv := os.Getenv("SYOSYO_JISYO_NAME")
	if jisyoNameEnv != "" {
		jisyoName = jisyoNameEnv
	} else {
		if len(args) == 0 {
			return errors.New("no jisyo name")
		}
		jisyoName = args[0]
	}
	jisyoCSVFileName := jisyoName + ".csv"

	f, err := os.OpenFile(filepath.Join("csv", jisyoCSVFileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	jisyoRows := []*JisyoCSV{}
	err = gocsv.UnmarshalFile(f, &jisyoRows)
	if err != nil {
		return errors.WithStack(err)
	}

	switch format {
	case "skk":
		// output: SKK-JISYO-<SYOSYO_JISYO_NAME>.txt
		jisyoAll, err := convertCsvToSkk(jisyoRows)
		if err != nil {
			return err
		}
		fileName := "SKK-JISYO-" + jisyoName + ".txt"
		err = export("skk", fileName, jisyoAll)
		if err != nil {
			return err
		}
		log.Printf("Done: %s => %s", jisyoCSVFileName, fileName)
		return nil
	case "contacts":
		// output: GContacts-JISYO-<SYOSYO_JISYO_NAME>.csv
		jisyoAll, err := convertCsvToGoogleContacts(jisyoRows, jisyoName)
		if err != nil {
			return err
		}
		fileName := "GContacts-JISYO-" + jisyoName + ".csv"
		err = export("contacts", fileName, jisyoAll)
		if err != nil {
			return err
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
		return err
	}
	err = os.WriteFile(filepath.Join(baseDir, fileName), []byte(all), 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func checkBaseDir(baseDir string) error {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.Mkdir(baseDir, 0755)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	return nil
}

func convertCsvToSkk(jisyoRows []*JisyoCSV) (string, error) {
	skkJisyoTmpl := "{{ .yomi }} /{{ range .word }}{{ . }}{{ with $.note }};{{ $.note }}{{ end }}/{{ end }}\n"
	skkJisyoAll := ";; okuri-nasi entries.\n"
	for _, row := range jisyoRows {
		// 雑に split
		yo := strings.Split(row.Yomi, ",")
		for _, y := range yo {
			t, err := template.New("SKKJisyo").Parse(skkJisyoTmpl)
			if err != nil {
				return "", errors.WithStack(err)
			}
			data := map[string]interface{}{
				"yomi": y,
				"word": strings.Split(row.Word, ","),
				"note": row.Note,
			}
			var buf bytes.Buffer
			err = t.Execute(&buf, data)
			if err != nil {
				return "", errors.WithStack(err)
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
			wo := strings.Split(row.Word, ",")
			for _, w := range wo {
				row := GoogleContactsCSV{
					GivenName:       w,
					FamilyName:      "_",
					GivenNameYomi:   y,
					GroupMembership: name + " ::: * myContacts",
				}
				rows = append(rows, row)
			}
		}
	}
	csv, err := gocsv.MarshalString(rows)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return csv, nil
}
