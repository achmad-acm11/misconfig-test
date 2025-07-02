package dto

import "time"

type ResultOutputFile struct {
	SchemaVersion int    `json:"SchemaVersion"`
	ArtifactName  string `json:"ArtifactName"`
	ArtifactType  string `json:"ArtifactType"`
	Metadata      struct {
		ImageConfig struct {
			Architecture string    `json:"architecture"`
			Created      time.Time `json:"created"`
			Os           string    `json:"os"`
			Rootfs       struct {
				Type    string      `json:"type"`
				DiffIds interface{} `json:"diff_ids"`
			} `json:"rootfs"`
			Config struct {
			} `json:"config"`
		} `json:"ImageConfig"`
	} `json:"Metadata"`
	Results []Results `json:"Results"`
}

type Results struct {
	Target         string `json:"Target"`
	Class          string `json:"Class"`
	Type           string `json:"Type"`
	MisconfSummary struct {
		Successes  int `json:"Successes"`
		Failures   int `json:"Failures"`
		Exceptions int `json:"Exceptions"`
	} `json:"MisconfSummary"`
	Misconfigurations []Misconfiguration `json:"Misconfigurations"`
}

type Misconfiguration struct {
	Type        string   `json:"Type" bson:"type,omitempty"`
	ID          string   `json:"ID" bson:"id,omitempty"`
	Title       string   `json:"Title" bson:"title,omitempty"`
	Description string   `json:"Description" bson:"description,omitempty"`
	Message     string   `json:"Message" bson:"message,omitempty"`
	Query       string   `json:"Query" bson:"query,omitempty"`
	Resolution  string   `json:"Resolution" bson:"resolution,omitempty"`
	Severity    string   `json:"Severity" bson:"severity,omitempty"`
	PrimaryURL  string   `json:"PrimaryURL" bson:"primary_url,omitempty"`
	References  []string `json:"References" bson:"references,omitempty"`
	Status      string   `json:"Status" bson:"status,omitempty"`
	Layer       struct {
	} `json:"Layer" bson:"layer,omitempty"`
	CauseMetadata struct {
		Resource  string `json:"Resource" bson:"resource,omitempty"`
		Provider  string `json:"Provider" bson:"provider,omitempty"`
		Service   string `json:"Service" bson:"service,omitempty"`
		StartLine int    `json:"StartLine" bson:"start_line,omitempty"`
		EndLine   int    `json:"EndLine" bson:"end_line,omitempty"`
		Code      struct {
			Lines []struct {
				Number     int    `json:"Number" bson:"number,omitempty"`
				Content    string `json:"Content" bson:"content,omitempty"`
				IsCause    bool   `json:"IsCause" bson:"is_cause,omitempty"`
				Annotation string `json:"Annotation" bson:"annotation,omitempty"`
				Truncated  bool   `json:"Truncated" bson:"truncated,omitempty"`
				FirstCause bool   `json:"FirstCause" bson:"first_cause,omitempty"`
				LastCause  bool   `json:"LastCause" bson:"last_cause,omitempty"`
			} `json:"Lines" bson:"lines,omitempty"`
		} `json:"Code" bson:"code,omitempty"`
	} `json:"CauseMetadata" bson:"cause_metadata,omitempty"`
}
