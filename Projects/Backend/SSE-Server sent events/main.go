// This Process of Event stream is called Sever Sent Events(SSE)



// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

func main() {
	http.HandleFunc("/event", events)
	http.ListenAndServe(":8080", nil)
}

func events(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/event-stream")

	tokens := []string{"This", "is", "munna", "mia", "Now", "Tell", "About", "you", "?"}

	for _, token := range tokens {
		content := fmt.Sprintf("Data : %s \n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush() // Flush one by one to the User 

		time.Sleep(time.Millisecond*420)
	}

}
