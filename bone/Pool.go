package bone

import "sync"

var boneObjectTrack struct {
	sync.Mutex
	m map[uintptr]interface{}
	// c int
}

func boneObjectAdd(c uintptr, v interface{}) {
	boneObjectTrack.Lock()
	defer boneObjectTrack.Unlock()
	if boneObjectTrack.m == nil {
		boneObjectTrack.m = make(map[uintptr]interface{})
	}
	if boneObjectTrack.m[c] != nil {
		panic("Bone Address Corruption")
	}
	boneObjectTrack.m[c] = v
}

func boneObjectLookup(c uintptr) interface{} {
	boneObjectTrack.Lock()
	defer boneObjectTrack.Unlock()
	ret := boneObjectTrack.m[c]
	if ret == nil {
		panic("Bone Object Not Found")
	}
	return ret
}

func boneObjectDelete(c uintptr) {
	boneObjectTrack.Lock()
	defer boneObjectTrack.Unlock()
	if boneObjectTrack.m[c] == nil {
		panic("Bone Object Not Found")
	}
	delete(boneObjectTrack.m, c)
}
