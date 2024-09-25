package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Laudo struct {
    IDSolicitacao         int    `json:"id_solicitacao"`
    CodigoExame           string `json:"codigo_exame"`
    Status                string `json:"status_exame"` 
    DataAtualizacaoStatus string `json:"data_atualizacao_status"`
}

type Payload struct {
    Status []Laudo `json:"status"`
}

func handleNotificacao(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var payload Payload
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
        return
    }

    fmt.Printf("Recebido: %+v\n", payload)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Notificação recebida com sucesso"}`))
}

func main() {
    http.HandleFunc("/api/notificacoes", handleNotificacao)

    fmt.Println("Servidor de notificações rodando em http://localhost:8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        fmt.Println("Erro ao iniciar servidor:", err)
    }
}