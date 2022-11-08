package onenote

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func (client *Client) GetTasks(doc *goquery.Document) ([]string, []string) {
	var notCompletedTasks []string
	var completedTasks []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		tag, ok := s.Attr("data-tag")
		if ok {
			if tag == "to-do:completed" {
				notCompletedTasks = append(notCompletedTasks, s.Text())
				completedTasks = append(completedTasks, s.Text())
			} else if tag == "to-do" {
				notCompletedTasks = append(notCompletedTasks, s.Text())
			}
		}
	})

	return notCompletedTasks, completedTasks
}

func (client *Client) GetTaskIDs(doc *goquery.Document) []string {
	var taskIDs []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		tag, ok := s.Attr("data-tag")
		if ok {
			if tag == "to-do:completed" || tag == "to-do" {
				id, ok := s.Attr("id")
				if ok {
					taskIDs = append(taskIDs, id)
				}
			}
		}
	})

	return taskIDs
}

func (client *Client) UpdateTask(data []PageContent, url string) (*http.Response, []byte, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, nil, err
	}
	return client.DoRequest(req)
}
