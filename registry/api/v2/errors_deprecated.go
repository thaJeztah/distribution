package v2

import (
	v2errs "github.com/docker/distribution/registry/api/v2/errors"
)

var (
	// ErrorCodeDigestInvalid is returned when uploading a blob if the
	// provided digest does not match the blob contents.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeDigestInvalid
	ErrorCodeDigestInvalid = v2errs.ErrorCodeDigestInvalid

	// ErrorCodeSizeInvalid is returned when uploading a blob if the provided
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeSizeInvalid
	ErrorCodeSizeInvalid = v2errs.ErrorCodeSizeInvalid

	// ErrorCodeNameInvalid is returned when the name in the manifest does not
	// match the provided name.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeNameInvalid
	ErrorCodeNameInvalid = v2errs.ErrorCodeNameInvalid

	// ErrorCodeTagInvalid is returned when the tag in the manifest does not
	// match the provided tag.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeTagInvalid
	ErrorCodeTagInvalid = v2errs.ErrorCodeTagInvalid

	// ErrorCodeNameUnknown when the repository name is not known.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeNameUnknown
	ErrorCodeNameUnknown = v2errs.ErrorCodeNameUnknown

	// ErrorCodeManifestUnknown returned when image manifest is unknown.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeManifestUnknown
	ErrorCodeManifestUnknown = v2errs.ErrorCodeManifestUnknown

	// ErrorCodeManifestInvalid returned when an image manifest is invalid,
	// typically during a PUT operation. This error encompasses all errors
	// encountered during manifest validation that aren't signature errors.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeManifestInvalid
	ErrorCodeManifestInvalid = v2errs.ErrorCodeManifestInvalid

	// ErrorCodeManifestUnverified is returned when the manifest fails
	// signature verification.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeManifestUnverified
	ErrorCodeManifestUnverified = v2errs.ErrorCodeManifestUnverified

	// ErrorCodeManifestBlobUnknown is returned when a manifest blob is
	// unknown to the registry.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeManifestBlobUnknown
	ErrorCodeManifestBlobUnknown = v2errs.ErrorCodeManifestBlobUnknown

	// ErrorCodeBlobUnknown is returned when a blob is unknown to the
	// registry. This can happen when the manifest references a nonexistent
	// layer or the result is not found by a blob fetch.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeBlobUnknown
	ErrorCodeBlobUnknown = v2errs.ErrorCodeBlobUnknown

	// ErrorCodeBlobUploadUnknown is returned when an upload is unknown.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeBlobUploadUnknown
	ErrorCodeBlobUploadUnknown = v2errs.ErrorCodeBlobUploadUnknown

	// ErrorCodeBlobUploadInvalid is returned when an upload is invalid.
	// Deprecated: use github.com/docker/distribution/registry/api/v2/errors.ErrorCodeBlobUploadInvalid
	ErrorCodeBlobUploadInvalid = v2errs.ErrorCodeBlobUploadInvalid
)
