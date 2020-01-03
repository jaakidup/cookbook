package proxy

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

// Proxy holds our configured client
// and BaseURL to proxy to
type Proxy struct {
	Client  *http.Client
	BaseURL string
}

// ServeHTTP means that proxy implments the Handler interface
// It manipulates the request, forwards it to BaseURL, then
// returns the response
func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := p.ProcessRequest(r); err != nil {
		log.Printf("error occurred during process request: %s",
			err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := p.Client.Do(r)
	if err != nil {
		log.Printf("error occurred during client operation: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	CopyResponse(w, resp)
}

// ProcessRequest modifies the request in accordance
// with Proxy settings
func (p *Proxy) ProcessRequest(r *http.Request) error {
	proxyURLRaw := p.BaseURL + r.URL.String()
	proxyURL, err := url.Parse(proxyURLRaw)
	if err != nil {
		return err
	}
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""
	return nil
}

// CopyResponse takes the client response and writes everything
// to the ResponseWriter in the original handler
func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	out.ReadFrom(resp.Body)
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	w.Write(out.Bytes())
}
