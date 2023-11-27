# testing

go test  
go test -v // to execute test functions and provide verbose output.  
go test -v -run="French|Canal"  // to specify which tests to run, matching the regular expression.  

go test -v -run=Coverage -coverprofile=c.out // to generate coverage report.  
go tool cover -html=c.out -o coverage.html // to generate html report from coverage report.  
go test -run=Coverage -coverprofile=c.out -covermode=count -v // to generate coverage report with count.

go test -bench . // to run benchmarks.
go test -bench . -benchmen // to run benchmarks with memory profiling.

