
package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// gitCommit is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	gitCommit string
	// buildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	//It should be set during build via -ldflags.
	buildDate string
	// version is the aws-auth package version
	pkgVersion string = "0.2.0"
)

// Info holds the information related to descheduler app version.
type Info struct {
	PackageVersion string `json:"pkgVersion"`
	GitCommit      string `json:"gitCommit"`
	BuildDate      string `json:"buildDate"`
	GoVersion      string `json:"goVersion"`
	Compiler       string `json:"compiler"`
	Platform       string `json:"platform"`
}

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() Info {
	return Info{
		GitCommit:      gitCommit,
		BuildDate:      buildDate,
		GoVersion:      runtime.Version(),
		Compiler:       runtime.Compiler,
		Platform:       fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		PackageVersion: pkgVersion,
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of aws-auth",
	Long:  `Prints the version of aws-auth.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("aws-auth version %+v\n", Get())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
