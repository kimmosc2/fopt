go build -ldflags "-X 'cmd.ProgramInfo.PlatForm=windows' -X 'cmd.ProgramInfo.Version=abc' -X 'cmd.ProgramInfo.GoVersion=$(go version)' -X 'cmd.ProgramInfo.Author=BuTn<github.com/kimmosc2> -X 'cmd.ProgramInfo.BuildTime=$(git show -s --format=%cd)'" -o fopt.exe main.go