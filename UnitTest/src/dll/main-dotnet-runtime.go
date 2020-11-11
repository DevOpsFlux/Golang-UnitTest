package main

import (
	"github.com/matiasinsaurralde/go-dotnet"

	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi, I'll initialize the .NET runtime.")

	/*
		If you don't set the TRUSTED_PLATFORM_ASSEMBLIES, it will use the default tpaList value.
		APP_PATHS & NATIVE_DLL_SEARCH_DIRECTORIES use the path of the current program,
		this makes it easier to load an assembly, just put the DLL in the same folder as your Go binary!
		You're free to override them to fit your needs.
	*/

	properties := map[string]string{
		// "TRUSTED_PLATFORM_ASSEMBLIES": "/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/mscorlib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Private.CoreLib.ni.dll",
		// "APP_PATHS":                     "/Users/matias/dev/dotnet/lib/HelloWorld",
		// "NATIVE_DLL_SEARCH_DIRECTORIES": "/Users/matias/dev/dotnet/lib/HelloWorld",
		//"TRUSTED_PLATFORM_ASSEMBLIES":   "C:/Program Files/dotnet/shared/Microsoft.NETCore.App/2.1.22/mscorlib.dll:C:/Program Files/dotnet/shared/Microsoft.NETCore.App/2.1.22/System.Private.CoreLib.dll",
		"TRUSTED_PLATFORM_ASSEMBLIES":   "C:/Program Files (x86)/Reference Assemblies/Microsoft/Framework/.NETFramework/v4.5.1/mscorlib.dll:C:/Program Files (x86)/Reference Assemblies/Microsoft/Framework/.NETFramework/v4.5.1/System.Private.CoreLib.dll",
		"APP_PATHS":                     "D:/Works/Go/src/dll",
		"NATIVE_DLL_SEARCH_DIRECTORIES": "D:/Works/Go/src/dll",
	}

	/*
		CLRFilesAbsolutePath sets the SDK path.
		In case you don't set this parameter, this package will try to find the SDK using a list of common paths.
		It seems to find the right paths under Linux & OSX, feel free to override this setting (like the commented line).
	*/

	runtime, err := dotnet.NewRuntime(dotnet.RuntimeParams{
		Properties:           properties,
		CLRFilesAbsolutePath: "C:/Program Files (x86)/Reference Assemblies/Microsoft/Framework/.NETFramework/v4.5.1",
		//CLRFilesAbsolutePath: "C:/Program Files/dotnet/shared/Microsoft.NETCore.App/2.1.22",
		// CLRFilesAbsolutePath: "/usr/share/dotnet/shared/Microsoft.NETCore.App/1.0.0"
	})
	//defer runtime.Shutdown()

	if err != nil {
		fmt.Println("Something bad happened! :(")
		os.Exit(1)
	}

	fmt.Println("Runtime loaded.")

	//SayHello := runtime.CreateDelegate("HelloWorld", "HelloWorld.HelloWorld", "Hello")
	SayHello := runtime.CreateDelegate("GoNetLibTest", "GoNetLibTest.GoNetLibCls", "Hello")
	//SayHello := runtime.CreateDelegate("CryptoNetCom", "CryptoNetCom.CryptLib", "SeedCBCPADEncrypt")

	// this will call HelloWorld.HelloWorld.Hello :)
	SayHello()
}

/*
package main

import (
	"fmt"
	"syscall"
)

func main() {
	var mod = syscall.NewLazyDLL("MathForGo.dll")
	var proc = mod.NewProc("Add")
	proc.Call(2, 3)
	fmt.Printf("%v", proc)
}
*/
