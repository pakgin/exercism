package erratum

import "errors"

func Use(opener ResourceOpener, input string) error {
	resource, err := open(opener)
	if err != nil {
		return err
	}
	defer resource.Close()

	return frob(resource, input)
}

func open(opener ResourceOpener) (Resource, error) {
	resource, err := opener()
	if err != nil {
		if errors.As(err, &TransientError{}) {
			return open(opener)
		}
		return nil, err
	}
	return resource, nil
}

func frob(resource Resource, input string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case FrobError:
				resource.Defrob(t.defrobTag)
				err = t.inner
			default:
				err = r.(error)
			}
		}
	}()

	resource.Frob(input)
	return
}
