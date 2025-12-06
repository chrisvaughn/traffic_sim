package main

type Road struct {
	ID        string     `json:"id"`
	Waypoints []Position `json:"waypoints"`
}

func NewStraightRoad(id string, start, end Position) Road {
	return Road{
		ID:        id,
		Waypoints: []Position{start, end},
	}
}
