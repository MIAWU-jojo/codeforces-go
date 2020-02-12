package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

// 生成比赛模板（需要在创建目录之后）
func GenContestTemplates(contestID string, overwrite bool) error {
	rootPath := fmt.Sprintf("../../dash/%s/", contestID)
	openedOneFile := false
	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := filepath.Join(path, fileName)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
			if !openedOneFile {
				openedOneFile = true
				open.Run(absPath(goFilePath))
			}
		}
		return nil
	}); err != nil {
		return err
	}

	_, err := strconv.Atoi(contestID)
	if isCF := err == nil; isCF {
		tips := fmt.Sprintf("cd %[1]s\ncf submit %[1]s a a/main.go\n", contestID)
		if err := ioutil.WriteFile(rootPath+"tips.txt", []byte(tips), 0644); err != nil {
			return err
		}
	}
	return nil
}

// 生成单道题目的模板（Codeforces）
// https://codeforces.com/problemset/problem/1293/C
// https://codeforces.com/problemset/status/1291/problem/D
// https://codeforces.com/gym/102253/problem/C
func GenCodeforcesNormalTemplates(problemURL string) error {
	contestID, problemID := parseProblemURL(problemURL)
	statusURL := fmt.Sprintf("https://codeforces.com/problemset/status/%s/problem/%s", contestID, problemID)
	open.Start(problemURL)
	open.Start(statusURL)

	problemID = contestID + problemID
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF%[1]s(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() { CF%[1]s(os.Stdin, os.Stdout) }
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF%[1]s(t *testing.T) {
	// just copy from website
	rawText := `+"`\n`"+`
	testutil.AssertEqualCase(t, rawText, 0, CF%[1]s)
}
`, problemID)

	const rootPath = "../../main/"
	mainFilePath := rootPath + problemID + ".go"
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		return fmt.Errorf("文件已存在！")
	}
	if err := ioutil.WriteFile(mainFilePath, []byte(mainStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(mainFilePath))
	testFilePath := rootPath + problemID + "_test.go"
	if err := ioutil.WriteFile(testFilePath, []byte(mainTestStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(testFilePath))
	return nil
}

// 生成单道题目的模板（非 Codeforces）
// rootPath like "../../nowcoder/2720/"
func GenNormalTemplates(rootPath string, overwrite bool) error {
	for i := 'a'; i <= 'h'; i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := dir + fileName
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
		}
	}
	return nil
}
