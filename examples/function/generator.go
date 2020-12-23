package function

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

//명명된 자료형 (Named Type)
//자료형에 새로 이름을 붙일 수 있음
type VertexID int
type EdgeID int

func NewVertexIDGenerator() func() VertexID {
	var next int
	return func() VertexID {
		next++
		return VertexID(next) //int -> VertexID 타입으로 변환함
	}
}
