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
			name: "yomi: 1",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						WordsBefore: "やくそくの",
						WordAfter:   "約束のアステリズム",
						Note:        "作詞・作曲・編曲：藤永龍太郎(Elements Garden)",
					},
				},
			},
			want: `;; okuri-nasi entries.
やくそくの /約束のアステリズム;作詞・作曲・編曲：藤永龍太郎(Elements Garden)/
`,
			wantErr: false,
		},
		{
			name: "yomi: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						WordsBefore: "きゃっち,ctr",
						WordAfter:   "Catch the Rainbow！",
						Note:        "作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY",
					},
				},
			},
			want: `;; okuri-nasi entries.
きゃっち /Catch the Rainbow！;作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY/
ctr /Catch the Rainbow！;作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY/
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
			name: "yomi: 1",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						WordsBefore: "やくそくの",
						WordAfter:   "約束のアステリズム",
						Note:        "作詞・作曲・編曲：藤永龍太郎(Elements Garden)",
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
			name: "yomi: 2",
			args: args{
				jisyoRows: []*JisyoCSV{
					{
						WordsBefore: "きゃっち,ctr",
						WordAfter:   "Catch the Rainbow！",
						Note:        "作詞：水瀬いのり　作曲：光増ハジメ　編曲：EFFY",
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
