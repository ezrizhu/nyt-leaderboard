package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type WordleLeaderboard struct {
	Leaderboard map[string]int `json:"wordle_leaderboard"`
}

type ConnectionsLeaderboard struct {
	Leaderboard map[string]int `json:"connections_leaderboard"`
}

func wordleLeaderboardInit() WordleLeaderboard {
	file, err := ioutil.ReadFile("wordle.json")
	if err != nil {
		panic(err)
	}

	var wordleLeaderboard WordleLeaderboard
	err = json.Unmarshal(file, &wordleLeaderboard)
	if err != nil {
		panic(err)
	}

	return wordleLeaderboard
}

func connectionsLeaderboardInit() ConnectionsLeaderboard {
	file, err := ioutil.ReadFile("connections.json")
	if err != nil {
		panic(err)
	}

	var connectionsLeaderboard ConnectionsLeaderboard
	err = json.Unmarshal(file, &connectionsLeaderboard)
	if err != nil {
		panic(err)
	}

	return connectionsLeaderboard
}

func (wl *WordleLeaderboard) addScore(user string, score int) int {
	if wl.Leaderboard == nil {
		wl.Leaderboard = make(map[string]int)
	}

	wl.Leaderboard[user] += score
	wl.flush()

	return wl.Leaderboard[user]
}

func (cl *ConnectionsLeaderboard) addScore(user string, score int) int {
	if cl.Leaderboard == nil {
		cl.Leaderboard = make(map[string]int)
	}

	cl.Leaderboard[user] += score
	cl.flush()

	return cl.Leaderboard[user]
}

func (wl *WordleLeaderboard) flush() {
	data, err := json.Marshal(wl)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile("wordle.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (cl *ConnectionsLeaderboard) flush() {
	data, err := json.Marshal(cl)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = ioutil.WriteFile("connections.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func (wl *WordleLeaderboard) showLeaderboard() string {
	if wl.Leaderboard == nil {
		return "Leaderboard is empty."
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range wl.Leaderboard {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	leaderboardText := ""
	i := 1
	for j, kv := range ss {
		if j != 0 {
			if ss[j-1].Value != kv.Value {
				i += 1
			}
		}
		leaderboardText += fmt.Sprintf("%d, <@%s>: %d\n", i, kv.Key, kv.Value)
	}

	return leaderboardText
}

func (cl *ConnectionsLeaderboard) showLeaderboard() string {
	if cl.Leaderboard == nil {
		return "Leaderboard is empty."
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range cl.Leaderboard {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	leaderboardText := ""
	i := 1
	for j, kv := range ss {
		if j != 0 {
			if ss[j-1].Value != kv.Value {
				i += 1
			}
		}
		leaderboardText += fmt.Sprintf("%d, <@%s>: %d\n", i, kv.Key, kv.Value)
	}

	return leaderboardText
}
