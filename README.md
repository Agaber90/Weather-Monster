# Weather-Monster
This a Golang project is going to be used by different institutions to monitor weather temperature and get forecasts.

# install the following packages 
<br />
<ol>
  <li> go get github.com/go-sql-driver/mysql </li>
  <li> go get github.com/labstack/echo </li>
  <li> go get -u github.com/spf13/viper </li>
  <li> go get github.com/pdfcrowd/pdfcrowd-go</li>
  <li> go get github.com/sirupsen/logrus</li>
  <li> go get github.com/stretchr/testify/mock</li>
  <li> go get github.com/bxcodec/faker</li>
  <li> go get github.com/stretchr/testify/assert</li>
  <li> go get gopkg.in/DATA-DOG/go-sqlmock.v2</li>
  <li> go get github.com/stretchr/testify/require</li>
  </ol>
<br />

<p>the architecture that is being used is a clean architecture with dependency injection.</p>
  <p>This project has 4 Domain layer as stated below
    
    city
├── delivery
│   └── http
│       ├── city_handler.go
│       └── city_test.go
├── mocks
│   ├── cityRepository.go
│   └── cityUsecase.go
├── repository //Encapsulated Implementation of Repository Interface
│   ├── sql_city.go
│   └── sql_city_test.go
├── repository.go // Repository Interface
├── usecase //Encapsulated Implementation of Usecase Interface
│   ├── city_usecase_test.go
│   └── city_usecase.go
└── usecase.go // Usecase Interface.
    
</p>
# Rule of Clean Architecture by Uncle Bob
<ol>
 <li> Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.</li>
<li>Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.</li>
<li>Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.</li>
<li>Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.</li>
</ol>

## How To Run This Project

Since the project already use Go Module, I recommend to put the source code in any folder but GOPATH. and make sure to run the mysql scripts



 
