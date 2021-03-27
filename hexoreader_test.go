package hexoreader

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestHexoReader_ReadAll(t *testing.T) {
	type fields struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    Post
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				r: strings.NewReader(`title: "技術ブログの年間連載予定を発表します"
date: 2021/01/12 00:00:00
tags:
  - タグ1
  - タグ2
category:
  - カテゴリ1
thumbnail: /images/20210112/thumbnail.jpg
author: 神奈川太郎
featured: false
lede: "あけましておめでとうございます。"
---

こんにちは。神奈川太郎です。

あけましておめでとうございます。本年もよろしくお願いいたします。

2021年に計画している連載についてご紹介します。`),
			},
			want: Post{
				FrontMatter: FrontMatter{
					Title:      "技術ブログの年間連載予定を発表します",
					Date:        "2021/01/12 00:00:00",
					Tags:       []string{"タグ1", "タグ2"},
					Categories: []string{"カテゴリ1"},
				},
				Content: `こんにちは。神奈川太郎です。

あけましておめでとうございます。本年もよろしくお願いいたします。

2021年に計画している連載についてご紹介します。`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HexoReader{
				r: tt.fields.r,
			}
			got, err := h.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func MustTime(t time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return t
}
