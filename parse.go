package onenote

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
)

func (client *Client) ParseResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}

func (client *Client) PageContent(url string) (*http.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	return client.DoRequest(req)
}

// ParseTasksWithIdAndTagAndText Parse tasks with id, tag and text
func (client *Client) ParseTasksWithIdAndTagAndText(doc *goquery.Document) []Task {
	var tasks []Task
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		tag, ok := s.Attr("data-tag")
		if ok {
			if tag == "to-do:completed" || tag == "to-do" {
				if id, ok := s.Attr("id"); ok {
					if tag == "to-do:completed" {
						tasks = append(tasks, Task{
							Tag:    tag,
							Text:   s.Text(),
							Id:     id,
							Status: "completed",
						})
					} else {
						tasks = append(tasks, Task{
							Tag:    tag,
							Text:   s.Text(),
							Id:     id,
							Status: "uncompleted",
						})
					}
				}
			}
		}
	})

	return tasks
}
