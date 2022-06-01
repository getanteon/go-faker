package faker

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

type Faker struct {
	Generator *rand.Rand
}

func (f Faker) RandomBoolean() bool {
	return faker.New().Bool()
}

func (f Faker) RandomInt() int {
	return faker.New().IntBetween(0, 1000)
}

func (f Faker) RandomStringMaxLenght(l int) string {
	return faker.New().RandomStringWithLength(l)
}

func (f Faker) RandomFloatBetween(maxDecimals, min, max int) float64 {
	return faker.New().RandomFloat(maxDecimals, min, max)
}

func (f Faker) RandomIntBetween(min, max int) int {
	return faker.New().IntBetween(min, max)
}

func (f Faker) RandomUUID() uuid.UUID {
	return uuid.New()
}

func (f Faker) CurrentISOTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}

func (f Faker) CurrentTimestamp() int64 {
	return time.Now().Unix()
}

func (f Faker) RandomGuid() uuid.UUID {
	return uuid.New()
}

func (f Faker) RandomAddressLongitude() float64 {
	return faker.New().Address().Longitude()
}

func (f Faker) RandomAddressLatitude() float64 {
	return faker.New().Address().Latitude()
}

func (f Faker) RandomAddressCountry() string {
	return faker.New().Address().Country()
}

func (f Faker) RandomAddressStreetAddress() string {
	return faker.New().Address().StreetAddress()
}

func (f Faker) RandomAddresStreetName() string {
	return faker.New().Address().StreetName()
}

func (f Faker) RandomAddressCity() string {
	return faker.New().Address().City()
}

func (f Faker) RandomPersonSuffix() string {
	return faker.New().Person().Suffix()
}

func (f Faker) RandomPersonTitle() string {
	return faker.New().Person().Title()
}

func (f Faker) RandomPersonName() string {
	return faker.New().Person().Name()
}

func (f Faker) RandomPersonLastName() string {
	return faker.New().Person().LastName()
}

func (f Faker) RandomPersonFirstName() string {
	return faker.New().Person().FirstName()
}

func (f Faker) RandomUserAgent() string {
	return faker.New().UserAgent().UserAgent()
}

func (f Faker) RandomLocale() string {
	return faker.New().Language().LanguageAbbr()
}

func (f Faker) RandomPassword() string {
	return faker.New().Internet().Password()
}

func (f Faker) RandomMACAddress() string {
	return faker.New().Internet().MacAddress()
}

func (f Faker) RandomSafeColorHex() string {
	return faker.New().Color().Hex()
}

func (f Faker) RandomSafeColorName() string {
	return faker.New().Color().SafeColorName()
}

func (f Faker) RandomDateFuture() string {
	min := time.Now().Unix()
	max := time.Now().Unix() * 10 / 9
	delta := max - min

	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)

	return randTime.Format(time.UnixDate)
}

func (f Faker) RandomDatePast() string {
	min := time.Now().Unix() - (time.Now().Unix() * 2 / 10)
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)

	return randTime.Format(time.UnixDate)
}

func (f Faker) RandomDateRecent() string {
	min := time.Now().Unix() - (time.Now().Unix() * 1 / 200)
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	randTime := time.Unix(sec, 0)

	return randTime.Format(time.UnixDate)
}

func (f Faker) RandomLoremWord() string {
	return LoremWords[f.Generator.Intn(len(LoremWords))]
}

func (f Faker) RandomLoremWords() string {
	words := ""
	size := f.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += f.RandomLoremWord() + " "
	}
	return words
}

func (f Faker) RandomLoremSentence() string {
	words := ""
	size := f.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += f.RandomLoremWords()
	}
	return words + "."
}

func (f Faker) RandomLoremSentences() string {
	words := ""
	size := f.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += f.RandomLoremSentence() + " "
	}
	return words
}

func (f Faker) RandomLoremLines() string {
	words := ""
	size := f.Generator.Intn(5-1) + 1
	for i := 0; i < size; i++ {
		words += f.RandomLoremSentence() + "\n"
	}
	return words
}

func (f Faker) RandomLoremParagraph() string {
	words := ""
	size := f.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += f.RandomLoremSentences()
	}
	return words
}

func (f Faker) RandomLoremParagraphs() string {
	words := ""
	size := 3
	for i := 0; i < size; i++ {
		words += f.RandomLoremSentences()
	}
	return words
}

func (f Faker) RandomLoremSlug() string {
	return f.RandomLoremWord() + "-" + f.RandomLoremWord() + "-" + f.RandomLoremWord()
}

func (f Faker) RandomNoun() string {
	return Nouns[f.Generator.Intn(len(Nouns))]
}

func (f Faker) RandomVerb() string {
	return Verbs[f.Generator.Intn(len(Verbs))]
}

func (f Faker) RandomIngVerb() string {
	return IngVerbs[f.Generator.Intn(len(IngVerbs))]
}

func (f Faker) RandomAdjective() string {
	return Adjectives[f.Generator.Intn(len(Adjectives))]
}

func (f Faker) RandomWord() string {
	words := append(append(append(Adjectives, IngVerbs...), Verbs...), Nouns...)
	return words[f.Generator.Intn(len(words))]
}

func (f Faker) RandomWords() string {
	words := ""
	size := f.Generator.Intn(6-2) + 2
	for i := 0; i < size; i++ {
		words += f.RandomWord() + " "
	}
	return words
}

func (f Faker) RandomDepartment() string {
	return StoreDepartments[f.Generator.Intn(len(StoreDepartments))]
}

func (f Faker) RandomProductName() string {
	return f.RandomProductAdjective() + " " + f.RandomProductMaterial() + " " + f.RandomProduct()
}

func (f Faker) RandomProductMaterial() string {
	return ProductMaterials[f.Generator.Intn(len(ProductMaterials))]
}

func (f Faker) RandomProductAdjective() string {
	return ProductAdjectives[f.Generator.Intn(len(ProductAdjectives))]
}

func (f Faker) RandomProduct() string {
	return Products[f.Generator.Intn(len(Products))]
}

func (f Faker) RandomPrice() string {
	return strconv.Itoa(f.Generator.Intn(1000)) + "." + strconv.Itoa(f.Generator.Intn(99-10)+10)
}

func (f Faker) RandomFilePath() string {
	return f.RandomDirectoryPath() + "/" + f.RandomDomainWord() + "." + f.RandomFileExtension()
}

func (f Faker) RandomMimeTypes() string {
	return CommonMimeTypes[f.Generator.Intn(len(CommonMimeTypes))]
}

func (f Faker) RandomDirectoryPath() string {
	return DirectoryPaths[f.Generator.Intn(len(DirectoryPaths))]
}

func (f Faker) RandomCommonFileExtension() string {
	return CommonFileExtensions[f.Generator.Intn(len(CommonFileExtensions))]
}

func (f Faker) RandomCommonFileType() string {
	return CommonFileTypes[f.Generator.Intn(len(CommonFileTypes))]
}

func (f Faker) RandomCommonFileName() string {
	return f.RandomDomainWord() + "." + f.RandomFileExtension()
}

func (f Faker) RandomFileType() string {
	return FileTypes[f.Generator.Intn(len(FileTypes))]
}

func (f Faker) RandomFileExtension() string {
	return FileExtensions[f.Generator.Intn(len(FileExtensions))]
}

func (f Faker) RandomFileName() string {
	return strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))]+"_"+
		LastNames[f.Generator.Intn(len(LastNames))]) + "." + f.RandomFileExtension()
}

func (f Faker) RandomUrl() string {
	return f.RandomProtocol() + "://" + strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))]+
		LastNames[f.Generator.Intn(len(LastNames))]) + "." + f.RandomDomainSuffix()
}

func (f Faker) RandomUsername() string {
	return FirstNames[f.Generator.Intn(len(FirstNames))] + "." + LastNames[f.Generator.Intn(len(LastNames))]
}

func (f Faker) RandomExampleEmail() string {
	return strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))]+"."+
		LastNames[f.Generator.Intn(len(LastNames))]) + "@" + ExampleEmails[f.Generator.Intn(len(ExampleEmails))]
}

func (f Faker) RandomEmail() string {
	return strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))]+"."+
		LastNames[f.Generator.Intn(len(LastNames))]) + "@" + Emails[f.Generator.Intn(len(Emails))]
}

func (f Faker) RandomDomainWord() string {
	return strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))] + LastNames[f.Generator.Intn(len(LastNames))])
}

func (f Faker) RandomDomainSuffix() string {
	return DomainSuffixes[f.Generator.Intn(len(DomainSuffixes))]
}

func (f Faker) RandomDomainName() string {
	return strings.ToLower(FirstNames[f.Generator.Intn(len(FirstNames))]+LastNames[f.Generator.Intn(len(LastNames))]) +
		"." + f.RandomDomainSuffix()
}

func (f Faker) RandomMonth() string {
	return Months[f.Generator.Intn(len(Months))]
}

func (f Faker) RandomWeekday() string {
	return WeekDays[f.Generator.Intn(len(WeekDays))]
}

func (f Faker) RandomDatabaseColumn() string {
	return DatabaseColumns[f.Generator.Intn(len(DatabaseColumns))]
}

func (f Faker) RandomDatabaseType() string {
	return DatabaseTypes[f.Generator.Intn(len(DatabaseTypes))]
}

func (f Faker) RandomDatabaseCollation() string {
	return DatabaseCollations[f.Generator.Intn(len(DatabaseCollations))]
}

func (f Faker) RandomDatabaseEngine() string {
	return DatabaseEngines[f.Generator.Intn(len(DatabaseEngines))]
}

func (f Faker) RandomCatchPhrase() string {
	return f.RandomCatchPhraseAdjective() + " " + f.RandomCatchPhraseDescriptor() + " " + f.RandomCatchPhraseNoun()
}

func (f Faker) RandomCatchPhraseAdjective() string {
	return CompanyAdjectives[f.Generator.Intn(len(CompanyAdjectives))]
}

func (f Faker) RandomCatchPhraseDescriptor() string {
	return CompanyDescriptors[f.Generator.Intn(len(CompanyDescriptors))]
}

func (f Faker) RandomCatchPhraseNoun() string {
	return CompanyNouns[f.Generator.Intn(len(CompanyNouns))]
}

func (f Faker) RandomBsNouns() string {
	return BusinessNouns[f.Generator.Intn(len(BusinessNouns))]
}

func (f Faker) RandomBsBuzzVerbs() string {
	return BusinessVerbs[f.Generator.Intn(len(BusinessVerbs))]
}

func (f Faker) RandomBsAdjective() string {
	return BusinessAdjectives[f.Generator.Intn(len(BusinessAdjectives))]
}

func (f Faker) RandomBs() string {
	return f.RandomBsBuzzVerbs() + " " + f.RandomBsAdjective() + " " + f.RandomBsNouns()
}

func (f Faker) RandomCompanySuffix() string {
	return CompanySuffixes[f.Generator.Intn(len(CompanySuffixes))]
}

func (f Faker) RandomCompanyName() string {
	return FirstNames[f.Generator.Intn(len(FirstNames))] + " " +
		LastNames[f.Generator.Intn(len(LastNames))]
}

func (f Faker) RandomBitcoin() string {
	const letters = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

	b := make([]byte, f.Generator.Intn(35-26)+26)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (f Faker) RandomCurrencySymbol() string {
	return CurrencySymbols[f.Generator.Intn(len(CurrencySymbols))]
}

func (f Faker) RandomCurrencyCode() string {
	return CurrencyCodes[f.Generator.Intn(len(CurrencyCodes))]
}

func (f Faker) RandomCurrencyName() string {
	return CurrencyNames[f.Generator.Intn(len(CurrencyNames))]
}

func (f Faker) RandomTransactionType() string {
	return BankTransactionTypes[f.Generator.Intn(len(BankTransactionTypes))]
}

func (f Faker) RandomCreaditCardMask() string {
	return strconv.Itoa(f.Generator.Intn(9999-1000) + 1000)
}

func (f Faker) RandomBankAccountName() string {
	return BankAccounts[f.Generator.Intn(len(BankAccounts))]
}

func (f Faker) RandomBankAccountBic() string {
	return BankAccountBics[f.Generator.Intn(len(BankAccountBics))]
}

func (f Faker) RandomBankAccountIban() string {
	randomIbanFormat := BankAccountIbans[f.Generator.Intn(len(BankAccountIbans))]
	randomIbanFormatList := strings.Split(randomIbanFormat, " ")

	randomIban := randomIbanFormatList[0]
	numberLen, _ := strconv.Atoi(randomIbanFormatList[1])
	for i := 0; i < numberLen; i++ {
		randomIban += strconv.Itoa(f.Generator.Intn(9-1) + 1)
	}
	return randomIban
}

func (f Faker) RandomBankAccount() string {
	return strconv.Itoa(f.Generator.Intn(99999999-10000000) + 10000000)
}

func (f Faker) RandomDataImageUri() string {
	return "data:image/svg+xml;charset=UTF-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20version%3D" +
		"%221.1%22%20baseProfile%3D%22full%22%20width%3D%22undefined%22%20height%3D%22undefined%22%3E%3Crect%20width%" +
		"3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22grey%22%2F%3E%3Ctext%20x%3D%22NaN%22%20y%3D%22NaN%22%20fo" +
		"nt-size%3D%2220%22%20alignment-baseline%3D%22middle%22%20text-anchor%3D%22middle%22%20fill%3D%22white%22%3Eu" +
		"ndefinedxundefined%3C%2Ftext%3E%3C%2Fsvg%3E"
}

func (f Faker) RandomAvatarImage() string {
	return "http://placeimg.com/640/480/people"
}

func (f Faker) RandomImageURL() string {
	return "http://placeimg.com/640/480"
}

func (f Faker) RandomAbstractImage() string {
	return "http://placeimg.com/640/480/abstract"
}

func (f Faker) RandomAnimalsImage() string {
	return "http://placeimg.com/640/480/animals"
}

func (f Faker) RandomBusinessImage() string {
	return "http://placeimg.com/640/480/business"
}

func (f Faker) RandomCatsImage() string {
	return "http://placeimg.com/640/480/cats"
}

func (f Faker) RandomCityImage() string {
	return "http://placeimg.com/640/480/city"
}

func (f Faker) RandomFoodImage() string {
	return "http://placeimg.com/640/480/food"
}

func (f Faker) RandomNightlifeImage() string {
	return "http://placeimg.com/640/480/nightlife"
}

func (f Faker) RandomFashionImage() string {
	return "http://placeimg.com/640/480/fashion"
}

func (f Faker) RandomPeopleImage() string {
	return "http://placeimg.com/640/480/people"
}

func (f Faker) RandomNatureImage() string {
	return "http://placeimg.com/640/480/nature"
}

func (f Faker) RandomSportsImage() string {
	return "http://placeimg.com/640/480/sports"
}

func (f Faker) RandomTransportImage() string {
	return "http://placeimg.com/640/480/transport"
}

func (f Faker) RandomCountryCode() string {
	return CountryCodes[f.Generator.Intn(len(CountryCodes))]
}

func (f Faker) RandomPhoneNumber() string {
	return strconv.Itoa(f.Generator.Intn(999-100)+100) +
		"-" + strconv.Itoa(f.Generator.Intn(999-100)+100) +
		"-" + strconv.Itoa(f.Generator.Intn(9999-1000)+1000)
}

func (f Faker) RandomPhoneNumberExt() string {
	return strconv.Itoa(f.Generator.Intn(99-10)+10) + "-" + f.RandomPhoneNumber()
}

func (f Faker) RandomJobArea() string {
	return JobAreas[f.Generator.Intn(len(JobAreas))]
}

func (f Faker) RandomJobDescriptor() string {
	return JobDescriptors[f.Generator.Intn(len(JobDescriptors))]
}

func (f Faker) RandomJobType() string {
	return JobTypes[f.Generator.Intn(len(JobTypes))]
}

func (f Faker) RandomJobTitle() string {
	return f.RandomJobDescriptor() + " " + f.RandomJobArea() + " " + f.RandomJobType()
}

func (f Faker) RandomSemver() string {
	return strconv.Itoa(f.Generator.Intn(9)) +
		"." + strconv.Itoa(f.Generator.Intn(9)) +
		"." + strconv.Itoa(f.Generator.Intn(9))
}

func (f Faker) RandomProtocol() string {
	return Protocols[f.Generator.Intn(len(Protocols))]
}

func (f Faker) RandomAbbreviation() string {
	return Abbreviations[f.Generator.Intn(len(Abbreviations))]
}

func (f Faker) RandomAlphanumeric() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

	seed := rand.NewSource(time.Now().Unix())
	generator := rand.New(seed)

	b := make([]byte, 1)
	for i := range b {
		b[i] = letters[generator.Intn(len(letters))]
	}
	return string(b)
}

func (f Faker) Ipv6() string {
	ips := []string{}
	ipv6Alphabet := []string{
		"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for j := 0; j < 8; j++ {
		block := ""
		for w := 0; w < 4; w++ {
			block = block + ipv6Alphabet[rand.Intn(len(ipv6Alphabet))]
		}

		ips = append(ips, block)
	}

	return strings.Join(ips, ":")
}

func (f Faker) RandomDigitNotNull() int {
	return f.Generator.Int()%8 + 1
}

func (f Faker) RandomDigit() int {
	return f.Generator.Int() % 10
}

func (f Faker) RandomFloat() float64 {
	return f.Generator.Float64()
}

func (f Faker) RandomString() string {
	return f.RandomStringWithLength(10)
}

func (f Faker) RandomStringWithLength(l int) string {
	r := []string{}
	for i := 0; i < l; i++ {
		r = append(r, f.RandomLetter())
	}
	return strings.Join(r, "")
}

func (f Faker) RandomLetter() string {
	return fmt.Sprintf("%c", f.IntBetween(97, 122))
}

func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff < 0 {
		diff = 0
	}

	if diff == 0 {
		return min
	}

	if diff == maxInt {
		return f.Generator.Intn(diff)
	}

	return f.Generator.Intn(diff+1) + min
}

func NewFaker() (f Faker) {
	seed := rand.NewSource(time.Now().Unix())
	generator := rand.New(seed)
	f = Faker{Generator: generator}
	return
}
