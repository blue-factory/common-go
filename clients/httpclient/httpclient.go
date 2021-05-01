package httpclient

import (
	"net/http"
	"time"
)

// ValidatorFunc ...
type ValidatorFunc func(*http.Response) (*http.Response, error)

// HTTPClient ...
type HTTPClient struct {
	Client   *http.Client
	MaxRetry int
	Delay    int
	Validate ValidatorFunc
	log      Logger
}

// Do ...
func (cl *HTTPClient) Do(req *http.Request) (res *http.Response, err error) {
	for i := 0; i < cl.MaxRetry; i++ {
		res, err = cl.Client.Do(req)
		if err != nil {
			cl.log.Infof("Failed request, attempt nummber %d. Retrying next time in %d seconds", i, cl.Delay)
			time.Sleep(time.Second * time.Duration(cl.Delay))
			continue
		}
		res, err = cl.Validate(res)
		if err != nil {
			cl.log.Infof("Failed validation strategy, attempt #%d Retrying next time in %d seconds", i, cl.Delay)
			cl.log.Infof("http status received %d", res.StatusCode)
			cl.log.Infof("Requested URL=%s", req.URL.String())
			time.Sleep(time.Second * time.Duration(cl.Delay))
			continue
		}

		return res, nil
	}

	cl.log.Infof("Max retry URL=%s", req.URL.String())
	return nil, err
}
