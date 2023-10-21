# GoSber
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)
## Overview

GoSber is a versatile GoLang script for web scraping. It offers multiple modes for scraping, the ability to specify search prompts, and the option to scrape specific links. Additionally, it can automatically connect to a PostgreSQL database using the `DB_URL` environment variable.

## Features

- Selection of different modes for customizable scraping
- Search prompts to target specific content
- Scraping of data from specific URLs
- Automatic connection to a PostgreSQL database using `DB_URL`

## Usage

To get started with sber-scrape, follow the instructions below:

### Using pre-build executables
1. Download appropriate executable depending on your platform from [releases](https://github.com/malvere/GoSber/releases)

2. Run with with flags if needed
    ```shell
    ./sber-scrape -m <mode> -s <search-prompt> -u <url>
    ```

### Building from source

1. Clone this repository:

   ```shell
   git clone https://github.com/malvere/GoSberScrape.git
   cd GoSberScrape
   ```

2. Build the project

    ```shell
        go build
    ```
3. Run
    ```shell
    ./sber-scrape -m <mode> -s <search-prompt> -u <url>
    ```
## PostgreSQL Connection
If you have a PostgreSQL database, sber-scrape can connect to it by setting the DB_URL environment variable. The script will use it to establish a connection.

```shell
export DB_URL="postgres://username:password@localhost/database"
./sber-scrape
```

## .csv support

If `DB_URL` is not specified, a .csv file with parse results will be generated near the executable.

## License
MIT

## Contribution
Contributions are welcome! Feel free to open issues and submit pull requests to help improve this project.
