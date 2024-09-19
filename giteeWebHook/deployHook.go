package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // 解析 Webhook 请求的内容
        // 这里只是简单的示例，实际情况可能需要更多的处理
        body := r.Body
        defer body.Close()
        // 你可以根据需要解析请求体中的信息

        // 执行 Shell 命令
        cmd := exec.Command("/bin/bash", "-c", "./deploy.sh")
        output, err := cmd.CombinedOutput()
        if err != nil {
            http.Error(w, "Failed to execute command", http.StatusInternalServerError)
            log.Println("Command execution error:", err)
            return
        }

        // 输出执行结果
        fmt.Fprintf(w, "Command executed successfully: %s", output)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/webhook", webhookHandler)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}