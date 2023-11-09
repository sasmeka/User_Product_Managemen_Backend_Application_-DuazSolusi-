
<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<div align="center">
  <a href="https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-">
    <img src="duaz-logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">User Product Managemen Back-End Application (PT. Duaz Solusi)</h3>

  <p align="center">
    <br />
    <a href="https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-"><strong>Explore the docs »</strong></a>
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation-&-running">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project aims to build a Rest-API built using Go-lang and other supporting modules which was completed in 2 weeks. This project is limited by the task requirements given by the trainer. The features include login, register, service user, product, size, and delivery method.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This project is based on the following packages:

* [![go][go.js]][go-url]
* [![postgresql][postgresql.js]][postgresql-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This project worker can follow the steps below:

### Prerequisites

1. Install [go-lang](https://go.dev/dl/)

### Installation & Running

1. Clone the repo
   ```sh
   git clone https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-.git
   ```
2. Install packages
   ```sh
   go mod download
   ```
3. Please configure .env
4. Please install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to migrate the database. command to migrate the database as follows.
   ```sh
   migrate -path ./migrations -database "postgresql://user:pass@localhost/namedatabase?port=5432&sslmode=disable&search_path=public" -verbose up
   ```
   What must be changed in the command above is user, pass, namedatabase. You can also manually migrate the SQL file in the migrations folder.
5. Run
   ```sh
   go run ./cmd/main.go
   ```
4. Below are comments to run all unit tests
   ```sh
   go test -v ./...
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Install [postman](https://www.postman.com/)
2. Download file workspace [Product_DuazSolusi.postman_collection.json](https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-/raw/main/Product_DuazSolusi.postman_collection.json)
3. Import the workspace into your Postman application.
4. Go to the workspace you imported -> auth -> register. Do registration and login.
5. Please try to get data on the endpoint that requires a login/token. To insert a token, you can do it on the authorization tab and select Bearer Token

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Verdi Sasmeka - [@vrd_meka](https://twitter.com/vrd_meka) - verdysas@gmail.com

Project Link: [https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-](https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-.svg?style=for-the-badge
[contributors-url]: https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-.svg?style=for-the-badge
[forks-url]: https://github.com/sasmeka/User_Product_Managemen_Backend_Application_-DuazSolusi-/network/members
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/verdi-sasmeka-62b91b132/
[go.js]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white
[go-url]: https://go.dev
[postgresql.js]: https://img.shields.io/badge/Postgresql-4169E1?style=for-the-badge&logo=postgresql&logoColor=white
[postgresql-url]: https://www.postgresql.org/