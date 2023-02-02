# COMPLETED TEST TASK

> golang version: 1.18 <br/>
> dev os: Pop!_OS 22.04

## How to run?
`> cd; git clone https://github.com/yabifurkator/YADRO_testtask.git` <br/>
`> cd YADRO_testtask` <br/>
`> go run main.go some_input_file.csv` <br/>
> many different examples of input files are in the directory **tests/csv**

## How to run unit tests?
`> cd; cd YADRO_testtask` <br/>
`> go test -v tests/unit_test.go` <br/> <br>

## The subtleties of program work.
1. Only EN (lower and upper case) letters are allowed in column names
2. Only digits (0-9) are allowed in row numbers
3. Duplicates in column names and row numbers are not allowed
4. In column names, in row numbers, in cells, space characters are skipped everywhere. Therefore, "= A1 + B2" turns into "=A1+B2", and the element named "  " is an empty element.
5. Empty column names (except for the first one), empty row numbers, empty cell contents are considered invalid.
6. Integer division is used at division points. <br><br>

###### Contact me
> telegram `https://t.me/leonidganenko` <br>
> gmail `yabifurkator@gmail.com`
