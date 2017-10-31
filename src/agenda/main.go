package main

import  "cmd"

func init() {
	log.SetFlags(log.Lshortfile)
	cmd.EnsureAgendaDir()
}

func main() {
	cmd.Execute()
}