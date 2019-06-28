# context

このサンプルは、複数のgoroutineが全て完了してから次の処理にうつる。中断されたらそれぞれのgoroutineを安全に終了させる。

# メモ

context自身は親から子へ終了指示を伝播させるのが目的のようで、子から親へ指示をだすものではなさそうである。そのためgoroutineの正常終了の処理は自力で実装する必要がある。

Context.Done() のチャンネルクローズで停止の通知をするのは、select case でブロックさせないようにするためだろうか？