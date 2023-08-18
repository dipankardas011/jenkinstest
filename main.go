package main

import (
	"github.com/dipankardas011/jenkinstest/api/azure"
	"github.com/dipankardas011/jenkinstest/api/civo"
)

func main() {
	civo.Civo()
	azure.Azure()
}
