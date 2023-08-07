module runnermodule

go 1.20

require (
	fbservice v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

require github.com/huandu/facebook/v2 v2.7.0 // indirect

replace fbservice => ../fbservice
