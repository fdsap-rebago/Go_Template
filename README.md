This is a GoLang Template that will be used 
as FDS ASYA PHILIPPINES INC. standard format.

#******************************************#
#*********** GoLang is Required ***********#
#******************************************#

1.  Download the template by using either:
        - git clone https://github.com/FDSAP-Git-Org/Go_Template.git
        or
        - Download the repository at Github
2.  Place the template in your GOROOT
        - To know your GOROOT, you can use this command:
                - go env
3.  Check or edit .env configuration before running the template
4.  Change the template to your desired application
    name by following this instruction:
        - Delete go.mod and go.sum
        - Initialize the application:
                - go mod init <Application name>
        - Change the imports from Template folder
          to your <Application name>
                - "Template/middleware"
                to
                - "<Application name>/middleware"
        - Get the application packages:
                - go mod tidy
        - Then run the application:
                - go run main.go
5.  Check if the application is running by using postman
    or visiting this link on a browser:
        - http://127.0.0.1:8000/api/public/v1
        or
        - localhost:8000/api/public/v1
6.  To deploy the application on a server,
    follow this instructions and command:
        - Upload the source code at the server
        - Navigate to the root folder
        - Run this command:
                - go build <Application Name> -o
                - nohup go run main.go
