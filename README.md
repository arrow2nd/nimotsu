# ![nimotsu](https://user-images.githubusercontent.com/44780846/143875126-623331d8-dc27-4145-95c7-5d26968d63bf.png)

荷物の配達状況を追跡する CLI ツール

[![test](https://github.com/arrow2nd/nimotsu/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/nimotsu/actions/workflows/test.yml)
[![release](https://github.com/arrow2nd/nimotsu/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/nimotsu/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/nimotsu)](https://goreportcard.com/report/github.com/arrow2nd/nimotsu)
![GitHub all releases](https://img.shields.io/github/downloads/arrow2nd/nimotsu/total)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/nimotsu)](https://github.com/arrow2nd/nimotsu/blob/main/LICENSE)

![screenshot](https://user-images.githubusercontent.com/44780846/141614236-b7ea80b9-e76f-4514-a0ce-d1fb19e5290a.png)

## 対応している配送業者

- 日本郵便
- ヤマト運輸
- 佐川急便

## インストール

### Homebrew

```sh
brew tap arrow2nd/tap
brew install nimotsu
```

### Go

```sh
go install github.com/arrow2nd/nimotsu@latest
```

### それ以外

[Releases](https://github.com/arrow2nd/nimotsu/releases) からお使いの環境にあったファイルをダウンロードしてください。

## コマンド

### get

`get [<追跡番号>] [配送業者フラグ]`

追跡番号から荷物を追跡します。

```txt
$ nimotsu get 1122334455 --sagawa
```

配送業者を指定しなかった場合、選択肢が表示されます。

```txt
$ nimotsu get 112233445566
Use the arrow keys to navigate: ↓ ↑ → ←
Carrier?
  > 日本郵便
    ヤマト運輸
    佐川急便
```

### get all

リスト内の荷物を全て追跡します。

```txt
$ nimotsu get all
```

### add

`add [<追跡番号>] [配送業者フラグ] [--comment "コメント"]`

リストに荷物を追加します。

コメントを省略した場合、"なし" が設定されます。

```txt
$ nimotsu add 112233445566 --japanpost --comment "🍺"
✔ Added!
```

配送業者・コメントを指定しなかった場合、対話形式で入力できます。

```txt
$ nimotsu add 112233445566
Carrier: 日本郵便
Comment: beer🍺
✔ Added!
```

### remove

リスト内の荷物を選択して削除します。

```txt
$ nimotsu remove
Use the arrow keys to navigate: ↓ ↑ → ←
Tracking number?
    112233445566
    223344556677
  > 123456789123
Carrier: 日本郵便
Comment: beer🍺
```

### remove all

リスト内の荷物を全て削除します。

```txt
$ nimotsu remove all
✔ Removed all!
```

### edit

リスト内の荷物を選択してコメントを変更します。

```txt
$ nimotsu edit
Tracking number: 444191245470
Comment: iPad
✔ Edited!
```

### list

リスト内の荷物を表示します

```txt
$ nimotsu list
```

### version

バージョンを表示します。

```txt
$ nimotsu version
📦  nimotsu ver.x.x.x
```
