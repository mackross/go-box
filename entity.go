package box

import (
	"encoding/json"
	"errors"
	"time"
)

//  Represents both mini folder and mini file.
type Entity struct {
	SequenceId string `json:"sequence_id,omitempty"` // A unique ID for use with the /events endpoint.
	Name       string `json:"name,omitempty"`        // The name of the entity.
	Id         string `json:"id,omitempty"`          // The id of the entity.
	ETag       string `json:"etag,omitempty"`        // A unique string identifying the version of this entity.
	Type       string `json:"type,omitempty"`        // Type of entity
}

// IsFolder checks if the given entity is a folder
func (e *Entity) IsFolder() bool {
	if e.Type == "folder" {
		return true
	} else {
		return false
	}
}

// toFolder converts the given entity to a folder. Only attributes present in
// the entity are populated rest are untouched.
func (e *Entity) toFolder(f *Folder) error {
	if !e.IsFolder() {
		return errors.New("Entity is not a folder")
	}
	f.Id = e.Id
	f.Name = e.Name
	f.ETag = e.ETag
	f.SequenceId = e.SequenceId
	return nil
}

// IsFile checks if the given entity is a file.
func (e *Entity) IsFile() bool {
	if e.Type == "file" {
		return true
	} else {
		return false
	}
}

// toFile converts the given entity to a file. Only attributes present in
// the entity are populated rest are untouched.
func (e *Entity) toFile(f *File) error {
	if !e.IsFile() {
		return errors.New("Entity is not a file")
	}
	f.Id = e.Id
	f.Name = e.Name
	f.ETag = e.ETag
	f.SequenceId = e.SequenceId
	return nil
}

type BoxTime time.Time

// UnmarshalJSON unmarshals a time according to the Dropbox format.
func (bt *BoxTime) UnmarshalJSON(data []byte) error {
	if data == nil || string(data) == "null" {
		return nil
	}
	var s string
	var err error
	var t time.Time
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}
	if t, err = time.ParseInLocation(time.RFC3339, s, time.UTC); err != nil {
		return err
	}
	if t.IsZero() {
		*bt = BoxTime(time.Time{})
	} else {
		*bt = BoxTime(t)
	}
	return nil
}

// MarshalJSON marshals a time according to the Dropbox format.
func (bt BoxTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(bt).Format(time.RFC3339))
}

type Permission struct {
	Download bool `json:"can_download,omitempty"`
	Preview  bool `json:"can_preview,omitempty"`
	Upload   bool `json:"can_upload,omitempty"`
	Comment  bool `json:"can_comment,omitempty"`
	Rename   bool `json:"can_rename,omitempty"`
	Delete   bool `json:"can_delete,omitempty"`
	Share    bool `json:"can_share,omitempty"`
	SetShare bool `json:"can_set_share_access,omitempty"`
}

type Collection struct {
	Count  int      `json:"total_count,omitempty"`
	Entry  []Entity `json:"entries,omitempty"`
	Limit  int      `json:"limit,omitempty"`
	Offset int      `json:"offset,omitempty"`
}

type BoxLock struct {
	Id        string   `json:"id,omitempty"`
	CreatedBy string   `json:"created_by,omitempty"`
	CreatedAt *BoxTime `json:"created_at,omitempty"`
	ExpiresAt *BoxTime `json:"expires_at,omitempty"`
	Download  bool     `json:"is_download_prevented,omitempty"`
}

type SharedObject struct {
	Url           string      `json:"url,omitempty"`
	DownloadUrl   string      `json:"download_url,omitempty"`
	VanityUrl     string      `json:"vanity_url,omitempty"`
	HasPassword   bool        `json:"is_password_enabled,omitempty"`
	UnsharedAt    *BoxTime    `json:"unshared_at,omitempty"`
	DownloadCount int         `json:"download_count,omitempty"`
	PreviewCount  int         `json:"preview_count,omitempty"`
	Access        string      `json:"access,omitempty"`
	Permission    *Permission `json:"permissions,omitempty"`
}

type UploadEmail struct {
	Access string `json:"access,omitempty"`
	Email  string `json:"email,omitempty"`
}
