package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"teepee78/reverse-proxy-go/config"
)

func ServeDynamic(w http.ResponseWriter, r *http.Request) {
	if config.Vars.Routes == nil {
		return
	}

	urlPath := "/" + cleanPath(r.URL.Path)

	for _, route := range config.Vars.Routes {
		path := route.Path
		targets := route.Targets

		if strings.HasPrefix(urlPath, path) {
			targetUrl := getUrl(targets, path, urlPath)
			err := makeRequest(w, r, targetUrl)
			if err != nil {
				// Try a new target
				config.Retrials--
				ServeDynamic(w, r)
			}

			if config.Retrials == 0 {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
}

func makeRequest(w http.ResponseWriter, r *http.Request, url *url.URL) error {
	req, reqErr := http.NewRequest(r.Method, url.String(), r.Body)
	if reqErr != nil {
		http.Error(w, reqErr.Error(), http.StatusBadRequest)
		return nil
	}

	req.Header = r.Header.Clone()

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		fmt.Println("Server unreachable:", resErr)
		return fmt.Errorf("server unreachable")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(res.StatusCode)

	_, err := io.Copy(w, res.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
	}

	config.Retrials = len(config.Vars.Routes)

	return nil
}
