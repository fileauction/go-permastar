package model

type DirInfo struct {
	Path string `json:"path"`
	Cid  string `json:"cid,omitempty"`
}

type FileInfo struct {
	Path string `json:"path"`
	Cid  string `json:"cid,omitempty"`
	Data []byte `json:"data,omitempty"`
}
