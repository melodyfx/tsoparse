package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main2() {
    var number Timestamp

    for {
        fmt.Print("输入TSO: ")
        _, err := fmt.Scanf("%d", &number)

        if err != nil {
            fmt.Println("无效的输入:", err)
            return
        }
        p, l := ParseTS(number)
        fmt.Printf("物理时间:%s\n", p)
        fmt.Printf("逻辑时间:%d\n", l)
    }

}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("输入TSO: ")
        text, _ := reader.ReadString('\n')

        text = strings.TrimSpace(text)
        if text == "exit" {
            fmt.Println("退出")
            break
        }

        num, err := strconv.ParseUint(text, 10, 64)
        if err != nil {
            fmt.Printf("无效的输入:%s\n", err)
            continue
        }
        p, l := ParseTS(num)
        fmt.Printf("物理时间:%s\n", p)
        fmt.Printf("逻辑时间:%d\n", l)
        fmt.Println()
    }
}
