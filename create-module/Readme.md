# Create modules and Call Main Function

- First of all, Create directory which has module names
- In this example, module name is greetings.
- And, run this command which is initialized .mod file;
```
    go mod init [package-name]/[module-name]
```
- Create main file and call function.
- But, in this stage, main file also have mod file which is using for add modules from directory
- Run this commands;
```
    go mod init [package-name]/main
    //If the main file and modules single directory
    go mod edit -replace [package-name]/[module-name]=./[module-name]
    //If the main file and modules same directory
    go mod edit -replace [package-name]/[module-name]=../[module-name]
```