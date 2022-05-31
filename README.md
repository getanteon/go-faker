# Go Faker

Generates random data compatible with [Postman dynamic variables](https://learning.postman.com/docs/writing-scripts/script-references/variables-list).

[Ddosify open-source](https://github.com/ddosify/ddosify) version uses this faker library for dynamic variables.

This library extends [jaswdr/faker](https://github.com/jaswdr/faker). There are also some constants from [faker-js](https://github.com/faker-js/faker) to produce random values.

## Usage 

```golang
package main

import (
	"fmt"

	"github.com/ddosify/go-faker/faker"
)

func main() {
	faker := faker.NewFaker()
	fmt.Println(faker.RandomBankAccountIban()) // DE15534524466712768735
	fmt.Println(faker.RandomMimeTypes())       // audio/mpeg
	fmt.Println(faker.RandomAdjective())       // open-source
	fmt.Println(faker.RandomAbstractImage())   // http://placeimg.com/640/480/abstract
	fmt.Println(faker.RandomCountryCode())     // GF
	fmt.Println(faker.RandomUsername())        // Charlie.Hansen
	fmt.Println(faker.RandomSemver())          // 3.5.3
}

```
