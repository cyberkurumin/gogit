package main

type Command struct {
	UsageLine string
	Short     string
	Long      string
}

type Helper interface {
	Run()
}

var help = &Command{
	UsageLine: "feature|hotfix",
	Short:     "You need to insert one option of execution",
	Long:      "Set here the long description for an empty command",
}

Print
