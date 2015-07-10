package store

import (
	"github.com/GabbyyLS/go-school/musicstore/model"
	"golang.org/x/net/context"
)

type Store interface {
	GetArtistByName(name string) (*model.Artist, error)
	// TODO: GetAlbumByTitle
	// TODO: GetTrackByTitle

	// TODO: GetArtistById
	// TODO: GetAlbumById
	// TODO: GetTrackById

	// TODO: GetArtistsByTrack
	// TODO: GetAlbumsByTrack

	// TODO: GetAlbumRating

	// TODO: CreateArtist
	// TODO: CreateAlbum
	// TODO: CreateTrack

	// TODO: UpdateArtist
	// TODO: UpdateAlbum
	// TODO: UpdateTrack

	// TODO: RemoveArtistByID
	// TODO: RemoveAlbumByID
	// TODO: RemoveTrackByID
}

// Пример:
func GetArtistByName(ctx context.Context, name string) (*model.Artist, error) {
	return FromContext(ctx).GetArtistByName(name)
}

// TODO: Прописать все остальные функции