### Simple html parser

This simplistic parser handles html.

The implementation validates html structure using an in-memory tree.

##### Run

To run, pipe input file to program via standard input.

Program returns zero on success and 1 otherwise.

i.e.

    go run main.go < test1.html
    echo $?

    # or

    curl -s https://autobad-sb.autodesk.com/swagger/index.html | go run main.go
    echo $?
