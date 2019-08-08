# Tools

Small, standalone, tools to assist in day to day workflows. 

## Contributing

Please do.

## Build

### .go

Language: https://golang.org

	go build myfile.go

## Examples

Examples are shown with a prompt to indicate the environment and the line after the command, including expected output.

Prompts key:

	>	Powershell (may not work under cmd)
	$	Bash (WSL)

### encode

URL-encode a file named .vimrc from stdin:

	$ encode < .vimrc
	colo+robpike%0A%0A
	$ 

URL-encode, with no trailing newline, a file named .gitconfig and copy to the clipboard:

	> encode -n .gitconfig | clip
	> 

### decode

URL-decode some text:

	> echo %23%21%2Fusr%2Fbin%2Fenv+bash | decode
	#!/usr/bin/env bash
	
	
	> 

### vmdk2vhdx.ps1

Convert foo.vmdk into bar.vhdx:

	> vmdk2vhdx foo.vmdk bar.vhdx
	... (output omitted) 
	> 
