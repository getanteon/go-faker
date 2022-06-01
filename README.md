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

## Supported Variables

### Common

| Variable Name         | Description                        | Type              | Method                   | Examples                                                                           |
| ------                | ------------------------------     | ----------------- | -------------            | ------------------------------------------------------------------------------     | 
| `_guid`               | UUID                               | uuid.UUID         | `RandomGuid()`           | `5066a748-9a72-404d-94f7-512c0779ff8e`, `adda49b1-7148-4a89-92de-6c4b756c1226`     |
| `_timestamp`          | Current timestamp in seconds       | int64             | `RandomTimestamp()`      | `1654037772`, `1654037847`                                                         |
| `_isoTimestamp`       | Current ISO timestamp in seconds   | string            | `RandomISOTimestamp()`   | `2022-05-31T22:58:40.653Z`, `2022-05-31T22:59:06.013Z`                             |
| `_randomUUID`         | UUID                               | uuid.UUID         | `RandomUUID()`           | `35334fa3-fa83-4589-97e7-7419c9a2173e`, `6661c4cb-ec3a-464c-9ecc-d9fd051e8def`     |


### Text, Numbers and Colors

| Variable Name         | Description                        | Type              | Method                   | Examples                                                                           |
| ------                | ------------------------------     | ----------------- | -------------            | ------------------------------------------------------------------------------     | 
| `randomAlphaNumeric`  | Alphanumberic character            | string            | `RandomAlphanumeric()`   | `b`,`6`                                                                            |
| `randomBoolean`       | Random boolean value               | bool              | `RandomBoolean()`        | `true`, `false`                                                                    |
| `randomInt`           | Random integer btw. 0 and 1000     | int               | `RandomInt()`            | `352`, `518`                                                                       |
| `randomColor`         | Random color                       | string            | `RandomSafeColorName()`  | `silver`, `yellow`                                                                 |
| `randomHexColor`      | Random HEX color                   | string            | `RandomSafeColorHex()`   | `#269423`, `#2859B0`                                                               |
| `randomAbbreviation`  | Rancom abbreviation                | string            | `RandomAbbreviation()`   | `THX`, `XML`                                                                       |
