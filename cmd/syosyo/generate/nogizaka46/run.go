package nogizaka46

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
type MemberAPIResponse struct {
	Count string   `json:"count"`
	Data  []Member `json:"data"`
}

type Member struct {
	// ???
	Code string `json:"code"`

	// 井上 和
	Name string `json:"name"`

	// nagi inoue
	EnglishName string `json:"english_name"`

	// いのうえ なぎ
	Kana string `json:"kana"`

	// 5期生
	Cate string `json:"cate"`

	// https://www.nogizaka46.com/images/46/540/61c907a38e56dced12a57ab544fd3.jpg
	Img string `json:"img"`

	// https://www.nogizaka46.com/s/n46/artist/55389?ima=1020
	Link string `json:"link"`

	// 選抜メンバー
	Pick string `json:"pick"`

	// 十一福神
	God string `json:"god"`

	// アンダー, ""
	Under string `json:"under"`

	// 2005/02/17
	Birthday string `json:"birthday"`

	// B型, 不明
	Blood string `json:"blood"`

	// みずがめ座
	Constellation string `json:"constellation"`

	// YES, NO
	Graduation string `json:"graduation"`

	// ""
	Groupcode string `json:"groupcode"`
}

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "nogizaka46",
		Short: "generate nogizaka46 csv file",
		RunE:  run,
	}

	return rootCmd
}

func run(*cobra.Command, []string) error {
	const memberURL = "https://www.nogizaka46.com/s/n46/api/list/member?callback=res"
	res, err := http.Get(memberURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBodyBuf := new(bytes.Buffer)
	resBodyBuf.ReadFrom(res.Body)
	resBodyStr := resBodyBuf.String()

	// JSON として解釈するときに邪魔となる `res(` と `);` を削除する
	resBodyTrimmed := resBodyStr[4 : len(resBodyStr)-2]

	var memberAPIResponse MemberAPIResponse
	err = json.Unmarshal([]byte(resBodyTrimmed), &memberAPIResponse)
	if err != nil {
		return err
	}

	jisyoRows := []*JisyoCSV{}
	for _, member := range memberAPIResponse.Data {
		// 先頭に "箱推し" がある
		if member.Kana == "箱推し" {
			continue
		}

		kanaSplit := strings.Split(member.Kana, " ")
		if len(kanaSplit) != 2 {
			return fmt.Errorf("性・名 区切りに失敗: %s", member.Kana)
		}
		sei := kanaSplit[0]
		mei := kanaSplit[1]

		jisyoRows = append(jisyoRows, &JisyoCSV{
			Yomi: strings.Join([]string{sei, mei}, ","),
			Word: member.Name,
			Note: strings.Join([]string{member.Cate, strings.ReplaceAll(member.Birthday, `/`, `-`)}, " "),
		})
	}

	dstFile, err := os.Create("csv/nogizaka46-member.csv")
	if err != nil {
		return err
	}

	err = gocsv.MarshalFile(jisyoRows, dstFile)
	return err
}
