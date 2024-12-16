Package : A folder that contains multiple GO files.
Module : A collection of packages.

** About GO ** 
1. Statically Typed : Data types have to be either declared or inferred by initialisation.
2. Strongly Typed : Can't do operations between different data type variables or objects.
3. GO is compiled.
4. Fast compile time.
5. Built in Concurrency. // Goroutines
6. Simplicity  (Easy, concise syntax and automatic garbage collection etc.)

>> Every GO project is a module that we track using 'go.mod' file.
>> Create 'go.mod' using -> ` go mod init <module_name> ` . Here <module_name> is generally the path of the project / module in github repo.

>> ` go mod tidy ` : It refreshes the go.mod file to track dependencies.
>> ` go build <file_path> ` : It compiles the specified file.
>> ` ./file.exe ` : It executes the .exe binary file to yield the output.
>> ` go run <file_path> ` : It builds and runs the file altogether to give output directly.

>> Every GO file belongs to a package that we specify at the top of the file using ` package <name> `
Every file in the folder belongs to the same package.
>> The main function's name shall be the same as of it's package name.



