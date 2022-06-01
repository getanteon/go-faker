<h1 align="center">
    Fake Data Generator for Golang
</h1>

<p align="center">
    <a href="https://github.com/ddosify/go-faker/releases" target="_blank"><img src="https://img.shields.io/github/v/release/ddosify/go-faker?style=for-the-badge&logo=github&color=orange" alt="go-faker latest version" /></a>&nbsp;
	<a href="https://pkg.go.dev/github.com/ddosify/go-faker" target="_blank"><img src="https://img.shields.io/github/go-mod/go-version/ddosify/go-faker?style=for-the-badge&logo=go" alt="golang version" /></a>&nbsp;
    <a href="https://github.com/ddosify/go-faker/blob/master/LICENSE" target="_blank"><img src="https://img.shields.io/github/license/ddosify/go-faker?style=for-the-badge&logo=none" alt="go-faker license" /></a>
    <a href="https://discord.gg/9KdnrSUZQg" target="_blank"><img src="https://img.shields.io/discord/898523141788287017?style=for-the-badge&logo=discord&label=DISCORD" alt="ddosify discord server" /></a>
</p>


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

## Supported Methods

### Common

| Method                    | Description                        | Type               | Examples                                                                           |
| -------------             | ------------------------------     | -----------------  | ------------------------------------------------------------------------------     | 
| `RandomGuid()`            | UUID                               | uuid.UUID          | `5066a748-9a72-404d-94f7-512c0779ff8e`, `adda49b1-7148-4a89-92de-6c4b756c1226`     |
| `RandomTimestamp()`       | Current timestamp in seconds       | int64              | `1654037772`, `1654037847`                                                         |
| `RandomISOTimestamp()`    | Current ISO timestamp in seconds   | string             | `2022-05-31T22:58:40.653Z`, `2022-05-31T22:59:06.013Z`                             |
| `RandomUUID()`            | UUID                               | uuid.UUID          | `35334fa3-fa83-4589-97e7-7419c9a2173e`, `6661c4cb-ec3a-464c-9ecc-d9fd051e8def`     |


### Text, Numbers and Colors

| Method                  | Description                        | Type                | Examples                                                                          |
| -------------           | ------------------------------     | -----------------   | ------------------------------------------------------------------------------    | 
| `RandomAlphanumeric()`  | Alphanumberic character            | string              | `b`,`6`                                                                           |
| `RandomBoolean()`       | Random boolean value               | bool                | `true`, `false`                                                                   |
| `RandomInt()`           | Random integer btw. 0 and 1000     | int                 | `352`, `518`                                                                      |
| `RandomSafeColorName()` | Random color                       | string              | `silver`, `yellow`                                                                |
| `RandomSafeColorHex()`  | Random HEX color                   | string              | `#269423`, `#2859B0`                                                              |
| `RandomAbbreviation()`  | Rancom abbreviation                | string              | `THX`, `XML`                                                                      |


### Phone, Address, and Location

| Method                         | Description                        | Type              | Examples                                                                          |
| -------------                  | ------------------------------     | ----------------- | ------------------------------------------------------------------------------    | 
| `RandomPhoneNumber()`          | Random phone number                | string            | `601-272-3813`,`208-658-9378`                                                     |
| `RandomPhoneNumberExt()`       | Random phone number with extension | string            | `50-859-544-9083`,`86-685-370-3469`                                               |
| `RandomAddressCity()`          | Random city name                   | string            | `Ethelfurt`,`Bergstromport`                                                       | 
| `RandomAddresStreetName()`     | Random street name                 | string            | `Ethel Roads`,`Bergstrom Valley`                                                  |
| `RandomAddressStreetAddress()` | Random street address              | string            | `88 Cassandra Trail`,`847147 Helmer Corners Suite 777`                            |
| `RandomAddressCountry()`       | Random country                     | string            | `Malta`,`Bermuda`                                                                 |
| `RandomCountryCode()`          | Random 2 letter country code       | string            | `UG`,`IE`                                                                         |
| `RandomAddressLatitude()`      | Random latitude                    | float64           | `85.939027`,`89.190634`                                                           |
| `RandomAddressLongitude()`     | Random longitude                   | float64           | `85.939027`,`89.190634`                                                           |


### Images

| Method                     | Description                        | Type                   | Examples                                                                          |
| -------------              | ------------------------------     | -----------------      | ------------------------------------------------------------------------------    | 
| `RandomAvatarImage()`      | Random avatar image                | string                 | `http://placeimg.com/640/480/people`                                              |
| `RandomImageURL()`         | Random image URL                   | string                 | `http://placeimg.com/640/480`                                                     |
| `RandomAbstractImage()`    | Random abstract image URL          | string                 | `http://placeimg.com/640/480/abstract`                                            |
| `RandomAnimalsImage()`     | Random animal image URL            | string                 | `http://placeimg.com/640/480/animals`                                             |
| `RandomBusinessImage()`    | Random business image URL          | string                 | `http://placeimg.com/640/480/business`                                            |
| `RandomCatsImage()`        | Random cat image URL               | string                 | `http://placeimg.com/640/480/cats`                                                |
| `RandomCityImage()`        | Random city image URL              | string                 | `http://placeimg.com/640/480/city`                                                |
| `RandomFoodImage()`        | Random food image URL              | string                 | `http://placeimg.com/640/480/food`                                                |
| `RandomNightlifeImage()`   | Random night life image URL        | string                 | `http://placeimg.com/640/480/nightlife`                                           |
| `RandomFashionImage()`     | Random fashion image URL           | string                 | `http://placeimg.com/640/480/fashion`                                             |
| `RandomPeopleImage()`      | Random people image URL            | string                 | `http://placeimg.com/640/480/people`                                              |
| `RandomNatureImage()`      | Random nature image URL            | string                 | `http://placeimg.com/640/480/nature`                                              |
| `RandomSportsImage()`      | Random sport image URL             | string                 | `http://placeimg.com/640/480/sports`                                              |
| `RandomTransportImage()`   | Random transport image URL         | string                 | `http://placeimg.com/640/480/transport`                                           |
| `RandomDataImageUri()`     | Random image data URI              | string                 | `data:image/svg+xml;charset=UTF-8,%3Csvg%20xmlns%3D%22htt...`                     |
 

### Finance

| Method                       | Description                           | Type                  | Examples                                                               |
| -------------                | ------------------------------        | -----------------     | ---------------------------------------------------------------------  | 
| `RandomBankAccount()`        | Random eight digit bank bumber        | string                | `44014606`, `21604863`                                                 |
| `RandomBankAccountName()`    | Random bank account name              | string                | `Checking`, `Home Loan`                                                |
| `RandomCreaditCardMask()`    | Random credit card mask number        | string                | `6520`, `5339`                                                         |
| `RandomBankAccountBic()`     | Random bank identifier code (swift)   | string                | `BKTRTRIS`, `SOGEFRPPRIG`                                              |
| `RandomBankAccountIban()`    | Random IBAN                           | string                | `HU63755471863115172345115723`, `EE268872311844634138`                 |
| `RandomTransactionType()`    | Random transaction type               | string                | `invoice`, `payment`                                                   |
| `RandomCurrencyCode()`       | Random three letter currency code     | string                | `BIF`, `QAR`                                                           |
| `RandomCurrencyName()`       | Random currency name                  | string                | `Uzbekistan Sum`, `Rupiah`                                             |
| `RandomCurrencySymbol()`     | Random currency symbol                | string                | `J$`, `£`                                                              |
| `RandomBitcoin()`            | Random bitcoin address                | string                | `XBvHFJHdGdszbgqwDi6yZy7QXgeU`, `XBvHFJHdGdszbgqwDi6yZy7QXgeUSiH5kT`   |


### Business

| Method                      | Description                           | Type                    | Examples                                                               |
| -------------               | ------------------------------        | -----------------       | ---------------------------------------------------------------------  | 
| `RandomCompanyName()`       | Random company name                   | string                  | `Daniel Bruen`, `Michael Cassin`                                       |
| `RandomCompanySuffix()`     | Random company suffix                 | string                  | `und Söhne`, `Inc.`                                                    |
| `RandomBs()`                | Random business speak words           | string                  | `synthesize efficient systems`, `matrix virtual functionalities`       |
| `RandomBsAdjective()`       | Random business speak adjective       | string                  | `interactive`, `plug-and-play`                                         |
| `RandomBsBuzzWord()`        | Random business speak buzzword        | string                  | `harness`, `morph`                                                     |
| `RandomBsNoun()`            | Random business speak noun            | string                  | `paradigms`, `infrastructures`                                         |

### Catchphrases

| Method                           | Description                      | Type               | Examples                                                               |
| -------------                    | ------------------------------   | -----------------  | ---------------------------------------------------------------------  | 
| `RandomCatchPhrase()`            | Random catchpharase              | string             | `Monitored explicit throughput`, `Secured stable budgetary management` |
| `RandomCatchPhraseAdjective()`   | Random catchpharase adjective    | string             | `Implemented`, `Secured`                                               |
| `RandomCatchPhraseDescriptor()`  | Random catchpharase descriptor   | string             | `bandwidth-monitored`, `hybrid`                                        |
| `RandomCatchPhraseNoun()`        | Random catchpharase noun         | string             | `array`, `installation`                                                |


### Databases

| Method                           | Description                      | Type                  | Examples                                 |
| -------------                    | ------------------------------   | -----------------     | ---------------------------------------  | 
| `RandomDatabaseColumn()`         | Random database column           | string                | `updatedAt`, `createdAt`                 |
| `RandomDatabaseType()`           | Random database type             | string                | `serial`, `datetime`                     |
| `RandomDatabaseCollation()`      | Random database collation        | string                | `cp1250_general_ci`, `ascii_bin`         |
| `RandomDatabaseEngine()`         | Random database engine           | string                | `ARCHIVE`, `BLACKHOLE`                   |


### Dates

| Method                            | Description                      | Type                     | Examples                                                           |
| -------------                     | ------------------------------   | -----------------        | ---------------------------------------                            | 
| `RandomDateFuture()`              | Random future datetime           | string                   | `Wed Mar 18 09:19:24 UTC 2026`, `Fri Aug  1 09:58:13 UTC 2025`     |
| `RandomDatePast()`                | Random past datetime             | string                   | `Tue Dec 20 21:06:56 UTC 2016`, `Sun Dec  5 07:14:23 UTC 2021`     |
| `RandomDateRecent()`              | Random recent datetime           | string                   | `Thu Apr 14 12:14:25 UTC 2022`, `Thu Apr 14 12:14:33 UTC 2022`     |
| `RandomWeekday()`                 | Random weekday                   | string                   | `Wednesday`, `Thursday`                                            |
| `RandomMonth()`                   | Random month                     | string                   | `March`, `August`                                                  |

### Domains, Emails, and Usernames

| Method                          | Description                      | Type                     | Examples                                                     |
| -------------                   | ------------------------------   | -----------------        | ---------------------------------------                      | 
| `RandomDomainName()`            | Random domain name               | string                   | `claudiabeier.com`, `chasescott.ac.uk`                       |
| `RandomDomainSuffix()`          | Random domain suffix             | string                   | `scot`, `info`                                               |
| `RandomDomainWord()`            | Random domain word               | string                   | `jessecormier`, `victoriafisher`                             |
| `RandomEmail()`                 | Random email                     | string                   | `max.smitham@hotmail.com`, `ivy.wood@outlook.com`            |
| `RandomExampleEmail()`          | Random email with example domain | string                   | `isabella.bauch@example.net`, `tahlia.baumbach@example.org`  |
| `RandomUsername()`              | Random username                  | string                   | `Hayden.Heaney`, `Lara.Durgan`                               |
| `RandomUrl()`                   | Random URL                       | string                   | `https://sophiemills.scot`, `https://harrisongoldner.ltd.uk` |


### Files and Directories


| Method                             | Description                      | Type                 | Examples                                                              |
| -------------                      | ------------------------------   | -----------------    | ---------------------------------------                               | 
| `RandomFileName()`                 | Random file name                 | string               | `matthew_moore.gdoc`, `angus_daniel.mov`                              |
| `RandomFileType()`                 | Random file type                 | string               | `video`, `application`                                                |
| `RandomFileExtension()`            | Random file extension            | string               | `xml`, `gif`                                                          |
| `RandomCommonFileName()`           | Random common file name          | string               | `marcuswaters.csv`, `elizaowen.html`                                  |
| `RandomCommonFileType()`           | Random common file tyoe          | string               | `application`, `image`                                                |
| `RandomCommonFileExtension()`      | Random common file extension     | string               | `m3a`, `gif`                                                          |
| `RandomFilePath()`                 | Random file path                 | string               | `/usr/X11R6/hugoschmidt.mp4`, `/usr/libexec/mitchelloreilly.mp4`      |
| `RandomDirectoryPath()`            | Random directory path            | string               | `/proc`, `/usr`                                                       |
| `RandomMimeTypes()`                | Random mime type                 | string               | `image/gif`, `image/png`                                              |


### Stores

| Method                            | Description                      | Type                  | Examples                                              |
| -------------                     | ------------------------------   | -----------------     | ---------------------------------------               | 
| `RandomPrice()`                   | Random price (0.00-1000.00)      | string                | `445.60`, `802.67`                                    |
| `RandomProduct()`                 | Random product                   | string                | `Hat`, `Shoes`                                        |
| `RandomProductAdjective()`        | Random product adjective         | string                | `Fantastic`, `Tasty`                                  |
| `RandomProductMaterial()`         | Random product material          | string                | `Cotton`, `Steel`                                     |
| `RandomProductName()`             | Random product name              | string                | `Gorgeous Frozen Chicken`, `Recycled Bronze Sausages` |
| `RandomDepartment()`              | Random stores department         | string                | `Tools`, `Health`                                     |


### Grammar

| Method                                | Description                      | Type                   | Examples                                            |
| -------------                         | ------------------------------   | -----------------      | ---------------------------------------             | 
| `RandomNoun()`                        | Random noun                      | string                 | `interface`, `port`                                 |
| `RandomVerb()`                        | Random verb                      | string                 | `calculate`, `parse`                                |
| `RandomIngVerb()`                     | Random ing verb                  | string                 | `indexing`, `hacking`                               |
| `RandomAdjective()`                   | Random adjective                 | string                 | `multi-byte`, `neural`                              |
| `RandomWord()`                        | Random word                      | string                 | `program`, `overriding`                             |
| `RandomWords()`                       | Random words                     | string                 | `virtual navigate 1080p`, `reboot online alarm`     |
| `RandomLoremSentences()`              | Random phrase                    | string                 | `nobis iste omnis iusto aut. aliquam illo maxime aperiam nobis unde labore quos qui laboriosam voluptas sint perferendis quis iusto rerum voluptas vero tempore iure. veniam sit sit velit nihil sint in. odit ipsam incidunt earum voluptates mollitia repellat et unde et tempora sit. aliquam ullam ipsum odit ut labore omnis voluptate minima. `<br><br> `modi perferendis consectetur perspiciatis vero natus laborum quidem in dolorem. quia architecto quis error ducimus nihil hic voluptas. minus accusamus velit quis praesentium molestias odit sequi laboriosam magnam sapiente commodi. quos corrupti eveniet id suscipit unde dolorem adipisci accusantium qui ipsam molestiae voluptatem qui dolores omnis ducimus aut. `                 |

### Lorem Ipsum

| Method                           | Description                      | Type                   |  Examples                                 |
| -------------                    | ------------------------------   | -----------------      | ---------------------------------------  | 
| `RandomLoremWord()`              | Random lorem ipsum word          | string                 | `iusto`, `quaerat`                 |
| `RandomLoremWords()`             | Random lorem ipsum words         | string                 | `officiis iusto`, `deleniti molestiae veritatis non `                 |
| `RandomLoremSentence()`          | Random lorem ipsum sentence      | string                 | `delectus iure qui qui veritatis sit accusantium eaque dicta tempore saepe qui corporis rerum nulla fugiat iure soluta esse.`<br><br> `aspernatur saepe harum ex minus molestias fugit minus amet et et fugiat impedit sit eum quasi.`                 |
| `RandomLoremSentences()`         | Random lorem ipsum sentences     | string                 | `ab perferendis cupiditate atque saepe sed nulla ut deserunt consequuntur architecto maiores qui quis a voluptatem possimus. sit quia vel perferendis impedit et similique dolorem omnis qui ut. minima quasi totam velit dolorem excepturi assumenda exercitationem ut aut. nulla fugit illo odio sequi et omnis quo officiis est aliquam enim quia harum. debitis corrupti iste enim voluptatem tenetur laborum id ut nobis earum. `<br><br> `reprehenderit reprehenderit sed vitae ut ipsam totam numquam vitae laboriosam ut ut voluptas aut. ut adipisci optio quia quam velit quisquam eum in est iure. natus voluptatem eius est beatae voluptas nihil aperiam nesciunt sit et aspernatur. `                 |
| `RandomLoremParagraph()`         | Random lorem ipsum paragraph     | string                 | `provident asperiores dolorum beatae totam ut culpa rem neque quae rerum voluptas. pariatur aut cum dolore laudantium cum et error placeat rerum eligendi. aspernatur voluptas inventore porro necessitatibus sit magnam fuga et adipisci. occaecati voluptates voluptatem non ut atque placeat beatae veniam aut perspiciatis ut numquam sapiente minima doloribus. hic impedit atque et tempora doloribus sit sint ipsam. ab rem officiis ut cum ipsum quia minus officia odit aut porro doloremque quibusdam eligendi ab atque cupiditate. vitae in nostrum nobis cum labore qui. corporis fugiat optio eum laborum quos maxime commodi reiciendis et in. quia magni ut laboriosam mollitia ullam voluptatem explicabo corporis veniam. quia possimus vitae rerum labore qui omnis. vel eum quod impedit fugiat mollitia dolorem deserunt quis. id quod alias repellat provident molestiae voluptate tempore iste minus quasi accusamus esse consequatur et voluptas debitis sunt deleniti. odio sit voluptas voluptatem excepturi. rerum sit velit molestiae impedit quia illo et quod. alias totam neque velit doloremque corporis. vel omnis dolores sit nam aut harum aut eaque dolores consectetur quaerat dolor facere deserunt. `<br><br> `porro eum minus quis unde ut temporibus nulla non asperiores fuga tenetur. molestias possimus totam in sapiente qui. qui architecto quis quas quis saepe quia illo est repudiandae deserunt neque est sit dolorum tempora est placeat ad quia non. quam deserunt quia iusto est harum provident. vero illo omnis rem sed consequatur qui ex cumque itaque sed non aut quia corrupti quia ratione qui modi et. `                 |
| `RandomLoremParagraphs()`        | Random lorem ipsum paragraphs    | string                 | `eum reiciendis modi ab minus vero exercitationem nisi voluptatibus nihil in suscipit ad soluta iure ut sunt. et aut est laborum sed et minima iste odit voluptatem consequatur commodi doloremque. possimus possimus voluptate adipisci sequi eum commodi dolorem sunt dolore totam. adipisci ducimus eos velit sed provident et aut modi architecto velit quia provident saepe enim sunt voluptas et. reiciendis velit magni voluptas ut delectus voluptatem qui dolorem ipsum. sunt temporibus sit eos voluptates. molestias distinctio autem nihil aperiam dolorem illo qui. architecto ratione et facilis accusamus incidunt voluptates officia deleniti ut praesentium saepe tempore illum sunt ut consectetur. natus ut et voluptatem quidem labore et eaque cumque voluptate non et non. eum natus vel minima et non et architecto sunt quis deserunt molestias voluptatem et nemo repellendus est in. architecto aliquid mollitia asperiores. facilis voluptatem distinctio voluptas et aliquam dolorum et qui ipsum alias aut et eum aut. `<br><br> `qui voluptatibus modi itaque nam a explicabo eos nihil excepturi porro omnis assumenda. nesciunt consequuntur labore et a dolorem adipisci. quos eligendi atque expedita assumenda qui. accusamus exercitationem omnis in doloribus sint ab a corporis tempore. explicabo placeat non repudiandae assumenda ut nihil. quo dolor qui illum libero velit iure et et aliquid rerum et praesentium. ipsum est iste occaecati reiciendis nesciunt voluptates id numquam ad non totam at est quo illo officiis. autem quia et quo aliquam qui qui blanditiis consequuntur dolor quo consequatur corporis dicta cumque autem architecto. est harum ipsum et sequi iure adipisci optio vero voluptatem et molestiae dolorem ea facilis sequi sit. `                 |
| `RandomLoremParagraphs()`        | Random lorem ipsum text          | string                 | `quia officiis et delectus ab maxime mollitia et enim laborum. aliquid ullam blanditiis voluptatem ipsam labore. perferendis aut necessitatibus eos aut numquam et dolorem nemo perspiciatis aut cupiditate reprehenderit aut natus. illum aspernatur explicabo veniam dignissimos fugiat necessitatibus maiores nostrum numquam repellendus laboriosam repudiandae et aspernatur similique. sequi doloribus consequatur possimus optio aliquam et quis quia eum tenetur reiciendis sed dicta voluptatem accusantium ex repudiandae eaque. cum et laudantium enim facere repellat voluptatum ullam aspernatur blanditiis aut incidunt dignissimos ipsa dolorem. reprehenderit voluptatem fugiat laboriosam ipsum sed voluptatibus qui inventore consequatur sunt quas. et eos asperiores sapiente tempore aut ea cum consequuntur. delectus et qui est et cumque enim dolorum et consequuntur dolorem nemo. `<br><br> `neque facilis enim aut et perspiciatis distinctio et delectus. hic voluptatem maxime libero incidunt deleniti corrupti error ducimus ex. omnis qui quod necessitatibus architecto voluptas molestias consequatur ipsum et vero id voluptatum. ullam dolore nemo inventore dolore voluptatem sed qui consectetur accusantium quo ab deleniti nemo et nihil atque corrupti. adipisci ullam modi facere consequatur sunt vel. maxime sapiente unde amet ipsa possimus ut voluptates nesciunt molestias. voluptatem est iure facilis animi quasi ex eos blanditiis consequatur. quis nemo sunt aut maxime illo nihil voluptatibus aliquid. dicta sint quia necessitatibus in eius adipisci dolores et. `                 |
| `RandomLoremSlug()`              | Random lorem ipsum slug          | string                 | `cupiditate-consectetur-et`, `adipisci-suscipit-qui`                 |
| `RandomLoremLines()`             | Random lorem ipsum lines         | string                 | `esse soluta voluptatem atque aut ut exercitationem optio est expedita animi qui in est reiciendis eum qui quis dolore.\n neque perspiciatis nihil ut iure qui libero voluptas nobis corrupti.\n beatae dolores sit rerum unde nesciunt rerum nesciunt adipisci sunt voluptatem recusandae est eum.\n quia velit qui necessitatibus perspiciatis neque rerum facilis architecto odio laudantium facere voluptas tenetur.`<br><br> `voluptatem in ut ut facilis et fugiat sunt suscipit facere perspiciatis est at suscipit repudiandae tempore.\n natus nostrum aut odio necessitatibus qui sunt et assumenda eos qui nihil.\n aut mollitia rerum quos non rerum inventore velit tempora libero aut nobis dolores eligendi quod aut qui amet rerum.\n esse architecto veritatis non atque dolorum sint maxime veniam a velit vel aut qui in totam saepe enim ad nemo omnis.`                  |


## Communication

You can join our [Discord Server](https://discord.gg/9KdnrSUZQg) for issues, feature requests, feedbacks or anything else. 
