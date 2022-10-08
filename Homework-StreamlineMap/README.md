# Homework-StreamlineMap
A CI process has the stages as below.  
Each key in the map is the name of a stage in a CI process.  
Each value in the map is a slice containing the names of other stages that the current stage depends on.  
```go
var buildStages = map[string][]string{
	"checkout": {},                          // Clone the repository
	"lint":     {"checkout"},                // Run a linter
	"build":    {"checkout"},                // Build binaries
	"test":     {"checkout"},                // Run tests
	"package":  {"build"},                   // Package the binaries into a package
	"publish":  {"test", "lint", "package"}, // Publish the packages to a server for storing build artifacts
	"deploy":   {"publish"},                 // Pick the package from where it was published and deploy it
}
```
The repo contains the logic to print the stage names in an order which is consistent with the dependencies between the stages.  
Time complexity: O(n)  
Space complexity: O(n)  
with n being the map length.