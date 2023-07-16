package main

type Movie struct{
	ID string`json:"id"`
	Name string`json:"name"`
	ReleaseYear int64`json:"releaseYear"`
	Director *Director`json:"director"`
}


