package handler

type writerMock struct{}

func (writerMock) Write(_ []byte) (n int, err error) {
	return 0, nil
}
