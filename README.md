Convert ascii maps into gifs.

```bash
$ go run cmd/txtimg/main.go  -delay 10 data/simple*.txt
 11% - Encoding data/simple-01.txt
 22% - Encoding data/simple-02.txt
 33% - Encoding data/simple-03.txt
 44% - Encoding data/simple-04.txt
 56% - Encoding data/simple-05.txt
 67% - Encoding data/simple-06.txt
 78% - Encoding data/simple-07.txt
 89% - Encoding data/simple-08.txt
100% - Encoding data/simple-09.txt
Done - Saved to out.gif with a delay of 10 between frames
```

![simple.gif](data/simple.gif)
