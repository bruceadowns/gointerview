### Simple xml parser

This simplistic parser handles element-normal xml. Attributes handled, but ignored.

The implementation validates xml structure by leveraging the properties of a LIFO stack.

##### Run

To run, pipe input file to program via standard input.

Program returns zero on success and 1 otherwise.

i.e.

    go run main.go < test1.xml
    echo $?

    # or

    cat test1.xml | go run main.go
    echo $?
