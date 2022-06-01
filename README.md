# Go Faker

Generates random data compatible with [Postman dynamic variables](https://learning.postman.com/docs/writing-scripts/script-references/variables-list).

[Ddosify open-source](https://github.com/ddosify/ddosify) version uses this faker library for dynamic variables.

This library uses [jaswdr/faker](https://github.com/jaswdr/faker). There are also some constants from [faker-js](https://github.com/faker-js/faker) to generate random values.

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

| Variable Name          | Description                        | Type              | Method                   | Examples                                                                          |
| ------                 | ------------------------------     | ----------------- | -------------            | ------------------------------------------------------------------------------    | 
| `_randomAlphaNumeric`  | Alphanumberic character            | string            | `RandomAlphanumeric()`   | `b`,`6`                                                                           |
| `_randomBoolean`       | Random boolean value               | bool              | `RandomBoolean()`        | `true`, `false`                                                                   |
| `_randomInt`           | Random integer btw. 0 and 1000     | int               | `RandomInt()`            | `352`, `518`                                                                      |
| `_randomColor`         | Random color                       | string            | `RandomSafeColorName()`  | `silver`, `yellow`                                                                |
| `_randomHexColor`      | Random HEX color                   | string            | `RandomSafeColorHex()`   | `#269423`, `#2859B0`                                                              |
| `_randomAbbreviation`  | Rancom abbreviation                | string            | `RandomAbbreviation()`   | `THX`, `XML`                                                                      |


### Phone, Address, and Location

| Variable Name           | Description                        | Type              | Method                         | Examples                                                                          |
| ------                  | ------------------------------     | ----------------- | -------------                  | ------------------------------------------------------------------------------    | 
| `_randomPhoneNumber`    | Random phone number                | string            | `RandomPhoneNumber()`          | `601-272-3813`,`208-658-9378`                                                     |
| `_randomPhoneNumberExt` | Random phone number with extension | string            | `RandomPhoneNumberExt()`       | `50-859-544-9083`,`86-685-370-3469`                                               |
| `_randomCity`           | Random city name                   | string            | `RandomAddressCity()`          | `Ethelfurt`,`Bergstromport`                                                       | 
| `_randomStreetName`     | Random street name                 | string            | `RandomAddresStreetName()`     | `Ethel Roads`,`Bergstrom Valley`                                                  |
| `_randomStreetAddress`  | Random street address              | string            | `RandomAddressStreetAddress()` | `88 Cassandra Trail`,`847147 Helmer Corners Suite 777`                            |
| `_randomCountry`        | Random country                     | string            | `RandomAddressCountry()`       | `Malta`,`Bermuda`                                                                 |
| `_randomCountryCode`    | Random 2 letter country code       | string            | `RandomCountryCode()`          | `UG`,`IE`                                                                         |
| `_randomLatitude`       | Random latitude                    | float64           | `RandomAddressLatitude()`      | `85.939027`,`89.190634`                                                           |
| `_randomLongitude`      | Random longitude                   | float64           | `RandomAddressLongitude()`     | `85.939027`,`89.190634`                                                           |


### Images

| Variable Name            | Description                        | Type              | Method                         | Examples                                                                          |
| ------                   | ------------------------------     | ----------------- | -------------                  | ------------------------------------------------------------------------------    | 
| `_randomAvatarImage`     | Random avatar image                | string            | `RandomAvatarImage()`          | `http://placeimg.com/640/480/people`                                              |
| `_randomImageUrl`        | Random image URL                   | string            | `RandomImageURL()`             | `http://placeimg.com/640/480`                                                     |
| `_randomAbstractImage`   | Random abstract image URL          | string            | `RandomAbstractImage()`        | `http://placeimg.com/640/480/abstract`                                            |
| `_randomAnimalsImage`    | Random animal image URL            | string            | `RandomAnimalsImage()`         | `http://placeimg.com/640/480/animals`                                             |
| `_randomBusinessImage`   | Random business image URL          | string            | `RandomBusinessImage()`        | `http://placeimg.com/640/480/business`                                            |
| `_randomCatsImage`       | Random cat image URL               | string            | `RandomCatsImage()`            | `http://placeimg.com/640/480/cats`                                                |
| `_randomCityImage`       | Random city image URL              | string            | `RandomCityImage()`            | `http://placeimg.com/640/480/city`                                                |
| `_randomFoodImage`       | Random food image URL              | string            | `RandomFoodImage()`            | `http://placeimg.com/640/480/food`                                                |
| `_randomNightlifeImage`  | Random night life image URL        | string            | `RandomNightlifeImage()`       | `http://placeimg.com/640/480/nightlife`                                           |
| `_randomFashionImage`    | Random fashion image URL           | string            | `RandomFashionImage()`         | `http://placeimg.com/640/480/fashion`                                             |
| `_randomPeopleImage`     | Random people image URL            | string            | `RandomPeopleImage()`          | `http://placeimg.com/640/480/people`                                              |
| `_randomNatureImage`     | Random nature image URL            | string            | `RandomNatureImage()`          | `http://placeimg.com/640/480/nature`                                              |
| `_randomSportsImage`     | Random sport image URL             | string            | `RandomSportsImage()`          | `http://placeimg.com/640/480/sports`                                              |
| `_randomTransportImage`  | Random transport image URL         | string            | `RandomTransportImage()`       | `http://placeimg.com/640/480/transport`                                           |
| `_randomImageDataUri`    | Random image data URI              | string            | `RandomDataImageUri()`         | `data:image/svg+xml;charset=UTF-8,%3Csvg%20xmlns%3D%22htt...`                     |
 

### Finance

| Variable Name            | Description                           | Type              | Method                          | Examples                                                               |
| ------                   | ------------------------------        | ----------------- | -------------                   | ---------------------------------------------------------------------  | 
| `_randomBankAccount`     | Random eight digit bank bumber        | string            | `RandomBankAccount()`           | `44014606`, `21604863`                                                 |
| `_randomBankAccountName` | Random bank account name              | string            | `RandomBankAccountName()`       | `Checking`, `Home Loan`                                                |
| `_randomCreditCardMask`  | Random credit card mask number        | string            | `RandomCreaditCardMask()`       | `6520`, `5339`                                                         |
| `_randomBankAccountBic`  | Random bank identifier code (swift)   | string            | `RandomBankAccountBic()`        | `BKTRTRIS`, `SOGEFRPPRIG`                                              |
| `_randomBankAccountIban` | Random IBAN                           | string            | `RandomBankAccountIban()`       | `HU63755471863115172345115723`, `EE268872311844634138`                 |
| `_randomTransactionType` | Random transaction type               | string            | `RandomTransactionType()`       | `invoice`, `payment`                                                   |
| `_randomCurrencyCode`    | Random three letter currency code     | string            | `RandomCurrencyCode()`          | `BIF`, `QAR`                                                           |
| `_randomCurrencyName`    | Random currency name                  | string            | `RandomCurrencyName()`          | `Uzbekistan Sum`, `Rupiah`                                             |
| `_randomCurrencySymbol`  | Random currency symbol                | string            | `RandomCurrencySymbol()`        | `J$`, `£`                                                              |
| `_randomBitcoin`         | Random bitcoin address                | string            | `RandomBitcoin()`               | `XBvHFJHdGdszbgqwDi6yZy7QXgeU`, `XBvHFJHdGdszbgqwDi6yZy7QXgeUSiH5kT`   |


### Business

| Variable Name            | Description                           | Type              | Method                          | Examples                                                               |
| ------                   | ------------------------------        | ----------------- | -------------                   | ---------------------------------------------------------------------  | 
| `_randomCompanyName`     | Random company name                   | string            | `RandomCompanyName()`           | `Daniel Bruen`, `Michael Cassin`                                       |
| `_randomCompanySuffix`   | Random company suffix                 | string            | `RandomCompanySuffix()`         | `und Söhne`, `Inc.`                                                    |
| `_randomBs`              | Random business speak words           | string            | `RandomBs()`                    | `synthesize efficient systems`, `matrix virtual functionalities`       |
| `_randomBsAdjective`     | Random business speak adjective       | string            | `RandomBsAdjective()`           | `interactive`, `plug-and-play`                                         |
| `_randomBsBuzz`          | Random business speak buzzword        | string            | `RandomBsBuzzWord()`            | `harness`, `morph`                                                     |
| `_randomBsNoun`          | Random business speak noun            | string            | `RandomBsNoun()`                | `paradigms`, `infrastructures`                                         |

### Catchphrases

| Variable Name                   | Description                      | Type              | Method                          | Examples                                                               |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------------------------------------  | 
| `_randomCatchPhrase`            | Random catchpharase              | string            | `RandomCatchPhrase()`           | `Monitored explicit throughput`, `Secured stable budgetary management` |
| `_randomCatchPhraseAdjective`   | Random catchpharase adjective    | string            | `RandomCatchPhraseAdjective()`  | `Implemented`, `Secured`                                               |
| `_randomCatchPhraseDescriptor`  | Random catchpharase descriptor   | string            | `RandomCatchPhraseDescriptor()` | `bandwidth-monitored`, `hybrid`                                        |
| `_randomCatchPhraseNoun`        | Random catchpharase noun         | string            | `RandomCatchPhraseNoun()`       | `array`, `installation`                                                |


### Databases

| Variable Name                   | Description                      | Type              | Method                          | Examples                                 |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------  | 
| `_randomDatabaseColumn`         | Random database column           | string            | `RandomDatabaseColumn()`        | `updatedAt`, `createdAt`                 |
| `_randomDatabaseType`           | Random database type             | string            | `RandomDatabaseType()`          | `serial`, `datetime`                     |
| `_randomDatabaseCollation`      | Random database collation        | string            | `RandomDatabaseCollation()`     | `cp1250_general_ci`, `ascii_bin`         |
| `_randomDatabaseEngine`         | Random database engine           | string            | `RandomDatabaseEngine()`        | `ARCHIVE`, `BLACKHOLE`                   |


### Dates

| Variable Name                   | Description                      | Type              | Method                          | Examples                                                           |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------                            | 
| `_randomDateFuture`             | Random future datetime           | string            | `RandomDateFuture()`            | `Wed Mar 18 09:19:24 UTC 2026`, `Fri Aug  1 09:58:13 UTC 2025`     |
| `_randomDatePast`               | Random past datetime             | string            | `RandomDatePast()`              | `Tue Dec 20 21:06:56 UTC 2016`, `Sun Dec  5 07:14:23 UTC 2021`     |
| `_randomDateRecent`             | Random recent datetime           | string            | `RandomDateRecent()`            | `Thu Apr 14 12:14:25 UTC 2022`, `Thu Apr 14 12:14:33 UTC 2022`     |
| `_randomWeekday`                | Random weekday                   | string            | `RandomWeekday()`               | `Wednesday`, `Thursday`                                            |
| `_randomMonth`                  | Random month                     | string            | `RandomMonth()`                 | `March`, `August`                                                  |

### Domains, Emails, and Usernames

| Variable Name                    | Description                      | Type              | Method                          | Examples                                                     |
| ------                           | ------------------------------   | ----------------- | -------------                   | ---------------------------------------                      | 
| `_randomDomainName`              | Random domain name               | string            | `RandomDomainName()`            | `claudiabeier.com`, `chasescott.ac.uk`                       |
| `_randomDomainSuffix`            | Random domain suffix             | string            | `RandomDomainSuffix()`          | `scot`, `info`                                               |
| `_randomDomainWord`              | Random domain word               | string            | `RandomDomainWord()`            | `jessecormier`, `victoriafisher`                             |
| `_randomEmail`                   | Random email                     | string            | `RandomEmail()`                 | `max.smitham@hotmail.com`, `ivy.wood@outlook.com`            |
| `_randomExampleEmail`            | Random email with example domain | string            | `RandomExampleEmail()`          | `isabella.bauch@example.net`, `tahlia.baumbach@example.org`  |
| `_randomUserName`                | Random username                  | string            | `RandomUsername()`              | `Hayden.Heaney`, `Lara.Durgan`                               |
| `_randomUrl`                     | Random URL                       | string            | `RandomUrl()`                   | `https://sophiemills.scot`, `https://harrisongoldner.ltd.uk` |


### Files and Directories


| Variable Name                   | Description                      | Type              | Method                          | Examples                                                              |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------                               | 
| `_randomFileName`               | Random file name                 | string            | `RandomFileName()`              | `matthew_moore.gdoc`, `angus_daniel.mov`                              |
| `_randomFileType`               | Random file type                 | string            | `RandomFileType()`              | `video`, `application`                                                |
| `_randomFileExt`                | Random file extension            | string            | `RandomFileExtension()`         | `xml`, `gif`                                                          |
| `_randomCommonFileName`         | Random common file name          | string            | `RandomCommonFileName()`        | `marcuswaters.csv`, `elizaowen.html`                                  |
| `_randomCommonFileType`         | Random common file tyoe          | string            | `RandomCommonFileType()`        | `application`, `image`                                                |
| `_randomCommonFileExt`          | Random common file extension     | string            | `RandomCommonFileExtension()`   | `m3a`, `gif`                                                          |
| `_randomFilePath`               | Random file path                 | string            | `RandomFilePath()`              | `/usr/X11R6/hugoschmidt.mp4`, `/usr/libexec/mitchelloreilly.mp4`      |
| `_randomDirectoryPath`          | Random directory path            | string            | `RandomDirectoryPath()`         | `/proc`, `/usr`                                                       |
| `_randomMimeType`               | Random mime type                 | string            | `RandomMimeTypes()`             | `image/gif`, `image/png`                                              |


### Stores

| Variable Name                   | Description                      | Type              | Method                          | Examples                                              |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------               | 
| `_randomPrice`                  | Random price (0.00-1000.00)      | string            | `RandomPrice()`                 | `445.60`, `802.67`                                    |
| `_randomProduct`                | Random product                   | string            | `RandomProduct()`               | `Hat`, `Shoes`                                        |
| `_randomProductAdjective`       | Random product adjective         | string            | `RandomProductAdjective()`      | `Fantastic`, `Tasty`                                  |
| `_randomProductMaterial`        | Random product material          | string            | `RandomProductMaterial()`       | `Cotton`, `Steel`                                     |
| `_randomProductName`            | Random product name              | string            | `RandomProductName()`           | `Gorgeous Frozen Chicken`, `Recycled Bronze Sausages` |
| `_randomDepartment`             | Random stores department         | string            | `RandomDepartment()`            | `Tools`, `Health`                                     |


### Grammar

| Variable Name                    | Description                      | Type              | Method                          | Examples                                            |
| ------                           | ------------------------------   | ----------------- | -------------                   | ---------------------------------------             | 
| `_randomNoun`                    | Random noun                      | string            | `RandomNoun()`                  | `interface`, `port`                                 |
| `_randomVerb`                    | Random verb                      | string            | `RandomVerb()`                  | `calculate`, `parse`                                |
| `_randomIngverb`                 | Random ing verb                  | string            | `RandomIngVerb()`               | `indexing`, `hacking`                               |
| `_randomAdjective`               | Random adjective                 | string            | `RandomAdjective()`             | `multi-byte`, `neural`                              |
| `_randomWord`                    | Random word                      | string            | `RandomWord()`                  | `program`, `overriding`                             |
| `_randomWords`                   | Random words                     | string            | `RandomWords()`                 | `virtual navigate 1080p`, `reboot online alarm`     |
| `_randomPhrase`                  | Random phrase                    | string            | `RandomLoremSentences()`        | `nobis iste omnis iusto aut. aliquam illo maxime aperiam nobis unde labore quos qui laboriosam voluptas sint perferendis quis iusto rerum voluptas vero tempore iure. veniam sit sit velit nihil sint in. odit ipsam incidunt earum voluptates mollitia repellat et unde et tempora sit. aliquam ullam ipsum odit ut labore omnis voluptate minima. `<br><br> `modi perferendis consectetur perspiciatis vero natus laborum quidem in dolorem. quia architecto quis error ducimus nihil hic voluptas. minus accusamus velit quis praesentium molestias odit sequi laboriosam magnam sapiente commodi. quos corrupti eveniet id suscipit unde dolorem adipisci accusantium qui ipsam molestiae voluptatem qui dolores omnis ducimus aut. `                 |

### Lorem Ipsum

| Variable Name                   | Description                      | Type              | Method                          | Examples                                 |
| ------                          | ------------------------------   | ----------------- | -------------                   | ---------------------------------------  | 
| `randomLoremWord`               | Random lorem ipsum word          | string            | `RandomLoremWord()`             | `iusto`, `quaerat`                 |
| `randomLoremWords`              | Random lorem ipsum words         | string            | `RandomLoremWords()`            | `officiis iusto`, `deleniti molestiae veritatis non `                 |
| `randomLoremSentence`           | Random lorem ipsum sentence      | string            | `RandomLoremSentence()`         | `delectus iure qui qui veritatis sit accusantium eaque dicta tempore saepe qui corporis rerum nulla fugiat iure soluta esse.`<br><br> `aspernatur saepe harum ex minus molestias fugit minus amet et et fugiat impedit sit eum quasi.`                 |
| `randomLoremSentences`          | Random lorem ipsum sentences     | string            | `RandomLoremSentences()`        | `ab perferendis cupiditate atque saepe sed nulla ut deserunt consequuntur architecto maiores qui quis a voluptatem possimus. sit quia vel perferendis impedit et similique dolorem omnis qui ut. minima quasi totam velit dolorem excepturi assumenda exercitationem ut aut. nulla fugit illo odio sequi et omnis quo officiis est aliquam enim quia harum. debitis corrupti iste enim voluptatem tenetur laborum id ut nobis earum. `<br><br> `reprehenderit reprehenderit sed vitae ut ipsam totam numquam vitae laboriosam ut ut voluptas aut. ut adipisci optio quia quam velit quisquam eum in est iure. natus voluptatem eius est beatae voluptas nihil aperiam nesciunt sit et aspernatur. `                 |
| `randomLoremParagraph`          | Random lorem ipsum paragraph     | string            | `RandomLoremParagraph()`        | `provident asperiores dolorum beatae totam ut culpa rem neque quae rerum voluptas. pariatur aut cum dolore laudantium cum et error placeat rerum eligendi. aspernatur voluptas inventore porro necessitatibus sit magnam fuga et adipisci. occaecati voluptates voluptatem non ut atque placeat beatae veniam aut perspiciatis ut numquam sapiente minima doloribus. hic impedit atque et tempora doloribus sit sint ipsam. ab rem officiis ut cum ipsum quia minus officia odit aut porro doloremque quibusdam eligendi ab atque cupiditate. vitae in nostrum nobis cum labore qui. corporis fugiat optio eum laborum quos maxime commodi reiciendis et in. quia magni ut laboriosam mollitia ullam voluptatem explicabo corporis veniam. quia possimus vitae rerum labore qui omnis. vel eum quod impedit fugiat mollitia dolorem deserunt quis. id quod alias repellat provident molestiae voluptate tempore iste minus quasi accusamus esse consequatur et voluptas debitis sunt deleniti. odio sit voluptas voluptatem excepturi. rerum sit velit molestiae impedit quia illo et quod. alias totam neque velit doloremque corporis. vel omnis dolores sit nam aut harum aut eaque dolores consectetur quaerat dolor facere deserunt. `<br><br> `porro eum minus quis unde ut temporibus nulla non asperiores fuga tenetur. molestias possimus totam in sapiente qui. qui architecto quis quas quis saepe quia illo est repudiandae deserunt neque est sit dolorum tempora est placeat ad quia non. quam deserunt quia iusto est harum provident. vero illo omnis rem sed consequatur qui ex cumque itaque sed non aut quia corrupti quia ratione qui modi et. `                 |
| `randomLoremParagraphs`         | Random lorem ipsum paragraphs    | string            | `RandomLoremParagraphs()`       | `eum reiciendis modi ab minus vero exercitationem nisi voluptatibus nihil in suscipit ad soluta iure ut sunt. et aut est laborum sed et minima iste odit voluptatem consequatur commodi doloremque. possimus possimus voluptate adipisci sequi eum commodi dolorem sunt dolore totam. adipisci ducimus eos velit sed provident et aut modi architecto velit quia provident saepe enim sunt voluptas et. reiciendis velit magni voluptas ut delectus voluptatem qui dolorem ipsum. sunt temporibus sit eos voluptates. molestias distinctio autem nihil aperiam dolorem illo qui. architecto ratione et facilis accusamus incidunt voluptates officia deleniti ut praesentium saepe tempore illum sunt ut consectetur. natus ut et voluptatem quidem labore et eaque cumque voluptate non et non. eum natus vel minima et non et architecto sunt quis deserunt molestias voluptatem et nemo repellendus est in. architecto aliquid mollitia asperiores. facilis voluptatem distinctio voluptas et aliquam dolorum et qui ipsum alias aut et eum aut. `<br><br> `qui voluptatibus modi itaque nam a explicabo eos nihil excepturi porro omnis assumenda. nesciunt consequuntur labore et a dolorem adipisci. quos eligendi atque expedita assumenda qui. accusamus exercitationem omnis in doloribus sint ab a corporis tempore. explicabo placeat non repudiandae assumenda ut nihil. quo dolor qui illum libero velit iure et et aliquid rerum et praesentium. ipsum est iste occaecati reiciendis nesciunt voluptates id numquam ad non totam at est quo illo officiis. autem quia et quo aliquam qui qui blanditiis consequuntur dolor quo consequatur corporis dicta cumque autem architecto. est harum ipsum et sequi iure adipisci optio vero voluptatem et molestiae dolorem ea facilis sequi sit. `                 |
| `randomLoremText`               | Random lorem ipsum text          | string            | `RandomLoremParagraphs()`       | `quia officiis et delectus ab maxime mollitia et enim laborum. aliquid ullam blanditiis voluptatem ipsam labore. perferendis aut necessitatibus eos aut numquam et dolorem nemo perspiciatis aut cupiditate reprehenderit aut natus. illum aspernatur explicabo veniam dignissimos fugiat necessitatibus maiores nostrum numquam repellendus laboriosam repudiandae et aspernatur similique. sequi doloribus consequatur possimus optio aliquam et quis quia eum tenetur reiciendis sed dicta voluptatem accusantium ex repudiandae eaque. cum et laudantium enim facere repellat voluptatum ullam aspernatur blanditiis aut incidunt dignissimos ipsa dolorem. reprehenderit voluptatem fugiat laboriosam ipsum sed voluptatibus qui inventore consequatur sunt quas. et eos asperiores sapiente tempore aut ea cum consequuntur. delectus et qui est et cumque enim dolorum et consequuntur dolorem nemo. `<br><br> `neque facilis enim aut et perspiciatis distinctio et delectus. hic voluptatem maxime libero incidunt deleniti corrupti error ducimus ex. omnis qui quod necessitatibus architecto voluptas molestias consequatur ipsum et vero id voluptatum. ullam dolore nemo inventore dolore voluptatem sed qui consectetur accusantium quo ab deleniti nemo et nihil atque corrupti. adipisci ullam modi facere consequatur sunt vel. maxime sapiente unde amet ipsa possimus ut voluptates nesciunt molestias. voluptatem est iure facilis animi quasi ex eos blanditiis consequatur. quis nemo sunt aut maxime illo nihil voluptatibus aliquid. dicta sint quia necessitatibus in eius adipisci dolores et. `                 |
| `randomLoremSlug`               | Random lorem ipsum slug          | string            | `RandomLoremSlug()`             | `cupiditate-consectetur-et`, `adipisci-suscipit-qui`                 |
| `randomLoremLines`              | Random lorem ipsum lines         | string            | `RandomLoremLines()`            | `esse soluta voluptatem atque aut ut exercitationem optio est expedita animi qui in est reiciendis eum qui quis dolore.\n neque perspiciatis nihil ut iure qui libero voluptas nobis corrupti.\n beatae dolores sit rerum unde nesciunt rerum nesciunt adipisci sunt voluptatem recusandae est eum.\n quia velit qui necessitatibus perspiciatis neque rerum facilis architecto odio laudantium facere voluptas tenetur.`<br><br> `voluptatem in ut ut facilis et fugiat sunt suscipit facere perspiciatis est at suscipit repudiandae tempore.\n natus nostrum aut odio necessitatibus qui sunt et assumenda eos qui nihil.\n aut mollitia rerum quos non rerum inventore velit tempora libero aut nobis dolores eligendi quod aut qui amet rerum.\n esse architecto veritatis non atque dolorum sint maxime veniam a velit vel aut qui in totam saepe enim ad nemo omnis.`                  |


## Communication

You can join our [Discord Server](https://discord.gg/9KdnrSUZQg) for issues, feature requests, feedbacks or anything else. 
