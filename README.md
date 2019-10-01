# cw-go-sayhelloName
This is a basic go app to personalize a greeting developed using the codewind extension.
When the app is running, your can append `/hello` to the URL to get a little more information. Or you can go straight into the fun part and append `/hello/<name>` to see the app return `Hello <name-input>. Today is <date> <time> <day>.`

### Pre-reqs
Have your machine set up for development with Golang.
https://hackernoon.com/basics-of-golang-for-beginners-6bd9b40d79ae
- have GOPATH in .bash_profile
- use min Go v1.12.10
- https://github.com/gin-gonic/gin

### References
gin-gonic / gin
https://github.com/gin-gonic/gin#parameters-in-path

makefile
https://tutorialedge.net/golang/makefiles-for-go-developers/
https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html

time formatting
https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html
