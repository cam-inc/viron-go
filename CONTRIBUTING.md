# Viron GOへの貢献

---

このプロジェクトへの貢献を皆さんに奨励し、サポートしていただきたいと考えています。

タスクは[GitHub Project](https://github.com/orgs/cam-inc/projects/2)で管理しており、進行中、完了、高優先度の課題を確認できます。

## どのように貢献できますか？

### バグの報告

バグを報告するには、[GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=bug&template=bug_report.md&title=)を提出してください。提出する前に、似たような問題が既に存在しないか確認してください。

### 機能改善の提案

機能改善を提案するには、[GitHub issue](https://github.com/cam-inc/viron-go/issues/new?assignees=&labels=enhancement&template=feature_request.md&title=)を提出してください。

### コードの貢献

#### コードを取得
```
$ git clone git@github.com:cam-inc/fensi-go.git
```

#### ツールセットアップ
##### taskfileインストール
```
# tips: 他のインストール方法でも大丈夫です。
# 詳しくは https://taskfile.dev/installation/ を参照してください。

$ brew install go-task/tap/go-task
```

##### asdf インストール
```
# tips: 他のインストール方法でも大丈夫です。
# 詳しくは https://asdf-vm.com/guide/getting-started.html を参照してください
$ brew install asdf

# 設定
$ echo 'export PATH="${ASDF_DATA_DIR:-$HOME/.asdf}/shims:$PATH"' >> ~/.zshrc
$ source ~/.zshrc
```

##### asdf plugin インストール
```
$ asdf plugin add golang
$ asdf plugin add ko https://github.com/zasdaym/asdf-ko.git
$ asdf plugin add lefthook https://github.com/jtzero/asdf-lefthook.git
$ asdf plugin add helm https://github.com/Antiarchitect/asdf-helm.git
```

##### asdf install
```
$ asdf install
```

#### アプリケーションセットアップ
##### task install
```
$ task install
```

##### go mod tidy
```
$ task tidy
```

#### テスト実行
```
$ task test
```

以下はコード貢献の簡単なガイドです。

1. リポジトリをフォークし、ローカルマシンにクローンします。
2. `main`ブランチから、新しいタスクに適した意味のある名前のブランチを作成します。
3. 次のコマンドを実行してセットアップします: `go mod tidy`。
4. ブランチをプッシュします。
5. アップストリームリポジトリにプルリクエストを提出します。

#### バージョンを管理
Go Modulesでバージョンを認識させるには、Gitタグを付ける必要があります。使用するバージョンは 必ず vX.Y.Z（[セマンティックバージョニング](https://semver.org)）形式にします。

[Module version numbering](https://go.dev/doc/modules/version-numbers)
```
例)

v1.0.0 を付けるには、

git tag v1.0.0
git push origin v1.0.0
```

📌 注意

タグ名に必ず v プレフィックスが必要です（例: v1.2.3）
Go Modulesはこれを検出して、そのバージョンをモジュールとして利用します。

#### 使用側プロジェクトから利用する
別のGoプロジェクトから、このライブラリを読み込むには go get を使います。
```
例:go get github.com/github.com/cam-inc/viron-go@v1.0.0

インポートも通常通り行います。

import "github.com/cam-inc/viron-go/lib/domains"
```

#### v2以降のバージョンについて
Go Modulesでは メジャーバージョンv2以降 は特別な扱いになります。ディレクトリ名やモジュールパスにもそのバージョンを含める必要があります。

[Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules)
```
例)

モジュール名: github.com/cam-inc/viron-go/v2
ディレクトリ構成: github.com/cam-inc/viron-go/v2/domains
実行するコマンド:

go mod init github.com/cam-inc/viron-go/v2

git tag v2.0.0

go get github.com/cam-inc/viron-go/v2@v2.0.0
```

#### リリースワークフロー
[Module release and versioning workflow](https://go.dev/doc/modules/release-workflow)

Go (Golang) の自作ライブラリのリリースワークフロー（GitHub連携を前提）は、以下のように設計できます。これにより、Gitでバージョン管理されたコードを正しくリリースし、Go Modules対応で外部プロジェクトから利用できるようになります。

**🔁 一般的なリリースワークフローの流れ**

1. 開発ブランチで機能実装・テスト
2. main ブランチへマージ
3. バージョン番号（セマンティックバージョニング）を決定
4. Git タグを付与（go modules 対応）
5. GitHub リリース作成（オプション）
6. 他プロジェクトから go get でインポート・利用

**📌 手順詳細**

- ① 開発・テスト
  - 機能追加やバグ修正をブランチ（feature/xxx、fix/yyyなど）で行い、ユニットテストやCIを通して検証します。

- ② main ブランチにマージ
  - Pull Request などでレビュー後、main ブランチにマージしてマイルストーンを完了とします。

- ③ バージョンを決める
  - 機能追加 → minor アップ (v1.1.0 → v1.2.0)
  - バグ修正のみ → patch アップ (v1.2.3 → v1.2.4)
  - 互換性のない変更 → major アップ (v1.x.x → v2.0.0)

- ④ Git タグを付けてリリース
  - CLIでタグを作成します（必ず v を付ける）:
    ```
    git tag v1.2.0
    git push origin v1.2.0

    これにより Go Modules がこのタグをバージョンとして認識します。
    ```

- ⑤ GitHub Releases でリリースノートを書く（任意）
  - GitHub の UI または CLI（gh release）を使ってリリースノートを追加します。これによりユーザーに変更点が明確に伝わります。
    ```
    例:

    gh release create v1.2.0 --title "v1.2.0 リリース" --notes "新機能とバグ修正を含むリリースです"
    ```


## 行動規範

[こちら](./CODE_OF_CONDUCT.md)に従ってください。

## ライセンス

このプロジェクトに貢献することで、あなたの貢献を[MITライセンス](./LICENSE)の下でライセンスすることに同意したことになります。

## 貢献者

このプロジェクトに貢献してくださった皆さんに感謝します。

<table>
  <tr>
    <td align="center"><a href="https://github.com/takoring"><img src="https://avatars.githubusercontent.com/u/24517668?v=4" width="100px;" alt=""/><br /><sub><b>takoring</b></sub></a><br />💻</td>
  </tr>
</table>
