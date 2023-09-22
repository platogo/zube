package main

import "github.com/platogo/zube/v2"

func main() {
	clientId := "b228eedc-a138-11ec-97ca-53e3bea7ba63"
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOiJiMjI4ZWVkYy1hMTM4LTExZWMtOTdjYS01M2UzYmVhN2JhNjMiLCJ1c2VyX2lkIjo0OTE5MywiYXVkIjoienViZV9hcGkiLCJpYXQiOjE2OTUzNzcwMTAsImV4cCI6MTY5NTQ2MzQxMH0.LyDgIhyRRg50G_1x86I1wt4OuLM0IzIeaYwS7yFRuRs"

	client, _ := zube.NewClient(clientId, accessToken)

	q := zube.Query{
		Filter: zube.Filter{
			Where: map[string]any{"workspace_id": 5689, "priority": 5, "status": "queued", "search_key": []string{"pas", "next", "release"}},
		},
	}

	cards := client.SearchCards(&q)

	for _, card := range cards {
		println(card.Title, card.GithubIssue.Number)
	}
}
