package main

import (
	"fmt"
	"log"
	"net/http"
)

const user string = "agent-e11"

func main() {
	http.HandleFunc("/{repo}", redirectRepo)
	http.HandleFunc("/{repo}/{tag}", redirectTag)
	http.HandleFunc("/{repo}/{tag}/{artifact}", redirectArtifact)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func redirectRepo(w http.ResponseWriter, r *http.Request) {
	repo := r.PathValue("repo")
	redirectUrl := fmt.Sprintf("https://github.com/%s/%s/releases/", user, repo)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

func redirectTag(w http.ResponseWriter, r *http.Request) {
	repo := r.PathValue("repo")
	tag := r.PathValue("tag")
	redirectUrl := fmt.Sprintf("https://github.com/%s/%s/releases/tag/%s", user, repo, tag)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

func redirectArtifact(w http.ResponseWriter, r *http.Request) {
	repo := r.PathValue("repo")
	tag := r.PathValue("tag")
	artifact := r.PathValue("artifact")
	redirectUrl := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s", user, repo, tag, artifact)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}
