# Step 1

## Model

* Начинаем работу с файла [artist.go](/musicstore/model/artist.go). Выполнить `go get gopkg.in/mgo.v2`, чтобы import нашел "gopkg.in/mgo.v2/bson"

* Нужно  создать 4 типа:
  * Genre - это  enums c жанрами музыки ([см. пример](/musicstore/model/enum_example))
    * К нему необходимо реализовать метод func (g Genre) String() string
  * Artist
    * Поля:
      * ID bson.ObjectId
      * Name string
      * Genres []Genre
      * Born *time.Time
    * Методы:
      * func NewArtist(name string) *Artist (Создать новый ID, остальные поля инициализируются по умолчанию)
      * func (a *Artist) AddGenre(g Genre)
      * func (a *Artist) SetBorn(t time.Time)
      * func (a *Artist) Validate() error (Логику валидации придумайте сами)
      * func (a *Artist) IsValid() bool (wrapper, который вызывает Validate())
  * Album
    * Поля:
      * ID
      * Title
      * Genres
      * Released
      * Rating       uint
    * Методы:
      * func NewAlbum(title string) *Album (Создать новый ID, остальные поля инициализируются по умолчанию)
      * func (a *Album) AddGenre(g Genre)
      * func (a *Album) SetReleased(t time.Time)
      * func (a *Album) Validate() error (Логику валидации придумайте сами)
      * func (a *Album) IsValid() bool (wrapper, который вызывает Validate())
  * Track
    * Поля:
      * ID
	  * Title
	  * Genres
	  * YearReleased int
	  * Duration     time.Duration
	  * Albums       []bson.ObjectId
	  * Artists      []bson.ObjectId
    * Методы:
      * func NewTrack(title string) *Track (Создать новый ID, остальные поля инициализируются по умолчанию)
      * func (t *Track) AddGenre(g Genre)
      * func (t *Track) AddAlbum(album *Album) error (проверить Validate() у Album и задать только его ID)
      * func (t *Track) AddArtist(artist *Artist) error (проверить Validate() у Artist и задать только его ID)
      * func (t *Track) SetYearReleased(t time.Time)
      * func (t *Track) Validate() error (Логику валидации придумайте сами)
      * func (t *Track) IsValid() bool (wrapper, который вызывает Validate())

* Предложить реализацию типа Stats, который будет содержать информацию, необходимую для вычисления рейтинга альбома,
а также сам алгоритм вычисления этого рейтинга. Можно параллельно учитывать и голоса обычных слушателей,
и голоса критиков -- как, например, это делается на [Metacritic](http://www.metacritic.com/) или [Allmusic](http://www.allmusic.com/).
  * [пример шкал и названий категорий](http://www.metacritic.com/about-metascores)
  * [еще один пример логики вычисления рейтинга альбома](http://www.tunequest.org/in-search-of-a-definitive-album-rating-formula/20070314/)
  * обсуждения формулы, которая используется в IMDb для топ-250: [1](http://www.quora.com/How-is-a-Movies-weighted-rating-calculated-on-IMDb) и [2](http://math.stackexchange.com/questions/169032/understanding-the-imdb-weighted-rating-function-for-usage-on-my-own-website)

* Также продумать логику валидации голосов

* Пишем табличные тесты (table-driven tests) при помощи пакета testing в файле [artist_test.go](/musicstore/model/artist_test.go) ([см. пример](/musicstore/model/table_driven_ test_example))

* Теперь нужно развести все написанное по соответствующим файлам:
  * artist.go
  * genre.go
  * album.go
  * track.go

* Тесты по-прежнему должны проходить