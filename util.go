package main

func templateList(names ...string) []string {
	fullNames := []string{}
	for _, name := range names {
		fullNames = append(fullNames, "./templates/"+name+".html")
	}
	return fullNames
}
