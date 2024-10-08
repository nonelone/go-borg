package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type NPCharacter struct {
	Name      string `json:"name"`
	Agility   int    `json:"agility"`
	Presence  int    `json:"presence"`
	Strength  int    `json:"strength"`
	Toughness int    `json:"toughness"`
	HitPoints int    `json:"hp"`
	Damage    int    `json:"dmg"`
	ImagePath string `json:"img"`
	Powers    string `json:"powers"`
}

func (npc *NPCharacter) save() error {
	filename := "npc/" + npc.Name + ".json"
	data, _ := json.Marshal(npc)
	return os.WriteFile(filename, data, 0600)
}

func loadNPC(name string) *NPCharacter {
	filename := "npc/" + name + ".json"
	body, _ := os.ReadFile(filename)
	var npc NPCharacter
	_ = json.Unmarshal(body, &npc)
	return &npc
}

func home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home%s!", r.URL.Path[1:])
}

func dm_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love DM %s!", r.URL.Path[1:])
}

func player_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func npc_handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/npc/"):]
	fmt.Println(name)
	npc := loadNPC(name)
	fmt.Fprint(w, "<title>%s</title><body>%s</body>", npc.Name)
}

func main() {
	enemy := NPCharacter{
		Name:      "Enemy",
		Agility:   1,
		Presence:  1,
		Strength:  1,
		Toughness: 1,
		HitPoints: 1,
		Damage:    1,
		ImagePath: "",
		Powers:    "",
	}
	enemy.save()

	http.HandleFunc("/", home_handler)

	http.HandleFunc("/dm", dm_handler)
	http.HandleFunc("/player", player_handler)

	http.HandleFunc("/npc/", npc_handler)

	fmt.Println("Welcome to GöBorg!")
	fmt.Println("Server is up and running at localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
