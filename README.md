
<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<br />
<div align="center">
  <a href="https://github.com/sasmeka/backend_coffeshop_with_go">
    <img src="https://camo.githubusercontent.com/72d4e416bd802a1abc16d86e9d7d7a62318fca378d103f97fda207ef7d61463d/68747470733a2f2f7974332e67677068742e636f6d2f7974632f414b65644f4c543759443978365069522d4366624262464333777a3257617469495a4672495f4930762d366b3d733930302d632d6b2d63307830306666666666662d6e6f2d726a" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">User Product Managemen Back-End Application (PT. Duaz Solusi)</h3>

  <p align="center">
    <br />
    <a href="https://github.com/sasmeka/backend_coffeshop_with_go"><strong>Explore the docs »</strong></a>
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
   git clone https://github.com/sasmeka/backend_coffeshop_with_go.git
   ```
2. Install NPM packages
   ```sh
   go mod download
   ```
3. please configure .env and config.yml
4. Run
   ```sh
   go run ./cmd/main.go
   ```
5. 

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Install [postman](https://www.postman.com/)
2. Visit the following link to export Coffee Shop postman workspace 
   ```sh
   https://www.postman.com/avionics-meteorologist-14374576/workspace/tickitz/collection/22380820-2a8492cd-b607-4943-b31d-9d8c50cc4543?action=share&creator=22380820
   ```
3. Import the workspace that you already have in stage 2 into the postman application
4. Go to Coffee Shop workspace -> auth -> register. Do registration and login.
5. Please try to do get data with the token. To insert a token, you can do it on the authorization tab and select Bearer Token

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Verdi Sasmeka - [@vrd_meka](https://twitter.com/vrd_meka) - verdysas@gmail.com

Project Link: [https://github.com/sasmeka/backend_coffeshop_with_go](https://github.com/sasmeka/backend_coffeshop_with_go)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/sasmeka/Tickitz_Backend.svg?style=for-the-badge
[contributors-url]: https://github.com/sasmeka/backend_coffeshop_with_go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/sasmeka/Tickitz_Backend.svg?style=for-the-badge
[forks-url]: https://github.com/sasmeka/backend_coffeshop_with_go/network/members
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/verdi-sasmeka-62b91b132/
[go.js]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white
[go-url]: https://go.dev
[postgresql.js]: https://img.shields.io/badge/Postgresql-4169E1?style=for-the-badge&logo=postgresql&logoColor=white
[postgresql-url]: https://www.postgresql.org/