# テンプレート

https://golang.org/pkg/text/template

* 構造体をつかったものと、Interface{} をつかったもの。
* Interface{} を使った場合、コード側に構造の定義が不要となる。

## 実行例

	2019/06/21 08:27:02  *** シンプル版 ***
	こんばんわ まめも です。
	まめも のホームページは https://github.com/mamemomonga です。
	
	好きなものは  なし  りんご  バナナ  です。
	
	今後ともよろしくお願いいたします。
	
	2019/06/21 08:27:02  *** YAML+Interface{}版 ***
	2019/06/21 08:27:02 Read: data.yaml
	(*main.MyData)(0xc00000e160)({
	 Template: (map[interface {}]interface {}) (len=3) {
	  (string) (len=4) "Name": (string) (len=12) "ぐぐーる",
	  (string) (len=8) "HomePage": (string) (len=19) "https://google.com/",
	  (string) (len=9) "Favorites": ([]interface {}) (len=3 cap=3) {
	   (string) (len=2) "Go",
	   (string) (len=4) "Java",
	   (string) (len=10) "JavaScript"
	  }
	 }
	})
	こんばんわ ぐぐーる です。
	ぐぐーる のホームページは https://google.com/ です。
	
	好きなものは  Go  Java  JavaScript  です。
	
	今後ともよろしくお願いいたします。

## 参考

* https://coderwall.com/p/ns60fq/simply-output-go-html-template-execution-to-strings
* https://qiita.com/kamina/items/58e5290ff0f569a76331


