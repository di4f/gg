package gg

import (
	//"github.com/hajimehoshi/ebiten/v2"
)

func keyDiff(s1, s2 []Key) []Key {
    combinedSlice := append(s1, s2...)
    dm := make(map[Key]int)
    for _, v := range combinedSlice {
        if _, ok := dm[v]; ok {
            // remove element later as it exist in both slice.
            dm[v] += 1
            continue
        }
        // new entry, add in map!
        dm[v] = 1
    }
    var retSlice []Key
    for k, v := range dm {
        if v == 1 {
            retSlice = append(retSlice, k)
        }
    }
    return retSlice
}

type KeyDown struct {
	Key
}

type KeyUp struct {
	Key
}

type MouseButtonDown struct {
	MouseButton
	P Vector
}

type MouseButtonUp struct {
	MouseButton
	P Vector
}

type MouseMove struct {
	Delta Vector
}

type EventChan chan any

