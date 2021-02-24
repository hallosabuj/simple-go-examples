1. GOROOT => Where go executable resides
2. GOPATH => Where all the package source code will be downloaded
             bin/ => For executables
             src/ => For source codes
    
    It can also be used for creating workspace,
    and there will be following folders....
        a. bin/
        b. pkg/
        c. src/

3. To run a go file
   go run filename.go

4. To build a go package
   go build package_name
   This package folder must present inside the src folder of GOROOT or GOPATH

   To build a single file
   go build filename.go

5. To make the executable as a command use following command
   go install package_name
   This package folder must present inside the src folder of GOROOT or GOPATH
   This command will place the executable inside the bin folder inside workspace