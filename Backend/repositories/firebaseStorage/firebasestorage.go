package firebasestorage

import "context"

type FirebaseStorage interface {
	GetResourceBucket(ctx context.Context)
}