# mulef

This tool performs OSINT by finding LinkedIn employees on GitHub. It has two modes: one for searching keywords on the GitHub profiles of the users you're searching for, and the other for scraping the location of the employee from LinkedIn, searching for the name of the employee, and then checking if their location on GitHub matches the one on LinkedIn.

## Installation

To use this tool, you need to have Go installed on your system. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

Once you have installed Go, you can download and install this tool by running the following command:

```
go install -v github.com/mux0x/mulef@latest
```

### Usage

To use this tool, you need to provide the following command-line arguments:

```
-keywords: comma-separated list of keywords
-mode: mode of finding employees (location, keywords)
-LinkedInRequest: path of the LinkedIn request file
-token: GitHub token
-output: path of the output file
```

### To get that LinkedIn request

1. you open the dev tools in the browser you’re using
2. open corp page on LinkedIn
3. click See all x,xxx employees on LinkedIn
4. Search for `/voyager/api/graphql?includeWebMetadata=true&variables=(start:0,origin:COMPANY_PAGE_CANNED_SEARCH,query` in the requests in the network tab
5. Right-Click on the request and Copy > Copy Request headers
6. save it in a txt file

### Example Usage

```
mulef -keywords indrive -mode keywords -LinkedInRequest /path/to/linkedin_request.html -token="ghp_xxxxxxxxxxxxxxxx" -output /path/to/output.txt
```

This will search for LinkedIn employees on GitHub who have the keywords "indrive” in their profiles,code, and repos details and output the results to the file `/path/to/output.txt`.

### Modes

### Keyword Mode

In this mode, the tool will search for the provided keywords on the GitHub profiles of the users you're searching for. To use this mode, set the `-mode` flag to "keywords" and provide the keywords as a comma-separated list using the `-keywords` flag.

### Location Mode

In this mode, the tool will scrape the location of the employee from LinkedIn, search for the name of the employee, and then check if their location on GitHub matches the one on LinkedIn. To use this mode, set the `-mode` flag to "location" and provide the path of the LinkedIn request file using the `-LinkedInRequest` flag.

Made by love from a Muslim <3
