Txtimg, given a set of frames, which are plain text files, generate a gif. Run
`go get github.com/minond/txtimg/cmd/txtimg` to install. Once installed, this
can be used to generate one-offs with the command line (`txtimg -delay <int>
<frameFiles>`) or you can start a server that handles requests made to `POST /`
and generates the image using the `frames` passed (as files). Sample usage:

```bash
# Creates out.gif using data/simples-0{1,9}.txt
> txtimg -delay 10 data/simple*.txt
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

```bash
# Sets up a server on localhost:8080
> txtimg -listen localhost:8080
Setting up server on localhost:8080
```

```bash
# Make a POST request to http://localhost:8080 with the same files we passed in
# the first command line example.
> curl -X POST http://localhost:8080 \
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
