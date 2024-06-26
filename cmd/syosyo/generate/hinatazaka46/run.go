package hinatazaka46

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
)

type JisyoCSV struct {
	Yomi string `csv:"yomi"`
	Word string `csv:"word"`
	Note string `csv:"note"`
}

// 基本的には bool や int なんてものはなくて、すべて string で表現されている
type GreetingAPIResponse struct {
	Greetings []Greeting `json:"greeting"`
}

type Greeting struct {
	// 加藤 史帆
	Name string `json:"name"`
	// かとう しほ
	NameFuri string `json:"name_furi"`
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "hinatazaka46",
		Short: "generate hinatazaka46 csv file",
		RunE:  run,
	}

	return rootCmd
}

func run(*cobra.Command, []string) error {
	const greetingURL = "https://www.hinatazaka46.com/s/official/api/list/greeting"
	res, err := http.Get(greetingURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBodyBuf := new(bytes.Buffer)
	resBodyBuf.ReadFrom(res.Body)
	resBodyStr := resBodyBuf.String()

	var greetingAPIResponse GreetingAPIResponse
	err = json.Unmarshal([]byte(resBodyStr), &greetingAPIResponse)
	if err != nil {
		return err
	}

	jisyoRows := []*JisyoCSV{}
	for _, gr := range greetingAPIResponse.Greetings {
		// 何名か name_furi の先頭にスペースが入ってしまっているので trim してから split
		kanaSplit := strings.Split(strings.TrimSpace(gr.NameFuri), " ")
		if len(kanaSplit) != 2 {
			return fmt.Errorf("性・名 区切りに失敗: %s", gr.NameFuri)
		}
		sei := kanaSplit[0]
		mei := kanaSplit[1]

		jisyoRows = append(jisyoRows, &JisyoCSV{
			Yomi: strings.Join([]string{sei, mei}, ","),
			Word: gr.Name,
			Note: "",
		})
	}

	dstFile, err := os.Create("csv/hinatazaka46-member.csv")
	if err != nil {
		return err
	}

	err = gocsv.MarshalFile(jisyoRows, dstFile)
	return err
}
