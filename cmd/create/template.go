package create

import "fmt"

// generateScript returns a template string for a new script file.
// It includes a shebang, YASM config block, and a placeholder for the main script.
// shebang: the interpreter directive (e.g., #!/bin/bash)
// comment: the comment prefix for the language (e.g., #, //)
// name: the script name to include in the config block
func generateScriptTemplate(shebang, comment, name string) string {
	return fmt.Sprintf(`%[1]s

%[2]s YASM CONFIG START >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
%[2]s @yasm.title %[3]s
%[2]s @yasm.description 
%[2]s @yasm.tags []
%[2]s @yasm.dependencies []
%[2]s <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<< YASM CONFIG END

%[2]s ################################################
%[2]s ### MAIN SCRIPT STARTS #########################
%[2]s ################################################

%[2]s Write your script here.
`, shebang, comment, name)
}
