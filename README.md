# Odie

Odie is a CLI wrapper of the [Oxford Dictionaries API](https://developer.oxforddictionaries.com/).
This is currently a toy project and is not intended to be fully compatible with the Oxford Dictionaries API.

> [Oxford Dictionaries Premium](https://premium.oxforddictionaries.com/) provides a greater experience in the browser with a lower cost.
> I would definitely recommend it as the alternative to the discontinued lexico.com.

## Usage

### Prerequisites

Create an account at [Oxford Dictionaries API](https://developer.oxforddictionaries.com/) to get a pair of Application ID and Application Key.

### Installation

Install the binary

```bash
go install github.com/central182/odie@latest
```

And create a config file at `$HOME/.odie.yaml`, containing the following content:

```yaml
appId: your_application_id
appKey: your_application_key
```

(The configuration step is optional as you can always provide `--appId` and `--appKey` flags directly when running the program.)

### Commands

- `odie WORD` looks a word up and exits.
- `odie repl` starts an interactive session where you can keep looking words up until you hit `Ctrl-D`.
- In case you wanted to look the word "repl" up, use `odie -- repl`. (No, it doesn't exist in the dictionary.)

### Screenshots

#### One-off mode

<img width="75%" src="https://user-images.githubusercontent.com/62089140/192131982-456706d7-382a-452c-8020-3b1645ab54a5.png">

#### Interactive mode

<img width="75%" src="https://user-images.githubusercontent.com/62089140/192131981-6ea6c196-4d88-41a9-b510-2db87ca91e8d.png">

## Design

### Architecture

The architecture of Odie is basically a reiteration of [Ports and Adapters](https://alistair.cockburn.us/hexagonal-architecture/) (or whatever fancy name the idea has got).
Dependency points inwards, from concreteness to abstractness, towards the `domain/dictionary` package at the centre.

Such an architecture emphasises that *Odie is a dictionary app whose data, among other possible options, comes from the Oxford Dictionaries API*,
and that *Odie is not a consumer app at the mercy of the Oxford Dictionaries API*.
(In an architectural sense, of course.)

<img width="75%" src="https://user-images.githubusercontent.com/62089140/192135438-2726debc-257e-431b-af94-325353733686.png">

### Fully-encapsulated-ness

Encapsulation of exported structs can't be enforced strictly in Go. Take this as a counterexample.
```golang
// OddInteger is a wrapper of an odd integer.
type OddInteger struct {
	value int
}

func NewOddInteger(value int) OddInteger {
	if value%2 == 0 {
		panic("can't wrap an even integer")
	}
	return OddInteger{value: value}
}

func (o OddInteger) Value() int {
	return o.value
}
```

The mere existence of the zero value breaks `OddInteger`'s invariant. `OddInteger{}.Value()` returns `0`, in contrary to what `OddInteger` purports to be, a wrapper of an odd integer.

In order to fix this flaw, every method of `OddInteger` needs to check if it's being called on a zero value and either return an error or panic. That's quite a bit boilerplate and such expected behaviour warrants unit-testing.

Now compare:
```golang
// OddInteger is a wrapper of an odd integer.
type OddInteger interface {
	Value() int
}

func NewOddInteger(value int) OddInteger {
	if value%2 == 0 {
		panic("can't wrap an even integer")
	}
	return oddInteger{value: value}
}

type oddInteger struct {
	value int
}

func (o oddInteger) Value() int {
	return o.value
}
```

This time, as far as the current package is concerned, `NewOddInteger` is the only way to instantiate an `OddInteger`, and it is guaranteed that any `OddInteger` is truly a wrapper of an odd integer.

> Cheating by providing an alternative implementation of `OddInteger` somewhere else is nevertheless possible. However, that's an unexpected behaviour from the current package's perspective and can not be unit-tested.

In Odie, every type that has some invariant is realised as an exported interface implemented by a non-exported struct.

### Structured errors

A return value of `error` type doesn't tell much by the signature. Given the following example,

```golang
type Fallible interface {
	Fail() error
}

var (
	ErrOne = errors.New("")
)
```
It's not clear how `ErrOne` is related to `Fallible`. Maybe some implementation returns `ErrOne` in the `Fail` method, but we don't know for sure.

Compare:
```golang
type Fallible interface {
	Fail() FailError
}

type FailError interface {
	error
	FailedForReasonA() bool
	FailedForReasonB() bool
}
```
Now we can conclude that `Fail` may fail for either reason A or reason B.
Any client of `Fallible` can act accordingly by checking the return values of `FailedForReasonA` and `FailedForReasonB`.

In Odie, errors that take place inside *domain* have their own structured types.

### Testing the testable (and leaving out the untestable)

Given the proper direction of dependency, everything **inside** and **in the vicinity of** the *domain* zone is easily testable, and has been tested accordingly.
Those that interact with the outer world are left untested, including:
- [internal/adapter/outbound/common/odapi/resty/client.go](internal/adapter/outbound/common/odapi/resty/client.go)
  - depends on the Oxford Dictionaries API being up and running
- [internal/program](internal/program)
  - depends on the OS doing its job
