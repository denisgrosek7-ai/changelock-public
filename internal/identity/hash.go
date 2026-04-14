package identity

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"sort"
	"strings"
)

type DecisionInput struct {
	PolicyBundleHash string
	ImageDigest      string
	RequestID        string
	Decision         string
	Component        string
	Repo             string
	Environment      string
}

func CanonicalFileSetHash(files map[string][]byte) string {
	if len(files) == 0 {
		return ""
	}

	paths := make([]string, 0, len(files))
	for path := range files {
		paths = append(paths, strings.TrimSpace(path))
	}
	sort.Strings(paths)

	hasher := sha256.New()
	for _, path := range paths {
		io.WriteString(hasher, path)
		hasher.Write([]byte{0})
		hasher.Write(files[path])
		hasher.Write([]byte{0})
	}

	return "sha256:" + hex.EncodeToString(hasher.Sum(nil))
}

func DecisionHash(input DecisionInput) string {
	hasher := sha256.New()
	writeField := func(name, value string) {
		io.WriteString(hasher, name)
		hasher.Write([]byte{0})
		io.WriteString(hasher, strings.TrimSpace(value))
		hasher.Write([]byte{0})
	}

	writeField("policy_bundle_hash", input.PolicyBundleHash)
	writeField("image_digest", input.ImageDigest)
	writeField("request_id", input.RequestID)
	writeField("decision", input.Decision)
	writeField("component", input.Component)
	writeField("repo", input.Repo)
	writeField("environment", input.Environment)

	return "sha256:" + hex.EncodeToString(hasher.Sum(nil))
}
