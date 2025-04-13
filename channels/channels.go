package main

import "fmt"

func main() {
	type result struct {
		res response
		err error
	}

	ch := make(chan result)

	for i := range 3 {
		go func() {
			res, err := process(request{id: i + 1, auth: fmt.Sprintf("psuedo_auth_%d", i)})
			ch <- result{res, err}
		}()
	}

	go func() {
		res, err := process(request{id: 4, auth: ""})
		ch <- result{res, err}
	}()

	for _ = range 4 {
		receiver := <-ch
		if receiver.err != nil {
			fmt.Printf("Error: request %d has error - %s\n", receiver.res.code, receiver.err)
		}
		fmt.Printf("Response: %s\n", &receiver.res)
	}
}

type ResponseError string

func (err ResponseError) Error() string{
	return fmt.Sprintf("HTTP error: %s", string(err))
}

const (
	UnauthErr = ResponseError("Unauthorized user request")
)

type request struct {
	auth string
	id int
}

type response struct {
	code int
	id int
}

func (r *response) String() string {
	return fmt.Sprintf("request %d has code %d", r.id, r.code)
}

func process(r request) (response, error) {
	if r.auth == "" {
		res := response{code: 401, id: r.id}
		return res, UnauthErr
	}
	return response{code: 200, id: r.id}, nil
}
