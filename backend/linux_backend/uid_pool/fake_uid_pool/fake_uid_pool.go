package fake_uid_pool

type FakeUIDPool struct {
	nextUID uint32

	AcquireError error

	Acquired []uint32
	Released []uint32
}

func New(start uint32) *FakeUIDPool {
	return &FakeUIDPool{
		nextUID: start,
	}
}

func (p *FakeUIDPool) Acquire() (uint32, error) {
	if p.AcquireError != nil {
		return 0, p.AcquireError
	}

	uid := p.nextUID
	p.nextUID++

	return uid, nil
}

func (p *FakeUIDPool) Release(uid uint32) {
	p.Released = append(p.Released, uid)
}
