# go-pluralize

Pluralize and singularize any word

## Acknowledgements, meta information

> This is my (Stuart Fletcher/ratrocket on github/~md0 on sr.ht)
> personal fork of
> [gertd/go-pluralize](https://github.com/gertd/go-pluralize).
>
> If you for some reason have stumbled upon this repo, note that the
> authoritative version is on
> [sourcehut](https://git.sr.ht/~md0/pluralize).
> The [github repo](https://github.com/ratrocket/go-pluralize) is a
> mirror and may (will) lag behind.
>
> Most of what follows is from the original README, but I did remove
> some things and slightly "correct" others.  Stuart out!
>
> The go-pluralize module is the  Golang adaptation of the great work
> from [Blake Embrey](https://www.npmjs.com/~blakeembrey) and other
> contributors who created and maintain the NPM JavaScript
> [pluralize](https://www.npmjs.com/package/pluralize) package.
>
> The originating Javascript implementation can be found on
> https://github.com/blakeembrey/pluralize
>
> Without their great work this module would have taken a lot more
> effort, **thank you all**!

## Version mapping

The latest go-pluralize version is compatible with
[pluralize](https://www.npmjs.com/package/pluralize) version 8.0.0
commit
[#36f03cd](https://github.com/blakeembrey/pluralize/commit/36f03cd2d573fa6d23e12e1529fa4627e2af74b4)

| go-pluralize version  | NPM Pluralize Package version |
| ------------- | ------------- |
| 0.2.0 - Jan 25, 2022 [v0.2.0](https://github.com/gertd/go-pluralize/releases/tag/v0.2.0) | 8.0.0 - Oct 6, 2021 [#36f03cd](https://github.com/blakeembrey/pluralize/commit/36f03cd2d573fa6d23e12e1529fa4627e2af74b4)
| 0.1.7 - Jun 23, 2020 [v0.1.7](https://github.com/gertd/go-pluralize/releases/tag/v0.1.7) | 8.0.0 - Mar 14, 2020 [#e507706](https://github.com/blakeembrey/pluralize/commit/e507706be779612c06ebfd6043163e063e791d79)
| 0.1.2 - Apr 1, 2020 [v0.1.2](https://github.com/gertd/go-pluralize/releases/tag/v0.1.2) | 8.0.0 - Mar 14, 2020 [#e507706](https://github.com/blakeembrey/pluralize/commit/e507706be779612c06ebfd6043163e063e791d79)
| 0.1.1 - Sep 15, 2019 [v0.1.1](https://github.com/gertd/go-pluralize/releases/tag/v0.1.1) | 8.0.0 - Aug 27, 2019 [#abb3991](https://github.com/blakeembrey/pluralize/commit/abb399111aedd1d62dd418d7e0217d85f5bf22c9)
| 0.1.0 - Jun 12, 2019 [v0.1.0](https://github.com/gertd/go-pluralize/releases/tag/v0.1.0) | 8.0.0 - May 24, 2019 [#0265e4d](https://github.com/blakeembrey/pluralize/commit/0265e4d131ecad8e11c420fa4be98b75dc92c33d)

## Installation

To install the go module:

    go get md0.org/pluralize

## Usage

*## Code

    import "md0.org/pluralize"

    word := "Empire"

    pluralize := pluralize.NewClient()

    fmt.Printf("IsPlural(%s)   => %t\n", input, pluralize.IsPlural(word))
    fmt.Printf("IsSingular(%s) => %t\n", input, pluralize.IsSingular(word))
    fmt.Printf("Plural(%s)     => %s\n", input, pluralize.Plural(word))
    fmt.Printf("Singular(%s)   => %s\n", input, pluralize.Singular(word))

### Result

	IsPlural(Empire)   => false
	IsSingular(Empire) => true
	Plural(Empire)     => Empires
	Singular(Empire)   => Empire

## Pluralize Command Line

### Installation

	go install md0.org/pluralize/cmd/pluralize

### Usage

#### Help

	pluralize -help
    Usage of ./bin/pluralize:
      -cmd string
            command [All|IsPlural|IsSingular|Plural|Singular] (default "All")
      -version
            display version info
      -word string
            input value

#### Word with All Commands

    pluralize -word Empire

	IsPlural(Empire)   => false
	IsSingular(Empire) => true
	Plural(Empire)     => Empires
	Singular(Empire)   => Empire

#### Is Word Plural?

    pluralize -word Cactus -cmd IsPlural

	IsPlural(Cactus)   => false

#### Is Word Singular?

    pluralize -word Cacti -cmd IsSingular

    IsSingular(Cacti)  => false

#### Word Make Plural

    pluralize -word Cactus -cmd Plural

	Plural(Cactus)     => Cacti

#### Word Make Singular

    pluralize -word Cacti -cmd Singular

	Singular(Cacti)    => Cactus
