package sample

type sample interface {
	sample()
}

type sampleImpl struct {

}

func (s sampleImpl) sample() {
	panic("implement me")
}
