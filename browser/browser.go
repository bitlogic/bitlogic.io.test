package browser

import (
	"fmt"
	"io/ioutil"
	"os"

	"strings"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

//Client is the browser we use to test our site
type Client struct {
	*browser.Browser
	base string
}

//Open makes the virtual browser navigate to path url in our site.
func (c *Client) Open(path string) error {
	return c.Browser.Open(fmt.Sprintf("%s%s", c.base, path))
}

func NewClient(path string) (*Client, error) {
	gse := os.Getenv("BITLOGIC_URL")
	b := surf.NewBrowser()
	client := &Client{b, gse}

	err := client.Open(path)
	if err != nil {
		return nil, err
	}
	return client, nil
}

//CheckPage ensure the browser status is ok and all page resources are downloaded
func (c *Client) CheckPage() error {
	if c.StatusCode() != 200 {
		return fmt.Errorf("request failed: '%s', bitlogic returned status: %d", c.Url().Path, c.StatusCode())
	}

	errors := []error{}
	// TODO ensure al resoruces are cacheable
	// TODO ensure the size of all resources is less than a max size
	//validate all scripsts can be downloaded
	for _, s := range c.Scripts() {
		_, err := s.Download(ioutil.Discard)
		if err != nil {
			errors = append(errors, fmt.Errorf("Can not download script %s: %v", s.Url().Path, err))
		}
	}

	//validate all scripsts can be downloaded
	for _, s := range c.Stylesheets() {
		_, err := s.Download(ioutil.Discard)
		if err != nil {
			errors = append(errors, fmt.Errorf("Can not download css %s: %v", s.Url().Path, err))
		}
	}

	//validate all scripsts can be downloaded
	for _, s := range c.Images() {
		_, err := s.Download(ioutil.Discard)
		if err != nil {
			errors = append(errors, fmt.Errorf("Can not download image %s: %v", s.Url().Path, err))
		}
	}
	s := len(errors)
	if s > 0 {
		msgs := make([]string, s, s)
		for i, err := range errors {
			msgs[i] = err.Error()
		}
		return fmt.Errorf("Found %d resource errors: \n %s ", s, strings.Join(msgs, "\n"))
	}
	return nil
}
