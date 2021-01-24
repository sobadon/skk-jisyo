package main

import (
	"testing"
)

func Test_convertCsvToSkk(t *testing.T) {
	type args struct {
		jisyoRows []*JisyoCSV
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "note: false",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "やくそくの",
						Word: "約束のアステリズム",
						Note: "",
					},
				},
			},
			want: `;; okuri-nasi entries.
やくそくの /約束のアステリズム/
`,
			wantErr: false,
		},
		{
			name: "yomi: 1, word: 1, note: true",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "やくそくの",
						Word: "約束のアステリズム",
						Note: "作詞・作曲・編曲：藤永龍太郎(Elements Garden)",
					},
				},
			},
			want: `;; okuri-nasi entries.
やくそくの /約束のアステリズム;作詞・作曲・編曲：藤永龍太郎(Elements Garden)/
`,
			wantErr: false,
		},
		{
			name: "yomi: 2, word: 1",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "きゃっち,ctr",
						Word: "Catch the Rainbow！",
						Note: "作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY",
					},
				},
			},
			want: `;; okuri-nasi entries.
きゃっち /Catch the Rainbow！;作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY/
ctr /Catch the Rainbow！;作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY/
`,
			wantErr: false,
		},
		{
			name: "yomi: 1, word: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "しゃこう",
						Word: "社会工学類,社工",
						Note: "",
					},
				},
			},
			want: `;; okuri-nasi entries.
しゃこう /社会工学類/
しゃこう /社工/
`,
			wantErr: false,
		},
		{
			name: "yomi: 2, word: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "こうしす,esys",
						Word: "工学システム学類,工シス",
						Note: "esys（いーしす：Engineering System）",
					},
				},
			},
			want: `;; okuri-nasi entries.
こうしす /工学システム学類;esys（いーしす：Engineering System）/
こうしす /工シス;esys（いーしす：Engineering System）/
esys /工学システム学類;esys（いーしす：Engineering System）/
esys /工シス;esys（いーしす：Engineering System）/
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertCsvToSkk(tt.args.jisyoRows)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertCsvToSkk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertCsvToSkk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertCsvToGoogleContacts(t *testing.T) {
	type args struct {
		jisyoRows []*JisyoCSV
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "yomi: 1, word: 1",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "やくそくの",
						Word: "約束のアステリズム",
						Note: "作詞・作曲・編曲：藤永龍太郎(Elements Garden)",
					},
				},
				name: "inoriminase",
			},
			want: `Given Name,Family Name,Given Name Yomi,Group Membership
約束のアステリズム,_,やくそくの,inoriminase ::: * myContacts
`,
			wantErr: false,
		},
		{
			name: "yomi: 2, word: 1",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "きゃっち,ctr",
						Word: "Catch the Rainbow！",
						Note: "作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY",
					},
				},
				name: "inoriminase",
			},
			want: `Given Name,Family Name,Given Name Yomi,Group Membership
Catch the Rainbow！,_,きゃっち,inoriminase ::: * myContacts
Catch the Rainbow！,_,ctr,inoriminase ::: * myContacts
`,
			wantErr: false,
		},
		{
			name: "yomi: 1, word: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "しゃこう",
						Word: "社会工学類,社工",
						Note: "",
					},
				},
				name: "itf",
			},
			want: `Given Name,Family Name,Given Name Yomi,Group Membership
社会工学類,_,しゃこう,itf ::: * myContacts
社工,_,しゃこう,itf ::: * myContacts
`,
			wantErr: false,
		},
		{
			name: "yomi: 2, word: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						Yomi: "こうしす,esys",
						Word: "工学システム学類,工シス",
						Note: "esys（いーしす：Engineering System）",
					},
				},
				name: "itf",
			},
			want: `Given Name,Family Name,Given Name Yomi,Group Membership
工学システム学類,_,こうしす,itf ::: * myContacts
工シス,_,こうしす,itf ::: * myContacts
工学システム学類,_,esys,itf ::: * myContacts
工シス,_,esys,itf ::: * myContacts
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertCsvToGoogleContacts(tt.args.jisyoRows, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertCsvToGoogleContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertCsvToGoogleContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}
