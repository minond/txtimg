package txtimg

import (
	"fmt"
	"image/gif"
	"io/ioutil"
	"net/http"
)

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	var frames []string
	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		http.Error(w, "Unable to parse request", http.StatusBadRequest)
		return
	}

	files, ok := r.MultipartForm.File["frames"]

	if !ok {
		http.Error(w, "Missing frames file in request", http.StatusBadRequest)
		return
	}

	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()

		if err != nil {
			http.Error(w, "Unable to open uploaded file", http.StatusBadRequest)
			return
		}

		content, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(w, "Unable to read uploaded file", http.StatusBadRequest)
			return
		}

		frames = append(frames, string(content))
	}

	out := &gif.GIF{}
	images, err := BuildGifFrames(frames)

	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to generate gif: %v", err),
			http.StatusInternalServerError)

		return
	}

	out.Image = append(out.Image, images...)

	for i := 0; i < len(images); i++ {
		out.Delay = append(out.Delay, 25)
	}

	w.Header().Set("Content-type", "image/gif")
	gif.EncodeAll(w, out)
}

func Service() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleFileUpload(w, r)
			break

		default:
			http.Error(w, "Unsupported method", http.StatusNotFound)
			break
		}
	})

	return mux
}
