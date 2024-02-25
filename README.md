# skk-jisyo

変な CSV ファイルから SKK で使えるような辞書を「雑に」出力してくれる子。skk-jisyo と名乗りつつも、Google Contacts 形式にも出力してくれる。
SKK 日本語入力 FEP で利用可能な辞書ファイルを作成する。他のソフトウェアまわりは考慮されていない。雑にということで、本来必要である（？）辞書のソートをしていない。

## SKK

```
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/skk/SKK-JISYO-inoriminase.txt
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/skk/SKK-JISYO-itf.txt
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/skk/SKK-JISYO-nogizaka46-member.txt
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/skk/SKK-JISYO-say-you.txt
```

## Google Contacts

iOS などのため。

```
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/contacts/GContacts-JISYO-inoriminase.csv
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/contacts/GContacts-JISYO-itf.csv
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/contacts/GContacts-JISYO-nogizaka46-member.csv
https://object-storage1.chasoba.net/skk-jisyo-chasoba-net/latest/contacts/GContacts-JISYO-say-you.csv
```

---

以下は SKK セットアップ用のメモ

## SKKFEP

### インストール

- [`skkfep.js`](http://coexe.web.fc2.com/js/skkfep.js)
- [`skkgate03_20190401.zip`](http://coexe.web.fc2.com/skkgate03_20190401.zip)

### `DICTS`

- System：`C:\Windows\IME\SKK0\DICTS`
- User：`%APPDATA%\SKKFEP\DICTS`

### 全般
https://skk-dev.github.io/dict/

SKKFEP が自動でダウンロードする [OpenLab の辞書](http://openlab.ring.gr.jp/skk/skk/dic/)は更新されていないため

- [SKK-JISYO.L.gz](https://skk-dev.github.io/dict/SKK-JISYO.L.gz)
- [SKK-JISYO.jinmei.gz](https://skk-dev.github.io/dict/SKK-JISYO.jinmei.gz)
- [SKK-JISYO.fullname.gz](https://skk-dev.github.io/dict/SKK-JISYO.fullname.gz)
- [SKK-JISYO.geo.gz](https://skk-dev.github.io/dict/SKK-JISYO.geo.gz)
- [SKK-JISYO.propernoun.gz](https://skk-dev.github.io/dict/SKK-JISYO.propernoun.gz)
- [SKK-JISYO.station.gz](https://skk-dev.github.io/dict/SKK-JISYO.station.gz)
- [SKK-JISYO.law.gz](https://skk-dev.github.io/dict/SKK-JISYO.law.gz)
- [SKK-JISYO.okinawa.gz](https://skk-dev.github.io/dict/SKK-JISYO.okinawa.gz)
- [SKK-JISYO.china_taiwan.gz](https://skk-dev.github.io/dict/SKK-JISYO.china_taiwan.gz)
- [SKK-JISYO.assoc.gz](https://skk-dev.github.io/dict/SKK-JISYO.assoc.gz)
- [SKK-JISYO.edict.tar.gz](https://skk-dev.github.io/dict/SKK-JISYO.edict.tar.gz)
- [zipcode.tar.gz](https://skk-dev.github.io/dict/zipcode.tar.gz)

### 絵文字

- [SKK-JISYO.emoji.utf8](https://raw.githubusercontent.com/uasi/skk-emoji-jisyo/master/SKK-JISYO.emoji.utf8)
    - [uasi/skk-emoji-jisyo: SKK 絵文字辞書](https://github.com/uasi/skk-emoji-jisyo)
