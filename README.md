**Resume**:
1. 3 goroutines are created (1 - sender, 2 - receiver, 3 - by means of a context and a time packet, completing the rest), communicating through a channel. There is also an additional waitCh channel, since I decided that I could do without the sync.waitgroup in this case.
2. The send interval was added to the .env file, since the opposite was not described in the task. We also used the common viper package, which allows you to work with other extensions too. In addition, the configuration file was not moved to a separate directory and was not implemented in any structure, since I don’t see the point in this small task - I decided not to complicate
3. However, it was not clear what was meant by custom text, incl. the generation of a random string from a set of characters was added

**What was not done (and probably not expected):**
1. No tests were written
2. No clean architecture approach was used
3. No makefile and configuration for docker, no CI/CD instructions etc.
4. Init() function was not used for configuration initialization.

P.S. I was not sure how meticulous to approach this small task :)

**Output example**
```
dsemenk@FVFGK3ERQ05N % go run main.go -race
Receiver got message. Time: 2022-08-04 19:34:13.075537 +0300 EEST m=+2.006143001, text: l
Receiver got message. Time: 2022-08-04 19:34:16.071565 +0300 EEST m=+5.002205793, text: xqDZDatzXZDRtFMsOBU
Receiver got message. Time: 2022-08-04 19:34:21.075403 +0300 EEST m=+10.006099834, text: bbaHnlUVGAyliq
Receiver got message. Time: 2022-08-04 19:34:26.075814 +0300 EEST m=+15.006565168, text: pgQlxrFTxahaxFgIecwLrFUDelHE
Receiver got message. Time: 2022-08-04 19:34:31.072023 +0300 EEST m=+20.002827418, text: UnktLEamNHdpx
Receiver got message. Time: 2022-08-04 19:34:36.070387 +0300 EEST m=+25.001244209, text: eunKQFcjEkEXccYPstXEvvbc
Receiver got message. Time: 2022-08-04 19:34:41.072041 +0300 EEST m=+30.002949334, text: atEGwKYvfzhqApdMrveJ
Receiver got message. Time: 2022-08-04 19:34:46.071381 +0300 EEST m=+35.002339834, text: jvAH
Receiver got message. Time: 2022-08-04 19:34:51.074462 +0300 EEST m=+40.005471834, text: frdbhiMPYkBgBBkWfKhbmiJX
Receiver got message. Time: 2022-08-04 19:34:56.07431 +0300 EEST m=+45.005370293, text: TbtIISKPVPExfDoEajjAqXMKwUTBx
Receiver got message. Time: 2022-08-04 19:35:01.072529 +0300 EEST m=+50.003639209, text: HdLkILnic
Receiver got message. Time: 2022-08-04 19:35:06.074935 +0300 EEST m=+55.006094918, text: SwWhyjmFCNkeMBE
dsemenk@FVFGK3ERQ05N % 
