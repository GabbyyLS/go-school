# Step 1

## Model

* Начинаем работу с файла [artist.go](/musicstore/model/artist.go). Выполнить go get gopkg.in/mgo.v2, чтобы import нашел "gopkg.in/mgo.v2/bson"

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
      * func (a *Artist) Exists() bool
      * func (a *Artist) AddGenre(g Genre)
      * func (a *Artist) SetBorn()
      * func (a *Artist) Validate() (Логику валидации придумайте сами)
      * func (a *Artist) IsValid() bool
  * Album
    * Поля:
      * ID
      * Title
      * Genres
      * Released
    * Методы:
      * func NewAlbum(title string) *Album (Создать новый ID, остальные поля инициализируются по умолчанию)
      * func (a *Album) Exists() bool
      * func (a *Album) AddGenre(g Genre)
      * func (a *Album) SetReleased()
      * func (a *Album) Validate() (Логику валидации придумайте сами)
      * func (a *Album) IsValid() bool
  * Track
    * Поля:
      * ID
	  * Title
	  * Genres
	  * YearReleased int
	  * Duration     time.Duration
	  * Rating       uint
	  * Albums       []bson.ObjectId
	  * Artists      []bson.ObjectId
    * Методы:
      * func NewTrack(title string) *Track (Создать новый ID, остальные поля инициализируются по умолчанию)
      * func (a *Track) Exists() bool
      * func (t *Track) SetRating()
      * func (t *Track) AddGenre(g Genre)
      * func (t *Track) AddAlbum(id bson.ObjectId)
      * func (t *Track) AddArtist(id bson.ObjectId)
      * func (t *Track) SetDuration()
      * func (t *Track) SetYearReleased()
      * func (a *Track) Validate() (Логику валидации придумайте сами)
      * func (a *Track) IsValid() bool

* Пишем табличные тесты (table-driven tests) при помощи пакета testing в файле [artist_test.go](/musicstore/model/artist_test.go) ([см. пример](/musicstore/model/table_driven_ test_example))

* Теперь нужно развести все написанное по соответствующим файлам:
  * artist.go
  * genre.go
  * album.go
  * track.go

* Тесты по-прежнему должны проходить