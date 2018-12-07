# byte_games - My first Golang adventure.
Playing with binary files.
* Included are both a script for creation and a script for reading the generated log.

## createlog
* Generates a transaction log of *count* length in *directory* with *version*.

>> ..\createlog\createlog.exe -h
> Usage of D:\VisualStudioProjects\GoProjects\byte_games\createlog\createlog.exe:
>   -count int
>         Number of records to generate. (default 67)
>   -directory string
>         Location to place file
>   -version int
>         Version to write to header. (default 206)

### TODO for createlog: 
 - [x] Remove debug and prints.
    
## readlog
* Reads a transaction log as formatted by createlog. 
* Several options are available for output.
  * autopaystart: tallies all autopays that were started.
  * autopayend: tallies all autopays that were ended.
  * totall: totals all monetary transactions.
  * totcredit: totals all credit transactions.
  * totdebit: totals all debit transactions.

> > .\readlog.exe -h
> Usage of D:\VisualStudioProjects\GoProjects\byte_games\readlog\readlog.exe:
>     readlog.exe [options] [/path/to/file or C:\path\to\file]
>   -autopayend
>         Total number of autopays ended. (default true)
>   -autopaystart
>         Total number of autopays started. (default true)
>   -totall
>         Total all monetary transactions.
>   -totcredit
>         Total all credit transactions. (default true)
>   -totdebit
>         Total of all debit transactions. (default true)
### TODO for readlog:
- [x] Remove debug and print statements.

## Building.
* Ensure that go is installed and properly configured. 
* Unzip files into your go project directory.
* From inside the directory, run:
>   go build .\readlog\readlog.go
>   go build .\createlog\createlog.go
* Run as demonstrated above or using help option.