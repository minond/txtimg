Txtimg, convert ascii maps into gifs. Use as a stand-alone command line tool or
as a server. Generate a Gif with a delay of a tenth of a second between frames:

```bash
go run cmd/txtimg/main.go  -delay 10 data/simple*.txt
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

Or, sping up a server that accepts POST requests with files (`frames`
parameter):

```bash
go run cmd/txtimg/main.go -listen localhost:8080
Setting up server on localhost:8080

curl -X POST http://localhost:8080 \
  -F frames=@data/simple-01.txt \
  -F frames=@data/simple-02.txt \
  -F frames=@data/simple-03.txt \
  -F frames=@data/simple-04.txt \
  -F frames=@data/simple-05.txt \
  -F frames=@data/simple-06.txt \
  -F frames=@data/simple-07.txt \
  -F frames=@data/simple-08.txt \
  -F frames=@data/simple-09.txt \
  --output out.gif
```

![simple.gif](data/simple.gif)
