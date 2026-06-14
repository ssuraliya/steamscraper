package tags

type TagCollector interface {
	All() ([]string, error)
}
