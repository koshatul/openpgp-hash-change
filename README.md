# openpgp-hash-change
Testing an issue with the secret hash changing

### Output

```shell
$ go run main.go; md5sum test*.key
feb63c754aefb907506720c140aea17d  test1.key
28abfb6a581a4ff162e8abfc666110a6  test2.key
feb63c754aefb907506720c140aea17d  test3.key
28abfb6a581a4ff162e8abfc666110a6  test4.key
```

- First key is loaded (or generated)
- Second key is written from the first
- Second key is loaded from the file
- Third key is written from the second
- Third key is loaded from the file
- Fourth key is written from the third
- Fourth key is loaded from the file
- First key is written from the fourth

diff of test1.key and test2.key using `gpg --list-packets`

```diff
--- sample.test1.packet	2020-02-27 15:05:54.000000000 +1000
+++ sample.test2.packet	2020-02-27 15:05:59.000000000 +1000
@@ -6,8 +6,8 @@
 	skey[2]: [2046 bits]
 	skey[3]: [1024 bits]
 	skey[4]: [1024 bits]
-	skey[5]: [1023 bits]
-	checksum: 50ff
+	skey[5]: [1024 bits]
+	checksum: 536e
 	keyid: 9EE9632F50616BBB
 # off=923 ctb=cd tag=13 hlen=2 plen=33 new-ctb
 :user ID packet: "name (comment) <email@domain.org>"
@@ -30,8 +30,8 @@
 	skey[2]: [2048 bits]
 	skey[3]: [1024 bits]
 	skey[4]: [1024 bits]
-	skey[5]: [1024 bits]
-	checksum: 35a0
+	skey[5]: [1023 bits]
+	checksum: 2ccc
 	keyid: 7CEA763A7F7C44BC
 # off=2180 ctb=c2 tag=2 hlen=3 plen=287 new-ctb
 :signature packet: algo 1, keyid 9EE9632F50616BBB
```
