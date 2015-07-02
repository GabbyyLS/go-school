# Step 2

## Store

* В [store.go](/musicstore/store/store.go) создаем интерфейс Store, который будет реализован при помощи Mongo DB и mock-объекта. Далее:
  * Выполнить `go get github.com/vektra/mockery`, затем `GOPATH/bin/mockery`
  * Перейти в go-school/musicstore/store и выполнить `mockery -all` (mockery генерирует mock-объекты для интерфейсов)
  * В результате в go-school/musicstore/store/mocks должен появиться Store.go, в котором сгенерированный mock-объект реализует наш интерфейс.

* Там же: прописать все функции (см. приведенный пример) ~~-- они понадобятся нам позже, когда мы будем работать с handler'ами~~
  * Выполнить `go get golang.org/x/net/context`
  * В [context.go](/musicstore/store/context.go) реализуем функцию FromContext(ctx context.Context) Store, которая возвращает Store из контекста (см. [документацию к net/context](https://godoc.org/golang.org/x/net/context)).

* Переходим в go-school/musicstore/store/mocks, где уже лежит сгенерированный ранее Store.go
  * Добавить artist.go, album.go и track.go и в них реализовать методы, создающие mock-объекты для различных сценариев (например, исполнителя с пустым/несуществующим/рабочим ID и т.д.)
  * Пишем record.go, в котором функция NewMockStore создает новый mock-объект Store и описывает ожидания (expectations) от вызова его методов (сгенерированных в Store.go) c mock-объектами из artist.go, album.go и track.go [см. описание Mock.On() в документации](https://github.com/stretchr/testify/blob/master/mock/mock.go)

* Переходим в go-school/musicstore/store/
  * Пишем в store_test.go тесты для функций из store.go (здесь также пригодятся artist.go, album.go и track.go, созданные ранее в go-school/musicstore/store/mocks)

* Переходим в go-school/musicstore/model/
  * В файле artist.go уже добавлены json-поля и bson-поле для работы с JSON-объектом, содержащим информацию об исполнителе, чтобы, используя методы [Marshal()](http://golang.org/pkg/encoding/json/#Marshal) и [Unmarshal()](http://golang.org/pkg/encoding/json/#Unmarshal), мы могли переводить JSON-объект в struct и наоборот
  * В файле artist_test.go протестировать работу Marshal() Unmarshal()
  * Прописать необходимые поля (ctrl+j) в файлах album.go, track.go, genre.go и тоже все протестировать

* Переходим в go-school/musicstore/store/mongo
  * **(работа с файлами .json - запись/считывание информации?)**
  * Нам понадобится драйвер для работы с mongoDB -- выполнить `go get gopkg.in/mgo.v2`. Примеры работы с mgo: [1](https://gist.github.com/border/3489566), [2](http://labix.org/mgo). [Документация](http://godoc.org/labix.org/v2/mgo)
  * В качестве url в Dial использовать "mongodb://localhost" или "127.0.0.1"
  * В файле artist.go (не забыть импортировать в него go-school/musicstore/model/artist) создаем тип artistStore, реализующий интерфейс Store, т.е. имеющий реализованные методы с необходимыми сигнатурами (например, GetArtistByName(name string) (*model.Artist, error) и т.д.). CreateArtist -- это wrapper, внутри которого мы используем Insert из mgo, внутри GetArtistById -- FindId и т.д.
  * Создать artist_test.go, в котором протестировать созданные CRUD-методы



