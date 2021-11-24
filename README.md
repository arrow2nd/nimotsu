# ![nimotsu](https://user-images.githubusercontent.com/44780846/143192856-74afca8a-d837-4c7e-96dc-239ef15960a5.png)
  
[![test](https://github.com/arrow2nd/nimotsu/actions/workflows/test.yml/badge.svg)](https://github.com/arrow2nd/nimotsu/actions/workflows/test.yml)
[![release](https://github.com/arrow2nd/nimotsu/actions/workflows/release.yml/badge.svg)](https://github.com/arrow2nd/nimotsu/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/arrow2nd/nimotsu)](https://goreportcard.com/report/github.com/arrow2nd/nimotsu)
[![GitHub license](https://img.shields.io/github/license/arrow2nd/nimotsu)](https://github.com/arrow2nd/nimotsu/blob/main/LICENSE)

荷物の配達状況を追跡する CLI ツール

## スクリーンショット

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

`get [<配送業者フラグ>] [<追跡番号>]`

追跡番号から荷物を追跡します。

```txt
$ nimotsu get --japanpost 112233445566
```

### get all

リスト内の荷物を全て追跡します。

```txt
$ nimotsu get all
```

### add

`add [<配送業者フラグ>] [<追跡番号>] [--comment "コメント"]`

リストに荷物を追加します。

コメントを省略した場合、"なし"が設定されます。

```txt
$ nimotsu add --japanpost 112233445566 --comment "🍺"
```

### remove

`remove [<追跡番号>]`

リストから荷物を削除します。

```txt
$ nimotsu remove 112233445566
```

### remove all

リスト内の荷物を全て削除します。

```txt
$ nimotsu remove all
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
```
