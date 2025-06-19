package graph

import (
	"fmt"
	"movie_api/graph/model"
	"movie_api/models"
	"strings"
)

func intToPtr32(i int) *int32 {
	val := int32(i)
	return &val
}

func MapMovieToGraphQL(m models.Movies) *model.Movie {
	isComingSoon := false
	for _, s := range m.Broadcast {
		if strings.ToLower(s.Name) == "coming soon" {
			isComingSoon = true
			break
		}
	}

	var genres []*model.Genre
	for _, g := range m.Genres {
		genres = append(genres, &model.Genre{ID: fmt.Sprint(g.ID), Name: g.Genres})
	}

	var actors []*model.Actor
	for _, a := range m.Actors {
		actors = append(actors, &model.Actor{ID: fmt.Sprint(a.ID), Name: a.Name})
	}

	movie := &model.Movie{
		ID:       fmt.Sprint(m.ID),
		Title:    m.Name,
		Synopsis: &m.Description,
		Genres:   genres,
		Actors:   actors,
		Poster:   m.Poster,
		Views: 	  intToPtr32(m.Views),
	}

	if !isComingSoon {
	r := float64(m.Rating)
	d := int32(m.Duration)
	movie.Rating = &r
	movie.Duration = &d

	var broadcasts []*model.Broadcast
	for _, b := range m.Broadcast {
		broadcasts = append(broadcasts, &model.Broadcast{
			ID:   fmt.Sprint(b.ID),
			Name: b.Name,
			Link: b.Link,
		})
	}
	movie.Broadcasts = broadcasts
	}
	var statuses []*model.Status
	for _, s := range m.Statuses {
		statuses = append(statuses, &model.Status{
			ID:   fmt.Sprint(s.ID),
			Name: s.Name,
		})
	}
	movie.Statuses = statuses

	return movie
}