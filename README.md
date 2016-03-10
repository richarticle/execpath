# execpath
This is a small utility for getting the path or directory of the executable 
binary. You can also change the working directory easily to the directory 
of the executable binary. Currentlly only Linux and Windows are supported.

# Usage
import "github.com/richarticle/execpath"
	
	path, err := ExecPath()
	dir, err := ExecDir()
	err := execpath.EnterExecDir()

