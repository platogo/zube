package main

import "github.com/platogo/zube/v2"

func main() {
	clientId := "b228eedc-a138-11ec-97ca-53e3bea7ba63"
	accessToken := "ad1d83b95cad448351abebb2ad1d83b95cad448351abebb2ad1d83b95cad448351abebb2ad1d83b95cad448351abebb2"

	client, _ := zube.NewClient(clientId, accessToken)

	projectId := 5953

	epics := client.FetchEpics(projectId)

	for _, epic := range epics {
		println(epic.Title)
	}
}
