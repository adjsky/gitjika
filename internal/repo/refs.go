package repo

import "time"

type Ref struct {
	Name    string
	Message string
	Author  string
	Date    time.Time
}

func (repo Repo) Refs() {

}
