[![License: MIT][license-shield]][license-url]
![Size](https://img.shields.io/github/repo-size/totoledao/brew-journal)
![Platform](https://img.shields.io/badge/platform-Web-7F00FF)

[![LinkedIn][linkedin-shield]][linkedin-url]

[![go][go-shield]][go-url][![htmx][htmx-shield]][htmx-url]

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/totoledao/brew-journal">
    <!-- <img src="web\src\assets\logo.svg" alt="SpaceTime Logo" width="250"> -->
  </a>
  
  <p align="center">
    Brew Journal: Brew, Log, Enjoy.
    <br />
    <a href="https://github.com/totoledao/brew-journal"><strong>Explore the docs »</strong></a>    
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
        <li><a href="#technologies">Technologies</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>    
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>    
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

**A journal to keep track of brewing experiences!**

The primary objective of this project is to learn about [htmx](https://htmx.org/), Go and it's ecosystem.<br>
The project uses Air for live reloading and a Web Socket to ping the server and reload the page when the server goes down at every reloading.<br>

### Built With

- [Go][go-url]
- [htmx][htmx-url]

### Technologies

- [Echo](https://echo.labstack.com/)
- [Air](https://github.com/cosmtrek/air)
- [Tailwind CSS](https://tailwindcss.com/)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Go [Download](https://go.dev/dl/)
- Tailwind CSS standalone CLI [Download](https://github.com/tailwindlabs/tailwindcss/releases/)

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/totoledao/brew-journal.git
   ```
2. Install the dependencies
   ```sh
   go mod download
   ```
3. Start the server
   ```sh
   air
   ```
4. Start tailwindcss _(optional - it listens to style changes)_
   ```sh
   tailwindcss -i ./input.css -o ./static/output.css --watch
   ```

<!-- USAGE EXAMPLES -->

<!-- ## Usage -->

<!-- ![web-login](https://github.com/totoledao/totoledao/assets/40635662/60743232-836d-4190-96bc-828b88c560ed)
Create an account or Login using your GitHub account -->

<!-- LICENSE -->

## License

Distributed under the MIT License. See [`LICENSE`][license-url] for more information.

<!-- CONTACT -->

## Contact

Guilherme Toledo - guilherme-toledo@live.com

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/guilhermemtoledo/)[![Facebook](https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white)](https://www.facebook.com/totoledao)[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://www.instagram.com/totoledao)[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=whit)](https://www.github.com/totoledao)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[license-shield]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: https://github.com/totoledao/brew-journal/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=0e76a8
[linkedin-url]: http://www.linkedin.com/in/guilhermemtoledo
[htmx-shield]: https://img.shields.io/badge/htmx-white.svg?style=for-the-badge&logo=htmx&logoColor=%23F7DF1E
[htmx-url]: https://htmx.org/
[go-shield]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/
