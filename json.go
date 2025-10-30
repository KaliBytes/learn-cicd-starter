func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    dat, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Error marshalling JSON: %v", err)
        w.WriteHeader(500)
        return
    }

    w.WriteHeader(code)
    if _, err := w.Write(dat); err != nil {
        log.Printf("error writing response body: %v", err)
    }
}

