package tst

type TestTimeSource struct {
	TimeToReturn int64
}

func (source *TestTimeSource) CurrentTimeMillis() int64 {
	return source.TimeToReturn
}