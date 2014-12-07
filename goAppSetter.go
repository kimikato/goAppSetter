package main

import (
    "fmt"
    "os"
    "strconv"
    "code.google.com/p/go.text/encoding/japanese"
    "code.google.com/p/go.text/transform"
    "github.com/jessevdk/go-flags"
    "path/filepath"
    "io/ioutil"
    "net/http"
    "gopkg.in/yaml.v2"
    "strings"
    "regexp"
)

type Options struct {
}

type Servers struct {
    Sever []Server
}

type Server struct {
    Scheme  string
    Host    string
    Port    int
    Path    string
}

var options Options

func main() {
    parser := flags.NewParser(&options, flags.Default)
    parser.Name = "goAppSetter"
    parser.Usage = "[OPTIONS] TARGET_FILE"

    args, _ := parser.Parse()

    // 引数がひとつもなければヘルプを表示する
    if len(args) == 0 {
        parser.WriteHelp(os.Stdout)
        os.Exit(1)
    }

    // 現在のパス
    currentDir, _ := filepath.Abs(".")
    targetFilePath := filepath.Join(currentDir, args[0])

    // ファイルの存在確認
    if !FileExists(targetFilePath) {
        // 存在していない
        fmt.Printf("Not found. [%s]\n", targetFilePath)
        os.Exit(1)
    }

    // YAMLファイルの読み込み
    buf, err := ioutil.ReadFile(targetFilePath)
    if err != nil {
        fmt.Println("Can not read yaml file.")
        os.Exit(1)
    }

    m := make(map[interface{}]interface{})
    err = yaml.Unmarshal(buf, &m)
    if err != nil {
        panic(err)
        os.Exit(1)
    }

    servers := m["servers"].([]interface {})
    for _, v := range servers {
        getString := func(key string) string {
            result, _ := v.(map[interface {}]interface {})["server"].(map[interface {}]interface {})[key].(string)
            return result
        }
        getValue := func(key string) int {
            result, _ := v.(map[interface {}]interface {})["server"].(map[interface {}]interface {})[key].(int)
            return result
        }

        if ! HttpRequest(getString("scheme"), getString("host"), getValue("port"), getString("path")) {
            fmt.Println("Access Failed.")
        } else {
            fmt.Println("Successed.")
        }
    }
}


func HttpRequest(scheme string, host string, port int, path string) bool {
    url := scheme + "://" + host + ":" + strconv.Itoa(port) + path
    fmt.Printf("URL : %s\n", url)

    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return false
    }

    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()

    // 文字コードの変換（ShiftJIS -> UTF-8）
    result, _ := SJIStoUTF8(string(body))

    re, err := regexp.Compile("Application変数を再セットしました。")
    if err != nil {
        fmt.Println(err)
        return false
    }
    match := re.MatchString(result)

    if match {
        return true
    } else {
        return false
    }
}

func SJIStoUTF8(base string) (string, error) {
    result, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(base), japanese.ShiftJIS.NewDecoder()))
    if err != nil {
        return "", err
    }
    return string(result), err
}

func FileExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}


func StringBuild(base, str string) string {
    buffer := []byte(base)
    buffer = append(buffer, []byte(str)...)
    return string(buffer)
}