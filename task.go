package onenote

import "github.com/PuerkitoBio/goquery"

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
